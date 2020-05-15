package interf

type Transform interface {
	AddOp(opType OperationType, params []interface{}) bool
	Execute(direct int8, input interface{}, output interface{}) error
	Reset()
}
