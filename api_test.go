package b3

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tests for the hand-written parts of the Go API. Some are ported from box3d test/test_*.c.

func TestHullAccessors(t *testing.T) {
	box := MakeBoxHull(1, 2, 3)
	hull := &box.Base

	assert.Len(t, hull.Points(), 8)
	assert.Len(t, hull.Vertices(), 8)
	assert.Len(t, hull.Edges(), 24)
	assert.Len(t, hull.Faces(), 6)
	assert.Len(t, hull.Planes(), 6)

	for _, p := range hull.Points() {
		assert.InDelta(t, 1, abs(p.X), 1e-6)
		assert.InDelta(t, 2, abs(p.Y), 1e-6)
		assert.InDelta(t, 3, abs(p.Z), 1e-6)
	}

	// every point of a face lies on the plane of the face
	edges := hull.Edges()
	for i, face := range hull.Faces() {
		plane := hull.Planes()[i]
		edge := face.Edge
		for {
			p := hull.Points()[edges[edge].Origin]
			assert.InDelta(t, plane.Offset, Dot(plane.Normal, p), 1e-5)

			edge = edges[edge].Next
			if edge == face.Edge {
				break
			}
		}
	}
}

func TestCreateHull(t *testing.T) {
	points := []Vec3{
		{X: -1, Y: -1, Z: -1},
		{X: 1, Y: -1, Z: -1},
		{X: -1, Y: 1, Z: -1},
		{X: 1, Y: 1, Z: -1},
		{X: -1, Y: -1, Z: 1},
		{X: 1, Y: -1, Z: 1},
		{X: -1, Y: 1, Z: 1},
		{X: 1, Y: 1, Z: 1},
		{}, // inside
	}

	hull := CreateHull(points, 32)
	require.NotNil(t, hull.HullData)
	defer hull.Destroy()

	assert.EqualValues(t, 8, hull.VertexCount)
	assert.InDelta(t, 8, hull.Volume, 1e-4)

	mass := ComputeHullMass(hull.HullData, 2)
	assert.InDelta(t, 16, mass.Mass, 1e-3)

	out := RayCastHull(hull.HullData, RayCastInput{Origin: Vec3{X: -5}, Translation: Vec3{X: 10}, MaxFraction: 1})
	assert.EqualValues(t, 1, out.Hit)
	assert.InDelta(t, 0.4, out.Fraction, 1e-5)
}

// Ported from box3d test/test_mesh.c
func TestCreateMesh(t *testing.T) {
	vertices := []Vec3{{}, {X: 1}, {X: 1, Z: 1}, {Z: 1}}
	indices := []int32{0, 2, 1, 0, 3, 2}

	def := MeshDef{IdentifyEdges: 1}
	mesh := CreateMesh(def, vertices, indices, []uint8{1, 2})
	require.NotNil(t, mesh.MeshData)
	defer mesh.Destroy()

	assert.EqualValues(t, 2, mesh.TriangleCount)
	assert.Equal(t, vertices, mesh.Vertices())
	assert.Len(t, mesh.Triangles(), 2)
	assert.Equal(t, []uint8{1, 2}, mesh.MaterialIndices())
	assert.Len(t, mesh.Flags(), 2)

	var triangles []int32
	QueryMesh(Mesh{Data: mesh.MeshData, Scale: Vec3{X: 1, Y: 1, Z: 1}}, AABB{UpperBound: Vec3{X: 1, Y: 1, Z: 1}}, func(a, b, c Vec3, triangleIndex int32) bool {
		triangles = append(triangles, triangleIndex)
		return true
	})
	assert.ElementsMatch(t, []int32{0, 1}, triangles)

	out := RayCastMesh(Mesh{Data: mesh.MeshData, Scale: Vec3{X: 1, Y: 1, Z: 1}}, RayCastInput{
		Origin:      Vec3{X: 0.25, Y: 1, Z: 0.5},
		Translation: Vec3{Y: -2},
		MaxFraction: 1,
	})
	assert.EqualValues(t, 1, out.Hit)
	assert.InDelta(t, 0.5, out.Fraction, 1e-5)
	assert.InDelta(t, 1, out.Normal.Y, 1e-5)
}

