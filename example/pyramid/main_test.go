package main

import (
	"testing"

	b3 "github.com/oliverbestmann/box3d-go"
)

func TestClipNear(t *testing.T) {
	cam := camera{fwd: b3.Vec3{Z: 1}}

	// a square from 1 behind to 1 in front of the camera, the half behind the near plane goes
	square := []b3.Vec3{{X: -1, Z: -1}, {X: 1, Z: -1}, {X: 1, Z: 1}, {X: -1, Z: 1}}
	got := clipNear(cam, square)

	want := []b3.Vec3{{X: 1, Z: near}, {X: 1, Z: 1}, {X: -1, Z: 1}, {X: -1, Z: near}}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	for i := range want {
		if b3.Length(b3.Sub(got[i], want[i])) > 1e-5 {
			t.Fatalf("got %v, want %v", got, want)
		}
	}

	if len(clipNear(cam, []b3.Vec3{{Z: -1}, {X: 1, Z: -1}, {Z: -2}})) != 0 {
		t.Fatal("a triangle behind the camera should be dropped")
	}
}
