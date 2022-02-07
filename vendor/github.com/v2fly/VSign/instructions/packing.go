package instructions

import "fmt"

func PackToString(instr Instruction, withHash bool) string {
	switch ins := instr.(type) {
	case FileIns:
		return packFileToString(ins, withHash)
	case AttributeIns:
		return packAttributeToString(ins, withHash)
	default:
		panic(nil)
	}
}

func packFileToString(instr FileIns, withHash bool) string {
	var content string

	content = instr.Filename()

	if !withHash {
		return content
	}

	return fmt.Sprintf("SHA256 (%v) = %v", content, instr.Hash())
}
func packAttributeToString(instr AttributeIns, withHash bool) string {
	inshash := instr.Hash()
	var attrtag string
	if instr.Critical() {
		if inshash != CriticalAttrHash {
			panic("Incorrect Hash")
		}
		attrtag = "!#"
	} else {
		if inshash != AttrHash {
			panic("Incorrect Hash")
		}
		attrtag = "##"
	}
	var extendedVersionNameDot = ""
	if instr.AttrExtendedName() != "" {
		extendedVersionNameDot = "."
	}
	content := fmt.Sprintf("%v%v%v%v%v=%v=%v", instr.AssociatedFile(), attrtag, instr.AttrMajorName(), extendedVersionNameDot, instr.AttrExtendedName(), instr.EncodingMethod(), instr.AttrDataEncoded())
	if instr.IsTail() {
		content = "|" + content
	}
	if !withHash {
		return content
	}
	return fmt.Sprintf("SHA256 (%v) = %v", content, instr.Hash())
}
