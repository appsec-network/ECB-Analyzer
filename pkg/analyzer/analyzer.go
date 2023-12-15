package analyzer

import (
	"ebc_analyzer/util/encryption"
	"fmt"
	"strings"
)

type Result struct {
	IsSuccess bool
}

func ProcessCracking(unecryptedData string, encryptedData string) (*Result, error) {

	if strings.TrimSpace(unecryptedData) == "" || strings.TrimSpace(encryptedData) == "" {
		return &Result{}, fmt.Errorf("Invalid encypted/unecrypted data.")
	}

	ecb := encryption.ECB(unecryptedData, encryptedData, true)
	_, err := ecb.CrackECB()

	if err != nil {
		fmt.Println(err)
		return &Result{
			IsSuccess: false,
		}, err
	}

	return &Result{}, err
}
