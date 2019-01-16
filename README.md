# Hide IDs

[![GoDoc](https://godoc.org/github.com/emvicom/hide?status.svg)](https://godoc.org/github.com/emvicom/hide)
[![CircleCI](https://circleci.com/gh/emvicom/hide.svg?style=svg)](https://circleci.com/gh/emvicom/hide)
[![Go Report Card](https://goreportcard.com/badge/github.com/emvicom/hide)](https://goreportcard.com/report/github.com/emvicom/hide)

Hide is a simple package to provide an ID type that is marshalled to/from a hash string.
This prevents sending technical IDs to clients and offers auto converting those IDs on the API layer.
Hide uses [hashids](https://github.com/speps/go-hashids) as its default hash function.
But you can provide your own by implementing the `Hash` interface and configuring it using `hide.UseHash`.

## Installation

```
go get github.com/emvicom/hide
```

## Example

Consider the following struct:

```
type User struct {
    Id       int64  `json:"id"`
    Username string `json:"username"`
}
```

When marshalling this struct to JSON, the ID will be represented by a number:

```
{
    "id": 123,
    "username": "foobar"
}
```

In this case you expose the technical user ID to your clients. By changing the type of the ID, you get a better result:

```
type User struct {
    Id       hide.ID `json:"id"`
    Username string  `json:"username"`
}
```

Notice that the `int64` ID got replaced by the `hide.ID`, which internally is represented as an `int64` as well, but implements the marshal interface.
This allows you to cast between them and use `hide.ID` as an replacement. The resulting JSON changes to the following:

```
{
  "id": "beJarVNaQM",
  "username": "foobar"
}
```

If you send the new ID (which is a string now) back to the server and unmarshal it into the `hide.ID` type, you'll get the original technical ID back.

## Contribute

[See CONTRIBUTING.md](CONTRIBUTING.md)

## License

MIT
