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

#### Introduction to Syscalls

Syscalls allow the User Space Programs to talk to the Kernel Space and perform the necessary operations related to handling the resources with which we can get the right output back to the User Space. (basic def)
[Golang UK Conference 2017 | Liz Rice - The Beginner's Guide to Linux Syscalls](https://www.youtube.com/watch?v=BdfNrs_oeko)

### 2> Debugger
### 3> Packet Capture Tool [GopherCon 2016: John Leon - Packet Capture, Analysis, and Injection with Go](https://www.youtube.com/watch?v=APDnbmTKjgM)
### 4> Client/Server : Unix Socket
### 5> Client/Server : TCP
### 6> Client/Server : RPC
### 7> Wall of Sheep [DEF CON 23 - Packet Hacking Village - Tools And Techniques Used At The Wall of Sheep](https://www.youtube.com/watch?v=o_OoUv_HPls)
### 8> Go on BareMetal [Bare Metal Gophers](https://github.com/achilleasa/bare-metal-gophers)
### 9> [CliveOS](http://lsub.org/export/clivesys.pdf)
