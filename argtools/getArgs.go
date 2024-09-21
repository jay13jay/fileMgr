package argtools

import (
	"fmt"
	"os"

	"github.com/jay13jay/fileMgr/handlers"
)

type Handler struct {
	Args []string
	Dirs handlers.NodeList
}

func ReadArgs() (handler Handler, err error) {
	handler = Handler{}
	handler.Args = os.Args[1:] // Skips the program name
	return handler, err
}

func (h *Handler) help() {
	fmt.Println("Help - Not Implemented yet")
	fmt.Println("Usage: program [options] arguments")
}

func (h *Handler) version() {
	fmt.Println("Version - Not Implemented yet")
}

func (h *Handler) ProcessArgs() {
	// Loop through all the arguments
	for i := 0; i < len(h.Args); i++ {
		flag := h.Args[i]

		// Check if the argument is a flag (starts with "-")
		if len(flag) > 1 && flag[0] == '-' {
			var data string

			// Check if the next argument exists and is not another flag
			if i+1 < len(h.Args) && len(h.Args[i+1]) > 0 && h.Args[i+1][0] != '-' {
				data = h.Args[i+1]
				i++ // Skip the next argument since it's used as data
			}

			// Handle the flags and their associated data
			switch flag {
			case "-h":
				h.help()
			case "-v":
				h.version()
			case "-d":
				err := h.Dirs.Add(data)
				if err != nil {
					handlers.HandleError(err, 1)
				}
			default:
				handlers.HandleError(fmt.Errorf("unknown argument provided: %s", flag), 1)
			}
		} else {
			// Handle cases where arguments are not flags
			fmt.Printf("Non-flag argument: %s\n", flag)
		}
	}
}
