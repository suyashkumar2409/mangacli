package types

import (
	"errors"
	"mangacli/internal/actions"
)

func GetAction(input string) (actions.Action, error) {
	switch input {
	case "latest":
		return &actions.LatestAction{}, nil
	default:
		return nil, errors.New("unknown action type passed")
	}
}

