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

//操作接口
type Operation interface {
	//Init(direct bool, params []interface{}) bool //direct 表示操作方向, true　表示　编码/压缩/加密，　false 表示　解码/解压/解密
	Operate(input interface{}, output interface{}) (bool, error)
}
