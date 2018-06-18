/*

Process Control Manipulation :
	Stop process at the required
	wait state in the life of a process
	to gain information on the process
	intermediately.
Line Table := PCounter to LineNumber
*/

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	// necessary for information on elf binaries
)

func main() {
	fmt.Printf("Debugging %v \n", os.Args[1:])

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}
	var regs syscall.PtraceRegs
	var ws syscall.WaitStatus
	var wpid int

	// Start Process
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	//Stop Process
	err = cmd.Wait()
	log.Printf("State : %v \n", err)
	log.Println("Restarting ...")
	//Read Registers
	syscall.PtraceGetRegs(cmd.Process.Pid, &regs)

	steps := 0 // Total steps taken
	syscall.PtraceSetOptions(cmd.Process.Pid, syscall.PTRACE_O_TRACECLONE)

	// Now we are going to move through process PIDs and PGIDs
	pgid, err := syscall.Getpgid(cmd.Process.Pid) //do error check
	if err != nil {
		log.Fatal(err)
	}
	syscall.PtraceSingleStep(cmd.Process.Pid) // do error check
	steps++                                   // increment the no. of steps taken
	for {

		// Wait a Step of the Process Group (waiting on a child to return)
		// Read Registers
		// Check if the child process exited and if the pid has changed
		// Take a step if the process hasn't exited
		// we have to break if we are in the parent process

		//wpid refers to the pid of the process in waitstatus
		wpid, err = syscall.Wait4(-1*pgid, &ws, syscall.WALL, nil)
		if err != nil {
			log.Fatal(err)
		}
		if wpid == -1 {
			log.Fatal(err)
		}
		if wpid == cmd.Process.Pid && ws.Exited() {
			break
		}

		if !ws.Exited() {
			//syscall.PtraceGetRegs(wpid, &regs) does not work on child processes directly
			err = syscall.PtraceSingleStep(wpid)
			if err != nil {
				log.Fatal(err)
			}
			steps++
		}

	}
	log.Printf("Total steps taken : %v", steps)
	log.Println("Exitting debugger..")

}
