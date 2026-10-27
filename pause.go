//go:build (!amd64 && !arm64) || tinygo

package b3

func pause() {
	// noop
}
