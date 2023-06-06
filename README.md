## What To Do

1. Start two terminals.
2. On one do:

`go run echo.go listen -v`

3. On the other do:

`go run echo.go network -v 123456`

4. For `go run echo.go listen -v`

```
2023/06/06 04:53:40 Running a networkSender listener on port '3031' via the 'tcp' protocol.
2023/06/06 04:53:40 Address: '127.0.0.1:3031'
2023/06/06 04:53:40 DEBUG: from main:
2023/06/06 04:53:40 Listening...
2023/06/06 04:53:44 Copied 14 bytes into the buffer.
2023/06/06 04:53:44 This is a test
2023/06/06 04:53:44 Wrote 14 bytes to the connection.
```

5. For `go run echo.go network -v "This is a test"`

```
2023/06/06 04:53:44 Running a networkSender echo to port '3031' via the 'tcp' protocol.
2023/06/06 04:53:44 Address: '127.0.0.1:3031'
2023/06/06 04:53:44 Got a successful connection:  &{{0xc00010ea00}}
2023/06/06 04:53:44 Wrote arg 0 (This is a test) with 14 bytes.
2023/06/06 04:53:44 Listening for response now...
2023/06/06 04:53:44 Got response 14 bytes: This is a test.
2023/06/06 04:53:44 Closing connection. &{{0xc00010ea00}}
```
6. You can run:
```
go run echo.go help
go run echo.go listen --help
go run echo.go network --help
```

^^ See them if you need to change port numbers for some reason.