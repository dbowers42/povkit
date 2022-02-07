package plugin

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

func compute(L *lua.LState) int {
	a := L.ToInt(1)
	b := L.ToInt(1)
	r := a + b
	ln := lua.LNumber(r)
	L.Push(ln)
	return 1
}

func generate() {
	var L = lua.NewState()
	defer L.Close()
	L.SetGlobal("compute", L.NewFunction(compute))
	err := L.DoFile("scripts/p1.lua")
	if err != nil {
		fmt.Println(err)
	}
}
