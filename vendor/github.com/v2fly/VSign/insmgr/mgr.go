package insmgr

import "github.com/v2fly/VSign/instructions"

type InstructionMgr interface {
	SubmitIns(instruction instructions.Instruction)
}

type InstructionYield interface {
	InstructionYield(instMgr InstructionMgr)
}

type yieldSingle struct {
	yields instructions.Instruction
}

func (y yieldSingle) InstructionYield(instMgr InstructionMgr) {
	instMgr.SubmitIns(y.yields)
}

func NewYieldSingle(yields instructions.Instruction) InstructionYield {
	return &yieldSingle{
		yields: yields,
	}
}
