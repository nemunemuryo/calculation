package main

import (
  "io"
  "strings"
  "strconv"
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

    //splitと数値変換
    word := strings.Split(line, " ")
    var i int
    var j int
    i, _ = strconv.Atoi(word[0])
    j, _ = strconv.Atoi(word[2])

    //演算
    switch word[1] {
    case "+":
      println(i + j)
    case "-":
      println(i - j)
    case "*":
      println(i * j)
    case "/":
      println(i / j)
    }
  }

}
