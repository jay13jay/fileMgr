package handlers

import (
	"fmt"
	"os"
)

func GetFiles(dir string) {
	out, err := os.ReadDir(dir)
	if err != nil {
		HandleError(err, 1)
	}
	fmt.Printf("%v:\n", dir)
	for _,f := range out {
		d,_ := f.Info()
		dFlag := " "
		if d.IsDir() {
			dFlag = "D"
		}
		fmt.Printf("\t%v %v\t%v\n", dFlag, d.Name(), d.Size())
	}
}