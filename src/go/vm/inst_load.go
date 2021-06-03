package vm

import . "golua/api"

func loadBool(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1
	vm.PushBoolean(b != 0)
	vm.Replace(a)
	if c != 0 {
		vm.AddPC(1)
	}
}

func loadK(i Instruction, vm LuaVM) {
	a, bx := i.ABx()
	a += 1
	vm.GetConst(bx)
	vm.Replace(a)
}

// iABx, Extraarg(iax)
func loadKX(i Instruction, vm LuaVM) {
	a, _ := i.ABx()
	a += 1
	ax := Instruction(vm.Fectch()).Ax()
	vm.GetConst(ax)
	vm.Replace(a)
}
