package git

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func StagedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", "--cached")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	files := make([]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(output)))

	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		files = append(files, txt)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

func RepoExist() error {
	cmd := exec.Command("git", "status")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.New(stderr.String())
	}

	return nil
}

func GetStagedFileContent(filePath string) (string, error) {
	unneededFiles := map[string]bool{"go.mod": true, "go.sum": true}

	if unneededFiles[filePath] {
		return "", fmt.Errorf("file path %s ignored because it is considered useless", filePath)

	}

	cmd := exec.Command("git", "show", ":"+filePath)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running git show: %v\nstderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}
