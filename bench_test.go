package b3

import (
	"fmt"
	"testing"
	"unsafe"
)

// BenchmarkPyramid steps a settled pyramid of 210 cubes.
func BenchmarkPyramid(b *testing.B) {
	world := CreateWorld(DefaultWorldDef())
	defer world.DestroyWorld()

	ground := world.CreateBody(DefaultBodyDef())
	groundBox := MakeOffsetBoxHull(50, 1, 50, Vec3{Y: -1})
	ground.CreateHullShape(DefaultShapeDef(), &groundBox.Base)

	cube := MakeCubeHull(0.5)
	bodyDef := DefaultBodyDef()
	bodyDef.Type1 = DynamicBody
	bodyDef.EnableSleep = 0

	const rows = 20
	for row := range rows {
		for i := range rows - row {
			bodyDef.Position = Pos{X: float32(i) - float32(rows-row)/2, Y: 0.5 + float32(row)}
			world.CreateBody(bodyDef).CreateHullShape(DefaultShapeDef(), &cube.Base)
		}
	}

	for range 60 {
		world.Step(1.0/60.0, 4)
	}

	b.ResetTimer()
	for range b.N {
		world.Step(1.0/60.0, 4)
	}
}

// upstreamBenchmarks are the benchmarks of box3d's benchmark/main.c. The scenes come from
// shared/benchmarks.c, translated into box3d_c-scenes_test.go.
var upstreamBenchmarks = []struct {
	name     string
	capacity func(*_Stack, uintptr)
	create   func(*_Stack, WorldId)
	destroy  func(*_Stack)
	step     func(*_Stack, WorldId, int32)
	steps    int
}{
	{"convex_pile", _GetConvexPileCapacity, _CreateConvexPile, nil, nil, 500},
	{"joint_grid", nil, _CreateJointGrid, nil, nil, 100},
	{"junkyard", _GetJunkyardCapacity, _CreateJunkyard, nil, _StepJunkyard, 500},
	{"large_pyramid", nil, _CreateLargePyramid, nil, nil, 200},
	{"large_world", _GetLargeWorldCapacity, _CreateLargeWorld, nil, _StepLargeWorld, 500},
	{"many_pyramids", nil, _CreateManyPyramids, nil, nil, 100},
	{"mesh_drop", nil, _CreateMeshDropBenchmark, _DestroyMeshDropBenchmark, nil, 170},
	{"rain", _GetRainCapacity, _CreateRain, _DestroyRain, _StepRain, 400},
	{"sleep", _GetSleepCapacity, _CreateSleep, nil, _StepSleep, 300},
	{"spinner", _GetSpinnerCapacity, _CreateSpinner, _DestroySpinner, nil, 800},
	{"trees100", nil, _CreateTrees100, _DestroyTrees, nil, 500},
	{"trees50", nil, _CreateTrees50, _DestroyTrees, nil, 500},
	{"trees25", nil, _CreateTrees25, _DestroyTrees, nil, 500},
	{"village", _GetVillageCapacity, _CreateVillage, _DestroyVillage, _StepVillage, _VILLAGE_STEP_COUNT},
	{"washer", _GetWasherCapacity, _CreateWasher, nil, nil, 1000},
}

// BenchmarkUpstream runs the box3d benchmarks like benchmark/main.c: the first step is not
// timed, so ns/op is the time of the other steps and compares with the milliseconds that
// upstream prints, see box3d-c/benchmark/*/<name>.csv. Some take seconds per run, select one
// with -bench, e.g. -bench 'Upstream/large_pyramid/workers=4'.
func BenchmarkUpstream(b *testing.B) {
	for _, bench := range upstreamBenchmarks {
		for _, workerCount := range []uint32{1, 4} {
			b.Run(fmt.Sprintf("%s/workers=%d", bench.name, workerCount), func(b *testing.B) {
				for range b.N {
					b.StopTimer()

					// the capacity function writes through a uintptr
					def := new(WorldDef)
					escapes(def)
					*def = DefaultWorldDef()
					def.WorkerCount = workerCount
					if bench.capacity != nil {
						bench.capacity(theStack, uintptr(unsafe.Pointer(&def.Capacity)))
					}

					world := CreateWorld(*def)
					bench.create(theStack, world.Id)

					if bench.step != nil {
						bench.step(theStack, world.Id, 0)
					}

					world.Step(1.0/60.0, 4)

					b.StartTimer()
					for step := 1; step < bench.steps; step++ {
						if bench.step != nil {
							bench.step(theStack, world.Id, int32(step))
						}

						world.Step(1.0/60.0, 4)
					}
					b.StopTimer()

					if bench.destroy != nil {
						bench.destroy(theStack)
					}

					world.DestroyWorld()
				}
			})
		}
	}
}

// malloc is used by the translated benchmarks of box3d's shared/benchmarks.c
func malloc(tls *_Stack, size uint64) uintptr {
	return __builtin_aligned_alloc(tls, 16, size)
}
