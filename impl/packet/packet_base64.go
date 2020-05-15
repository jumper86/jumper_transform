package packet

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/jumper86/jumper_transform/interf"
	"github.com/jumper86/jumper_transform/log"
	"reflect"
)

type packetOpBase64 struct {
}

func NewpacketOpBase64(params []interface{}) interf.PacketOp {
	var op packetOpBase64
	op.init(params)
	return &op
}

func (self *packetOpBase64) init(params []interface{}) bool {
	return true
}

func (self *packetOpBase64) Operate(direct int8, input interface{}, output interface{}) (bool, error) {

	if direct == interf.Forward {
		tmpOutput, err := self.Pack(input)
		if err != nil {
			fmt.Printf("pack failed. err: %s", err)
			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil

	} else {
		err := self.Unpack(input.([]byte), output)
		if err != nil {
			fmt.Printf("unpack failed. err: %s", err)
			return false, err
		}
		return true, nil
	}

	return true, nil
}

//此函数中需要检查入参是否为 string / []byte
func (*packetOpBase64) Pack(originData interface{}) ([]byte, error) {
	defer log.TraceLog("packetOpBase64.Pack")()
	//需要判断入参类型为 string 或者 []byte
	vod := reflect.ValueOf(originData)
	tod := reflect.TypeOf(originData)
	if vod.IsValid() == false {
		fmt.Println("originData's value is nil.")
		return nil, errors.New("originData's value is nil")
	}

	if vod.Kind() == reflect.String {

		rst := make([]byte, base64.StdEncoding.EncodedLen(vod.Len()))
		base64.StdEncoding.Encode(rst, []byte(originData.(string)))
		return rst, nil

	}

	if vod.Kind() == reflect.Slice && tod.Elem().Kind() == reflect.Uint8 {

		rst := make([]byte, base64.StdEncoding.EncodedLen(vod.Len()))
		base64.StdEncoding.Encode(rst, originData.([]byte))
		return rst, nil
	}

	fmt.Println("invalid param type.")
	return nil, errors.New("invalid param type, use string or []byte.")
}

func (*packetOpBase64) Unpack(packData []byte, obj interface{}) error {

	defer log.TraceLog("packetOpBase64.Unpack")()
	//判断接收结果的入参是一个*[]byte
	tod := reflect.TypeOf(obj)
	vod := reflect.ValueOf(obj)

	//这里需要注意区别 reflect.Value.Elem() 和 reflect.Type.Elem() 两个函数
	//要想查看 指针/数组/切片 等的元素类型应该使用 reflect.Type.Elem() 函数
	if vod.Kind() != reflect.Ptr || tod.Elem().Kind() != reflect.Slice || tod.Elem().Elem().Kind() != reflect.Uint8 {
		return errors.New("invalid param, should use *[]byte")
	}

	rst := make([]byte, base64.StdEncoding.DecodedLen(len(packData)))
	_, err := base64.StdEncoding.Decode(rst, packData)
	if err != nil {
		fmt.Println("decode failed, err:", err)
		return errors.New("decode failed.")
	}

	//类型是指针，需要使用Elem获得指针指向的内存，然后进行值的设置
	reflect.ValueOf(obj).Elem().SetBytes(rst)
	return nil
}
