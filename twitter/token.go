package twitter

import (
	"os"
)

type Tokens struct {
	cs string
	ck string
	At string
	As string
}

func Token() *Tokens {
	t := Tokens{}
	t.cs = os.Getenv("CS")
	t.ck = os.Getenv("CK")
	return &t
}
