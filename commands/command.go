// Package commands define the CLI commands for Lenny
package commands

type Command interface {
	Name() string
	Execute(args []string) error
}

var commandsMap = map[string]Command{}

/* RegisterCommand every command should register itself by calling this action
 */
func RegisterCommand(c Command) {
	commandsMap[c.Name()] = c
}
