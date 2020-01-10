package version

import (
  "fmt"
)

// NUMBER should be overridden by a build flag.
var NUMBER = "unknown"

// COMMIT should be overridden by a build flag.
var COMMIT = "unknown"

func Identifier() string {
  return fmt.Sprintf("version: %s / commit: %s", NUMBER, COMMIT)
}