func TestCreateMeshPanicsOnMaterialCount(t *testing.T) {
	vertices := []Vec3{{}, {X: 1}, {Z: 1}}
	assert.Panics(t, func() { CreateMesh(MeshDef{}, vertices, []int32{0, 2, 1}, []uint8{1, 2}) })
}

// Ported from RayCastFlatField in box3d test/test_height_field.c
func TestCreateHeightField(t *testing.T) {
	def := HeightFieldDef{
		Scale:               Vec3{X: 1, Y: 1, Z: 1},
		CountX:              4,
		CountZ:              4,
		GlobalMinimumHeight: -1,
		GlobalMaximumHeight: 1,
	}

	hf := CreateHeightField(def, make([]float32, 16), make([]uint8, 9))
	require.NotNil(t, hf.HeightFieldData)
	defer hf.Destroy()

	out := RayCastHeightField(hf.HeightFieldData, RayCastInput{
		Origin:      Vec3{X: 1.25, Y: 10, Z: 1.25},
		Translation: Vec3{Y: -20},
		MaxFraction: 1,
	})

	assert.EqualValues(t, 1, out.Hit)
	assert.InDelta(t, 0.5, out.Fraction, 1e-5)
	assert.InDelta(t, 1, out.Normal.Y, 1e-5)

	var count int
	QueryHeightField(hf.HeightFieldData, AABB{LowerBound: Vec3{X: 0.1, Y: -1, Z: 0.1}, UpperBound: Vec3{X: 0.9, Y: 1, Z: 0.9}}, func(a, b, c Vec3, triangleIndex int32) bool {
		count++
		return true
	})
	assert.Equal(t, 2, count)

	assert.Panics(t, func() { CreateHeightField(def, make([]float32, 15), nil) })
}

// Ported from CompoundCreateMixed in box3d test/test_compound.c
func TestCreateCompound(t *testing.T) {
	mat := DefaultSurfaceMaterial()
	mat.UserMaterialId = 42

	box := MakeBoxHull(0.5, 0.5, 0.5)
	meshData := CreateBoxMesh(Vec3{}, Vec3{X: 0.5, Y: 0.5, Z: 0.5}, 0)
	defer meshData.Destroy()

	compound := CreateCompound(
		[]CompoundCapsuleDef{{Capsule: Capsule{Center1: Vec3{X: -1}, Center2: Vec3{X: 1}, Radius: 0.25}, Material: mat}},
		[]CompoundHullDef{{Hull: &box.Base, Transform: TransformIdentity, Material: mat}},
		[]CompoundMeshDef{{MeshData: meshData.MeshData, Transform: TransformIdentity, Scale: Vec3{X: 1, Y: 1, Z: 1}, Materials: &mat, MaterialCount: 1}},
		[]CompoundSphereDef{
			{Sphere: Sphere{Center: Vec3{X: 5}, Radius: 0.5}, Material: mat},
			{Sphere: Sphere{Center: Vec3{X: -5}, Radius: 0.5}, Material: mat},
		},
	)
	require.NotNil(t, compound.CompoundData)
	defer compound.Destroy()

	assert.EqualValues(t, 1, compound.CapsuleCount)
	assert.EqualValues(t, 1, compound.HullCount)
	assert.EqualValues(t, 1, compound.MeshCount)
	assert.EqualValues(t, 2, compound.SphereCount)

	require.Len(t, compound.Materials(), 1)
	assert.EqualValues(t, 42, compound.Materials()[0].UserMaterialId)

	shapeTypes := map[ShapeType]int{}
	for i := range compound.CapsuleCount + compound.HullCount + compound.MeshCount + compound.SphereCount {
		child := GetCompoundChild(compound.CompoundData, i)
		shapeTypes[child.Type1]++

		switch child.Type1 {
		case CapsuleShape:
			assert.InDelta(t, 0.25, child.Capsule().Radius, 1e-6)
		case HullShape:
			assert.EqualValues(t, 8, child.Hull().VertexCount)
		case MeshShape:
			assert.EqualValues(t, meshData.TriangleCount, child.Mesh().Data.TriangleCount)
		case SphereShape:
			assert.InDelta(t, 5, abs(child.Sphere().Center.X), 1e-6)
			assert.Panics(t, func() { child.Hull() })
		}
	}

	assert.Equal(t, map[ShapeType]int{CapsuleShape: 1, HullShape: 1, MeshShape: 1, SphereShape: 2}, shapeTypes)

	// a baked compound on a static body
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	body := world.CreateBody(DefaultBodyDef())
	def := DefaultShapeDef()
	shape := body.CreateBakedCompoundShape(def, compound.CompoundData)
	assert.Equal(t, CompoundShape, shape.GetType())
	assert.Equal(t, compound.CompoundData, shape.GetCompound())

	result := world.CastRayClosest(Pos{X: 5, Y: 5}, Vec3{Y: -10}, DefaultQueryFilter())
	assert.EqualValues(t, 1, result.Hit)
	assert.InDelta(t, 0.5, result.Point.Y, 1e-4)
}

