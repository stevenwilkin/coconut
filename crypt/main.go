package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"strings"
)

var (
	ssoPubKey  *rsa.PublicKey
	ssoPrivKey *rsa.PrivateKey
	label      = []byte("")
)

const ssoPubKey64 = `
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0247/7u+eNDLeVbb4Q9tJWPXl
w424Cjn7rtMGm47exDjVOBbIzTnXMqgk+DC4xlkstYo98tCcjcaztrnjmtlCNXok
5pgp57ikwYRWLiyJ1QW/WJUcuMDweIJUEhszny0a6mAgfEDM2PH852xCQ5hYH4uu
pB+2f2ayWNNJ/rz6ZQIDAQAB
`

const ssoPrivKey64 = `
MIICWwIBAAKBgQC0247/7u+eNDLeVbb4Q9tJWPXlw424Cjn7rtMGm47exDjVOBbI
zTnXMqgk+DC4xlkstYo98tCcjcaztrnjmtlCNXok5pgp57ikwYRWLiyJ1QW/WJUc
uMDweIJUEhszny0a6mAgfEDM2PH852xCQ5hYH4uupB+2f2ayWNNJ/rz6ZQIDAQAB
AoGAMnnw0EdsgxgIdnsvxNyTcEYu4wCQJiRciH0Dkp2J42yafA/adBKrKP+PZDoM
xxU8wDiuq5mAVaFQKso92WNXbSWEiC1w/UwwJqSgP2fJkQFl1tJ5I03KxFHnYi4w
kJlElWNBdVSNJ701u7vB+jjhleZBbIwxQYHRAhK+Q3x/1wUCQQDEDJrEv7+Mzh3c
NS7QrhAdmrFetPXf1OXQOgIxS+7QtiVG+jZQ5A6v/9THA0awcY69hyCvVDY4572S
VpeuXcm3AkEA7Cm0mKeZ6TFVt5Jq+qtD8cwny1NbBnbBA+d3dpkQvm/P+jeYCn4z
/J03Ub/3xRd9/SiOC6rEzG4UHaKbveZMwwJAXhq2x65K4emmR6d3m0+SCMPSU+WF
CDYHQhY2KzeJoMFtz04XeGif7DdfCVA3REad/7e3JrHDfTkvs+jA0j/OrQJAT5y7
md6ePwN3nvvH/prvu7qUC7Ic9G/iH3vPRPbSszAkT3igU6E5y0YAmRl64EFMIqSi
RgKad0QAmgDwObNCWwJAB5JX9bBccxv2btVHCzjHGzd88Xcy9IDCk8V2zQ82APOs
i8nyQsW22Jn0yw11HgHgOgDFq5TPZKcvUunnyT6iRA==
`

func decodeBase64Key(key string) []byte {
	b, err := base64.StdEncoding.DecodeString(strings.TrimSpace(key))
	if err != nil {
		panic(err)
	}

	return b
}

func init() {
	key, err := x509.ParsePKIXPublicKey(decodeBase64Key(ssoPubKey64))
	if err != nil {
		panic(err)
	}
	ssoPubKey = key.(*rsa.PublicKey)

	ssoPrivKey, err = x509.ParsePKCS1PrivateKey(decodeBase64Key(ssoPrivKey64))
	if err != nil {
		panic(err)
	}
}

func Encrypt(plaintext []byte) []byte {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, ssoPubKey, plaintext, label)
	if err != nil {
		panic(err)
	}

	return ciphertext
}

func Decrypt(ciphertext []byte) []byte {
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, ssoPrivKey, ciphertext, label)
	if err != nil {
		panic(err)
	}

	return plaintext
}
