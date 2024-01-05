package ask

import "github.com/AlecAivazis/survey/v2"

func Confirm(message string) bool {
	var ok bool

	prompt := &survey.Confirm{
		Message: message + " Are you sure?",
	}

	survey.AskOne(prompt, &ok)

	return ok
}
