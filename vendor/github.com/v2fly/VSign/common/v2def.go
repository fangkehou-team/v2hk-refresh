package common

type Command interface {
	Name() string
	DescriptionShort() string
	DescriptionUsage() []string
	Execute(args []string) error
}
