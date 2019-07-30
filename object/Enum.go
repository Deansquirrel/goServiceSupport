package object

type ErrTypeCode int

const (
	ErrTypeCodeNoError ErrTypeCode = 200
)

type ErrTypeMsg string

const (
	ErrTypeMsgNoError ErrTypeMsg = "success"
)
