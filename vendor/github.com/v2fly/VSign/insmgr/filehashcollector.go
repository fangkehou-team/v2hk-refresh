package insmgr

import (
	"github.com/v2fly/VSign/instructions"
)

type HashCollectorMgr struct {
	hash     map[string]string
	inverted bool
}

func (o HashCollectorMgr) SubmitIns(instruction instructions.Instruction) {
	switch e := instruction.(type) {
	case instructions.FileIns:
		filename := e.Filename()
		hash := e.Hash()
		if !o.inverted {
			o.hash[filename] = hash
		} else {
			o.hash[hash] = filename
		}
	}
}
func (o HashCollectorMgr) Result() map[string]string {
	return o.hash
}
func NewHashCollectorMgr(inverted bool) *HashCollectorMgr {
	return &HashCollectorMgr{make(map[string]string), inverted}
}
