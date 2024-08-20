package tool

import (
	"testing"
)

func TestEncryptWithBcrypt(t *testing.T) {
	password := "mysecretpassword"
	hashedPassword, err := EncryptWithBcrypt(password)
	if err != nil {
		t.Fatalf("Error hashing password: %v", err)
	}

	if hashedPassword == "" {
		t.Fatalf("Hashed password is empty")
	}
}

func TestCheckStrHash(t *testing.T) {
	password := "mysecretpassword"
	hashedPassword, err := EncryptWithBcrypt(password)
	if err != nil {
		t.Fatalf("Error hashing password: %v", err)
	}

	match := CheckStrHash(password, hashedPassword)
	if !match {
		t.Fatalf("Password does not match")
	}

	wrongPassword := "wrongpassword"
	match = CheckStrHash(wrongPassword, hashedPassword)
	if match {
		t.Fatalf("Wrong password should not match")
	}
}
