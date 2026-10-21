// A pyramid of cubes, rendered with ebiten's DrawTriangles and a simple perspective camera.
package main

import (
	"cmp"
	"fmt"
	"image"
	"image/color"
	"math"
	"slices"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	b3 "github.com/oliverbestmann/box3d-go"
)

const help = `arrows: rotate camera
enter: shoot a cube
r: reset
space: pause`

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(1200, 800)
	ebiten.SetWindowTitle("box3d-go pyramid")

	game := &game{yaw: 0.6, pitch: 0.35}
	game.reset()

	if err := ebiten.RunGame(game); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

type object struct {
	body  b3.Body
	hull  *b3.HullData
	color b3.Vec3
}

type game struct {
	world   b3.World
	objects []object
	ground  object

	paused      bool
	physicsTime time.Duration

	yaw, pitch float64

	vertices []ebiten.Vertex
	indices  []uint16
	faces    []face
}

func (g *game) reset() {
	if g.world.Id != (b3.WorldId{}) {
		g.world.DestroyWorld()
	}

	def := b3.DefaultWorldDef()
	def.Gravity = b3.Vec3{Y: -10}
	g.world = b3.CreateWorld(def)
	g.objects = nil

	groundHull := b3.MakeOffsetBoxHull(20, 1, 20, b3.Vec3{Y: -1})
	g.ground = g.addBody(b3.StaticBody, b3.Pos{}, &groundHull.Base, b3.Vec3{X: 0.4, Y: 0.45, Z: 0.4})

	// a square pyramid, layer l has (layers-l)^2 cubes
	const layers = 8
	cube := b3.MakeCubeHull(0.5)
	for l := range layers {
		n := layers - l
		for i := range n {
			for k := range n {
				p := b3.Pos{
					X: float32(i) - float32(n-1)/2,
					Y: 0.5 + float32(l),
					Z: float32(k) - float32(n-1)/2,
				}

				c := b3.Vec3{X: 0.9, Y: 0.5 + 0.4*float32(l)/layers, Z: 0.2}
				g.objects = append(g.objects, g.addBody(b3.DynamicBody, p, &cube.Base, c))
			}
		}
	}
}

func (g *game) addBody(bodyType b3.BodyType, position b3.Pos, hull *b3.HullData, color b3.Vec3) object {
	bodyDef := b3.DefaultBodyDef()
	bodyDef.Type1 = bodyType
	bodyDef.Position = position
	body := g.world.CreateBody(bodyDef)

	shape := body.CreateHullShape(b3.DefaultShapeDef(), hull)

	// box3d copies the hull, render the copy, it outlives the hull passed in
	return object{body: body, hull: shape.GetHull(), color: color}
}

func (g *game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.yaw -= 0.02
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.yaw += 0.02
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.pitch = min(g.pitch+0.01, 1.4)
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.pitch = max(g.pitch-0.01, 0.05)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.reset()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.paused = !g.paused
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		cam := g.camera(1, 1)
		cube := b3.MakeCubeHull(0.75)
		obj := g.addBody(b3.DynamicBody, cam.eye, &cube.Base, b3.Vec3{X: 0.2, Y: 0.5, Z: 0.9})
		obj.body.SetLinearVelocity(b3.MulSV(40, b3.Normalize(b3.Sub(cam.target, cam.eye))))
		g.objects = append(g.objects, obj)
	}

	if !g.paused {
		start := time.Now()
		g.world.Step(1.0/60.0, 4)
		g.physicsTime = (g.physicsTime*19 + time.Since(start)) / 20
	}

	return nil
}

type camera struct {
	eye, target    b3.Vec3
	right, up, fwd b3.Vec3
	focal, cx, cy  float32
}

func (g *game) camera(w, h float32) camera {
	const distance = 22

	target := b3.Vec3{Y: 3}
	eye := b3.Add(target, b3.Vec3{
		X: float32(distance * math.Cos(g.pitch) * math.Sin(g.yaw)),
		Y: float32(distance * math.Sin(g.pitch)),
		Z: float32(distance * math.Cos(g.pitch) * math.Cos(g.yaw)),
	})

	fwd := b3.Normalize(b3.Sub(target, eye))
	right := b3.Normalize(b3.Cross(fwd, b3.Vec3{Y: 1}))
	up := b3.Cross(right, fwd)

	return camera{eye: eye, target: target, right: right, up: up, fwd: fwd, focal: h, cx: w / 2, cy: h / 2}
}

