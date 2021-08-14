package models

type NodeJsweeklyNameNotFoundError struct{}

func (k *NodeJsweeklyNameNotFoundError) Error() string {
	return "NodeJsweekly name not found"
}
