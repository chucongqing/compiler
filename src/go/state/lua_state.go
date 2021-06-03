package state

import (
	"golua/binchunk"
)

type luaState struct {
	stack *luaStack

	proto *binchunk.Prototype
	pc    int //program counter
}

func New(stackSize int, proto *binchunk.Prototype) *luaState {
	return &luaState{
		stack: newLuaStack(20),
		proto: proto,
		pc:    0,
	}
}
