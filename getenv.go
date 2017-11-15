package getenv

import (
	"os"
	"strconv"
)

type envstr string

func GetEnv(key string, def string) *envstr {
	if val, ok := os.LookupEnv(key); ok {
		return newEnvStr(val)
	}
	return newEnvStr(def)
}

func newEnvStr(s string) *envstr {
	env := envstr(s)
	return &env
}

func (e *envstr) String() string {
	return string(*e)
}

func (e *envstr) Int() int {
	i, err := strconv.ParseInt(string(*e), 10, 0)
	if err != nil {
		return 0
	}
	return int(i)
}

func (e *envstr) Bool() bool {
	b, err := strconv.ParseBool(string(*e))
	if err != nil {
		return false
	}
	return b
}