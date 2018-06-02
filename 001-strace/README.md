# strace(1)

### Introduction to Syscalls

Syscalls (system calls) allow the User Space Programs to talk to the Kernel Space and perform the necessary operations related to handling the resources with which we can get the right output back to the User Space. (basic def)


### ptrace (system call)
* The ptrace() system call provides a means by which one process (the "tracer") may observe and control the execution of another process (the "tracee").
* We can examine and change the tracee's memory and registers. 
* It is primarily used to implement breakpoint debugging and system call tracing.

[...] While being traced, the child will stop each time a signal is delivered, even if the signal is being ignored. The parent will be notified at its next wait(2) and may inspect and modify the child process while it is stopped. [...]

##### PROCESS STATE CODES
   - R  running or runnable (on run queue)
   - D  uninterruptible sleep (usually IO)
   - S  interruptible sleep (waiting for an event to complete)
   - Z  defunct/zombie, terminated but not reaped by its parent
   - _T  stopped, either by a job control signal or because it is being traced_
	Ctrl+Z SIGSTOP Signal

##### LIST OF SIGNALS
```sh
$ kill -l
 1) SIGHUP	 2) SIGINT	 3) SIGQUIT	 4) SIGILL	 5) SIGTRAP
 6) SIGABRT	 7) SIGBUS	 8) SIGFPE	 9) SIGKILL	10) SIGUSR1
11) SIGSEGV	12) SIGUSR2	13) SIGPIPE	14) SIGALRM	15) SIGTERM
16) SIGSTKFLT	17) SIGCHLD	18) SIGCONT	19) SIGSTOP	20) SIGTSTP
21) SIGTTIN	22) SIGTTOU	23) SIGURG	24) SIGXCPU	25) SIGXFSZ
26) SIGVTALRM	27) SIGPROF	28) SIGWINCH	29) SIGIO	30) SIGPWR
31) SIGSYS	34) SIGRTMIN	35) SIGRTMIN+1	36) SIGRTMIN+2	37) SIGRTMIN+3
38) SIGRTMIN+4	39) SIGRTMIN+5	40) SIGRTMIN+6	41) SIGRTMIN+7	42) SIGRTMIN+8
43) SIGRTMIN+9	44) SIGRTMIN+10	45) SIGRTMIN+11	46) SIGRTMIN+12	47) SIGRTMIN+13
48) SIGRTMIN+14	49) SIGRTMIN+15	50) SIGRTMAX-14	51) SIGRTMAX-13	52) SIGRTMAX-12
53) SIGRTMAX-11	54) SIGRTMAX-10	55) SIGRTMAX-9	56) SIGRTMAX-8	57) SIGRTMAX-7
58) SIGRTMAX-6	59) SIGRTMAX-5	60) SIGRTMAX-4	61) SIGRTMAX-3	62) SIGRTMAX-2
63) SIGRTMAX-1	64) SIGRTMAX	
```

### SysProcAttr

The attributes of a process. We enable tracing of child processes by setting the Ptrace value in this to true and pass it on to the SysProcAttr of the command (exec.Command) which we are executing (exec.Cmd.SysProcAttr).

```go
	/*
		{
			Chroot:
			Credential:<nil>
			Ptrace:true // set tracing
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
            AmbientCaps:[]
            }
	*/
```

## Notes

We first execute the command and set the stdin, stdout, stderr on it. After that we allow Ptrace to work on the command by changing the ProcessAttributes. After that we start the program and put a breakpoint by making it wat or by giving a SIGSTOP signal. At this moment we can read their process registers. The registers contain information on the syscall executing at that particular moment. the eax/rax register contains the number of the syscall which is running.

## Progress

* Returns syscall numbers in order of execution.

## Future Work

* Print syscall names instead of a number
* Return Arguments being passed to the syscall

## References
(1)[How does strace work ?](https://blog.packagecloud.io/eng/2016/02/29/how-does-strace-work/)
(2)[Linux Process States](https://idea.popcount.org/2012-12-11-linux-process-states/)
(3)[Modifying System Call Arguments With ptrace](https://www.alfonsobeato.net/c/modifying-system-call-arguments-with-ptrace/) 
(4)[Syscalls and Registers](https://stackoverflow.com/questions/2535989/what-are-the-calling-conventions-for-unix-linux-system-calls-on-i386-and-x86-6)
