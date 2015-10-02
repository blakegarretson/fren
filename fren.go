/*
fren (File REName) renames any number of files by replacing a given string with another string.

Usage:

    fren <original> <new> <files>

Example:

    fren job jobfile *.txt
    fren 1 one test*.txt
    fren 1 one *doc *txt *.asc

*/
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	s "strings"
)

func main() {

	origStr := os.Args[1]
	newStr := os.Args[2]
	rawfilelist := os.Args[3:]
	var filelist []string

	// expand globs first
	for _, filename := range rawfilelist {
		newfiles, _ := filepath.Glob(filename)
		for _, v := range newfiles {
			filelist = append(filelist, v)
		}
	}
	if len(filelist) == 0 {
		fmt.Println("No files were found.")
		return
	}
	var newfilename string
	var err error
	for _, filename := range filelist {
		// do something
		if s.Contains(filename, origStr) {
			newfilename = s.Replace(filename, origStr, newStr, 1)
			err = os.Rename(filename, newfilename)
			if err != nil {
				log.Println("ERROR on", filename, "==>", err)
			} else {
				fmt.Println(filename, "renamed to", newfilename)
			}
		} else {
			fmt.Printf("%s skipped. It doesn't contain '%s'\n", filename, origStr)
		}
	}
}
