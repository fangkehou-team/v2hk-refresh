package sign

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/v2fly/VSign/sign/signify"
	"golang.org/x/crypto/sha3"
	"io"
	"io/ioutil"
)

func GenerateKeyFromSeed(seed string, password string) ([]byte, []byte) {
	shaw := sha3.NewShake256()
	shaw.Write([]byte(seed))
	pub, prv, err := signify.GenerateKey(shaw)
	if err != nil {
		panic(err)
	}

	pubb := signify.MarshalPublicKey(pub)
	prvb, err := signify.MarshalPrivateKey(prv, rand.Reader, []byte(password), 42)
	if err != nil {
		panic(err)
	}

	return prvb, pubb

}

func Sign(key []byte, password string, msg []byte) ([]byte, error) {
	pvkey, err := signify.ParsePrivateKey(key, []byte(password))
	if err != nil {
		return nil, err
	}
	out := bytes.NewBuffer(nil)
	outb := base64.NewEncoder(base64.StdEncoding, out)
	outb.Write(signify.MarshalSignature(signify.Sign(pvkey, msg)))
	outb.Close()
	return out.Bytes(), nil
}

func VerifyAndReturn(key []byte, data io.Reader) ([]byte, error) {
	signatureData, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, err
	}
	_, datav, err, readed := signify.ReadFile(bytes.NewReader(signatureData))
	if err != nil {
		return nil, err
	}

	signature, err := signify.ParseSignature(datav)
	if err != nil {
		return nil, err
	}
	publickey, err := signify.ParsePublicKey(key)
	if err != nil {
		return nil, err
	}
	ok := signify.Verify(publickey, signatureData[readed:], signature)
	if ok {
		return signatureData[readed:], nil
	}
	return nil, ErrSignatureMismatch
}

var ErrSignatureMismatch = errors.New("ErrSignatureMismatch")
