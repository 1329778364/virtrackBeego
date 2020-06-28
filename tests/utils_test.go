package test

import (
	"fmt"
	"gobeetestpro/utils"
	"testing"
)

func TestGetUUID(t *testing.T) {
	data, _ := utils.GetUUID()
	fmt.Printf(data)
}
