package main

import (
	"fmt"
	"io"
	"log"
	"os"

	lua "github.com/yuin/gopher-lua"
)

func deps(filePath string) (*string, error) {
	L := lua.NewState()
	defer L.Close()

	err := L.DoFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("deps:: error running lua file")
	}

	luaFn := L.GetGlobal("deps")
	if luaFn == lua.LNil {
		return nil, fmt.Errorf("deps:: lua function not found")
	}

	err = L.CallByParam(lua.P{
		Fn:      luaFn,
		NRet:    1,
		Protect: true,
	})
	if err != nil {
		return nil, fmt.Errorf("deps:: error calling Lua function: %v", err)
	}
	ret := L.Get(-1)
	L.Pop(1)
	println(ret)

	if str, ok := ret.(lua.LString); ok {
		retStr := string(str)
		return &retStr, nil
	}

	return nil, fmt.Errorf("deps:: invalid return value type")
}

func run(filePath string) error {

	log.Println("Running script ", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("run:: error opening Lua file: %v", err)
	}
	defer file.Close()

	script, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("run:: error reading Lua file: %v", err)
	}

	L := lua.NewState()
	defer L.Close()

	if err := L.DoString(string(script)); err != nil {
		return fmt.Errorf("run:: error running Lua file: %v", err)
	}

	return nil
}
