package humanid

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func Generate() (string, error) {
	num, err := rand.Int(rand.Reader, big.NewInt(999_999))
	if err != nil {
		return "", err
	}

	adjective, err := Adjectives.GetRandomWord()
	if err != nil {
		return "", err
	}

	noun, err := Nouns.GetRandomWord()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s-%d", adjective, noun, num.Int64()), nil
}
