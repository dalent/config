# config
like this 
```
[server]
name=test

[db]
addr = localhost

[test]
h=123
```
example
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
