package folder

import (
	"fmt"
	"log"
)

func CheckError(err error) {
	fmt.Println("CheckError")
	if err != nil {
		log.Fatal("herror here", err)
	}
}
