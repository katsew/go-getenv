# Add to GOPATH

```go

go get github.com/katsew/go-getenv

```

# Usage

```go

package main

import (
  ge "github.com/katsew/go-getenv"
)


// If a environment variable (first arg) is defined, specified value return.
// Otherwise, return default value (second arg).
func main() {

  v := ge.GetEnv("FROM_ENVIRONMENT_VARIABLE", "default value").String()
  log.Print(v)

}

```