func TestNamesAndUserData(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	world.SetUserData(1)
	assert.EqualValues(t, 1, world.GetUserData())

	body := world.CreateBody(DefaultBodyDef())
	assert.Equal(t, "", body.GetName())

	body.SetName("ground")
	body.SetUserData(2)
	assert.Equal(t, "ground", body.GetName())
	assert.EqualValues(t, 2, body.GetUserData())

	cube := MakeCubeHull(1)
	shape := body.CreateHullShape(DefaultShapeDef(), &cube.Base)
	shape.SetName("cube")
	shape.SetUserData(3)
	assert.Equal(t, "cube", shape.GetName())
	assert.EqualValues(t, 3, shape.GetUserData())
	assert.Equal(t, []Shape{shape}, body.GetShapes(nil))
}

func TestQueries(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	cube := MakeCubeHull(1)
	var shapes []Shape
	for i := range 3 {
		bodyDef := DefaultBodyDef()
		bodyDef.Position = Pos{X: float32(i) * 3}
		shapes = append(shapes, world.CreateBody(bodyDef).CreateHullShape(DefaultShapeDef(), &cube.Base))
	}

	var overlaps []Shape
	world.OverlapAABB(AABB{LowerBound: Vec3{X: -0.5, Y: -0.5, Z: -0.5}, UpperBound: Vec3{X: 3.5, Y: 0.5, Z: 0.5}}, DefaultQueryFilter(), func(shape Shape) bool {
		overlaps = append(overlaps, shape)
		return true
	})
	assert.ElementsMatch(t, shapes[:2], overlaps)

	overlaps = nil
	world.OverlapShape(Pos{X: 6}, MakeProxy([]Vec3{{}}, 0.5), DefaultQueryFilter(), func(shape Shape) bool {
		overlaps = append(overlaps, shape)
		return true
	})
	assert.Equal(t, shapes[2:], overlaps)

	// a ray along x hits all cubes, the callback returns the fraction to find all hits
	var hits []Shape
	world.CastRay(Pos{X: -5}, Vec3{X: 20}, DefaultQueryFilter(), func(shape Shape, point Pos, normal Vec3, fraction float32, userMaterialId uint64, triangleIndex, childIndex int32) float32 {
		hits = append(hits, shape)
		return 1
	})
	assert.ElementsMatch(t, shapes, hits)

	closest := world.CastRayClosest(Pos{X: -5}, Vec3{X: 20}, DefaultQueryFilter())
	assert.EqualValues(t, 1, closest.Hit)
	assert.Equal(t, shapes[0].Id, closest.ShapeId)
	assert.InDelta(t, -1, closest.Point.X, 1e-5)

	hits = nil
	world.CastShape(Pos{Y: 5}, MakeProxy([]Vec3{{}}, 0.5), Vec3{Y: -10}, DefaultQueryFilter(), func(shape Shape, point Pos, normal Vec3, fraction float32, userMaterialId uint64, triangleIndex, childIndex int32) float32 {
		hits = append(hits, shape)
		assert.InDelta(t, 1, normal.Y, 1e-4)
		return fraction
	})
	assert.Equal(t, shapes[:1], hits)
}

func TestCustomFilterCallback(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	var calls int
	world.SetCustomFilterCallback(func(shapeA, shapeB Shape) bool {
		calls++
		return false
	})

	shapeDef := DefaultShapeDef()
	shapeDef.EnableCustomFiltering = 1

	ground := world.CreateBody(DefaultBodyDef())
	groundBox := MakeBoxHull(10, 1, 10)
	ground.CreateHullShape(shapeDef, &groundBox.Base)

	bodyDef := DefaultBodyDef()
	bodyDef.Type1 = DynamicBody
	bodyDef.Position = Pos{Y: 2}
	body := world.CreateBody(bodyDef)
	cube := MakeCubeHull(0.5)
	body.CreateHullShape(shapeDef, &cube.Base)

	for range 60 {
		world.Step(1.0/60.0, 4)
	}

	// the filter rejects the contact, the cube falls through the ground
	assert.Positive(t, calls)
	assert.Less(t, body.GetPosition().Y, float32(-1))
}

