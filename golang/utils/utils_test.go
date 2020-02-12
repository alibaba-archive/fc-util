package utils

import "testing"

func Test_GetUrlFormedMap(t *testing.T) {
	tmp := map[string]string{
		"key": "value",
	}
	encode := GetUrlFormedMap(tmp)
	AssertEqual(t, encode, "key=value")
}
