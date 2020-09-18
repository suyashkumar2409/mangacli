package actions

type Action interface {
	Execute(args []string) (string, error)
}
