package handlers

import (
	"errors"
	"fmt"
	"os"
)

type Node struct {
	Name  string // link name
	DFlag bool   // is link directory
	Links int    // # of links referenced
	Files []Node // list of files
	Dirs  []Node // list of directories
	Size  int64  // size of link in Bytes
}

type NodeList []Node

// Len returns the number of directories in the list
func (d *NodeList) Len() int {
	return len(*d)
}

// Add adds a new directory to the DirList
func (d *NodeList) Add(dir string) error {
	if len(dir) == 0 {
		return errors.New("no directory provided")
	}

	// Check if the directory already exists
	for _, directory := range *d {
		if directory.Name == dir {
			return errors.New("directory already exists")
		}
	}

	newDir := Node{Name: dir}

	// Populate file list for the directory
	if err := newDir.GetFiles(); err != nil {
		return err
	}

	*d = append(*d, newDir)

	return nil
}

// GetFiles populates the file list for a directory
func (d *Node) GetFiles() error {
	out, err := os.ReadDir(d.Name)
	if err != nil {
		return fmt.Errorf("error reading directory %s: %w", d.Name, err)
	}

	fmt.Printf("%v:\n", d.Name)

	// Iterate over directory entries
	for _, f := range out {
		info, err := f.Info()
		if err != nil {
			return fmt.Errorf("error getting file info: %w", err)
		}

		// Create a new File object and add it to the directory's file list
		file := Node{
			Name: f.Name(),
			Size: info.Size(),
		}

		// Check if it's a directory or file
		dFlag := " "
		if f.IsDir() {
			dFlag = "D"
			d.Dirs = append(d.Dirs, file)
		}

		fmt.Printf("\t%v %v\t%v bytes\n", dFlag, f.Name(), info.Size())
	}

	return nil
}
