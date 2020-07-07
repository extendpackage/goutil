package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/extendpackage/goutil/encryption"
)

func main() {
	secret := "test_secret" //密钥
	m := map[string]string{
		"test1": "test1",
		"test2": "test2",
	}
	bytes, err := json.Marshal(m) //转为json
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	message := "test_message" + string(bytes)

	fmt.Println(encryption.HmacSha256_Base64(message, secret))
}
