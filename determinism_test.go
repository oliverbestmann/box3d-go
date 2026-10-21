package b3

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The scenes of box3d's test/test_determinism.c, translated from shared/ into box3d_c-scenes_test.go.
// The expected values are the goldens of the C test, so this checks that the translated solver
// produces bit-identical results.

type sceneResult struct {
	SleepStep     int32
	Hash          uint32
	QueryHitCount int32
	QueryHash     uint32
}

// runScene steps a world with the scene until update reports that it is done, at most maxSteps times.
func runScene[T any](t *testing.T, workerCount uint32, maxSteps int,
	create func(*_Stack, WorldId) T,
	update func(*_Stack, WorldId, uintptr) uint8,
	destroy func(*_Stack, uintptr),
	result func(*T) sceneResult,
) sceneResult {
	def := DefaultWorldDef()
	def.WorkerCount = workerCount
	world := CreateWorld(def)

	// the scene code holds the data as a uintptr
	data := new(T)
	escapes(data)
	*data = create(theStack, world.Id)

	done := false
	for range maxSteps {
		world.Step(1.0/60.0, 4)

		if update(theStack, world.Id, uintptr(unsafe.Pointer(data))) != 0 {
			done = true
			break
		}
	}

	world.DestroyWorld()

	require.True(t, done, "scene did not finish within %d steps", maxSteps)

	r := result(data)
	destroy(theStack, uintptr(unsafe.Pointer(data)))
	return r
}

func TestDeterminismFallingRagdolls(t *testing.T) {
	for workers := uint32(1); workers < 6; workers++ {
		r := runScene(t, workers, 500, _CreateFallingRagdolls, _UpdateFallingRagdolls, _DestroyFallingRagdolls,
			func(d *_FallingRagdollData) sceneResult { return sceneResult{SleepStep: d.SleepStep, Hash: d.Hash} })

		assert.Equal(t, sceneResult{SleepStep: 291, Hash: 0xFF4C605B}, r, "workers=%d hash=%#x", workers, r.Hash)
	}
}

func TestDeterminismWavePile(t *testing.T) {
	for workers := uint32(1); workers <= 4; workers++ {
		r := runScene(t, workers, 500, _CreateWavePile, _UpdateWavePile, _DestroyWavePile,
			func(d *_WavePileData) sceneResult { return sceneResult{SleepStep: d.SleepStep, Hash: d.Hash} })

		assert.Equal(t, sceneResult{SleepStep: 272, Hash: 0x266F83B0}, r, "workers=%d hash=%#x", workers, r.Hash)
	}
}

func TestDeterminismQuerySpawn(t *testing.T) {
	for workers := uint32(1); workers <= 4; workers++ {
		var spawnCount int32
		r := runScene(t, workers, 1000, _CreateQuerySpawn, _UpdateQuerySpawn, _DestroyQuerySpawn,
			func(d *_QuerySpawnData) sceneResult {
				spawnCount = d.SpawnCount
				return sceneResult{SleepStep: d.SleepStep, Hash: d.Hash, QueryHitCount: d.QueryHitCount, QueryHash: d.QueryHash}
			})

		assert.EqualValues(t, _QUERY_SPAWN_COUNT, spawnCount)

		expected := sceneResult{SleepStep: 242, Hash: 0xF3AD9BA6, QueryHitCount: 59, QueryHash: 0xE3271F3D}
		assert.Equal(t, expected, r, "workers=%d hash=%#x queryHash=%#x", workers, r.Hash, r.QueryHash)
	}
}

func TestDeterminismMeshDrop(t *testing.T) {
	createMeshDrop := func(tls *_Stack, worldId WorldId) _MeshDropData {
		return _CreateMeshDrop(tls, worldId, Pos{})
	}

	for _, workers := range []uint32{1, 4} {
		r := runScene(t, workers, 400, createMeshDrop, _UpdateMeshDrop, _DestroyMeshDrop,
			func(d *_MeshDropData) sceneResult { return sceneResult{SleepStep: d.SleepStep, Hash: d.Hash} })

		assert.Equal(t, sceneResult{SleepStep: 226, Hash: 0xD2FF069D}, r, "workers=%d hash=%#x", workers, r.Hash)
	}
}

