package ask

import "github.com/AlecAivazis/survey/v2"

func InputName(defaultName string) (string, error) {

	prompt := &survey.Input{
		Message: "Project Name: ",
		Default: defaultName,
	}

	var answer string

	err := survey.AskOne(prompt, &answer)

	if err != nil {
		return "", err
	}

	return answer, nil
}
