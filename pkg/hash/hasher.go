package hash

import "golang.org/x/crypto/bcrypt"

type Hasher struct{}

func (Hasher) Hash(str string) ([]byte, error) {
	cost := 5
	return bcrypt.GenerateFromPassword([]byte(str), cost)
}

func (Hasher) VerifuHash(hashedStr []byte, str string) bool {
	if err := bcrypt.CompareHashAndPassword(hashedStr, []byte(str)); err != nil {
		return false
	}

	return true
}
