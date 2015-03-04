# config
After installing Go and setting up your GOPATH, create your first .go file. We'll call it server.go.

package main

import (
"github.com/dalent/config"
"fmt"
)

func main() {
  iniConfig, err := NewConfiger("ini", "app.conf")
    if err != nil {
      fmt.Println(err)
    }   

    v, _ := iniConfig.String("test", "h")
    fmt.Println(v)
}
Then install the config package

go get github.com/dalent/config
