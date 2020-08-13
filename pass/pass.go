package pass

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"strings"
)

func readInPasswords() []string {
	b, err := ioutil.ReadFile("top-10000-passwords.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)
	return strings.Split(str, "\n")
}

func readInSalts() []string {
	b, err := ioutil.ReadFile("known-salts.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)
	return strings.Split(str, "\n")
}

func HashString(str string) string {
	bs := sha1.Sum([]byte(str))
	return fmt.Sprintf("%x", bs)
}

func CrackSha1Hash(str string) string {
	passwords := readInPasswords()
	for _, pass := range passwords {
		if HashString(pass) == str {
			return pass
		}
	}

	return "PASSWORD NOT IN DATABASE"
}