// Ported from SingleRollingMixTest in test/test_determinism.c, using the Go API. The C test has no
// golden hash, every worker count must give the same result.
func TestDeterminismRollingMix(t *testing.T) {
	var reference uint32
	for workers := uint32(1); workers <= 4; workers++ {
		hash := runRollingMix(workers)
		if workers == 1 {
			reference = hash
		}

		assert.Equal(t, reference, hash, "workers=%d", workers)
	}
}

func runRollingMix(workerCount uint32) uint32 {
	def := DefaultWorldDef()
	def.WorkerCount = workerCount
	world := CreateWorld(def)
	defer world.DestroyWorld()

	ground := world.CreateBody(DefaultBodyDef())
	groundBox := MakeBoxHull(40, 1, 40)
	ground.CreateHullShape(DefaultShapeDef(), &groundBox.Base)

	const rowCount = 8
	const columnCount = 8
	const bodyCount = rowCount * columnCount

	box := MakeBoxHull(0.4, 0.4, 0.4)
	sphere := Sphere{Radius: 0.4}
	capsule := Capsule{Center1: Vec3{X: -0.3}, Center2: Vec3{X: 0.3}, Radius: 0.25}

	var bodies []Body

	seed := uint32(271828)
	for i := range bodyCount {
		row := i / columnCount
		column := i % columnCount

		bodyDef := DefaultBodyDef()
		bodyDef.Type1 = DynamicBody
		bodyDef.Position = Pos{X: 2*float32(column) - 7, Y: 1.45, Z: 2*float32(row) - 7}
		body := world.CreateBody(bodyDef)
		bodies = append(bodies, body)

		seed = seed*1664525 + 1013904223
		shapeDef := DefaultShapeDef()
		shapeDef.Density = 1
		if seed>>28 < 5 {
			shapeDef.BaseMaterial.RollingResistance = 0.2
		}

		switch i % 3 {
		case 0:
			body.CreateSphereShape(shapeDef, sphere)
		case 1:
			body.CreateCapsuleShape(shapeDef, capsule)
		default:
			body.CreateHullShape(shapeDef, &box.Base)
		}

		sign := float32(1)
		if i&1 == 1 {
			sign = -1
		}

		negZero := float32(-0.0)
		body.SetLinearVelocity(Vec3{X: negZero, Y: -0.5, Z: sign * 0.5})
		body.SetAngularVelocity(Vec3{X: sign * 3, Y: negZero, Z: -2 * sign})
	}

	hash := uint32(_B3_HASH_INIT)
	hashOf := func(value any, size uintptr) {
		buf := theStack.Alloc(int(size))
		defer theStack.Free(int(size))

		switch value := value.(type) {
		case WorldTransform:
			*(*WorldTransform)(unsafe.Pointer(buf)) = value
		case Vec3:
			*(*Vec3)(unsafe.Pointer(buf)) = value
		case float32:
			*(*float32)(unsafe.Pointer(buf)) = value
		}

		hash = b3Hash(theStack, hash, buf, int32(size))
	}

	var contacts []ContactData
	for range 120 {
		world.Step(1.0/60.0, 4)

		for _, body := range bodies {
			hashOf(body.GetTransform(), unsafe.Sizeof(WorldTransform{}))

			contacts = body.GetContactData(contacts)
			for _, contact := range contacts[:min(len(contacts), 4)] {
				for _, manifold := range contact.GetManifolds() {
					hashOf(manifold.RollingImpulse, 12)
					hashOf(manifold.FrictionImpulse, 12)
					hashOf(manifold.TwistImpulse, 4)

					for _, point := range manifold.Points[:manifold.PointCount] {
						hashOf(point.NormalImpulse, 4)
						hashOf(point.TotalNormalImpulse, 4)
					}
				}
			}
		}
	}

	return hash
}
