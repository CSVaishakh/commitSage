package internal

import (
	"os"
	"os/exec"
)

func Changes(path string) (status bool, oldContent string, updatedContent string, err error) {
	newBytes, err := os.ReadFile(path)
	if err != nil {
		return false,"","",err
	}

	updatedContent = string(newBytes)

	cmd := exec.Command("git", "show", "HEAD:"+path)
	oldBytes, err := cmd.Output()

	if err != nil {
		oldBytes = []byte{}
	}

	oldContent = string(oldBytes)

	if oldContent != updatedContent {
		return true, oldContent,updatedContent,nil
	}

	return false,"","",nil
}