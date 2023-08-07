package main

import (
	"fmt"
	"net/http"

	lhttp "github.com/cjoudrey/gluahttp"
	lua "github.com/yuin/gopher-lua"
	ljson "layeh.com/gopher-json"
)

type LuaRunner struct {
	stdout []interface{}
}

func (lr *LuaRunner) printToGo(L *lua.LState) int {
	top := L.GetTop()
	args := make([]interface{}, top)

	for i := 1; i <= top; i++ {
		args[i-1] = L.Get(i).String()
	}
	lr.stdout = append(lr.stdout, args...)
	return 0
}

func (lr *LuaRunner) loadModules(L *lua.LState) {
	L.PreloadModule("http", lhttp.NewHttpModule(&http.Client{}).Loader)
	ljson.Preload(L)
}

func (lr *LuaRunner) RunLuaScript(luaFile string) (*string, error) {

	L := lua.NewState()
	defer L.Close()

	// Preload modules
	lr.loadModules(L)

	L.SetGlobal("print", L.NewFunction(lr.printToGo))

	err := L.DoFile(luaFile)
	if err != nil {
		fmt.Println("Error executing Lua script:", err)
		return nil, err
	}

	fmt.Println(lr.stdout...)
	lr.stdout = make([]interface{}, 0)

	return nil, nil
}

func (lr *LuaRunner) RunLuaFunc(luaFile string) (*string, error) {
	L := lua.NewState()
	defer L.Close()

	// Preload modules
	lr.loadModules(L)

	err := L.DoFile(luaFile)
	if err != nil {
		return nil, fmt.Errorf("Error loading Lua string: %v ", err)
	}

	err = L.CallByParam(lua.P{
		Fn:      L.GetGlobal("main"),
		NRet:    1,
		Protect: true,
	})
	if err != nil {
		return nil, fmt.Errorf("Error calling Lua function: %v ", err)
	}

	result := L.Get(-1)
	L.Pop(1)
	resultStr := result.String()
	return &resultStr, nil
}
