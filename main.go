package main

import (
	"fmt"
	"log"

	"github.com/s4malve/commit-me-daddy/git"
)

func main() {
	stagedFiles, err := git.StagedFiles()

	if err != nil {
		log.Fatal(err)
	}

	if len(stagedFiles) == 0 {
		log.Println("No staged files")
		return
	}

	fileContent := "Input: Generate a single, short and comprehensive git commit message based on the context (filenames and file contents). Following the standards (feat,fix,etc...). Dont provide any explanation, issue template or any of that, just the commit message\n\n"

	fileContent += "Files\n\n"
	for _, filename := range stagedFiles {
		stagedFileContent, err := git.GetStagedFileContent(filename)

		if err != nil {
			log.Println(err)
		}

		fileContent += fmt.Sprintf("filename: %s\n\n", filename)
		fileContent += fmt.Sprintf("content: %s\n\n", stagedFileContent)
	}

	fileContent = fileContent[:len(fileContent)-3]

	fmt.Println(fileContent)
}
