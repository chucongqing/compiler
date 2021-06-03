package api

type LuaVM interface {
	LuaState
	PC() int          //get current PC (for test)
	AddPC(n int)      //modify PC (to imply jump)
	Fectch() uint32   //get current instruction,let pc point to next instruction
	GetConst(idx int) //push const at idx position to stack top
	GetRK(rk int)     //push const at idx or stack value to stack top
}
