//go:build vulkan

package rgfw

// #cgo CFLAGS: -DRGFW_BUILD_VULKAN
import "C"

// Build with: go build -tags vulkan
