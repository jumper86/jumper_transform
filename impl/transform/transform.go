package transform

import (
	"fmt"
	"github.com/jumper86/jumper_transform/interf"

	//"../compress"
	//"../encrypt"
	//"../packet"
	"github.com/jumper86/jumper_transform/impl/compress"
	"github.com/jumper86/jumper_transform/impl/encrypt"
	"github.com/jumper86/jumper_transform/impl/packet"
)

// 使用一系列　Operation 来构造一条操作链，
//　对于封包过程来说，　输入为　interface{}, 输出为　[]byte
//　对于解包过程来说，　输入为 []byte, 输出为 interface{}

//!!!note:　需要明确一点,　一个单独的PackageOpLink是支持并发调用的，即可以在多个协程中互不干扰的使用
// 这是因为　所有涉及到的数据只有两种情况
// 1. 入参传入，因此不同协程不会具有共用数据　2. 初始化的成员如direct或者密钥，他们都是初始化之后不再变化，即只进行读取，不进行写入，因此可以并发访问

//因此在实际项目使用中可以考虑定义一个全局的 map[string]Transform　不同的需求定义不同的　Transform　放入其中．

type transform struct {
	opLink []interf.Operation
}

func Newtransform() interf.Transform {
	var tf transform
	return &tf
}

func (self *transform) checkParam(opType interf.OperationType, direct bool, params []interface{}) bool {
	if opType <= interf.PackageOpMin || opType >= interf.PackageOpMax {
		fmt.Println("invalid opType.")
		return false
	}

	if (opType == interf.EncryptMd5 || opType == interf.EncryptSha1) && !direct {
		fmt.Println("md5 | sha1 couldn't direct false.")
		return false
	}

	if (opType == interf.EncryptAes || opType == interf.EncryptDes) && len(params) != 1 {
		fmt.Println("aes | des should set key.")
		return false
	}

	if (opType == interf.EncryptRsa) && len(params) != 2 {
		fmt.Println(" rsa should set key.")
		return false
	}

	return true
}

func (self *transform) AddOp(opType interf.OperationType, direct bool, params []interface{}) bool {
	if !self.checkParam(opType, direct, params) {
		return false
	}

	var op interf.Operation

	switch opType {
	//编码
	case interf.PacketBase64:
		op = packet.NewpacketOpBase64(direct, nil)
		break

	case interf.PacketJson:
		op = packet.NewpacketOpJson(direct, nil)
		break

	case interf.PacketXml:
		op = packet.NewpacketOpXml(direct, nil)
		break

	case interf.PacketProtobuf:
		op = packet.NewpacketOpProtobuf(direct, nil)
		break

	case interf.PacketBinary:
		op = packet.NewpacketOpBinary(direct, nil)
		break

		//压缩
	case interf.CompressGzip:
		op = compress.NewcompressOpGzip(direct, nil)
		break

	case interf.CompressZlib:
		op = compress.NewcompressOpZlib(direct, nil)
		break

		//加密
	case interf.EncryptMd5:
		op = encrypt.NewencryptOpMd5(direct, nil)
		break

	case interf.EncryptSha1:
		op = encrypt.NewencryptOpSha1(direct, nil)
		break

	case interf.EncryptAes:
		op = encrypt.NewencryptOpAes(direct, params)
		break

	case interf.EncryptDes:
		op = encrypt.NewencryptOpDes(direct, params)
		break

	case interf.EncryptRsa:
		op = encrypt.NewencryptOpRsa(direct, params)
		break

	default:
		return false
		break
	}

	self.opLink = append(self.opLink, op)
	return true

}

func (self *transform) Execute(input interface{}, output interface{}) error {

	//这里是一个链式反应，因此需要根据op类型来构建中间类型
	//中间过程的输出都是 []byte
	var tmpOutput []byte
	var tmpInput interface{}
	tmpInput = input

	for k := range self.opLink {
		if k == len(self.opLink)-1 {
			rst, err := self.opLink[k].Operate(tmpInput, output)
			if !rst {
				fmt.Printf("op error: %s", err)
				return err
			}
			fmt.Printf("\ntmpInput: %v,\noutput: %v\n\n", tmpInput, output)

		} else {
			self.opLink[k].Operate(tmpInput, &tmpOutput)
			fmt.Printf("\ntmpInput: %v,\ntmpOutput: %v\n\n", tmpInput, tmpOutput)
			tmpInput = tmpOutput
		}
	}

	return nil
}

func (self *transform) Reset() {
	if len(self.opLink) != 0 {
		self.opLink = self.opLink[:0]
	}
	return
}
