# Hide IDs

[![Go Reference](https://pkg.go.dev/badge/github.com/emvi/hide?status.svg)](https://pkg.go.dev/github.com/emvi/hide?status)
[![CircleCI](https://circleci.com/gh/emvi/hide.svg?style=svg)](https://circleci.com/gh/emvi/hide)
[![Go Report Card](https://goreportcard.com/badge/github.com/emvi/hide)](https://goreportcard.com/report/github.com/emvi/hide)
<a href="https://discord.gg/fAYm4Cz"><img src="https://img.shields.io/discord/739184135649886288?logo=discord" alt="Chat on Discord"></a>

Hide is a simple package to provide an ID type that is marshaled to/from a hash string.
This prevents sending technical IDs to clients and converts them on the API layer.
Hide uses [hashids](https://github.com/speps/go-hashids) as its default hash function.
But you can provide your own by implementing the `Hash` interface and configuring it using `hide.UseHash`.

[Read our full article on Medium.](https://medium.com/emvi/golang-transforming-ids-to-a-userfriendly-representation-in-web-applications-85bf2f7d71c5)

## Installation

```
go get github.com/emvi/hide/v2
```

## Example

Consider the following struct:

```
type User struct {
    Id       uint64  `json:"id"`
    Username string `json:"username"`
}
```

When marshaling this struct to JSON, the ID will be represented by a number:

```
{
    "id": 123,
    "username": "foobar"
}
```

In this case, you expose the technical user ID to your clients. By changing the type of the ID, you get a better result:

```
type User struct {
    Id       hide.ID `json:"id"`
    Username string  `json:"username"`
}
```

Notice that the `uint64` ID got replaced by the `hide.ID`, which internally is represented as an `uint64` as well, but implements the marshal interface.
This allows you to cast between them and use `hide.ID` as a replacement. The resulting JSON changes to the following:

```
{
  "id": "beJarVNaQM",
  "username": "foobar"
}
```

If you send the new ID (which is a string now) back to the server and unmarshal it into the `hide.ID` type, you'll get the original technical ID back.
It's also worth mentioning that a value of 0 is translated to `null` when an ID is marshaled to JSON or stored in a database.

[View the full demo](https://github.com/emvi/hide-example)

## Contribute

[See CONTRIBUTING.md](CONTRIBUTING.md)

## License

MIT

<p align="center">
    <img src="hidegopher.png" width="300px" />
</p>
