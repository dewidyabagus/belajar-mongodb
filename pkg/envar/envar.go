package envar

import (
	"os"
	"strconv"
	"strings"
)

type Typ interface {
	string | int | interface{}
}

// TODO: Mengambil nilai dari variable env sesuai dengan key (kunci) yang dimasukan
//       ketika key tidak diketemukan maka menggunakan default value yang sudah di set.
func GetEnv[T Typ](key string, defValue T) T {
	var result interface{} = defValue

	valEnv := strings.TrimSpace(os.Getenv(key))
	if valEnv == "" {
		return result.(T)
	}

	switch result.(type) {
	case int:
		if res, err := strconv.Atoi(valEnv); err == nil {
			result = res
		}

	case string:
		result = valEnv

	}

	return result.(T)
}
