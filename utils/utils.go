package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

/*-----------获取UUID--------------------*/
func GetUUID() (string, error) {
	u2, err := uuid.NewV1()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return "", err
	}
	return u2.String(), nil
}
