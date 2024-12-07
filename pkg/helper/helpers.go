package helper

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func ToString(value any) string {
	result := fmt.Sprintf("%v", value)
	return result
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
