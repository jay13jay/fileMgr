package handlers

import "fmt"

func HandleError(err error, level int) {
	if err != nil {
		switch level {
			case 0:
				fmt.Println("INFO: ", err)
			case 1:
				fmt.Println("ERROR: ", err)
			case 9:
				panic(err)
		default:
			fmt.Println("Unknown error: ", err)
		}
	}
}

