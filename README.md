# GoUdoToStdoutServer

A very simple UDP Dumper written in golang

Very much based on https://github.com/hyrmn/GoTcpEchoServer

## Running

By default, the server will listen on port 2007. You can override this via either an environment variable or command-line switch.

### Command-Line

From a command prompt, run the main.go file

```
$> go run main.go
```

This will start the TCP server listening on port 7. 

You can specify another port using the `-p` switch

```
$> go run main.go -p 8080
```

Or, set the environment variable `PORT`. In Powershell, you would run

```
$> $env:PORT = 7000
$> go run main.go
```

This will start the UDP Listener on port 7000

Note, if you are on Windows, you will be prompted to allow network access. Select yes.

## Connecting

To connect to a UDP Port I used 
```

```

## License

This software is release free and unencumbered under the unlicense. Details in [LICENSE](https://github.com/christianwoehrle/udllistener/blob/master/LICENSE)
