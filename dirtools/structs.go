package dirtools

import (
	"errors"

	"github.com/jay13jay/fileMgr/handlers"
)



type File struct {
	FileName string
	Size int64
}

type Dir struct {
	Name string
	// Files []File
	Files []string
	Size int64
}

type DirList []Dir

func (d *DirList) Len() int {
	return len(*d)
}

func (d *DirList) Add(dir string) error {
	if len(dir) == 0 {
		return errors.New("no directory provided")
	}
	// Check if directory already exists
	for _, directory := range *d {
		if directory.Name == dir {
			return errors.New("directory already exists")
		}
	}
	newDir := Dir{Name: dir}
	*d = append(*d, newDir)

	// populate file list
	newDir.GetFiles()

	return nil
}

func (d *Dir) GetFiles() {
	handlers.GetFiles(d.Name)
}