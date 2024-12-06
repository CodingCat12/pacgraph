package helper

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

func ToString(value any) string {
	result := fmt.Sprintf("%v", value)
	return result
}

func RemoveContents(dirName string) error {
	dir, err := os.Open(dirName)
	if err != nil {
		return err
	}
	defer dir.Close()
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dirName, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func Confirm(question string, def bool) (bool, error) {
	prompt := promptui.Prompt{
		Label:     question,
		IsConfirm: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return def, err
	}

	return result == "y", nil
}
