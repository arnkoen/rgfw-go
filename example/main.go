package main

import (
	"fmt"
	rgfw "rgfw-go"
)

func main() {
	win := rgfw.CreateWindow(
		"RGFW Go Example",
		0, 0,
		800, 600,
		rgfw.WindowCenter,
	)

	if win == nil {
		fmt.Println("Failed to create window")
		return
	}
	defer win.Close()

	win.SetExitKey(rgfw.Escape)

	fmt.Println("Window created successfully!")
	fmt.Println("Press ESC to exit")

	for !win.ShouldClose() {
		for {
			event, ok := win.CheckEvent()
			if !ok {
				break
			}

			switch event.Type() {
			case rgfw.Quit:
				fmt.Println("Quit event received")
				win.SetShouldClose(true)

			case rgfw.KeyPressed:
				key := event.KeyCode()
				fmt.Printf("Key pressed: %d\n", key)

				if key == rgfw.KeyF {
					isFullscreen := win.IsFullscreen()
					win.SetFullscreen(!isFullscreen)
					if !isFullscreen {
						fmt.Println("Entering fullscreen")
					} else {
						fmt.Println("Exiting fullscreen")
					}
				}

			case rgfw.KeyReleased:
				key := event.KeyCode()
				fmt.Printf("Key released: %d\n", key)

			case rgfw.MouseButtonPressed:
				button := event.MouseButton()
				x, y := event.MouseX(), event.MouseY()
				fmt.Printf("Mouse button %d pressed at (%d, %d)\n", button, x, y)

			case rgfw.MouseButtonReleased:
				button := event.MouseButton()
				fmt.Printf("Mouse button %d released\n", button)

			case rgfw.MousePosChanged:
				x, y := event.MouseX(), event.MouseY()
				fmt.Printf("Mouse moved to (%d, %d)\n", x, y)

			case rgfw.MouseScroll:
				scrollX, scrollY := event.ScrollX(), event.ScrollY()
				fmt.Printf("Mouse scroll: (%f, %f)\n", scrollX, scrollY)

			case rgfw.WindowResized:
				w, h := win.GetSize()
				fmt.Printf("Window resized to %dx%d\n", w, h)

			case rgfw.WindowMoved:
				x, y := win.GetPosition()
				fmt.Printf("Window moved to (%d, %d)\n", x, y)

			case rgfw.FocusIn:
				fmt.Println("Window gained focus")

			case rgfw.FocusOut:
				fmt.Println("Window lost focus")

			case rgfw.WindowMaximized:
				fmt.Println("Window maximized")

			case rgfw.WindowMinimized:
				fmt.Println("Window minimized")

			case rgfw.WindowRestored:
				fmt.Println("Window restored")

			case rgfw.DataDrop:
				files := event.DropFiles()
				fmt.Printf("Files dropped: %v\n", files)
			}
		}
	}

	fmt.Println("Window closed")
}
