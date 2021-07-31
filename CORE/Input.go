package CORE

import (
	"fmt"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

// ----------------------------------- Keyboard ----------------------------------- //

func Keyboard(target *pixelgl.Window) {

	// Debug
	if target.JustPressed(pixelgl.Key9) {

		Debug = !Debug
		target.UpdateInputWait(time.Second)

	}

	// Reset
	if target.JustPressed(pixelgl.Key0) {

		Initialize()
		Reset()

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

			fmt.Printf("\t\tStep Forward\n")

			target.UpdateInput()
			// Runs the interpreter
			if CPU_Enabled {
				CPU_Interpreter()
			}
		}

		// Control repetition
		target.UpdateInputWait(time.Second)
	}

}
