### A repository to practice goroutins and channels in GO
:rocket: :rocket: :rocket:


### 1. Spinner 
It shows just a goroutin functioning

### 2. Clock1 
It shows an inefficient net.listener that serves one connection at a time. The second client must wait until the first is finished.
The server is *SEQUENTIAL*.
(use nc localhost 8080 twice in a separate terminals to see that the second connection is server only when the first one is killed)

### 3. Netcat1
It can be used instead of nc to test the previous example from section num 2

### 4. Clock2
It shows how to create a concurrent net.Listener simply putting the **handleConn** function into goroutine:

:arrow_lower_right:
``` go
    go handleConn(cn)
```

### 5. Enhanced netcat
Using the ***"make 5_clock"*** command we can run both: 
- 3 servers with 3 timezones
- a client consulting the 3 timezones within 3 goroutines during 15 secs (the clients are run after 5 sec sleep)


### 6. Interactive client
Using ***"make 6_server"*** and ***"make 6_interactive_client"*** in a separate terminal
we can play in the client terminal by writing a phrase and receiving 3 responses (uppercase, normal and lowercase) every second for each phrase we write


### 7. Interactive client - 2
Using ***"make 7_server"*** and ***"make 7_interactive_client"*** 
the same result as in pt. 6but the order of responses is not guaranteed 
due to gorouting for the echo command

### 8. Interactive client - 3
The same as previous examples, but the client closes only write operations, 
so it can keep on reading from the server.

### 9. Infinite pipeline
An example of an infinite pipeline where the first goroutine generates natural numbers and sends them to "naturals" channel (blocking it).
The second goroutine (squarer) reads from channel "naturals" unblocking it, squares the natural number and sends the squared number to the "squares" channel (blocking it).
The main goroutine then reads the "squares" channel (once there is something in it and unblocking it). All the goroutines (including the main one) work in infinite loops.
Ommitting the for loop (i.e. in the Squarer goroutine) results in fatal error: "all goroutines are asleep - deadlock".