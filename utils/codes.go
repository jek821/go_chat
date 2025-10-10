package utils

type Code int

const (
	RegCli Code = iota + 1
	Msg
	ReqCon
	AccCon
	GiveClientNewId
)
