# config
After installing Go and setting up your GOPATH, create your first .go file. We'll call it server.go.


```
package main
import (
	"fmt"
	"github.com/dalent/config"
)
func main(){
    iniConfig, err := config.NewConfiger("ini", "app.conf")
    if err != nil {
        fmt.Println(err)
    }   

    v, _ := iniConfig.String("test", "h")
    fmt.Println(v)
}
```
Then install the config package
```
go get github.com/dalent/config
```
