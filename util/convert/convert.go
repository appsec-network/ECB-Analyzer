package convert

import (
	"encoding/base64"
	"reflect"
)

func B64ToDec(b64Value []byte) []int {
	decArray := make([]int, len(b64Value))
	for i, v := range b64Value {
		decArray[i] = int(v)
	}
	return decArray
}

func B64ToBytes(b64Value []byte) []byte {
	decArray := make([]byte, len(b64Value))
	copy(decArray, b64Value)
	return decArray
}

func Encrypt(value, key []int) []int {
	encrypted := make([]int, len(value))
	for i := 0; i < len(value); i++ {
		encrypted[i] = value[i] ^ key[i%len(key)]
	}
	return encrypted
}

func SafeFormatHexToASCIIString(hexDecValues []int) string {
	safeString := ""
	for _, char := range hexDecValues {
		if char >= 32 && char <= 255 {
			safeString += string(rune(char))
		} else {
			safeString += " "
		}
	}
	return safeString
}

func AddEncryptionKeyFromB64(b64key string) []int {
	decodedKey, _ := base64.RawStdEncoding.DecodeString(b64key)
	encryptionKey := make([]int, len(decodedKey))
	for i := 0; i < len(decodedKey); i++ {
		encryptionKey[i] = int(decodedKey[i])
	}
	return encryptionKey
}

func SlicesEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func ToBase64(ecbKey []int) string {
	byteSlice := make([]byte, len(ecbKey))

	for i := range ecbKey {
		byteSlice[i] = byte(ecbKey[i])
	}

	base64EncodedKey := base64.RawStdEncoding.EncodeToString(byteSlice)
	return base64EncodedKey
}

func FromBase64(base64EncodedKey string) (string, error) {
	// Base64'den çözme
	decodedBytes, err := base64.RawStdEncoding.DecodeString(base64EncodedKey)
	if err != nil {
		return "", err
	}

	// Byte dizisini ASCII formatına dönüştürme
	rawASCII := string(decodedBytes)
	return rawASCII, nil
}

func StringToCharArray(str string) []rune {
	var charArray []rune
	for _, char := range str {
		charArray = append(charArray, char)
	}
	return charArray
}

func IntArrayToString(intArr []int) string {
	var stringEcbKey string
	for _, v := range intArr {
		stringEcbKey += string(rune(v))
	}
	return stringEcbKey
}

func ConvertFromBase64(base64EncodedKey string) ([]int, error) {
	byteSlice, err := base64.RawStdEncoding.DecodeString(base64EncodedKey)
	if err != nil {
		return nil, err
	}

	intSlice := make([]int, len(byteSlice))
	for i := range byteSlice {
		intSlice[i] = int(byteSlice[i])
	}

	return intSlice, nil
}

func HasRepeatBlock(values []int, blockSize int) bool {
	for i := 0; i < len(values)-blockSize; i += blockSize {
		for j := i + blockSize; j < len(values)-blockSize; j += blockSize {
			if reflect.DeepEqual(values[i:i+blockSize], values[j:j+blockSize]) {
				return true
			}
		}
	}
	return false
}
