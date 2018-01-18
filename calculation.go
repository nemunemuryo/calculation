package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/prataprc/goparsec"
	"github.com/prataprc/goparsec/expr"
	"io"
)

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			}
		} else if err == io.EOF {
			break
		}

		result := doExpr(line)

		fmt.Println(result)
	}
}

func doExpr(text string) parsec.ParsecNode {
	s := parsec.NewScanner([]byte(text))
	v, _ := expr.Y(s)
	return v
}
