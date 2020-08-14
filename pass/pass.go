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

func hashString(str string) string {
	bs := sha1.Sum([]byte(str))
	return fmt.Sprintf("%x", bs)
}

func hashWithSalts(str string) []string {
	var s []string
	for _, salt := range readInSalts() {
		s = append(s, hashString(str+salt))
		s = append(s, hashString(salt+str))
	}
	return s
}

func CrackSha1Hash(str string, useSalt bool) string {
	for _, pass := range readInPasswords() {
		if useSalt {
			for _, salted := range hashWithSalts(pass) {
				if salted == str {
					return pass
				}
			}
		} else if hashString(pass) == str {
			return pass
		}
	}

	return "PASSWORD NOT IN DATABASE"
}
