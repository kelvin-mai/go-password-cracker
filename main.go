package main

import (
	"fmt"

	"github.com/kelvin-mai/go-password-cracker/pass"
)

func main() {
	str := pass.CrackSha1Hash("18c28604dd31094a8d69dae60f1bcd347f1afc5a", false)
	fmt.Println(str) // superman

	str = pass.CrackSha1Hash("53d8b3dc9d39f0184144674e310185e41a87ffd5", true)
	fmt.Println(str) // superman
}
