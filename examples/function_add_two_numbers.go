package main

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func main() {
	m := ir.NewModule() // Think of a module like a source file in your project (top-level containers of LLVM IR objects)

	globalG := m.NewGlobalDef("g", constant.NewInt(types.I32, 2)) // Global variable "g" of type I32 with value 2

	funcAdd := m.NewFunc("add", types.I32, // Method "add" that returns I32
		ir.NewParam("x", types.I32), // Parameter "x" of type I32
		ir.NewParam("y", types.I32), // Parameter "y" of type I32
	)
	ab := funcAdd.NewBlock("")                                 // Block (Body) of the "add" function
	ab.NewRet(ab.NewAdd(funcAdd.Params[0], funcAdd.Params[1])) // First and last statement for the function "add", basically returns the sum of 'a' and 'b'

  // Main function
	funcMain := m.NewFunc(
		"main",
		types.I32,
	)
  // Block (body) of main function
	mb := funcMain.NewBlock("")
  // Return the call of function "add" with parameters 1 and global variable "g" (which has value of 2)
	mb.NewRet(mb.NewCall(funcAdd, constant.NewInt(types.I32, 1), mb.NewLoad(types.I32, globalG)))

	println(m.String())
}
