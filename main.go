package main

import (
	"fmt"
	//utils "github.com/dkmos2016/go-sdk-utils"
	"golang.org/x/crypto/md4"
	_ "gopkg.in/yaml.v3"
)

func main() {
	fmt.Println("Hello World!")

	//utils.Test()

	//crypto
	md4.New()
}
