# A debugger for Go

To understand how a debugger works we would ahve to completely understand how a linux process functions and the states it is found in.
An application has to be executed so that it turns into a process with states and can be loaded into memory for safe execution, with its own stack space to store its data.


![Fork Diagram](http://www.bogotobogo.com/Linux/images/process/fork_diagram.png)

###### Ptrace is a kernel hook into the task dispatch logic.

### wait (system call)

* Primary goal of this call is syncronization with its children.
* Whenever a child terminates or is stopped by a signal the parent needs to know.

### syscall.WALL (options)

In the man page for wait(2) we can see that __WALL basically is an option which we set for the wait() where we wait for all the children of the particular process.

### syscall.PTRACE_O_TRACECLONE (options)

* Stop the tracee at the next clone(2) and automatically start tracing the newly cloned process, which will start with a SIGSTOP, or PTRACE_EVENT_STOP if PTRACE_SEIZE was used. 
* A waitpid(2) by the tracer will return a status value such that status>>8 == (SIGTRAP | (PTRACE_EVENT_CLONE<<8))


### The "debug" package

THe debug pacakge at a glance.

```go
    func FreeOSMemory() // Collect garbage and return it to. Usually runs in the background
    func PrintStack() // runtime.Stack // Print stack to standard error 
    func ReadGCStats(stats *GCStats) // GC stats
    func SetGCPercent(percent int) int // ( ratio of freshly allocated data to live data remaining after the last collection ) returns last one
    func SetMaxStack(bytes int) int
    func SetMaxThreads(threads int) int // initial 10k
    func SetPanicOnFault(enabled bool) bool //panic on faults such as runtime memory corruption
    func SetTraceback(level string)
    func Stack() []byte // formatted stack trace of the goroutine that calls it.
    func WriteHeapDump(fd uintptr) // writes a description of the heap and the objects in it to the given file descriptor.  https://golang.org/s/go15heapdump. 
    type GCStats 
```

### The debug/gosym

The gosym package allows us to access the symbol table and line number table embeded in go binaries.


## Progress

* Get the Process to stop executing at the Wait State at every step.

## Future Work

* Set breakpoint at the end of an instruction by putting the 0xCC byte in the place of the first byte of the next instruction.
* Be able to read the number of bytes in the instruction after which we would like to place a breakpoint so that we can take that into consideration and would not have to specify the number of bytes to be able to poke data into the fist byte of the instruction.

## References
(1) [Linux Processes and Signals](http://www.bogotobogo.com/Linux/linux_process_and_signals.php)
