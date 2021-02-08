# go-cel

Package **cel** provide tools for creating software that models systems based on **Commands**, **Events**, and **Logs**.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-cel

[![GoDoc](https://godoc.org/github.com/reiver/go-cel?status.svg)](https://godoc.org/github.com/reiver/go-cel)


## Messages

Collectively we call _commands_, _events_, and _logs_: **messages**.

Each message has:

* version
* kind
* name
* payload

In the code, (using the Go proscribed style of naming interfaces) a **message** is represented as a:
```
cel.Messenger
```

A `cel.Messenger`` has the following methods:

* cel.Messenger.Version()
* cel.Messenger.Kind()
* cel.Messenger.Name()
* cel.Messenger.Payload()

In addition to `cel.Messenger`, this package provides an implementation of `cel.Messenger` with the struct:
```
cel_message.Type
```
