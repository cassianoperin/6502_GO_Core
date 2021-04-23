package VGS

import (
	"Atari2600/VGS"
	"fmt"
	"time"

	"github.com/faiface/pixel/pixelgl"
	// "github.com/faiface/pixel/imdraw"
)

// ----------------------------------- Keyboard ----------------------------------- //

func Keyboard(target *pixelgl.Window) {

	// Debug
	if target.JustPressed(pixelgl.Key9) {

		if Debug {
			Debug = false
		} else {
			Debug = true

		}
	}

	// Reset
	if target.JustPressed(pixelgl.Key0) {

		VGS.Reset()

	}

	// Pause Key
	if target.Pressed(pixelgl.KeyP) {
		if Pause {
			Pause = false
			fmt.Printf("\t\tPAUSE mode Disabled\n")
			// Control repetition
			target.UpdateInputWait(time.Second)
		} else {
			Pause = true
			fmt.Printf("\t\tPAUSE mode Enabled\n")
			target.UpdateInputWait(time.Second)
		}
	}

	// Step Forward
	if target.Pressed(pixelgl.KeyI) {
		if Pause {
			// if Debug {
			// for dbg_running_opc == true {
			fmt.Printf("\t\tStep Forward\n")

			target.UpdateInput()
			// Runs the interpreter
			if CPU_Enabled {
				CPU_Interpreter()
			}

		}

		// After being paused by the end of opcode, set again to start the new one
		// dbg_running_opc = true

		// Control repetition
		target.UpdateInputWait(time.Second)
		// }

		// }

	}

}
