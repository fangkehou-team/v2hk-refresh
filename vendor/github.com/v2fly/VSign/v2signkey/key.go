package v2signkey

import (
	"bytes"
	"github.com/v2fly/VSign/sign/signify"
)

const V2FlySigningKey = `untrusted comment: V2Fly Signing Key
RWTe6SReSmJUeqoA8cq1MxX7ycL06DAMAJcAgQ8dCN3kFtnWBHYDpTnx
`

func GetSignKeyAsByte() []byte {
	_, data, err, _ := signify.ReadFile(bytes.NewReader([]byte(V2FlySigningKey)))
	if err != nil {
		panic(err)
	}
	return data
}
