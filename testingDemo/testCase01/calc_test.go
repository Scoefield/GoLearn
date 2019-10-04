package testCase01

import "testing"

func TestAddUpper(t *testing.T) {
	ret := addUpper(10)
	if ret != 55 {
		t.Fatalf("test fail:expection:%v, actual:%v", 55, ret)
	}
	t.Logf("addUpper(10) exec correct")
}

func TestGetMulti(t *testing.T) {
	ret := getMulti(3, 8)
	if ret != 24 {
		t.Fatalf("test getMulti(3, 8) fail: %v", ret)
	}
	t.Logf("getMulti test success")
}
