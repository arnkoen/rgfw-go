//go:build directx && windows

package rgfw

// #cgo CFLAGS: -DRGFW_BUILD_DIRECTX
import "C"

// Build with: go build -tags directx
// Note: DirectX is only available on Windows
