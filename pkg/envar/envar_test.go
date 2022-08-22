package envar

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	tests := []struct {
		Key       string
		Default   interface{}
		SetEnvVal string
		Want      interface{}
	}{
		{
			Key:       "KEY_TEST_01",
			Default:   0,
			SetEnvVal: "15",
			Want:      15,
		},
		{
			Key:       "KEY_TEST_02",
			Default:   "",
			SetEnvVal: "localhost",
			Want:      "localhost",
		},
		{
			Key:       "KEY_TEST_03",
			Default:   3307,
			SetEnvVal: "",
			Want:      3307,
		},
		{
			Key:       "KEY_test_04",
			Default:   "http://",
			SetEnvVal: "",
			Want:      "http://",
		},
	}
	for _, test := range tests {
		if test.SetEnvVal != "" {
			os.Setenv(test.Key, test.SetEnvVal)
		}
		assert.Equal(t, test.Want, GetEnv(test.Key, test.Default))
	}
}
