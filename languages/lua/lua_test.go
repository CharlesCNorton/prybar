package main

import (
	"testing"
)

// Mostly a debugging aid, not sure if it'll actually fail if something goes wrong
func TestLuaStuff(t *testing.T) {
	var lua Lua
	lua.Open()
	lua.Version()
	for i := 0; i < 100; i++ {
		lua.EvalFile("testdata/test.lua", make([]string, 0))
	}
}
