package ask

import (
	"github.com/AlecAivazis/survey/v2"
)

type DirectoryInput struct {
	Cmd     string
	Options string
}

func InputCmd(message string) (string, error) {
	prompt := &survey.Input{
		Message: message,
		Default: "code",
	}

	var answer string

	err := survey.AskOne(prompt, &answer)

	if err != nil {
		return "", err
	}

	return answer, nil
}

// Directory에 대해 커맨드와 옵션을 입력받음
func InputCmdEachDirectories(directories []string, defaultCmd string) (map[string]DirectoryInput, error) {

	answer := make(map[string]DirectoryInput)

	var err error = nil
	for _, dir := range directories {
		questions := []*survey.Question{
			{
				Name:   "Cmd",
				Prompt: &survey.Input{Message: "Commands for directory '" + dir + "': ", Default: defaultCmd},
			},
			{
				Name:   "Options",
				Prompt: &survey.Input{Message: "Options for directory '" + dir + "': "},
			},
		}

		var inputs DirectoryInput
		err = survey.Ask(questions, &inputs)

		if err != nil {
			return nil, err
		}
		defaultCmd = inputs.Cmd
		newDirectoryInput := DirectoryInput{
			Cmd:     inputs.Cmd,
			Options: inputs.Options,
		}
		answer[dir] = newDirectoryInput

	}

	if err != nil {
		return nil, err
	}
	return answer, nil
}
