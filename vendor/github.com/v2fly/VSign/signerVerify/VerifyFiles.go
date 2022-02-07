package signerVerify

import (
	"fmt"
	"github.com/v2fly/VSign/insmgr"
	"github.com/v2fly/VSign/instimp"
	"github.com/v2fly/VSign/v2signkey"
	"io"
)

func CheckSignaturesFromFile(publicKey []byte, signatureFile io.Reader, project string, files []string) (string, map[string]string, error) {

	//Compute the hash first
	hashcoll := insmgr.NewHashCollectorMgr(false)
	for _, f := range files {
		instimp.NewFileBasedInsYield(f).InstructionYield(hashcoll)
	}
	res := hashcoll.Result()
	version, err := CheckSignature(publicKey, signatureFile, res, project)
	if err != nil {
		return version, res, err
	}
	return version, res, nil
}

func CheckSignaturesV2Fly(signatureFile io.Reader, files []string) (string, map[string]string, error) {
	key := v2signkey.GetSignKeyAsByte()
	return CheckSignaturesFromFile(key, signatureFile, "v2fly", files)
}

func OutputAndJudge(version string, result map[string]string, errResult error) error {
	if errResult != nil {
		if errResult == ErrNotFound {
			fmt.Printf("-TAINTED: file(s) have been tainted.\n")
		} else {
			fmt.Printf("-TAINTED: verification did not complete: %v\n", errResult)
			return errResult
		}
	} else {
		fmt.Printf("+OK: file(s) being checked is/are unblemished\n")
	}
	fmt.Printf("version: %v\n", version)
	for i, v := range result {
		fmt.Printf("%v : %v\n", i, v)
	}
	return errResult
}
