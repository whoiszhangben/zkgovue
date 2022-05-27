package base

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math"
)

func EncodeDsn(dsn string) (string, error) {
	e, err := TripleDesEncrypt([]byte(dsn), getKey())
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(e), nil
}

func DecodeDsn(dsn string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(dsn)
	if err != nil {
		return "", err
	}

	e, err := TripleDesDecrypt([]byte(data), getKey())
	if err != nil {
		return "", err
	}

	return string(e), nil
}

func getKey() []byte {
	f := math.Log2(10.0)
	s := fmt.Sprintf("%.10f", f)
	k := sha256.Sum256([]byte(s))
	return k[:24]
}
