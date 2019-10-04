package testCase01

import "testing"

func TestGetSub(t *testing.T) {
	ret := getSub(20, 12)
	if ret != 8 {
		t.Fatalf("getSub fail: epection:%v, actual:%v", 8, ret)
	}
	t.Logf("getSub test success")
}
