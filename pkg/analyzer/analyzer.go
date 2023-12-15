package analyzer

import (
	"ebc_analyzer/util/encryption"
	"fmt"
	"strings"
)

type Result struct {
	IsSuccess bool
}

func ProcessCracking(unecryptedData string, encryptedData string, blockSize string) (*Result, error) {

	if strings.TrimSpace(unecryptedData) == "" || strings.TrimSpace(encryptedData) == "" {
		return &Result{}, fmt.Errorf("Invalid encypted/unecrypted data.")
	}

	ecb := encryption.ECB(unecryptedData, encryptedData, blockSize, true)
	_, err := ecb.CrackECB()

	if err != nil {
		return &Result{
			IsSuccess: false,
		}, err
	}

	return &Result{}, err
}
