package def

import (
	"github.com/jumper86/jumper_error"
)

const (
	ErrMd5NoDecryptCode    = 11001
	ErrSha1NoDecryptCode   = 11002
	ErrSha256NoDecryptCode = 11003

	ErrInvalidAesKeyCode = 11004
	ErrInvalidDesKeyCode = 11005

	ErrInvalidRsaPrivateKeyCode = 11006
	ErrInvalidRsaPublicKeyCode  = 11007

	ErrParamShouldNotNilCode           = 11008
	ErrParamShouldStringOrBytesCode    = 11009
	ErrParamShouldPointOfByteSliceCode = 11010
	ErrParamShouldImplInterfMsgCode    = 11011
	ErrParamShouldImplProtoMsgCode     = 11012
	ErrDecodeFailedCode                = 11013
)

var (
	ErrMd5NoDecrypt                = jumper_error.New(ErrMd5NoDecryptCode, "md5 no decrypt.")
	ErrSha1NoDecrypt               = jumper_error.New(ErrSha1NoDecryptCode, "sha1 no decrypt.")
	ErrSha256NoDecrypt             = jumper_error.New(ErrSha256NoDecryptCode, "sha256 no decrypt.")
	ErrInvalidAesKey               = jumper_error.New(ErrInvalidAesKeyCode, "invalid aes key.")
	ErrInvalidDesKey               = jumper_error.New(ErrInvalidDesKeyCode, "invalid des key.")
	ErrInvalidRsaPrivateKey        = jumper_error.New(ErrInvalidRsaPrivateKeyCode, "invalid rsa private key.")
	ErrInvalidRsaPublicKey         = jumper_error.New(ErrInvalidRsaPublicKeyCode, "invalid rsa public key.")
	ErrParamShouldNotNil           = jumper_error.New(ErrParamShouldNotNilCode, "param should not be nil.")
	ErrParamShouldStringOrBytes    = jumper_error.New(ErrParamShouldStringOrBytesCode, "param should be string or []byte.")
	ErrParamShouldPointOfByteSlice = jumper_error.New(ErrParamShouldPointOfByteSliceCode, "param should be *[]byte.")
	ErrParamShouldImplInterfMsg    = jumper_error.New(ErrParamShouldImplInterfMsgCode, "param should be pointer of interf.Message.")
	ErrParamShouldImplProtoMsg     = jumper_error.New(ErrParamShouldImplProtoMsgCode, "param should implement proto.Message.")
	ErrDecodeFailed                = jumper_error.New(ErrDecodeFailedCode, "decode failed.")
)
