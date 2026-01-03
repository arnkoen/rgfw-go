//go:build webgpu

package rgfw

// #cgo CFLAGS: -DRGFW_BUILD_WEBGPU
import "C"

// Build with: go build -tags webgpu
