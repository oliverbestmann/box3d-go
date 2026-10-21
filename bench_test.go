package b3

import "testing"

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
