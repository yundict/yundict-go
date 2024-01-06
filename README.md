# yundict-go

Yundict Golang client library for accessing the [Yundict API](https://yundict.com/docs/api/).

## Installation

```bash
go get github.com/yundict/yundict-go
```

## Usage

```go
import "github.com/yundict/yundict-go"
```

Construct a new Yundict client, then use the various services on the client to access different parts of the Yundict API. For example:

```go
client := yundict.NewClient(nil)

// export all keys & translations
keys, err := client.Keys.Export("org/project-name")
```