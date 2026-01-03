package rgfw

/*
#cgo darwin CFLAGS: -I.
#cgo darwin LDFLAGS: -framework Cocoa -framework CoreVideo -framework IOKit
#cgo linux CFLAGS: -I.
#cgo linux LDFLAGS: -lX11 -lXrandr
#cgo windows CFLAGS: -I.
#cgo windows LDFLAGS: -lgdi32

// Graphics API linking (controlled by build tags)
#cgo opengl,darwin LDFLAGS: -framework OpenGL
#cgo opengl,linux LDFLAGS: -lGL
#cgo opengl,windows LDFLAGS: -lopengl32
#cgo vulkan,darwin LDFLAGS: -lvulkan
#cgo vulkan,linux LDFLAGS: -lvulkan
#cgo vulkan,windows LDFLAGS: -lvulkan-1

#include <stdint.h>

#define RGFW_IMPORT

// Graphics API selection (controlled by build tags)
// Optional: Build with -tags opengl OR -tags vulkan OR -tags webgpu OR -tags directx
// You can also build without any graphics backend if you don't need rendering
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

// Helper functions to access union fields
uint8_t getEventType(RGFW_event* e) {
	return e->type;
}

uint8_t getKeyValue(RGFW_event* e) {
	return e->key.value;
}

uint8_t getKeySym(RGFW_event* e) {
	return e->key.sym;
}

uint8_t getKeyMod(RGFW_event* e) {
	return e->key.mod;
}

uint8_t getKeyRepeat(RGFW_event* e) {
	return e->key.repeat;
}

uint8_t getButtonValue(RGFW_event* e) {
	return e->button.value;
}

int32_t getMouseX(RGFW_event* e) {
	return e->mouse.x;
}

int32_t getMouseY(RGFW_event* e) {
	return e->mouse.y;
}

float getMouseVecX(RGFW_event* e) {
	return e->mouse.vecX;
}

float getMouseVecY(RGFW_event* e) {
	return e->mouse.vecY;
}

float getScrollX(RGFW_event* e) {
	return e->scroll.x;
}

float getScrollY(RGFW_event* e) {
	return e->scroll.y;
}

char** getDropFiles(RGFW_event* e) {
	return e->drop.files;
}

size_t getDropCount(RGFW_event* e) {
	return e->drop.count;
}

// Wrapper functions for native handle getters (window-specific)
// These are needed because the RGFW functions are defined as inline functions
// in the header file, and CGo cannot directly call inline functions from headers.
void* getHWND(RGFW_window* win) {
	return RGFW_window_getHWND(win);
}

void* getHDC(RGFW_window* win) {
	return RGFW_window_getHDC(win);
}

uint64_t getWindow_X11(RGFW_window* win) {
	return RGFW_window_getWindow_X11(win);
}

void* getWindow_Wayland(RGFW_window* win) {
	return RGFW_window_getWindow_Wayland(win);
}

void* getWindow_OSX(RGFW_window* win) {
	return RGFW_window_getWindow_OSX(win);
}

void* getView_OSX(RGFW_window* win) {
	return RGFW_window_getView_OSX(win);
}

void* getSrc(RGFW_window* win) {
	return RGFW_window_getSrc(win);
}

// Wrapper functions for global native handle getters
static inline void* getLayer_OSX_wrapper(void) {
	return RGFW_getLayer_OSX();
}

static inline void* getDisplay_X11_wrapper(void) {
	return RGFW_getDisplay_X11();
}

static inline void* getDisplay_Wayland_wrapper(void) {
	return RGFW_getDisplay_Wayland();
}
*/
import "C"
import (
	"runtime"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

// Type Aliases

type (
	Bool   C.u8
	Key    C.u8
	Keymod C.u8
)

const (
	FALSE Bool = 0
	TRUE  Bool = 1
)

// Enums: Format

type Format C.u8

const (
	FormatRGB8  Format = C.RGFW_formatRGB8
	FormatBGR8  Format = C.RGFW_formatBGR8
	FormatRGBA8 Format = C.RGFW_formatRGBA8
	FormatARGB8 Format = C.RGFW_formatARGB8
	FormatBGRA8 Format = C.RGFW_formatBGRA8
	FormatABGR8 Format = C.RGFW_formatABGR8
)

// Enums: Keys

const (
	KeyNULL      Key = C.RGFW_keyNULL
	Escape       Key = C.RGFW_escape
	Backtick     Key = C.RGFW_backtick
	Key0         Key = C.RGFW_0
	Key1         Key = C.RGFW_1
	Key2         Key = C.RGFW_2
	Key3         Key = C.RGFW_3
	Key4         Key = C.RGFW_4
	Key5         Key = C.RGFW_5
	Key6         Key = C.RGFW_6
	Key7         Key = C.RGFW_7
	Key8         Key = C.RGFW_8
	Key9         Key = C.RGFW_9
	Minus        Key = C.RGFW_minus
	Equal        Key = C.RGFW_equal
	BackSpace    Key = C.RGFW_backSpace
	Tab          Key = C.RGFW_tab
	Space        Key = C.RGFW_space
	KeyA         Key = C.RGFW_a
	KeyB         Key = C.RGFW_b
	KeyC         Key = C.RGFW_c
	KeyD         Key = C.RGFW_d
	KeyE         Key = C.RGFW_e
	KeyF         Key = C.RGFW_f
	KeyG         Key = C.RGFW_g
	KeyH         Key = C.RGFW_h
	KeyI         Key = C.RGFW_i
	KeyJ         Key = C.RGFW_j
	KeyK         Key = C.RGFW_k
	KeyL         Key = C.RGFW_l
	KeyM         Key = C.RGFW_m
	KeyN         Key = C.RGFW_n
	KeyO         Key = C.RGFW_o
	KeyP         Key = C.RGFW_p
	KeyQ         Key = C.RGFW_q
	KeyR         Key = C.RGFW_r
	KeyS         Key = C.RGFW_s
	KeyT         Key = C.RGFW_t
	KeyU         Key = C.RGFW_u
	KeyV         Key = C.RGFW_v
	KeyW         Key = C.RGFW_w
	KeyX         Key = C.RGFW_x
	KeyY         Key = C.RGFW_y
	KeyZ         Key = C.RGFW_z
	Period       Key = C.RGFW_period
	Comma        Key = C.RGFW_comma
	Slash        Key = C.RGFW_slash
	Bracket      Key = C.RGFW_bracket
	CloseBracket Key = C.RGFW_closeBracket
	Semicolon    Key = C.RGFW_semicolon
	Apostrophe   Key = C.RGFW_apostrophe
	BackSlash    Key = C.RGFW_backSlash
	Return       Key = C.RGFW_return
	Delete       Key = C.RGFW_delete
	F1           Key = C.RGFW_F1
	F2           Key = C.RGFW_F2
	F3           Key = C.RGFW_F3
	F4           Key = C.RGFW_F4
	F5           Key = C.RGFW_F5
	F6           Key = C.RGFW_F6
	F7           Key = C.RGFW_F7
	F8           Key = C.RGFW_F8
	F9           Key = C.RGFW_F9
	F10          Key = C.RGFW_F10
	F11          Key = C.RGFW_F11
	F12          Key = C.RGFW_F12
	F13          Key = C.RGFW_F13
	F14          Key = C.RGFW_F14
	F15          Key = C.RGFW_F15
	F16          Key = C.RGFW_F16
	F17          Key = C.RGFW_F17
	F18          Key = C.RGFW_F18
	F19          Key = C.RGFW_F19
	F20          Key = C.RGFW_F20
	F21          Key = C.RGFW_F21
	F22          Key = C.RGFW_F22
	F23          Key = C.RGFW_F23
	F24          Key = C.RGFW_F24
	F25          Key = C.RGFW_F25
	CapsLock     Key = C.RGFW_capsLock
	ShiftL       Key = C.RGFW_shiftL
	ControlL     Key = C.RGFW_controlL
	AltL         Key = C.RGFW_altL
	SuperL       Key = C.RGFW_superL
	ShiftR       Key = C.RGFW_shiftR
	ControlR     Key = C.RGFW_controlR
	AltR         Key = C.RGFW_altR
	SuperR       Key = C.RGFW_superR
	Up           Key = C.RGFW_up
	Down         Key = C.RGFW_down
	Left         Key = C.RGFW_left
	Right        Key = C.RGFW_right
	Insert       Key = C.RGFW_insert
	Menu         Key = C.RGFW_menu
	End          Key = C.RGFW_end
	Home         Key = C.RGFW_home
	PageUp       Key = C.RGFW_pageUp
	PageDown     Key = C.RGFW_pageDown
	NumLock      Key = C.RGFW_numLock
	KpSlash      Key = C.RGFW_kpSlash
	KpMultiply   Key = C.RGFW_kpMultiply
	KpPlus       Key = C.RGFW_kpPlus
	KpMinus      Key = C.RGFW_kpMinus
	KpEqual      Key = C.RGFW_kpEqual
	Kp1          Key = C.RGFW_kp1
	Kp2          Key = C.RGFW_kp2
	Kp3          Key = C.RGFW_kp3
	Kp4          Key = C.RGFW_kp4
	Kp5          Key = C.RGFW_kp5
	Kp6          Key = C.RGFW_kp6
	Kp7          Key = C.RGFW_kp7
	Kp8          Key = C.RGFW_kp8
	Kp9          Key = C.RGFW_kp9
	Kp0          Key = C.RGFW_kp0
	KpPeriod     Key = C.RGFW_kpPeriod
	KpReturn     Key = C.RGFW_kpReturn
	ScrollLock   Key = C.RGFW_scrollLock
	PrintScreen  Key = C.RGFW_printScreen
	Pause        Key = C.RGFW_pause
)

// Enums: Mouse Buttons

type MouseButton C.u8

const (
	MouseLeft   MouseButton = C.RGFW_mouseLeft
	MouseMiddle MouseButton = C.RGFW_mouseMiddle
	MouseRight  MouseButton = C.RGFW_mouseRight
	MouseMisc1  MouseButton = C.RGFW_mouseMisc1
	MouseMisc2  MouseButton = C.RGFW_mouseMisc2
	MouseMisc3  MouseButton = C.RGFW_mouseMisc3
	MouseMisc4  MouseButton = C.RGFW_mouseMisc4
	MouseMisc5  MouseButton = C.RGFW_mouseMisc5
)

// Enums: Key Modifiers

const (
	ModCapsLock   Keymod = C.RGFW_modCapsLock
	ModNumLock    Keymod = C.RGFW_modNumLock
	ModControl    Keymod = C.RGFW_modControl
	ModAlt        Keymod = C.RGFW_modAlt
	ModShift      Keymod = C.RGFW_modShift
	ModSuper      Keymod = C.RGFW_modSuper
	ModScrollLock Keymod = C.RGFW_modScrollLock
)

// Enums: Event Types

type EventType C.u8

const (
	EventNone           EventType = C.RGFW_eventNone
	KeyPressed          EventType = C.RGFW_keyPressed
	KeyReleased         EventType = C.RGFW_keyReleased
	MouseButtonPressed  EventType = C.RGFW_mouseButtonPressed
	MouseButtonReleased EventType = C.RGFW_mouseButtonReleased
	MouseScroll         EventType = C.RGFW_mouseScroll
	MousePosChanged     EventType = C.RGFW_mousePosChanged
	WindowMoved         EventType = C.RGFW_windowMoved
	WindowResized       EventType = C.RGFW_windowResized
	FocusIn             EventType = C.RGFW_focusIn
	FocusOut            EventType = C.RGFW_focusOut
	MouseEnter          EventType = C.RGFW_mouseEnter
	MouseLeave          EventType = C.RGFW_mouseLeave
	WindowRefresh       EventType = C.RGFW_windowRefresh
	Quit                EventType = C.RGFW_quit
	DataDrop            EventType = C.RGFW_dataDrop
	DataDrag            EventType = C.RGFW_dataDrag
	WindowMaximized     EventType = C.RGFW_windowMaximized
	WindowMinimized     EventType = C.RGFW_windowMinimized
	WindowRestored      EventType = C.RGFW_windowRestored
	ScaleUpdated        EventType = C.RGFW_scaleUpdated
)

// Enums: Window Flags

type WindowFlags C.u32

const (
	WindowNoBorder        WindowFlags = C.RGFW_windowNoBorder
	WindowNoResize        WindowFlags = C.RGFW_windowNoResize
	WindowAllowDND        WindowFlags = C.RGFW_windowAllowDND
	WindowHideMouse       WindowFlags = C.RGFW_windowHideMouse
	WindowFullscreen      WindowFlags = C.RGFW_windowFullscreen
	WindowTransparent     WindowFlags = C.RGFW_windowTransparent
	WindowCenter          WindowFlags = C.RGFW_windowCenter
	WindowRawMouse        WindowFlags = C.RGFW_windowRawMouse
	WindowScaleToMonitor  WindowFlags = C.RGFW_windowScaleToMonitor
	WindowHide            WindowFlags = C.RGFW_windowHide
	WindowMaximize        WindowFlags = C.RGFW_windowMaximize
	WindowCenterCursor    WindowFlags = C.RGFW_windowCenterCursor
	WindowFloating        WindowFlags = C.RGFW_windowFloating
	WindowFocusOnShow     WindowFlags = C.RGFW_windowFocusOnShow
	WindowMinimize        WindowFlags = C.RGFW_windowMinimize
	WindowFocus           WindowFlags = C.RGFW_windowFocus
	WindowCaptureMouse    WindowFlags = C.RGFW_windowCaptureMouse
	WindowOpenGL          WindowFlags = C.RGFW_windowOpenGL
	WindowEGL             WindowFlags = C.RGFW_windowEGL
	WindowedFullscreen    WindowFlags = C.RGFW_windowedFullscreen
	WindowCaptureRawMouse WindowFlags = C.RGFW_windowCaptureRawMouse
)

// Enums: Mouse Icons

type MouseIcon C.u8

const (
	MouseNormal       MouseIcon = C.RGFW_mouseNormal
	MouseArrow        MouseIcon = C.RGFW_mouseArrow
	MouseIbeam        MouseIcon = C.RGFW_mouseIbeam
	MouseCrosshair    MouseIcon = C.RGFW_mouseCrosshair
	MousePointingHand MouseIcon = C.RGFW_mousePointingHand
	MouseResizeEW     MouseIcon = C.RGFW_mouseResizeEW
	MouseResizeNS     MouseIcon = C.RGFW_mouseResizeNS
	MouseResizeNWSE   MouseIcon = C.RGFW_mouseResizeNWSE
	MouseResizeNESW   MouseIcon = C.RGFW_mouseResizeNESW
	MouseResizeAll    MouseIcon = C.RGFW_mouseResizeAll
	MouseNotAllowed   MouseIcon = C.RGFW_mouseNotAllowed
)

// Enums: Icon Type

type Icon C.u8

const (
	IconTaskbar Icon = C.RGFW_iconTaskbar
	IconWindow  Icon = C.RGFW_iconWindow
	IconBoth    Icon = C.RGFW_iconBoth
)

// Structs: Monitor

type Monitor struct {
	cMonitor C.RGFW_monitor
}

func (m *Monitor) X() int32 {
	return int32(m.cMonitor.x)
}

func (m *Monitor) Y() int32 {
	return int32(m.cMonitor.y)
}

func (m *Monitor) Name() string {
	return C.GoString(&m.cMonitor.name[0])
}

func (m *Monitor) ScaleX() float32 {
	return float32(m.cMonitor.scaleX)
}

func (m *Monitor) ScaleY() float32 {
	return float32(m.cMonitor.scaleY)
}

func (m *Monitor) Width() int32 {
	return int32(m.cMonitor.mode.w)
}

func (m *Monitor) Height() int32 {
	return int32(m.cMonitor.mode.h)
}

func (m *Monitor) RefreshRate() uint32 {
	return uint32(m.cMonitor.mode.refreshRate)
}

// Structs: Window

type Window struct {
	cWindow *C.RGFW_window
}

// Structs: Event

type Event struct {
	cEvent C.RGFW_event
}

func (e *Event) Type() EventType {
	return EventType(C.getEventType(&e.cEvent))
}

func (e *Event) KeyCode() Key {
	return Key(C.getKeyValue(&e.cEvent))
}

func (e *Event) KeySym() uint8 {
	return uint8(C.getKeySym(&e.cEvent))
}

func (e *Event) KeyMod() Keymod {
	return Keymod(C.getKeyMod(&e.cEvent))
}

func (e *Event) KeyRepeat() bool {
	return C.getKeyRepeat(&e.cEvent) != 0
}

func (e *Event) MouseButton() MouseButton {
	return MouseButton(C.getButtonValue(&e.cEvent))
}

func (e *Event) MouseX() int32 {
	return int32(C.getMouseX(&e.cEvent))
}

func (e *Event) MouseY() int32 {
	return int32(C.getMouseY(&e.cEvent))
}

func (e *Event) MouseVecX() float32 {
	return float32(C.getMouseVecX(&e.cEvent))
}

func (e *Event) MouseVecY() float32 {
	return float32(C.getMouseVecY(&e.cEvent))
}

func (e *Event) ScrollX() float32 {
	return float32(C.getScrollX(&e.cEvent))
}

func (e *Event) ScrollY() float32 {
	return float32(C.getScrollY(&e.cEvent))
}

func (e *Event) DropFiles() []string {
	count := int(C.getDropCount(&e.cEvent))
	if count == 0 {
		return nil
	}

	files := make([]string, count)
	cFiles := (*[1 << 30]*C.char)(unsafe.Pointer(C.getDropFiles(&e.cEvent)))[:count:count]
	for i := 0; i < count; i++ {
		files[i] = C.GoString(cFiles[i])
	}
	return files
}

// Window Functions: Creation & Destruction

func CreateWindow(name string, x, y, w, h int32, flags WindowFlags) *Window {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cWin := C.RGFW_createWindow(
		cName,
		C.int32_t(x),
		C.int32_t(y),
		C.int32_t(w),
		C.int32_t(h),
		C.uint32_t(flags),
	)

	if cWin == nil {
		return nil
	}

	return &Window{cWindow: cWin}
}

func (w *Window) Close() {
	if w.cWindow != nil {
		C.RGFW_window_close(w.cWindow)
		w.cWindow = nil
	}
}

// Window Functions: Event Handling

func (w *Window) CheckEvent() (*Event, bool) {
	var cEvent C.RGFW_event
	hasEvent := C.RGFW_window_checkEvent(w.cWindow, &cEvent)
	if hasEvent == 0 {
		return nil, false
	}
	return &Event{cEvent: cEvent}, true
}

func (w *Window) ShouldClose() bool {
	return C.RGFW_window_shouldClose(w.cWindow) != 0
}

func (w *Window) SetShouldClose(shouldClose bool) {
	val := C.u8(0)
	if shouldClose {
		val = 1
	}
	C.RGFW_window_setShouldClose(w.cWindow, val)
}

func (w *Window) SetExitKey(key Key) {
	C.RGFW_window_setExitKey(w.cWindow, C.u8(key))
}

func (w *Window) GetExitKey() Key {
	return Key(C.RGFW_window_getExitKey(w.cWindow))
}

// Window Functions: Properties

func (w *Window) GetPosition() (x, y int32) {
	var cx, cy C.int32_t
	C.RGFW_window_getPosition(w.cWindow, &cx, &cy)
	return int32(cx), int32(cy)
}

func (w *Window) Move(x, y int32) {
	C.RGFW_window_move(w.cWindow, C.int32_t(x), C.int32_t(y))
}

func (w *Window) GetSize() (width, height int32) {
	var cw, ch C.int32_t
	C.RGFW_window_getSize(w.cWindow, &cw, &ch)
	return int32(cw), int32(ch)
}

func (w *Window) Resize(width, height int32) {
	C.RGFW_window_resize(w.cWindow, C.int32_t(width), C.int32_t(height))
}

func (w *Window) SetName(name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.RGFW_window_setName(w.cWindow, cName)
}

// Window Functions: State

func (w *Window) Maximize() {
	C.RGFW_window_maximize(w.cWindow)
}

func (w *Window) Minimize() {
	C.RGFW_window_minimize(w.cWindow)
}

func (w *Window) Restore() {
	C.RGFW_window_restore(w.cWindow)
}

func (w *Window) SetFullscreen(fullscreen bool) {
	val := C.u8(0)
	if fullscreen {
		val = 1
	}
	C.RGFW_window_setFullscreen(w.cWindow, val)
}

func (w *Window) IsFullscreen() bool {
	return C.RGFW_window_isFullscreen(w.cWindow) != 0
}

func (w *Window) Hide() {
	C.RGFW_window_hide(w.cWindow)
}

func (w *Window) Show() {
	C.RGFW_window_show(w.cWindow)
}

func (w *Window) IsHidden() bool {
	return C.RGFW_window_isHidden(w.cWindow) != 0
}

func (w *Window) Focus() {
	C.RGFW_window_focus(w.cWindow)
}

func (w *Window) IsInFocus() bool {
	return C.RGFW_window_isInFocus(w.cWindow) != 0
}

func (w *Window) ShowMouse(show bool) {
	val := C.u8(0)
	if show {
		val = 1
	}
	C.RGFW_window_showMouse(w.cWindow, val)
}

func (w *Window) IsMouseHidden() bool {
	return C.RGFW_window_isMouseHidden(w.cWindow) != 0
}

func (w *Window) MoveMouse(x, y int32) {
	C.RGFW_window_moveMouse(w.cWindow, C.int32_t(x), C.int32_t(y))
}

func (w *Window) GetMouse() (x, y int32) {
	var cx, cy C.int32_t
	C.RGFW_window_getMouse(w.cWindow, &cx, &cy)
	return int32(cx), int32(cy)
}

func (w *Window) IsMinimized() bool {
	return C.RGFW_window_isMinimized(w.cWindow) != 0
}

func (w *Window) IsMaximized() bool {
	return C.RGFW_window_isMaximized(w.cWindow) != 0
}

// Window Functions: Native Handles

// GetHWND returns the Windows window handle (HWND).
// Only available on Windows platform.
func (w *Window) GetHWND() unsafe.Pointer {
	return unsafe.Pointer(C.getHWND(w.cWindow))
}

// GetHDC returns the Windows device context (HDC).
// Only available on Windows platform.
func (w *Window) GetHDC() unsafe.Pointer {
	return unsafe.Pointer(C.getHDC(w.cWindow))
}

// GetWindow_X11 returns the X11 window ID.
// Only available on X11-based Linux systems.
func (w *Window) GetWindow_X11() uint64 {
	return uint64(C.getWindow_X11(w.cWindow))
}

// GetWindow_Wayland returns the Wayland surface handle.
// Only available on Wayland-based Linux systems.
func (w *Window) GetWindow_Wayland() unsafe.Pointer {
	return unsafe.Pointer(C.getWindow_Wayland(w.cWindow))
}

// GetWindow_OSX returns the macOS NSWindow handle.
// Only available on macOS platform.
func (w *Window) GetWindow_OSX() unsafe.Pointer {
	return unsafe.Pointer(C.getWindow_OSX(w.cWindow))
}

// GetView_OSX returns the macOS NSView handle.
// Only available on macOS platform.
func (w *Window) GetView_OSX() unsafe.Pointer {
	return unsafe.Pointer(C.getView_OSX(w.cWindow))
}

// GetSrc returns a pointer to the underlying RGFW_window C structure.
// This provides direct access to the native window structure for advanced use cases.
func (w *Window) GetSrc() unsafe.Pointer {
	return unsafe.Pointer(C.getSrc(w.cWindow))
}

// Global Native Handle Functions

// GetLayer_OSX returns the global CALayer used by RGFW on macOS.
// Only available on macOS platform. Returns NULL on other platforms.
func GetLayer_OSX() unsafe.Pointer {
	return unsafe.Pointer(C.getLayer_OSX_wrapper())
}

// GetDisplay_X11 returns the global X11 Display pointer.
// Only available on X11-based Linux systems. Returns NULL on other platforms.
func GetDisplay_X11() unsafe.Pointer {
	return unsafe.Pointer(C.getDisplay_X11_wrapper())
}

// GetDisplay_Wayland returns the global Wayland display handle.
// Only available on Wayland-based Linux systems. Returns NULL on other platforms.
func GetDisplay_Wayland() unsafe.Pointer {
	return unsafe.Pointer(C.getDisplay_Wayland_wrapper())
}

// Monitor Functions

func GetMonitors() []Monitor {
	var count C.size_t
	cMons := C.RGFW_getMonitors(&count)
	if cMons == nil || count == 0 {
		return nil
	}

	monitors := make([]Monitor, int(count))
	cMonSlice := (*[1 << 30]C.RGFW_monitor)(unsafe.Pointer(cMons))[:count:count]
	for i := 0; i < int(count); i++ {
		monitors[i] = Monitor{cMonitor: cMonSlice[i]}
	}

	return monitors
}

func GetPrimaryMonitor() Monitor {
	cMon := C.RGFW_getPrimaryMonitor()
	return Monitor{cMonitor: cMon}
}

func (w *Window) GetMonitor() Monitor {
	cMon := C.RGFW_window_getMonitor(w.cWindow)
	return Monitor{cMonitor: cMon}
}
