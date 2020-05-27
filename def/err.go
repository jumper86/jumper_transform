package def

import (
	"github.com/jumper86/jumper_error"
)

const (
	ErrMd5NoDecryptCode  = 11001
	ErrSha1NoDecryptCode = 11002

	ErrInvalidAesKeyCode = 11003
	ErrInvalidDesKeyCode = 11004

	ErrInvalidRsaPrivateKeyCode = 11005
	ErrInvalidRsaPublicKeyCode  = 11006

	ErrParamShouldNotNilCode           = 11007
	ErrParamShouldStringOrBytesCode    = 11008
	ErrParamShouldPointOfByteSliceCode = 11009
	ErrParamShouldImplInterfMsgCode    = 11010
	ErrParamShouldImplProtoMsgCode     = 11011
	ErrDecodeFailedCode                = 11012
)

var (
	ErrMd5NoDecrypt                = jumper_error.New(ErrMd5NoDecryptCode, "md5 no decrypt.")
	ErrSha1NoDecrypt               = jumper_error.New(ErrSha1NoDecryptCode, "sha1 no decrypt.")
	ErrInvalidAesKey               = jumper_error.New(ErrInvalidAesKeyCode, "invalid aes key.")
	ErrInvalidDesKey               = jumper_error.New(ErrInvalidDesKeyCode, "invalid des key.")
	ErrInvalidRsaPrivateKey        = jumper_error.New(ErrInvalidRsaPrivateKeyCode, "invalid rsa private key.")
	ErrInvalidRsaPublicKey         = jumper_error.New(ErrInvalidRsaPublicKeyCode, "invalid rsa public key.")
	ErrParamShouldNotNil           = jumper_error.New(ErrParamShouldNotNilCode, "param should not be nil.")
	ErrParamShouldStringOrBytes    = jumper_error.New(ErrParamShouldStringOrBytesCode, "param should be string or []byte.")
	ErrParamShouldPointOfByteSlice = jumper_error.New(ErrParamShouldPointOfByteSliceCode, "param should be *[]byte.")
	ErrParamShouldImplInterfMsg    = jumper_error.New(ErrParamShouldImplInterfMsgCode, "param should implement interf.Message.")
	ErrParamShouldImplProtoMsg     = jumper_error.New(ErrParamShouldImplProtoMsgCode, "param should implement proto.Message.")
	ErrDecodeFailed                = jumper_error.New(ErrDecodeFailedCode, "decode failed.")
)
