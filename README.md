# ardielle-hello
A simple hello world client/server example

Server:
```
$ make regenerate-go && make go-server
```

Client:
```
$ make go-client
Hello, <username>

$ USER="" make go-client
$ Hello, human

```
