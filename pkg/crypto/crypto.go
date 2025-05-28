package crypto

import "golang.org/x/crypto/bcrypt"

func BCrypto(val string) string {
	// TODO: skip the error for speed
	res, _ := bcrypt.GenerateFromPassword([]byte(val), bcrypt.DefaultCost)

	// TODO: alloc
	return string(res)
}

func CheckPasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
