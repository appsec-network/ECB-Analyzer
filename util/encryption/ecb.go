package encryption

import (
	"ebc_analyzer/util/convert"
	"ebc_analyzer/util/ui/message"
	"errors"
	"fmt"
	"os"
	"strconv"

	"encoding/base64"
	"net/url"
	"reflect"
)

type ECBType struct {
	UrlDecoded      string
	B64Decoded      []byte
	DecValues       []int
	ByteValues      []byte
	UnencryptedData string
	EncryptedData   string
	ECBCrack        bool
	ECBRepeatBlock  []byte
	EncryptionKey   []int
	ecbKey          []string
	BlockSize       string
}

func ECB(unencryptedData string, encryptedData string, blockSize string, ecbCrack bool) *ECBType {
	urlDecoded, _ := url.QueryUnescape(encryptedData)
	b64Decoded, _ := base64.RawStdEncoding.DecodeString(urlDecoded)
	decValues := convert.B64ToDec(b64Decoded)
	byteValues := convert.B64ToBytes(b64Decoded)

	return &ECBType{
		UrlDecoded:      urlDecoded,
		B64Decoded:      b64Decoded,
		DecValues:       decValues,
		ByteValues:      byteValues,
		UnencryptedData: unencryptedData,
		EncryptedData:   encryptedData,
		ECBCrack:        ecbCrack,
		BlockSize:       blockSize,
	}
}

func (c *ECBType) CrackECB() (*ECBType, error) {

	message.Println(message.Info, 0, "Finding repeated blocks in ECB encrypted string\n")

	blockSize := 0
	repeatBlock := []byte{}

	if c.BlockSize == "auto" {

		for idx, dec := range c.ByteValues {
			if idx == 0 {
				repeatBlock = append(repeatBlock, byte(dec))
			} else {
				if byte(dec) == repeatBlock[0] {
					checkRepeat := c.ByteValues[idx : idx+len(repeatBlock)]
					if reflect.DeepEqual(repeatBlock, checkRepeat) {
						message.Println(message.Success, 1, "Found repeat block: %v\n", repeatBlock)
						c.ECBRepeatBlock = repeatBlock
						c.CalculateECBKey()
						break
					}
				} else {
					repeatBlock = append(repeatBlock, dec)
				}
			}
		}

	} else {

		var err error
		blockSize, err = strconv.Atoi(string(c.BlockSize))
		if err != nil || blockSize <= 0 {
			return nil, errors.New("Invalid block size")
		}

		for idx, dec := range c.ByteValues {
			if idx%blockSize == 0 {
				repeatBlock = append(repeatBlock[:0], byte(dec))
			} else {
				if byte(dec) == repeatBlock[idx%blockSize] {
					checkRepeat := c.ByteValues[idx-blockSize : idx]
					if reflect.DeepEqual(repeatBlock, checkRepeat) {
						message.Println(message.Success, 1, "Found repeat block: %v\n", repeatBlock)
						c.ECBRepeatBlock = repeatBlock
						c.CalculateECBKey()
						break
					}
				} else {
					repeatBlock = append(repeatBlock, dec)
				}
			}
		}

	}

	if len(c.ECBRepeatBlock) != 0 {
		return nil, nil
	}

	return nil, fmt.Errorf("No repeated blocks found")
}

func (c *ECBType) CalculateECBKey() ([]string, error) {
	message.Println(message.Info, 0, "Calculating ECB key (%d-byte block key)\n", len(c.ECBRepeatBlock))
	ecbKey := make([]int, 0)

	var TmpUnencryptedData []rune
	TmpUnencryptedData = convert.StringToCharArray(c.UnencryptedData)

	if len(c.ECBRepeatBlock) > len(TmpUnencryptedData) {
		message.Println(message.Error, 1, "Length mismatch between ecbRepeatBlock and unencryptedString")
		os.Exit(0)
	}

	for idx, idecValue := range c.ECBRepeatBlock {
		unencryptedByte := int(TmpUnencryptedData[idx])
		message.Println(message.Notify, 1, "XOR : %4d with %4d for key-byte %2d\n", idecValue, unencryptedByte, idx)
		keyVal := unencryptedByte ^ int(idecValue)
		ecbKey = append(ecbKey, keyVal)
	}

	message.Println(message.Success, 1, "Found potential encryption key: %v\n", ecbKey)
	encryptedValues := c.DecValues
	unencryptedValues := convert.Encrypt(encryptedValues, ecbKey)
	reEncryptedValues := convert.Encrypt(unencryptedValues, ecbKey)
	if convert.SlicesEqual(encryptedValues, reEncryptedValues) {
		base64EncodedKey := convert.ToBase64(ecbKey)
		rawEncryptionKey, _ := convert.FromBase64(base64EncodedKey)
		message.Println(message.Success, 2, "Dec encryption key: %v\n", ecbKey)
		message.Println(message.Success, 2, "Raw encryption key: %v\n", rawEncryptionKey)
		message.Println(message.Success, 2, "Base64 encoded key: %v\n", base64EncodedKey)
	} else {
		message.Println(message.Error, 1, "ERROR. Oven broken")
	}

	return nil, nil
}

/*
func (c *ECBType) Test() ([]string, error) {
	dKey := "testtestesttest1"
	dData := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaadmin"

	// Key'i byte dizisine dönüştürme
	key := []byte(dKey)

	// Data'yı byte dizisine dönüştürme
	data := []byte(dData)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)
	padder := padding.NewPkcs7Padding(mode.BlockSize())
	data, err = padder.Pad(data) // pad last block of plaintext if block size less than block cipher size
	if err != nil {
		panic(err.Error())
	}
	ct := make([]byte, len(data))
	mode.CryptBlocks(ct, data)

	// Şifrelenmiş veriyi base64'e dönüştürme
	encrypted := base64.StdEncoding.EncodeToString(ct)
	fmt.Println("Encrypted:", encrypted)

	// URL encode işlemi
	urlEncoded := url.QueryEscape(encrypted)
	fmt.Println("URL Encoded:", dData+":"+urlEncoded)

	fmt.Println()
	fmt.Println()

	return nil, nil
}*/
