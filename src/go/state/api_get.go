package state

import . "golua/api"

func (ls *luaState) CreateTable(nArr, nRec int) {
	t := newLuaTable(nArr, nRec)
	ls.stack.push(t)
}

func (ls *luaState) NewTable() {
	ls.CreateTable(0, 0)
}

// get value from key (pop from stack)
func (ls *luaState) GetTable(idx int) LuaType {
	t := ls.stack.get(idx)
	k := ls.stack.pop()
	return ls.getTable(t, k)
}

func (ls *luaState) getTable(t, k luaValue) LuaType {
	if tbl, ok := t.(*luaTable); ok {
		v := tbl.get(k)
		ls.stack.push(v)
		return typeOf(v)
	}
	panic("not a table!")
}

func (ls *luaState) GetField(idx int, k string) LuaType {
	t := ls.stack.get(idx)
	return ls.getTable(t, k)
}

func (ls *luaState) GetI(idx int, i int64) LuaType {
	t := ls.stack.get(idx)
	return ls.getTable(t, i)
}
