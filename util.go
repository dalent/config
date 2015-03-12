package config

import (
	"strconv"
)

const (
	gB = 1 << (iota * 10)
	gK
	gM
	gG
)

//我们可以读出 例如 10M这样的尺寸
func string2int(str string) (int64, error) {
	if len(str) == 0 {
		return 0, nil
	}

	//base 16 or base 8
	if str[0] == '0' {
		return strconv.ParseInt(str, 0, 0)
	}

	var integer int64 = 0
	for _, v := range str {
		if v <= '9' && v >= '0' {
			integer = integer*10 + int64(v-'0')
			continue
		}

		switch v {
		case 'k', 'K':
			integer *= gK
		case 'M', 'm':
			integer *= gM
		case 'G', 'g':
			integer *= gG
		}
		break
	}

	return integer, nil
}
