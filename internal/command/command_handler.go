package command

type commandHandler func(s *State, cmd Command) error
