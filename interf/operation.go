package interf

//transform 表示一个操作，　包含三种类型：
//1. encode
//2. compress
//3. encrypt

//定义打包和加密类型
type OperationType int

const (
	PackageOpMin OperationType = 0 + iota
	//封包
	PacketBase64
	PacketJson
	PacketXml
	PacketProtobuf
	PacketBinary

	//压缩
	CompressGzip
	CompressZlib

	//加密
	EncryptMd5
	EncryptSha1
	EncryptAes
	EncryptDes
	EncryptRsa

	PackageOpMax
)

const (
	Forward  int8 = 1 //打包->压缩->加密
	Backward int8 = 2 //解密->解压->解包
)

//操作接口
type Operation interface {
	Operate(direct int8, input interface{}, output interface{}) (bool, error)
}
