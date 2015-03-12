package config

import (
	"testing"
)

func TestUtil(t *testing.T) {
	integer, _ := string2int("10M")
	if 10*1024*1024 != integer {
		t.Fail()
	}

	integer, _ = string2int("10m")
	if 10*1024*1024 != integer {
		t.Fail()
	}

	integer, _ = string2int("123456")
	if 123456 != integer {
		t.Fail()
	}
}