func TestRevoluteJoint(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	ground := world.CreateBody(DefaultBodyDef())

	bodyDef := DefaultBodyDef()
	bodyDef.Type1 = DynamicBody
	bodyDef.Position = Pos{X: 1}
	body := world.CreateBody(bodyDef)
	cube := MakeCubeHull(0.25)
	body.CreateHullShape(DefaultShapeDef(), &cube.Base)

	def := DefaultRevoluteJointDef()
	def.Base.BodyIdA = ground.Id
	def.Base.BodyIdB = body.Id
	def.Base.LocalFrameA.Q = QuatIdentity
	def.Base.LocalFrameB = Transform{P: Vec3{X: -1}, Q: QuatIdentity}
	def.Base.UserData = 7

	joint := world.CreateRevoluteJoint(def)
	assert.Equal(t, JointTypeRevolute, joint.GetType())
	assert.EqualValues(t, 7, joint.GetUserData())
	assert.Equal(t, []Joint{joint.Joint}, body.GetJoints(nil))
	assert.Equal(t, joint, joint.Joint.AsRevoluteJoint())
	assert.Panics(t, func() { joint.Joint.AsWeldJoint() })

	hertz, damping := joint.GetConstraintTuning()
	assert.Positive(t, hertz)
	assert.Positive(t, damping)

	for range 60 {
		world.Step(1.0/60.0, 4)
	}

	// the body swings around the anchor at the origin
	assert.InDelta(t, 1, Length(body.GetPosition()), 1e-2)
	assert.Less(t, body.GetPosition().Y, float32(-0.1))
}

func TestDraw(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	body := world.CreateBody(DefaultBodyDef())
	cube := MakeCubeHull(1)
	shape := body.CreateHullShape(DefaultShapeDef(), &cube.Base)

	var drawn []Shape
	draw := DefaultDebugDraw()
	draw.DrawShapes = true
	draw.DrawingBounds = AABB{LowerBound: Vec3{X: -100, Y: -100, Z: -100}, UpperBound: Vec3{X: 100, Y: 100, Z: 100}}
	draw.DrawShape = func(shape Shape, transform WorldTransform, color HexColor) {
		drawn = append(drawn, shape)
	}

	world.Draw(draw, ^uint64(0))
	assert.Equal(t, []Shape{shape}, drawn)
	assert.Equal(t, MakeBoxHull(1, 1, 1).Base.Volume, drawn[0].GetHull().Volume)
}

func TestMathWrappers(t *testing.T) {
	q := MakeQuatFromAxisAngle(Vec3{Y: 1}, 0.5)
	axis, angle := GetAxisAngle(q)
	assert.InDelta(t, 1, axis.Y, 1e-5)
	assert.InDelta(t, 0.5, angle, 1e-2) // b3Atan2 is an approximation

	v := RotateVector(q, Vec3{X: 1})
	back := InvRotateVector(q, v)
	assert.InDelta(t, 1, back.X, 1e-5)

	n, length := GetLengthAndNormalize(Vec3{X: 3, Y: 4})
	assert.InDelta(t, 5, length, 1e-6)
	assert.InDelta(t, 0.6, n.X, 1e-6)

	aabb := MakeAABB([]Vec3{{X: -1}, {Y: 2}}, 0.5)
	assert.Equal(t, AABB{LowerBound: Vec3{X: -1.5, Y: -0.5, Z: -0.5}, UpperBound: Vec3{X: 0.5, Y: 2.5, Z: 0.5}}, aabb)

	assert.Equal(t, Vec3{X: 1, Y: 1}, Add(Vec3{X: 1}, Vec3{Y: 1}))
	assert.Equal(t, Vec3{Z: 1}, Cross(Vec3{X: 1}, Vec3{Y: 1}))
}

func abs(x float32) float32 {
	if x < 0 {
		return -x
	}

	return x
}
