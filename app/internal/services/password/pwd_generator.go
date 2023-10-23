package password

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func GeneratePasswordHash(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))

	return fmt.Sprintf("%x", hash.Sum([]byte("")))
}

func GenerateSalt() string {
	var b = make([]byte, 32)

	s := rand.NewSource(time.Now().Unix() * rand.Int63())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		log.Fatalf("Salt generation: %s", err.Error())
	}

	return fmt.Sprintf("%x", b)
}
