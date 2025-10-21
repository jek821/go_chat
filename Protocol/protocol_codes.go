package Protocol

type Code int

const (
	TestCode Code = iota
	RequestClientIdCode
	GiveClientIdCode
	EndClientCode
)
