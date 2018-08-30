package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
)

func formatPage(path string, page int, sorted bool) (s string) {
	var b bytes.Buffer //to hold page contents
	
	//read file info from directory, while handling sorted option
	var files []os.FileInfo
	var err error
	if sorted {
		files, err = ioutil.ReadDir(path)
	} else {
		f, err := os.Open(path)
		if err != nil {
			b.WriteString("Directory not found")
			return b.String()
		}
		files, err = f.Readdir(-1)
		f.Close()
	}
	if err != nil {
		b.WriteString("Directory not found")
		return b.String()
	}

	//page out the directory 
	var pages [][]os.FileInfo
	pageSize := 20
	for i := 0; i < len(files); i += pageSize {
		end := i + pageSize
		if end > len(files) {
			end = len(files)
		}
		pages = append(pages, files[i:end])
	}
	numPages := len(pages)

	//make sure requested page number is valid
	if page < 1 || page > numPages {
		b.WriteString("Invalid page number")
		return b.String()
	}
	
	//write page contents to buffer
	b.WriteString(path[2:] + "\n")
	for _, file := range pages[page-1] {
		b.WriteString("\n")
		b.WriteString(file.Name())
		if file.IsDir(){
			b.WriteString(" (Directory)")
		}
	}
	b.WriteString(fmt.Sprintf("\n\nPage %d of %d\n", page, numPages))

	return b.String() //return buffer as string
}