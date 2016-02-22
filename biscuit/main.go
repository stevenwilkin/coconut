package biscuit

import (
	"encoding/base64"
	"crypto/rand"
	"net/url"

	"gopkg.in/macaroon.v1"

	"github.com/stevenwilkin/coconut/crypt"
)

func MustNew(key []byte, identity, location string) *macaroon.Macaroon {
	m, err := macaroon.New(key, identity, location)
	if err != nil {
		panic(err)
	}

	return m
}

func MustDecodeFromString(encoded string) *macaroon.Macaroon {
	b, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}

	var m macaroon.Macaroon
	err = m.UnmarshalBinary(b)
	if err != nil {
		panic(err)
	}

	return &m
}

func MustEncodeToString(m *macaroon.Macaroon) string {
	b, err := m.MarshalBinary()
	if err != nil {
		panic(err)
	}

	s := base64.StdEncoding.EncodeToString(b)

	return url.QueryEscape(s)
}

func MustAddFirstPartyCaveat(m *macaroon.Macaroon, caveat string) {
	err := m.AddFirstPartyCaveat(caveat)
	if err != nil {
		panic(err)
	}
}

func MustAddThirdPartyCaveat(m *macaroon.Macaroon, location string) {
	randomKey := make([]byte, 32)
	_, err := rand.Read(randomKey)
	if err != nil {
		panic(err)
	}

	encryptedRandomKey := string(crypt.Encrypt(randomKey)) // uses SSO pub key

	err = m.AddThirdPartyCaveat(randomKey, encryptedRandomKey, location)
	if err != nil {
		panic(err)
	}
}
