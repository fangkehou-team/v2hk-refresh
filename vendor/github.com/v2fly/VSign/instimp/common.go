package instimp

import "github.com/v2fly/VSign/instructions"

func NewVersionIns(version string) instructions.Instruction {
	return NewSimpleFilenameKeyValueInst("", "version", version, true)
}

func NewProjectIns(project string) instructions.Instruction {
	return NewSimpleFilenameKeyValueInst("", "project", project, true)
}
