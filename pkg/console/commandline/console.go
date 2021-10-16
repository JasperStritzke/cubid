package commandline

import (
	"github.com/chzyer/readline"
	"github.com/jasperstritzke/cubid/pkg/console/color"
)

type CommandLine struct {
	Completer *readline.PrefixCompleter
	Line      *readline.Instance
}

func NewCommandLine(completer *readline.PrefixCompleter, promptPrefix string) *CommandLine {
	commandLine := &CommandLine{
		Completer: completer,
	}

	var err error
	commandLine.Line, err = readline.NewEx(&readline.Config{
		Prompt:          color.Blue + promptPrefix + color.Gray + "@" + color.Blue + "Cubid " + color.Gray + "»" + color.Reset + " ",
		AutoComplete:    completer,
		InterruptPrompt: "quit",
		EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})

	if err != nil {
		panic(err)
	}

	return commandLine
}

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}

	return r, true
}
