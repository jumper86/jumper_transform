package interf

type Transform interface {
	AddOp(opType OperationType, direct bool, params []interface{}) bool
	Execute(input interface{}, output interface{}) error
	Reset()
}
