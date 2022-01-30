package core

type CommandHandler interface {
	Run(CommandContext) error
}
