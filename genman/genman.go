package main

import (
	"../cmd"
	"fmt"
	"github.com/spf13/cobra/doc"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

var version = "2.0.0"

func usage() {
	fmt.Println("usage: genman <output_path>")
}

func main() {
	destinationPath, _ := os.Getwd()
	if len(os.Args) > 2 {
		usage()
		os.Exit(1)
	}
	if len(os.Args) == 2 {
		destinationPath = os.Args[1]
	}

	header := &doc.GenManHeader{
		Title:   "GIT-USER",
		Section: "1",
		Manual:  "Git Manual",
		Source:  fmt.Sprintf("git-user %s", version),
	}
	cmd.RootCmd.DisableAutoGenTag = true
	fmt.Printf("Generating documentation for git-user\n")
	err := doc.GenManTree(cmd.RootCmd, header, destinationPath)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(2)
	}

	//Remove all of the double blank lines from output docs
	err = filepath.Walk(destinationPath, func(path string, f os.FileInfo, err error) error {
		stripFile(path)
		return nil
	})

	if err != nil {
		fmt.Printf("Could not clean up all the files\n")
		fmt.Printf("%s", err)
	}
}

func stripFile(path string) error {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	regex, err := regexp.Compile("\n{2,}")
	if err != nil {
		return err
	}
	output := regex.ReplaceAllString(string(input), "\n")

	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
