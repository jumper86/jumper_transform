package main

import (
	"fmt"

	"github.com/jumper86/jumper_transform/def"

	"github.com/jumper86/jumper_transform"

	"github.com/jumper86/jumper_transform/interf"

	"github.com/golang/protobuf/proto"
)

func main() {
	////////////////base64////////////////

	//var op packet.packetOpBase64
	//op.Init(true, nil)
	//originData := []byte("this is a test.")
	//var rstData []byte
	//op.Operate(originData, &rstData)
	//fmt.Printf("rst: %s\n", rstData)
	//
	//var opr packet.packetOpBase64
	//opr.Init(false, nil)
	//
	//var rstrData []byte
	//opr.Operate(rstData, &rstrData)
	//fmt.Printf("rstr: %s\n", rstrData)

	//var op packet.packetOpBinary
	//op.Init(true, nil)
	//originData := &packet.Message{
	//	Type: 8,
	//	Content: []byte{1,2,3,4,5,6},
	//}
	//var rstData []byte
	//op.Operate(originData, &rstData)
	//fmt.Printf("rst: %v\n", rstData)
	//
	//var opr packet.packetOpBinary
	//opr.Init(false, nil)
	//var rstrData packet.Message
	//opr.Operate(rstData, &rstrData)
	//fmt.Printf("rstr: %v\n", rstrData)

	////////////////json////////////////

	//var op packet.packetOpJson
	//op.Init(true, nil)
	//test1 := struct{
	//	Name string
	//	Age int
	//}{
	//	Name: "wang",
	//	Age: 1,
	//}
	//
	//var rst1 []byte
	//op.Operate(test1, &rst1)
	//fmt.Printf("rst1: %s", rst1)
	//
	//var test2 struct{
	//	Name string
	//	Age int
	//}
	//
	//var op2 packet.packetOpJson
	//op2.Init(false, nil)
	//op2.Operate(rst1, &test2)
	//fmt.Printf("rst2: %v", test2)

	////////////////protobuf////////////////

	//bodyData := "guangzhou/fangcun/vip/company"
	//
	//p := &packet.StringMessage{
	//	Body: proto.String(bodyData),
	//	Header: &packet.Header{
	//		MessageId: proto.String("20-05"),
	//		Topic:     proto.String("golang"),
	//	},
	//}
	//
	//var op packet.packetOpProtobuf
	//op.Init(true, nil)
	//
	//var rst1 []byte
	//op.Operate(p, &rst1)
	//fmt.Printf("rst1: %v", rst1)
	//
	//
	//var test2 packet.StringMessage
	//var op2 packet.packetOpProtobuf
	//op2.Init(false, nil)
	//op2.Operate(rst1, &test2)
	////注意此处　打印　&test2 和　test2　是不一样的，因为 *(packet.StringMessage) 是定义了 String() 函数的
	//fmt.Println("rst2: ", &test2)

	////////////////xml////////////////

	//type info struct {
	//	Name string
	//	Age  int
	//	Male bool
	//}
	//
	////注意：xml　比较受限, 不能使用map 并且如下匿名结构体类型也是不能支持编码的
	////test1 := struct{
	////	Name string
	////	Age int
	////}{
	////	Name: "wang",
	////	Age: 1,
	////}
	//
	//var op packet.packetOpXml
	//op.Init(true, nil)
	//var test1 info
	//test1.Name = "wang"
	//test1.Age = 1
	//test1.Male = true
	//
	//var rst1 []byte
	//op.Operate(&test1, &rst1)
	//fmt.Printf("rst1: %s", rst1)
	//
	//
	//var test2 info
	//var op2 packet.packetOpXml
	//op2.Init(false, nil)
	//op2.Operate(rst1, &test2)
	//fmt.Printf("rst2: %v", test2)

	////////////////aes////////////////

	//var op encrypt.encryptOpAes
	//params := make([]interface{}, 0)
	//params = append(params, []byte("abcdefghijklmnop"))
	//op.Init(true, params)
	//
	//param1 := []byte("this is a aes test")
	//var rst1 []byte
	//op.Operate(param1, &rst1)
	//fmt.Printf("rst1: %x\n", rst1)
	//
	//var test2 []byte
	//var op2 encrypt.encryptOpAes
	//op2.Init(false, params)
	//op2.Operate(rst1, &test2)
	//fmt.Printf("rst2: %s", test2)

	////////////////des////////////////

	//var op encrypt.encryptOpDes
	//params := make([]interface{}, 0)
	//params = append(params, []byte("ijklmnop"))
	//op.Init(true, params)
	//
	//param1 := []byte("this is a des test")
	//var rst1 []byte
	//op.Operate(param1, &rst1)
	//fmt.Printf("rst1: %x\n", rst1)
	//
	//var test2 []byte
	//var op2 encrypt.encryptOpDes
	//op2.Init(false, params)
	//op2.Operate(rst1, &test2)
	//fmt.Printf("rst2: %s", test2)

	////////////////rsa////////////////

	// !!!此处注意，不要在密钥公钥中的某行前面留空格，并且应该使用``，不是""

	//	PrivKeyLocal := []byte(`
	//-----BEGIN RSA PRIVATE KEY-----
	//MIICWwIBAAKBgQCnCnuWcNacRnqwDfSNLx7bbJLJM+foyxqSzp/M0fYqjhMp8voe
	//51PUEGetCvM2kAakmRue6MXQ3TKrV7L6d3XTYGabBPzwDd0KoucklVVOS2vi1E7U
	//V1bZhB60YdayCb9dcnEdA0uyA+qQgk2VhMtP1fER8lll5EiUUT+T0vnq9wIDAQAB
	//AoGAfZo9Seb5CLNaR42GyK6Y1kdyrEYSaJJoHeGueTWbk24XbOCeQKSS/Q+E1bI5
	//JVrxE81o3nmLXT0mf35HP1yaRCrofCV7a4QBlD9CNkMfy68fJEA6gMFuVVAES6Fa
	//Zt1ENZ81NeENURUC+lLFSlUWm2Xbf+MZtCFIRE5Tj1HxvQkCQQDPTDZKpyqZ/1yg
	//PO1/Quu0iisDYROJMm4sHQowIYXkHA/pUQMEveomBGRLavWrN9t4oEotFAPi0qYW
	//847m7TmDAkEAzkkNyoz08+Dg4+SfwbjEyglyX7OkmOOGnCvEJldQm0wLZvrpJS6i
	//n24UiYx2Cg93BZrvD9Ce7oNEnwbnHG3yfQJAJtOce6ER3qQwwiaHSUXMhhU29zwQ
	//f6r9ba/Gv7sXq+EBre6phRLZL2O1MVcISph8t/w1yHmuPKa9yyC1TFV0ZwJADbeh
	//6SQybb04dy8OyI0G2QCD0IVbnqcSnnPymTIZNBp8b56jvks5mSxyxSrH9qdMnNzO
	//pNiUmPu1pnWJDMTq6QJAHIToUuuAN2z3pLpUJsM40T6sEwgbxiFPZ3iT4/T2Tgpy
	//BKLqQxR7jXKdl0iWYteC96pQ0bqytFse4lnmPMUCew==
	//-----END RSA PRIVATE KEY-----
	//`)
	//	//fmt.Printf("%s", PrivKeyLocal)
	//
	//	PubKeyRemote := []byte(`
	//-----BEGIN PUBLIC KEY-----
	//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCnCnuWcNacRnqwDfSNLx7bbJLJ
	//M+foyxqSzp/M0fYqjhMp8voe51PUEGetCvM2kAakmRue6MXQ3TKrV7L6d3XTYGab
	//BPzwDd0KoucklVVOS2vi1E7UV1bZhB60YdayCb9dcnEdA0uyA+qQgk2VhMtP1fER
	//8lll5EiUUT+T0vnq9wIDAQAB
	//-----END PUBLIC KEY-----
	//`)
	//
	//		var op encrypt.encryptOpRsa
	//		params := make([]interface{}, 0)
	//		params = append(params, PubKeyRemote)
	//		params = append(params, PrivKeyLocal)
	//		op.Init(true, params)
	//
	//		param1 := []byte("this is a rsa test")
	//		var rst1 []byte
	//		op.Operate(param1, &rst1)
	//		fmt.Printf("rst1: %x\n", rst1)
	//
	//		var test2 []byte
	//		var op2 encrypt.encryptOpRsa
	//		op2.Init(false, params)
	//		op2.Operate(rst1, &test2)
	//		fmt.Printf("rst2: %s", test2)

	////////////////md5////////////////

	//var op encrypt.encryptOpMd5
	//op.Init(true, nil)
	//
	//param1 := []byte("this is a md5 test")
	//var rst1 []byte
	//op.Operate(param1, &rst1)
	//fmt.Printf("rst1: %x\n", rst1)

	////////////////sha1////////////////

	//var op encrypt.encryptOpSha1
	//op.Init(true, nil)
	//
	//param1 := []byte("this is a sha1 test")
	//var rst1 []byte
	//op.Operate(param1, &rst1)
	//fmt.Printf("rst1: %x\n", rst1)

	////////////////gzip////////////////

	//var op compress.compressOpGzip
	//op.Init(true, nil)
	//
	//param1 := []byte("this is a gzip test")
	//var rst1 []byte
	//op.Operate(param1, &rst1)
	//fmt.Printf("rst1: %x\n", rst1)
	//
	//param2 := make([]byte, len(rst1))
	//copy(param2, rst1)
	//
	////fmt.Printf("param2: %s", param2)
	//var op2 compress.compressOpGzip
	//op2.Init(false, nil)
	//
	//var test2 []byte
	//op2.Operate(param2, &test2)
	//fmt.Printf("rst2: %s\n", test2)

	////////////////zlib////////////////

	//var op compress.compressOpZlib
	//op.Init(true, nil)
	//
	//param1 := []byte("this is a zlib test")
	//var rst1 []byte
	//op.Operate(param1, &rst1)
	//fmt.Printf("rst1: %x\n", rst1)
	//
	//param2 := make([]byte, len(rst1))
	//copy(param2, rst1)
	//
	//var op2 compress.compressOpZlib
	//op2.Init(false, nil)
	//
	//var test2 []byte
	//op2.Operate(param2, &test2)
	//fmt.Printf("rst2: %s\n", test2)

	///////////////////////////////////////// 链接测试 /////////////////////////////////////////
	///////////////////////// json/xml ///////////////////////////
	PrivKeyLocal := []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCnCnuWcNacRnqwDfSNLx7bbJLJM+foyxqSzp/M0fYqjhMp8voe
51PUEGetCvM2kAakmRue6MXQ3TKrV7L6d3XTYGabBPzwDd0KoucklVVOS2vi1E7U
V1bZhB60YdayCb9dcnEdA0uyA+qQgk2VhMtP1fER8lll5EiUUT+T0vnq9wIDAQAB
AoGAfZo9Seb5CLNaR42GyK6Y1kdyrEYSaJJoHeGueTWbk24XbOCeQKSS/Q+E1bI5
JVrxE81o3nmLXT0mf35HP1yaRCrofCV7a4QBlD9CNkMfy68fJEA6gMFuVVAES6Fa
Zt1ENZ81NeENURUC+lLFSlUWm2Xbf+MZtCFIRE5Tj1HxvQkCQQDPTDZKpyqZ/1yg
PO1/Quu0iisDYROJMm4sHQowIYXkHA/pUQMEveomBGRLavWrN9t4oEotFAPi0qYW
847m7TmDAkEAzkkNyoz08+Dg4+SfwbjEyglyX7OkmOOGnCvEJldQm0wLZvrpJS6i
n24UiYx2Cg93BZrvD9Ce7oNEnwbnHG3yfQJAJtOce6ER3qQwwiaHSUXMhhU29zwQ
f6r9ba/Gv7sXq+EBre6phRLZL2O1MVcISph8t/w1yHmuPKa9yyC1TFV0ZwJADbeh
6SQybb04dy8OyI0G2QCD0IVbnqcSnnPymTIZNBp8b56jvks5mSxyxSrH9qdMnNzO
pNiUmPu1pnWJDMTq6QJAHIToUuuAN2z3pLpUJsM40T6sEwgbxiFPZ3iT4/T2Tgpy
BKLqQxR7jXKdl0iWYteC96pQ0bqytFse4lnmPMUCew==
-----END RSA PRIVATE KEY-----
`)
	//fmt.Printf("%s", PrivKeyLocal)

	PubKeyRemote := []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCnCnuWcNacRnqwDfSNLx7bbJLJ
M+foyxqSzp/M0fYqjhMp8voe51PUEGetCvM2kAakmRue6MXQ3TKrV7L6d3XTYGab
BPzwDd0KoucklVVOS2vi1E7UV1bZhB60YdayCb9dcnEdA0uyA+qQgk2VhMtP1fER
8lll5EiUUT+T0vnq9wIDAQAB
-----END PUBLIC KEY-----
`)

	type s1 struct {
		Name string
		Age  int
		Male bool
	}

	test1 := s1{
		Name: "wang",
		Age:  1,
		Male: true,
	}

	polink := jumper_transform.Newtransform()

	polink.AddOp(def.PacketJson, nil)
	//polink.AddOp(jumper_transform.PacketXml, true, nil)

	polink.AddOp(def.CompressGzip, nil)
	//polink.AddOp(jumper_transform.CompressZlib, true, nil)

	//polink.AddOp(jumper_transform.EncryptAes, []interface{}{[]byte("abcdefghijklmnop")})
	//polink.AddOp(jumper_transform.EncryptDes, []interface{}{[]byte("ijklmnop")})
	polink.AddOp(def.EncryptRsa, []interface{}{PubKeyRemote, PrivKeyLocal})
	var rst1 []byte
	err := polink.Execute(def.Forward, test1, &rst1)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	var test2 s1
	err = polink.Execute(def.Backward, rst1, &test2)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	fmt.Printf("result: %v\n", test2)
	fmt.Printf("origin data: %v\n", test1)

	fmt.Println("--------------------------------------------------------")
	/////////////////////////// protobuf ///////////////////////////

	polink.Reset()
	bodyData := "guangzhou/fangcun/vip/company"

	p := &StringMessage{
		Body: proto.String(bodyData),
		Header: &Header{
			MessageId: proto.String("20-05"),
			Topic:     proto.String("golang"),
		},
	}

	//polink.AddOp(jumper_transform.PacketJson,  nil)
	//polink.AddOp(jumper_transform.PacketXml,  nil)
	polink.AddOp(def.PacketProtobuf, nil)
	polink.AddOp(def.CompressGzip, nil)
	polink.AddOp(def.EncryptRsa, []interface{}{PubKeyRemote, PrivKeyLocal})
	var rst11 []byte
	err = polink.Execute(def.Forward, p, &rst11)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	var protobufrst StringMessage
	err = polink.Execute(def.Backward, rst11, &protobufrst)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	fmt.Printf("result: %s\n", protobufrst.String())
	fmt.Printf("origin data: %s\n", p.String())

	fmt.Println("--------------------------------------------------------")
	msg := &interf.Message{
		Type:    1,
		Content: []byte("this is a Message."),
	}
	polink.Reset()
	polink.AddOp(def.PacketBinary, nil)
	polink.AddOp(def.CompressGzip, nil)
	polink.AddOp(def.EncryptRsa, []interface{}{PubKeyRemote, PrivKeyLocal})
	var rst12 []byte
	err = polink.Execute(def.Forward, msg, &rst12)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	var msgrst interf.Message
	err = polink.Execute(def.Backward, rst12, &msgrst)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	fmt.Printf("result: %v\n", msgrst)
	fmt.Printf("origin data: %v\n", msg)

}
