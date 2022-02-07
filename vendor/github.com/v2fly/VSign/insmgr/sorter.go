package insmgr

import (
	"github.com/v2fly/VSign/instructions"
	"sort"
)

type SortableInstructions struct {
	instructions.Instruction
	sortkey string
}

type SortableInstructionsHolder struct {
	insa []SortableInstructions
	ins  []instructions.Instruction
}

func (s SortableInstructionsHolder) Len() int {
	return len(s.insa)
}

func (s SortableInstructionsHolder) Less(i, j int) bool {
	return s.insa[i].sortkey < s.insa[j].sortkey
}

func (s SortableInstructionsHolder) Swap(i, j int) {
	s.insa[i], s.insa[j] = s.insa[j], s.insa[i]
	s.ins[i], s.ins[j] = s.ins[j], s.ins[i]
}

func SortIns(ins []instructions.Instruction) {
	inss := make([]SortableInstructions, 0, len(ins))
	for _, w := range ins {
		si := SortableInstructions{w, instructions.PackToString(w, false)}
		inss = append(inss, si)
	}
	sort.Stable(SortableInstructionsHolder{inss, ins})
}
