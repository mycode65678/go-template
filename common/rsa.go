//参考内容来源于https://github.com/polaris1119/myblog_article_code
package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	//"flag"
	//"log"
	//"os"
	"bytes"
	"errors"
)

//创建公钥与私钥
func GenRsaKey(bits int) (string, string, error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	w := bytes.NewBuffer([]byte(nil)) //bufio.NewWriter()
	err = pem.Encode(w, block)
	if err != nil {
		return "", "", err
	}
	prikey := string(w.Bytes())
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	w2 := bytes.NewBuffer(nil) //bufio.NewWriter()
	err = pem.Encode(w2, block)
	if err != nil {
		return "", "", err
	}
	pubkey := string(w2.Bytes())
	return pubkey, prikey, nil
}

func EncyptogRSA(origData []byte, publicKey []byte) (res []byte, err error) {
	block, _ := pem.Decode(publicKey)
	blocks := Pkcs1Padding(origData, 1024/8)
	// 使用X509将解码之后的数据 解析出来
	//x509.MarshalPKCS1PublicKey(block):解析之后无法用，所以采用以下方法：ParsePKIXPublicKey
	keyInit, err := x509.ParsePKIXPublicKey(block.Bytes) //对应于生成秘钥的x509.MarshalPKIXPublicKey(&publicKey)
	//keyInit, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		println("EncyptogRSA err")
		println(err)
		return
	}
	//4.使用公钥加密数据
	//pubKey := keyInit.(*rsa.PublicKey)
	buffer := bytes.Buffer{}
	for _, block := range blocks {
		ciphertextPart, err := rsa.EncryptPKCS1v15(rand.Reader, keyInit.(*rsa.PublicKey), block)
		if err != nil {
			return nil, err
		}
		buffer.Write(ciphertextPart)
	}

	return buffer.Bytes(), nil
}

// 解密
func RsaDecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func UnPadding(src []byte, keySize int) [][]byte {

	srcSize := len(src)

	blockSize := keySize

	var v [][]byte

	if srcSize == blockSize {
		v = append(v, src)
	} else {
		groups := len(src) / blockSize
		for i := 0; i < groups; i++ {
			block := src[:blockSize]

			v = append(v, block)
			src = src[blockSize:]
		}
	}
	return v
}

func Pkcs1Padding(src []byte, keySize int) [][]byte {

	srcSize := len(src)

	blockSize := keySize - 11

	var v [][]byte

	if srcSize <= blockSize {
		v = append(v, src)
	} else {
		groups := len(src) / blockSize
		for i := 0; i < groups; i++ {
			block := src[:blockSize]

			v = append(v, block)
			src = src[blockSize:]

			if len(src) < blockSize {
				v = append(v, src)
			}
		}
	}
	return v
}
