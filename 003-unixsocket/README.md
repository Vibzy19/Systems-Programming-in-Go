# Unix Sockets

Unix domain sockets or IPC sockets are used for interprocess communication between the processes running on the same host.

Unix Socket support and other Network Programming including low level functions are present in golang. The "net" package helps us out with this. 

### net.Conn

Generic Stream oriented network connection.
This connection is setup with a buffer for the stream. 
The creation of this connection takes place when we have a socketfile and we dial a unix domain socket conneciton onto this.

#### net.Conn.Read()
Read to the Conn;

#### net.Conn.Write()
Write to the Conn;

#### net.Listen()
returns Listener which upon accepting any new request returns a Conn struct.

#### net.Dial()
Using this we can connect to already eshtablished connections.



