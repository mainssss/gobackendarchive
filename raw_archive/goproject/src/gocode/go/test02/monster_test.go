package monster

import (
	"testing"
)

//测试用例，用于测试TestStore方法
func TestStore(t *testing.T) {
	//先创建一个monster实例
	monster := &Monster{
		Name:  "牛魔王",
		Age:   500,
		Skill: "牛魔拳",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误 期望=%v,实际返回的是=%v", true, res)
	}
	t.Fatal("monster.Store()正确")
}

func TestReStore(t *testing.T) {
	//先创建一个monster实例
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore()错误 期望=%v,实际返回的是=%v", true, res)
	}

	//进一步判断
	if monster.Name != "牛魔王" {
		t.Fatalf("monster.Store()错误 期望=%v,实际返回的是=%v", "牛魔王", monster.Name)
	}
	t.Fatal("monster.ReStore()正确")
}
