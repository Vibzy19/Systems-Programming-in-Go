package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Println("Running ", os.Args[1:])
	cmd := exec.Command(os.Args[1], os.Args[2:]...) // returns *Cmd
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//io.Reader, io.Writer, io.Writer
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true} // Enable Tracing
	//cmd.Process = &os.Process{}                          // takes in PID
	//cmd.ProcessState = &os.ProcessState{}                // state of the process

	//var info syscall.Sysinfo_t
	//syscall.Sysinfo(&info)
	//fmt.Printf("%+v \n", info)
	//debug.PrintStack()
	/*
		{
			Chroot:
			Credential:<nil>
			Ptrace:true
			Setsid:false
			Setpgid:false
			Setctty:false
			Noctty:false
			Ctty:0
			Foreground:false
			Pgid:0
			Pdeathsig:signal 0
			Cloneflags:0
			Unshareflags:0
			UidMappings:[]
			GidMappings:[]
			GidMappingsEnableSetgroups:false
			AmbientCaps:[]}
	*/

	//Run the command
	cmd.Start()
	trace(cmd)

	//fmt.Printf("exec.Cmd : %+v \nSysProcAttr:%+v \n",*cmd, *cmd.SysProcAttr)
	/*
			What happens when making a syscall?

		    Set registers up with syscall ID and parameters
		    Trap - transition to kernel - run syscall code
		    Result returned in %rax (x86)
	*/
	// Linux system calls are passed using registers
	// Orig_rax has the number of the syscall being used

}

func trace(cmd *exec.Cmd) {
	var exit bool = true // As process would be on wait
	// Breakpoint //
	err := cmd.Wait()
	fmt.Printf("\nWait Returned : %v\n", err)
	PID := cmd.Process.Pid
	var regs syscall.PtraceRegs
	for {
		//fmt.Printf("%v\n", printexit(exit))
		if exit { //when exit is false process would be in entry more and would wait for the next process state change
			err := syscall.PtraceGetRegs(PID, &regs)
			if err != nil {
				break
			}
			fmt.Printf("%+v\n", regs.Orig_rax)
		}

		syscall.PtraceSyscall(PID, 0)   // No signal sent
		syscall.Wait4(PID, nil, 0, nil) //wait for the process to change state
		exit = !exit
	}
}

func printexit(ex bool) string {
	if ex {
		return "Entry"
	} else {
		return "Exit"
	}
}
