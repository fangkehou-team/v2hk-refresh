package instructions

import (
	"regexp"
	"strings"
)

type unpackedFile struct {
	hash     string
	filename string
}

func (u unpackedFile) Instruction() {
	panic("implement me")
}

func (u unpackedFile) Hash() string {
	return u.hash
}

func (u unpackedFile) Filename() string {
	return u.filename
}

func (u unpackedFile) File() {
	panic("implement me")
}

type unpackedAttrib struct {
	hash     string
	filename string

	associatedFile   string
	encodingMethod   string
	critical         bool
	attrMajorName    string
	attrExtendedName string
	attrDataEncoded  string

	isTail bool
}

func (u unpackedAttrib) Instruction() {
	panic("implement me")
}

func (u unpackedAttrib) Hash() string {
	return u.hash
}

func (u unpackedAttrib) Filename() string {
	return u.filename
}

func (u unpackedAttrib) Attribute() {
	panic("implement me")
}

func (u unpackedAttrib) AssociatedFile() string {
	return u.associatedFile
}

func (u unpackedAttrib) EncodingMethod() string {
	return u.encodingMethod
}

func (u unpackedAttrib) Critical() bool {
	return u.critical
}

func (u unpackedAttrib) AttrMajorName() string {
	return u.attrMajorName
}

func (u unpackedAttrib) AttrExtendedName() string {
	return u.attrExtendedName
}

func (u unpackedAttrib) AttrDataEncoded() string {
	return u.attrDataEncoded
}

func (u unpackedAttrib) IsTail() bool {
	return u.isTail
}
func (u *unpackedAttrib) Parse() *unpackedAttrib {
	var re = regexp.MustCompile(`^(\|)?((?:[a-zA-Z0-9\-./_@]){0,128})(?:(#|!)#)((?:[a-zA-Z0-9]){0,16})(?:\.((?:[a-zA-Z0-9.]){0,16}))?=((?:[a-z0-9]){0,5}?)=((?:.){0,650})$`)
	res := re.FindStringSubmatch(u.filename)

	u.isTail = res[1] == "|"
	u.associatedFile = res[2]
	u.critical = res[3] == "!"
	u.attrMajorName = res[4]
	u.attrExtendedName = res[5]
	u.encodingMethod = res[6]
	u.attrDataEncoded = res[7]

	return u
}

const Initial = "SHA256 ("
const End1 = ") = "
const End2 = AttrHash

func UnpackInstruction(instr string) Instruction {
	if !strings.HasPrefix(instr, Initial) {
		return nil
	}
	instr = instr[len(Initial):]

	hash := instr[len(instr)-len(End2):]

	content := instr[:len(instr)-len(End2)-len(End1)]

	switch hash {
	case AttrHash:
		fallthrough
	case CriticalAttrHash:
		return unpackAttribInstr(content, hash)
	default:
		return unpackFileInstr(content, hash)

	}
	return nil
}

func unpackFileInstr(ct, ha string) FileIns {
	return &unpackedFile{hash: ha, filename: ct}
}

func unpackAttribInstr(ct, ha string) AttributeIns {
	return (&unpackedAttrib{hash: ha, filename: ct}).Parse()
}
