package b3

import (
	"runtime"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func queryWorld(t *testing.T) (World, []Shape) {
	world := CreateWorld(DefaultWorldDef())
	t.Cleanup(world.DestroyWorld)

	cube := MakeCubeHull(1)
	var shapes []Shape
	for i := range 3 {
		bodyDef := DefaultBodyDef()
		bodyDef.Position = Pos{X: float32(i) * 3}
		shapes = append(shapes, world.CreateBody(bodyDef).CreateHullShape(DefaultShapeDef(), &cube.Base))
	}

	return world, shapes
}

// growStack recurses to force the goroutine stack to grow and move, then calls fn.
//
//go:noinline
func growStack(depth int, fn func()) int {
	var pad [256]byte
	if depth == 0 {
		fn()
		return int(pad[0])
	}

	return growStack(depth-1, fn) + int(pad[depth%len(pad)])
}

// The proxy points live in a local array. Inside the callback the stack grows and the GC runs,
// box3d must still see the right points afterwards.
func TestProxySurvivesStackGrowthAndGC(t *testing.T) {
	world, shapes := queryWorld(t)

	for range 20 {
		points := [4]Vec3{{X: 6}, {X: 6.1}, {X: 6, Y: 0.1}, {X: 6, Z: 0.1}}
		proxy := MakeProxy(points[:], 0.1)

		var hits []Shape
		world.OverlapShape(Pos{}, proxy, DefaultQueryFilter(), func(shape Shape) bool {
			growStack(2000, runtime.GC)
			hits = append(hits, shape)
			return true
		})
		require.Equal(t, shapes[2:], hits)

		// a second query with the same proxy after the stack moved
		growStack(2000, func() {
			out := ShapeDistance(DistanceInput{
				ProxyA:    proxy,
				ProxyB:    MakeProxy([]Vec3{{X: 9}}, 0),
				Transform: TransformIdentity,
			}, nil, nil)
			require.InDelta(t, 2.9, out.Distance, 1e-4)
		})
	}
}

// A panicking callback must leave the box3d stack usable.
func TestCallbackPanicKeepsStackUsable(t *testing.T) {
	world, shapes := queryWorld(t)

	func() {
		defer func() { assert.Equal(t, "boom", recover()) }()

		world.OverlapShape(Pos{X: 6}, MakeProxy([]Vec3{{}}, 0.5), DefaultQueryFilter(), func(shape Shape) bool {
			panic("boom")
		})
	}()

	theStack.CheckEmpty()

	var hits []Shape
	world.OverlapShape(Pos{X: 6}, MakeProxy([]Vec3{{}}, 0.5), DefaultQueryFilter(), func(shape Shape) bool {
		hits = append(hits, shape)
		return true
	})
	assert.Equal(t, shapes[2:], hits)
}

func TestCopyToStackAsRejectsSizeMismatch(t *testing.T) {
	assert.Panics(t, func() { copyToStackAs[Vec3](theStack, ShapeProxy{}) })
	theStack.CheckEmpty()
}

// copyToStackAs must move the pointees to the heap. If points stayed on the goroutine stack,
// growStack would move it and the address in the copy would be stale.
func TestCopyToStackAsMovesPointeesToHeap(t *testing.T) {
	points := [4]Vec3{{X: 1}, {X: 2}, {X: 3}, {X: 4}}
	proxy := ShapeProxy{Points: &points[0], Count: 4}

	a := copyToStackAs[cShapeProxy](theStack, proxy)
	defer a.Free()

	growStack(2000, func() {})

	stored := (*cShapeProxy)(unsafe.Pointer(a.Addr)).Points
	require.Equal(t, uintptr(unsafe.Pointer(&points[0])), stored)
}
