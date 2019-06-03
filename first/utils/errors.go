package utils

import "fmt"

func PanicIfErr(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%v : %v", msg, err))
	}
}
