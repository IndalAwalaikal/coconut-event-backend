

package main


// This tool was moved to cmd/tools to avoid interfering with `go build ./...`.
// The original implementation is intentionally excluded from normal builds.

import "fmt"

func main() {
    fmt.Println("gen_admin tool excluded from build; see cmd/tools for runnable tools")
}
    
