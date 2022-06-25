package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/spf13/viper"
	"io"
)

//func testAes() {
//	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
//	result, err := AesEncrypt([]byte("polaris@studygolang"))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(base64.StdEncoding.EncodeToString(result))
//	origData, err := AesDecrypt(result)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(origData))
//}

func AesDecrypt(text1 string) []byte {
	k := viper.GetString("base.session_key")
	key := []byte(k)
	ciphertext, _ := base64.StdEncoding.DecodeString(string(text1))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = PKCS5UnPadding(ciphertext)
	return ciphertext
}

func AesEncrypt(plaintext []byte) string {
	k := viper.GetString("base.session_key")
	key := []byte(k)
	plaintext = PKCS5Padding(plaintext, 16)
	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
