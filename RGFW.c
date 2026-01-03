// RGFW implementation file
// Handles all graphics backends based on Go build tags

#define RGFW_IMPLEMENTATION

// Graphics API selection based on build tags
#if defined(RGFW_BUILD_OPENGL)
	#define RGFW_OPENGL
#elif defined(RGFW_BUILD_VULKAN)
	#define RGFW_VULKAN
#elif defined(RGFW_BUILD_WEBGPU)
	#define RGFW_WEBGPU
#elif defined(RGFW_BUILD_DIRECTX)
	#define RGFW_DIRECTX
#endif

#include "RGFW.h"
