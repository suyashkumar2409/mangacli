package actions

import "errors"

type LatestAction struct{

}

func (action *LatestAction) Execute(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("only one manga should be passed")
	}

}
