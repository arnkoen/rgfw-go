//go:build opengl

package rgfw

/*
#include <stdint.h>
#define RGFW_OPENGL
#include "RGFW.h"
*/
import "C"

func (w *Window) SwapBuffers() {
	C.RGFW_window_swapBuffers_OpenGL(w.cWindow)
}

func (w *Window) MakeCurrent() {
	C.RGFW_window_makeCurrentWindow_OpenGL(w.cWindow)
}

func (w *Window) SwapInterval(interval int32) {
	C.RGFW_window_swapInterval_OpenGL(w.cWindow, C.int32_t(interval))
}
