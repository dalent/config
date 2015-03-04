package config

import (
	"os"
	"testing"
)

var iniContext = `
;sdkfj
#sdfs
test=sdf
[test]
h=hello
t=test
i=12
te tst
m
[test1]
`

func TestIni(t *testing.T) {
	f, err := os.Create("app.conf")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(iniContext)
	f.Close()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("app.conf")
	iniConfig, err := NewConfiger("ini", "app.conf")
	if err != nil {
		t.Fatal(err)
	}

	v, _ := iniConfig.String("test", "h")
	if v != "hello" {
		t.Fatal("h shoud be hello")
	}

	i, _ := iniConfig.Int("test", "i")
	if i != 12 {
		t.Fatal("i shoud be 12")
	}
	et, _ := iniConfig.String("test", "t")
	if et != "test" {
		t.Fatal("t shoud be test")
	}

	m, _ := iniConfig.String("test", "m")
	if m != "" {
		t.Fatal("t shoud be ")
	}
	te, _ := iniConfig.String("test", "te")
	if te != "tst" {
		t.Fatal("t shoud be tst")
	}
}
