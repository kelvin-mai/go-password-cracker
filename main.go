package main

import (
	"fmt"

	"github.com/kelvin-mai/go-password-cracker/pass"
)

func printSlice(s []string) {
	for _, str := range s {
		fmt.Println(str)
	}
}

func main() {
	str := pass.CrackSha1Hash("5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8")
	fmt.Println(str) // password

	str = pass.CrackSha1Hash("18c28604dd31094a8d69dae60f1bcd347f1afc5a")
	fmt.Println(str) // superman
}
