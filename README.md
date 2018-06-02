# Systems Programming in Go

#### Breaking Down Basic Systems Programming and things to do

* Files and directories (5)(Mihalis Tsoukalos' Chapter no.)
> pwd(1), which(1), find(1)
* File I/O (6)
> wc(1), dd(1)
* System Files and Configuration (7)
> editing files, regex, password generation
* IPC, Process Control/ Management (8)
> cat(1), signal handling, Plotting Data, Creating Client for Unix Socket
* Threads/ Goroutines (9,10)
> goroutines, channels, syncronization, mutexes
* Network Programming (12)
> Unix Socket Server and Client, TCP Server and Client, RPC Server and Client
> Packet Capture and Analysis

### 1> Strace

(Introduction from manpage)
In the simplest case strace runs the specified command until it exits. It  intercepts and records the system calls which are called by a process and the signals which are received by a process. The name of each system call, its arguments and its return value are printed on standard error or to the file specified with the -o option.


### 2> Debugger
### 3> Raw Sockets [Using Raw Sockets in Go](https://css.bz/2016/12/08/go-raw-sockets.html)
### 4> Packet Capture Tool [GopherCon 2016: John Leon - Packet Capture, Analysis, and Injection with Go](https://www.youtube.com/watch?v=APDnbmTKjgM)
### 5> Client/Server : Unix Socket
### 6> Client/Server : TCP
### 7> Client/Server : RPC
### 8> Wall of Sheep [DEF CON 23 - Packet Hacking Village - Tools And Techniques Used At The Wall of Sheep](https://www.youtube.com/watch?v=o_OoUv_HPls)
### 9> Go on BareMetal [Bare Metal Gophers](https://github.com/achilleasa/bare-metal-gophers)
### 10> [CliveOS](http://lsub.org/export/clivesys.pdf)
