# go-cel

Package **cel** provide tools for creating software that models systems based on **Commands**, **Events**, and **Logs**.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-cel

[![GoDoc](https://godoc.org/github.com/reiver/go-cel?status.svg)](https://godoc.org/github.com/reiver/go-cel)


## Quick Example

Here is a quick example:

```go
const apiVersion string = "1"

c := cel_message.Command(
		apiVersion,
		"LOCK_DOOR",
		map[string]interface{}{
			"from":"john",
			"doorid":"abc123",
		},
)

e := cel_message.Event(
		apiVersion,
		"DOOR_LOCKED",
		map[string]interface{}{
			"from":"jane",
			"doorid":"abc123",
		},
)

l := cel_message.Log(
		apiVersion,
		"DOORS",
		map[string]interface{}{
			"doors": []map[string]interface{}{
				map[string]interface{}{
					"doorid":"abc123",
					"is_locked":true,
				},
				map[string]interface{}{
					"doorid":"def456",
					"is_locked":false,
				},
				map[string]interface{}{
					"doorid":"ghi789",
					"is_locked":false,
				},
			},
		},
)
```

You could then send this over some transport, such as a websocket, or message queue, IP over avian carriers, or whatever.

You could also receive a CEL message with code similar to the following:
```go
var msg cel_mesage.Type

err := json.Unmarshal(data, &msg)

```

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

The serialized JSON will look like:

An example **command**:
```json
{
	"magic"   : "CEL/1",
	"version" : "1",
	"kind"    : "EVENT",
	"name"    : "LOCK_DOOR",
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
				"is_locked":true,
			},
			{
				"doorid":"def456",
				"is_locked":false,
			},
			{
				"doorid":"ghi789",
				"is_locked":false,
			},
		],
	},
}
```

`"magic"` is (for now) always `"CEL/1"`.

`"version"` can be set to whatever you want it to be. It is a means for your to (future-proof yourself and) version your API. (If you aren't sure what to set for this, just use (the string value) `"1"`.)

The only valid values for `"kind"` are
`"COMMAND"`,
`"EVENT"`, and
`"LOG"`.

You will come up with the values of `"name"`.

And the value of `"payload"` is dependent on what the value for `"name"` is.

## See Also

A similar library has also been created for JavaScript:

* https://github.com/reiver/cel.js
