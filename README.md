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
Using the **"make 5_clock"** command we can run both: 
- 3 servers with 3 timezones
- a client consulting the 3 timezones within 3 goroutines during 15 secs (the clients are run after 5 sec sleep)

