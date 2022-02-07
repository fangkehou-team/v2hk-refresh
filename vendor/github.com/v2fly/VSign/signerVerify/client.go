package signerVerify

import "github.com/v2fly/VSign/instructions"

func CheckAsClient(ins []instructions.Instruction, project string, invert bool) (map[string]string, string, bool) {
	var correctProject bool
	var hasVersion bool

	var version string

	filehashmap := make(map[string]string)

	for _, v := range ins {
		switch e := v.(type) {
		case instructions.AttributeIns:
			if e.Critical() {
				switch e.AttrMajorName() {
				case "version":
					if hasVersion == false {
						if e.EncodingMethod() == "" {
							hasVersion = true
							version = e.AttrDataEncoded()
						}
					} else {
						return nil, "", false
					}
				case "project":
					if correctProject == false {
						if e.AttrDataEncoded() == project && e.EncodingMethod() == "" {
							correctProject = true
						}
					} else {
						return nil, "", false
					}
				default:
					return nil, "", false
				}
			}
		case instructions.FileIns:
			hash := e.Hash()
			filename := e.Filename()
			if !invert {
				filehashmap[filename] = hash
			} else {
				if _, foundAlso := filehashmap[hash]; foundAlso {
					filehashmap[hash] += filename + ";"
				} else {
					filehashmap[hash] = filename + ";"
				}

			}

		}
	}
	return filehashmap, version, correctProject && hasVersion
}
