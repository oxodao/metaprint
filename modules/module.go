package modules

type Module interface {
	Print(args []string) string
	GetPrefix() string
	GetSuffix() string
}
