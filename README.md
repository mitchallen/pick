# github.com/mitchallen/pick

## Usage

### Initialize your module

```
$ go mod init example.com/my-demo
```

### Get the module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/mitchallen/pick@v0.1.0
```

### pick.One()

```go
package main

import (
	"fmt"
	"github.com/mitchallen/pick"
)

func main() {
    list := []int{10, 20, 30}
    fmt.Println(pick.OneInt(list))
}
```

## Init

How this repo was intialized as a **go** (golang) module

```
$ go mod init github.com/mitchallen/pick
```

## Testing

```
$ go test
```

```
$ cd (package)
$ go test
```

## Tagging

Substitute the light weight tag value with the latest tag to publish

```
$ git tag v0.1.0
$ git push origin --tags
```

## Publishing instructions

* https://go.dev/doc/modules/publishing

## Package

* https://pkg.go.dev/search?q=mitchallen/
