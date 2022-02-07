package instimp

import "github.com/v2fly/VSign/instructions"

type SimpleFilenameKeyValueInst struct {
	assocFilename string
	key           string
	keyExt        string
	value         string
	critical      bool
}

func (vi SimpleFilenameKeyValueInst) Filename() string {
	panic("implement me")
}

func (vi SimpleFilenameKeyValueInst) Instruction() {
	panic("implement me")
}

func (vi SimpleFilenameKeyValueInst) Hash() string {
	if vi.critical {
		return instructions.CriticalAttrHash
	}
	return instructions.AttrHash
}

func (vi SimpleFilenameKeyValueInst) Attribute() {
	panic("implement me")
}

func (vi SimpleFilenameKeyValueInst) AssociatedFile() string {
	return vi.assocFilename
}

func (vi SimpleFilenameKeyValueInst) EncodingMethod() string {
	return ""
}

func (vi SimpleFilenameKeyValueInst) Critical() bool {
	return vi.critical
}

func (vi SimpleFilenameKeyValueInst) AttrMajorName() string {
	return vi.key
}

func (vi SimpleFilenameKeyValueInst) AttrExtendedName() string {
	return vi.keyExt
}

func (vi SimpleFilenameKeyValueInst) AttrDataEncoded() string {
	return vi.value
}

func (vi SimpleFilenameKeyValueInst) IsTail() bool {
	return false
}

func (vi *SimpleFilenameKeyValueInst) AsIns() instructions.AttributeIns {
	return vi
}

func NewSimpleFilenameKeyValueInst(filename, key, value string, critical bool) *SimpleFilenameKeyValueInst {
	return &SimpleFilenameKeyValueInst{
		assocFilename: filename,
		key:           key,
		value:         value,
		critical:      critical,
	}
}

func NewSimpleFilenameKeyValueInst5(filename, key, keyext, value string, critical bool) *SimpleFilenameKeyValueInst {
	return &SimpleFilenameKeyValueInst{
		assocFilename: filename,
		key:           key,
		keyExt:        keyext,
		value:         value,
		critical:      critical,
	}
}
