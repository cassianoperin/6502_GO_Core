package VGS

import (
	// "fmt"
	// "os"
	// "time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	// "golang.org/x/image/colornames"
	// "github.com/faiface/pixel/imdraw"
)

func Run() {

	// Set up render system
	// Initial Pixel Size
	width = screenWidth / sizeX
	height = screenHeight / sizeY

	cfg := pixelgl.WindowConfig{
		Title:       "Pixel Rocks!",
		Bounds:      pixel.R(0, 0, 320, 240),
		VSync:       false,
		Resizable:   false,
		Undecorated: false,
		NoIconify:   false,
		AlwaysOnTop: false,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Disable Smooth
	win.SetSmooth(true)

	// renderGraphics(win)

	for !win.Closed() {

		// Esc to quit program
		if win.JustPressed(pixelgl.KeyEscape) {
			break
		}

		// Internal Loop to avoid slowness of !win.Closed() loop
		// for i := 0; i < 50000; i++ {

		// // Esc to quit program
		// if win.JustPressed(pixelgl.KeyEscape) {
		// 	os.Exit(0)
		// }

		select {
		case <-second_timer: // Second
			if !Pause {
				// Runs the interpreter
				if CPU_Enabled {
					CPU_Interpreter()
				}
			}
			// win.Update()

		default:
			// No timer to handle
		}

		Keyboard(win)

		// }
		win.Update()

	}

}
