package CORE

import (
	// "fmt"
	// "os"
	// "time"

	"fmt"
	"os"

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
		Title:       "6502 Emulator",
		Bounds:      pixel.R(0, 0, 640, 480),
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
		for i := 0; i < 50000; i++ {

			// Esc to quit program
			if win.JustPressed(pixelgl.KeyEscape) {
				os.Exit(0)
			}

			select {
			case <-Second_timer: // Second
				win.SetTitle(fmt.Sprintf("%s CPS: %d| IPS: %d | Cycles: %d", cfg.Title, CPS, IPS, Cycle))
				CPS = 0
				IPS = 0

			default:
				// No timer to handle
			}

			select {
			case <-clock_timer.C:
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

		}

		select {
		case <-screenRefresh_timer.C: // Second
			win.Update()

		default:
			// No timer to handle
		}

	}

}
