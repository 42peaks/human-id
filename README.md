# `@42peaks/human-id`

## Introduction
This package generates short, human-readable IDs of the form:
`<adjective>-<noun>-<number>` where `number` is a random number between 1 and 999999.

## Usage
Quite simple:

```go
import (
    humanid "go.42peaks.co/human-id"
    "fmt"
)

func main() {
	id, err := humanid.Generate()
}
```