package argsHandler

import (
	"errors"
	"mangacli/internal/actions"
	"mangacli/internal/types"
)

func Consume(args []string) (actions.Action, []string, error){
	if len(args) < 2 {
		return nil, nil, errors.New("action not passed")
	}
	action, err := types.GetAction(args[1])
	if err != nil {
		return nil, nil, err
	}
	return action, args[2:], nil
}

