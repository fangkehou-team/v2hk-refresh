package signerVerify

import (
	"bytes"
	"errors"
	"github.com/v2fly/VSign/insmgr"
	"github.com/v2fly/VSign/sign"
	"io"
)

func CheckSignature(publicKey []byte, signatureFile io.Reader, hash map[string]string, project string) (string, error) {

	data, err := sign.VerifyAndReturn(publicKey, signatureFile)
	if err != nil {
		return "", err
	}
	v := insmgr.ReadAllIns(bytes.NewReader(data))
	if v == nil {
		return "", io.ErrUnexpectedEOF
	}
	hashFromManifest, version, ok := CheckAsClient(v, project, true)
	if !ok {
		return "", io.ErrUnexpectedEOF
	}

	var tempered bool

	for Filename, hashFromFile := range hash {
		mftName, found := hashFromManifest[hashFromFile]
		if !found {
			tempered = true
			hash[Filename] = "TAINTED"
		} else {
			hash[Filename] = mftName
		}
	}
	if tempered {
		return version, ErrNotFound
	}
	return version, nil
}

var ErrNotFound = errors.New("tainted file")
