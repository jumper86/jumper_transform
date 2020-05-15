package interf

type PacketOp interface {
	Operation
	Pack(originData interface{}) ([]byte, error)
	Unpack(packData []byte, obj interface{}) error
}
