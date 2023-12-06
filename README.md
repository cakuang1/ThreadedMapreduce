# gochat
Wanted to test out goroutines so I made a simple multithreaded chat application. No special UI, Terminal based. 

## How this works
Server has a main thread which handles listening to connections.When a TCP connection is made, a connection thread is made, which listens to the clients and broadcasts these messages to all clients.

Client also has a main thread that handles user input and message sending, while also spawns another thread to receive new chats from the servers broadcasts.
 
## How to run 
Run the server

```
go run server.go

```

Register client(s)

```
go run client.go
```








