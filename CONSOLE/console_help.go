package CONSOLE

import (
	"fmt"
)

// Print Help Menu
func Console_Command_Help() {
	fmt.Printf("\n\tquit\t\t\t\t\t\tQuit console\n\thelp\t\t\t\t\t\tPrint help menu\n\treset\t\t\t\t\t\tReinitialize CPU and reload program to memory\n\t-")
	fmt.Printf("\n\tstep\t\t\t\t\t\tExecute current opcode\n\tstep <value>\t\t\t\t\tExecute <value> opcodes\n\tstep_limit <value>\t\t\t\tDefine the maximum steps allowed\n\tstep_debug_start <value>\t\t\tSet the cycle to step start showing the debug messages\n\t-")
	fmt.Printf("\n\tadd_breakpoint <PC|A|X|Y|CYCLE>=<Value>\t\tAdd a breakpoint\n\tdel_breakpoint <index>\t\t\t\tDelete a breakpoint\n\tshow_breakpoints\t\t\t\tShow breakpoints\n\t-")
	fmt.Printf("\n\tregisters  <PC|A|X|Y>=<Value>\t\t\tChange registers values")
	fmt.Printf("\n\tprocessor_status  <N|V|B|D|I|Z|C>=<Value>\tChange processor status registers values\n\t-")
	fmt.Printf("\n\tdebug  <on|off>\t\t\t\t\tEnable or Disable Debug mode\n\t-")
	fmt.Printf("\n\tmem\t\t\t\t\t\tDump full memory\n\tmem <address>\t\t\t\t\tDump memory address\n\tmem <start address> <end address>\t\tDump memory address range\n\t-")
	fmt.Printf("\n\tdisassemble <address>\t\t\t\tDisassemble memory address\n\tdisassemble <start address> <end address>\tDisassemble memory address range\n\t-")
	fmt.Printf("\n\tgoto <address>\t\t\t\t\tRun until PC=<address>\n\tgoto_limit <value>\t\t\t\tDefine the maximum steps allowed in GOTO\n\t-")
	fmt.Printf("\n\trun\t\t\t\t\t\tRun the emulator\n\trun_limit <value>\t\t\t\tDefine the maximum steps allowed in RUN\n\n")
}
