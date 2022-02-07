package insmgr

import (
	"fmt"
	"github.com/v2fly/VSign/instructions"
	"io"
)

type OutputInsMgr struct {
	out io.Writer
}

func (o OutputInsMgr) SubmitIns(instruction instructions.Instruction) {
	out := instructions.PackToString(instruction, true)
	fmt.Fprintf(o.out, "%v\n", out)
}

func NewOutputInsMgr(out io.Writer) InstructionMgr {
	return &OutputInsMgr{out}
}