// project returns the screen position and the depth of p. The depth is negative behind the camera.
func (c camera) project(p b3.Vec3) (x, y, depth float32) {
	d := b3.Sub(p, c.eye)
	depth = b3.Dot(d, c.fwd)
	x = c.cx + c.focal*b3.Dot(d, c.right)/depth
	y = c.cy - c.focal*b3.Dot(d, c.up)/depth
	return x, y, depth
}

// face is a visible hull face, sorted by depth to draw back to front.
type face struct {
	depth  float32
	points []b3.Vec3
	color  b3.Vec3
}

var lightDir = b3.Normalize(b3.Vec3{X: 0.3, Y: 1, Z: 0.5})

// collectFaces adds the faces of obj that point towards the camera.
func (g *game) collectFaces(cam camera, obj object) {
	tr := obj.body.GetTransform()
	if obj.body.IsAwake() == 0 {
		obj.color = b3.MulSV(0.6, obj.color)
	}

	points := obj.hull.Points()
	edges := obj.hull.Edges()
	planes := obj.hull.Planes()

	for i, hullFace := range obj.hull.Faces() {
		normal := b3.RotateVector(tr.Q, planes[i].Normal)

		var corners []b3.Vec3
		var center b3.Vec3

		edge := hullFace.Edge
		for {
			p := b3.TransformPoint(tr, points[edges[edge].Origin])
			corners = append(corners, p)
			center = b3.Add(center, p)

			edge = edges[edge].Next
			if edge == hullFace.Edge {
				break
			}
		}

		center = b3.MulSV(1/float32(len(corners)), center)
		if b3.Dot(normal, b3.Sub(cam.eye, center)) <= 0 {
			continue
		}

		light := 0.35 + 0.65*max(0, b3.Dot(normal, lightDir))
		g.faces = append(g.faces, face{
			depth:  b3.Dot(b3.Sub(center, cam.eye), cam.fwd),
			points: corners,
			color:  b3.MulSV(light, obj.color),
		})
	}
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x20, G: 0x24, B: 0x30, A: 0xff})

	w, h := float32(screen.Bounds().Dx()), float32(screen.Bounds().Dy())
	cam := g.camera(w, h)

	// the ground is below everything else, draw it first instead of sorting its large faces
	g.faces = g.faces[:0]
	g.collectFaces(cam, g.ground)
	groundFaces := len(g.faces)

	for _, obj := range g.objects {
		g.collectFaces(cam, obj)
	}

	slices.SortFunc(g.faces[groundFaces:], func(a, b face) int { return cmp.Compare(b.depth, a.depth) })

	g.vertices, g.indices = g.vertices[:0], g.indices[:0]
	for _, f := range g.faces {
		// DrawTriangles takes uint16 indices
		if len(g.vertices)+len(f.points) > math.MaxUint16 {
			g.flush(screen)
		}

		base := uint16(len(g.vertices))
		for _, p := range f.points {
			x, y, depth := cam.project(p)
			if depth <= 0.1 {
				// shortcut: faces crossing the near plane are dropped, not clipped
				g.vertices = g.vertices[:base]
				break
			}

			g.vertices = append(g.vertices, ebiten.Vertex{
				DstX: x, DstY: y, SrcX: 1, SrcY: 1,
				ColorR: f.color.X, ColorG: f.color.Y, ColorB: f.color.Z, ColorA: 1,
			})
		}

		for i := base + 2; i < uint16(len(g.vertices)); i++ {
			g.indices = append(g.indices, base, i-1, i)
		}
	}

	g.flush(screen)

	text := fmt.Sprintf("step: %1.2fms\nbodies: %d", float64(g.physicsTime.Microseconds())/1000, len(g.objects))
	if g.paused {
		text += "\npaused"
	}

	ebitenutil.DebugPrintAt(screen, text, 16, 16)
	ebitenutil.DebugPrintAt(screen, help, 16, int(h)-16*5)
}

func (g *game) flush(screen *ebiten.Image) {
	screen.DrawTriangles(g.vertices, g.indices, whiteImage, &ebiten.DrawTrianglesOptions{AntiAlias: true})
	g.vertices, g.indices = g.vertices[:0], g.indices[:0]
}

// whiteImage is the texture for DrawTriangles, the vertex colors tint it.
var whiteImage = func() *ebiten.Image {
	img := ebiten.NewImage(3, 3)
	img.Fill(color.White)
	return img.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
}()

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
