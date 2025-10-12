package utils

type Code int

const (
	GiveClientNewIdCode Code = iota + 1
	MsgCode
	ConnectionRequestCode
	SessionRequestAcceptedCode
	SessionRequestRejectedCode
)
