package b3

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Ported from box3d test/test_world.c

func TestHelloWorld(t *testing.T) {
	worldDef := DefaultWorldDef()
	worldDef.Gravity = Vec3{Y: -10}
	world := CreateWorld(worldDef)
	require.EqualValues(t, 1, world.IsValid())

	groundDef := DefaultBodyDef()
	groundDef.Position = Pos{Y: -10}
	ground := world.CreateBody(groundDef)
	require.EqualValues(t, 1, ground.IsValid())

	groundBox := MakeBoxHull(50, 10, 50)
	ground.CreateHullShape(DefaultShapeDef(), &groundBox.Base)

	bodyDef := DefaultBodyDef()
	bodyDef.Type1 = DynamicBody
	bodyDef.Position = Pos{Y: 4}
	body := world.CreateBody(bodyDef)

	dynamicBox := MakeCubeHull(1)
	shapeDef := DefaultShapeDef()
	shapeDef.Density = 1
	shapeDef.BaseMaterial.Friction = 0.3
	body.CreateHullShape(shapeDef, &dynamicBox.Base)

	for range 90 {
		world.Step(1.0/60.0, 4)
	}

	position := body.GetPosition()
	rotation := body.GetRotation()

	world.DestroyWorld()
	assert.EqualValues(t, 0, world.IsValid())

	assert.InDelta(t, 1, position.Y, 0.01)
	assert.InDelta(t, 0, rotation.V.X, 0.01)
	assert.InDelta(t, 0, rotation.V.Z, 0.01)
}

func TestDestroyAllBodiesWorld(t *testing.T) {
	const bodyCount = 10

	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	bodyDef := DefaultBodyDef()
	bodyDef.Type1 = DynamicBody
	cube := MakeCubeHull(0.5)

	var bodies []Body
	creating := true

	for range 2*bodyCount + 10 {
		switch {
		case creating && len(bodies) < bodyCount:
			body := world.CreateBody(bodyDef)
			body.CreateHullShape(DefaultShapeDef(), &cube.Base)
			bodies = append(bodies, body)

		case creating:
			creating = false

		case len(bodies) > 0:
			bodies[len(bodies)-1].DestroyBody()
			bodies = bodies[:len(bodies)-1]
		}

		world.Step(1.0/60.0, 3)
	}

	assert.EqualValues(t, 0, world.GetCounters().BodyCount)
}

func TestSensorEvents(t *testing.T) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	// wall from x = 1 to x = 2
	bodyDef := DefaultBodyDef()
	bodyDef.Position = Pos{X: 1.5, Y: 11}
	wall := world.CreateBody(bodyDef)
	box := MakeBoxHull(0.5, 10, 1)
	shapeDef := DefaultShapeDef()
	shapeDef.EnableSensorEvents = 1
	wall.CreateHullShape(shapeDef, &box.Base)

	// bullet fired towards the wall
	bodyDef = DefaultBodyDef()
	bodyDef.Type1 = DynamicBody
	bodyDef.IsBullet = 1
	bodyDef.GravityScale = 0
	bodyDef.Position = Pos{X: 7.39814, Y: 4}
	bodyDef.LinearVelocity = Vec3{X: -20}
	bullet := world.CreateBody(bodyDef)

	shapeDef = DefaultShapeDef()
	shapeDef.IsSensor = 1
	shapeDef.EnableSensorEvents = 1
	bullet.CreateSphereShape(shapeDef, Sphere{Radius: 0.1})

	var beginCount, endCount int
	for bullet.GetPosition().X >= -1 {
		world.Step(1.0/60.0, 4)

		events := world.GetSensorEvents()
		if len(events.BeginEvents) > 0 {
			beginCount++
		}

		if len(events.EndEvents) > 0 {
			endCount++
		}
	}

	assert.Equal(t, 1, beginCount)
	assert.Equal(t, 1, endCount)
}

func TestContactAndHitEvents(t *testing.T) {
	worldDef := DefaultWorldDef()
	worldDef.HitEventThreshold = 1
	world := CreateWorld(worldDef)
	defer world.DestroyWorld()

	groundDef := DefaultBodyDef()
	groundDef.Position = Pos{Y: -0.5}
	ground := world.CreateBody(groundDef)
	groundBox := MakeBoxHull(10, 0.5, 10)
	groundShape := ground.CreateHullShape(DefaultShapeDef(), &groundBox.Base)

	bodyDef := DefaultBodyDef()
	bodyDef.Type1 = DynamicBody
	bodyDef.Position = Pos{Y: 2}
	bodyDef.LinearVelocity = Vec3{Y: -30}
	body := world.CreateBody(bodyDef)

	shapeDef := DefaultShapeDef()
	shapeDef.Density = 1
	shapeDef.EnableContactEvents = 1
	shapeDef.EnableHitEvents = 1
	shapeDef.BaseMaterial.UserMaterialId = 7
	shapeDef.BaseMaterial.Restitution = 0.6
	sphere := body.CreateSphereShape(shapeDef, Sphere{Radius: 0.5})

	var begin []ContactBeginTouchEvent
	var hits []ContactHitEvent
	endCount := 0

	for range 120 {
		world.Step(1.0/60.0, 4)

		events := world.GetContactEvents()
		begin = append(begin, events.BeginEvents...)
		hits = append(hits, events.HitEvents...)
		endCount += len(events.EndEvents)
	}

	require.NotEmpty(t, begin)
	assert.ElementsMatch(t, []ShapeId{groundShape.Id, sphere.Id}, []ShapeId{begin[0].ShapeIdA, begin[0].ShapeIdB})
	assert.GreaterOrEqual(t, endCount, 1)

	require.NotEmpty(t, hits)
	assert.Greater(t, hits[0].ApproachSpeed, float32(1))
	assert.InDelta(t, 0, hits[0].Normal.X, 0.01)
	assert.InDelta(t, 0, hits[0].Normal.Z, 0.01)
	assert.Contains(t, []uint64{hits[0].UserMaterialIdA, hits[0].UserMaterialIdB}, uint64(7))
}

func TestSetWorkerCount(t *testing.T) {
	def := DefaultWorldDef()
	def.WorkerCount = 1
	world := CreateWorld(def)
	defer world.DestroyWorld()

	ground := world.CreateBody(DefaultBodyDef())
	groundBox := MakeBoxHull(20, 1, 20)
	ground.CreateHullShape(DefaultShapeDef(), &groundBox.Base)

	cube := MakeCubeHull(0.5)
	for idx := range 50 {
		bodyDef := DefaultBodyDef()
		bodyDef.Type1 = DynamicBody
		bodyDef.Position = Pos{X: float32(idx%10) - 5, Y: 2 + float32(idx/10)}
		world.CreateBody(bodyDef).CreateHullShape(DefaultShapeDef(), &cube.Base)
	}

	for _, workers := range []int32{4, 4, 0, -5, _B3_MAX_WORKERS + 10, 2} {
		world.SetWorkerCount(workers)
		assert.EqualValues(t, min(max(workers, 1), _B3_MAX_WORKERS), world.GetWorkerCount())

		for range 30 {
			world.Step(1.0/60.0, 4)
		}
	}
}
