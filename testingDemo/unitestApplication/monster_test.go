package unitestApplication

import (
	"testing"
)

// 测试用例，测试Store方法
func TestMonster_Store(t *testing.T) {
	// 创建一个Monster实例
	monster := Monster{"Router", 23, "Basketball"}
	ret := monster.Store()
	if !ret {
		t.Fatalf("test monster.Store() fail:%v", ret)
	}
	t.Logf("test monster.Store() success")
}

// 测试用例二，测试ReStore方法
func TestMonster_ReStore(t *testing.T) {
	// 创建一个Monster实例，不需要传字段
	monster := Monster{}
	ret := monster.ReStore()
	if !ret {
		t.Fatalf("test monster.ReStore() fail:%v", ret)
	}
	// 进一步判断
	if monster.Name != "Router" {
		t.Fatalf("test monster name has change")
	}

	t.Logf("test monster.ReStore() success")
}
