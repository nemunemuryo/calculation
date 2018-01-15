package main

import (
  "io"
  "github.com/chzyer/readline"
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
      } else {
        continue
      }
    } else if err == io.EOF {
      break
    }
  }
}
