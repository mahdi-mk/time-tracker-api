package hash

import "golang.org/x/crypto/bcrypt"

// Generate generates a hashed password from the given plaintext
func Generate(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		panic("Something went wrong while generating a hashed password.")
	}

	return string(hashed)
}

// Verify compares a bcrypt hashed password with its possible plaintext equivalent
func Verify(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
