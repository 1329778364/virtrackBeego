package test

import (
	"fmt"
	"gobeetestpro/controllers"
	_ "gobeetestpro/utils"
	"testing"
	_ "testing"
)

func TestMD5V(t *testing.T) {

}

func TestCrypto(t *testing.T) {
	hash := controllers.Crypto("wlqlwl")
	fmt.Printf("hash:%s", hash)
}

func TestValidatePassword(t *testing.T) {
	res := controllers.ValidatePassword("$2a$10$lYhZhKfRnAwGBhfTgaktB.7BqjUTubr4Swrp50vNS7ld94a6KOoBy", "wlqlwl")
	fmt.Print(res)
}
