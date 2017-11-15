package getenv

import (
	"testing"
	"reflect"
)

func TestGetEnv(t *testing.T) {

	env := GetEnv("TEST_STR_ENV", "NG")
	envType := reflect.TypeOf(env)
	if envType.String() != "*getenv.envstr" {
		t.Errorf("Type mismatch: %s", envType.String())
	}
	t.Logf("Ok: %s", envType)

}

func TestEnvstr_String(t *testing.T) {
	str := GetEnv("TEST_STR_ENV", "NG").String()
	if str != "hello" {
		t.Errorf("Failed to get env with key: %s, result: %s", "TEST_STR_ENV", str)
		return
	}
	t.Logf("Ok: %s", str)
}

func TestEnvstr_Int(t *testing.T) {
	intval := GetEnv("TEST_INT_ENV", "0").Int()
	if intval != 128 {
		t.Errorf("Failed to get env with key: %s, result: %d", "TEST_INT_ENV", intval)
		return
	}
	t.Logf("Ok: %d", intval)
}

func TestEnvstr_Bool(t *testing.T) {
	boolval := GetEnv("TEST_BOOL_ENV", "false").Bool()
	if boolval != true {
		t.Errorf("Failed to get env with key: %s, result: %d", "TEST_BOOL_ENV", boolval)
		return
	}
	t.Logf("Ok: %v", boolval)
}