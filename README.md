# go-cel

Package **cel** provide tools for creating software that models systems based on **Commands**, **Events**, and **Logs**.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-cel

[![GoDoc](https://godoc.org/github.com/reiver/go-cel?status.svg)](https://godoc.org/github.com/reiver/go-cel)


## Messages

Collectively we call _commands_, _events_, and _logs_: **messages**.

Each message has:

* magic
* version
* kind
* name
* payload

(Note that ‘magic’ is hidden.)

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

## JSON

`cel_message.Type` supports marshaling and unmarshaling to and from JSON.

An example **command**:
```json
{
	"magic"   : "CEL/1",
	"version" : "1",
	"kind"    : "EVENT",
	"name"    : "LOOK_DOOR",
	"payload" : {
		"from":"john",
		"doorid":"abc123",
	},
}
```

An example **event**:
```json
{
	"magic"   : "CEL/1",
	"version" : "1",
	"kind"    : "EVENT",
	"name"    : "DOOR_LOCKED",
	"payload" : {
		"from":"jane",
		"doorid":"abc123",
	},
}
```

An example **log**:
```json
{
	"magic"   : "CEL/1",
	"version" : "1",
	"kind"    : "LOG",
	"name"    : "DOORS",
	"payload" : {
		"doors": [
			{
				"doorid":"abc123",
			},
			{
				"doorid":"def456",
			},
			{
				"doorid":"ghi789",
			},
		],
	},
}
```
