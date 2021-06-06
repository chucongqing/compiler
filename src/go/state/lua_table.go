package state

import (
	"golua/number"
)

type luaTable struct {
	arr  []luaValue
	_map map[luaValue]luaValue
}

func newLuaTable(nArr, nRec int) *luaTable {
	t := &luaTable{}

	if nArr > 0 {
		t.arr = make([]luaValue, 0, nArr)
	}

	if nRec > 0 {
		t._map = make(map[luaValue]luaValue, nRec)
	}

	return t
}

func (lt *luaTable) get(key luaValue) luaValue {
	key = _floatToIntger(key)
	if idx, ok := key.(int64); ok {
		if idx >= 1 && idx <= int64(len(lt.arr)) {
			return lt.arr[idx-1]
		}
	}

	return lt._map[key]
}

func _floatToIntger(key luaValue) luaValue {
	if f, ok := key.(float64); ok {
		if i, ok := number.FloatToInteger(f); ok {
			return i
		}
	}
	return key
}
