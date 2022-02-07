package signerVerify

import (
	"github.com/v2fly/VSign/instructions"
)

func CheckVersionAndProject(ins []instructions.Instruction, version, project string) bool {
	var correctProject bool
	var correctVersion bool
	for _, v := range ins {
		switch e := v.(type) {
		case instructions.AttributeIns:
			if e.Critical() {
				switch e.AttrMajorName() {
				case "version":
					if correctVersion == false {
						if e.AttrDataEncoded() == version && e.EncodingMethod() == "" {
							correctVersion = true
						}
					} else {
						return false
					}
				case "project":
					if correctProject == false {
						if e.AttrDataEncoded() == project && e.EncodingMethod() == "" {
							correctProject = true
						}
					} else {
						return false
					}
				default:
					return false
				}
			}
		}
	}
	return correctProject && correctVersion
}
