package instructions

type Instruction interface {
	Instruction()
	Hash() string
	Filename() string
}

type FileIns interface {
	Instruction()
	Hash() string
	Filename() string

	File()
}

type AttributeIns interface {
	Instruction()
	Filename() string
	Hash() string

	Attribute()

	AssociatedFile() string
	EncodingMethod() string
	Critical() bool
	AttrMajorName() string
	AttrExtendedName() string
	AttrDataEncoded() string

	IsTail() bool
}

const CriticalAttrHash = "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
const AttrHash = "0000000000000000000000000000000000000000000000000000000000000000"
