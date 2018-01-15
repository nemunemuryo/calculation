package main

import (
  "flag"
  "fmt"
  "io"
  "io/ioutil"
  "os"
  // "strings"
  // "strconv"
  "github.com/chzyer/readline"
  "github.com/prataprc/goparsec"
  "github.com/prataprc/goparsec/expr"
)

var options struct {
  expr string
}

func argParse() {
  flag.StringVar(&options.expr, "expr", "",
  "Specify input file or arithmetic expression string")
  flag.Parse()
}

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
    // word := strings.Split(line, " ")
    // var i int
    // var j int
    // i, _ = strconv.Atoi(word[0])
    // j, _ = strconv.Atoi(word[2])

    argParse()
    if options.expr != "" {
      doExpr(getText(options.expr))
    }
    //演算

    // switch word[1] {
    // case "+":
    //   println(i + j)
    // case "-":
    //   println(i - j)
    // case "*":
    //   println(i * j)
    // case "/":
    //   println(i / j)
    // }
  }

}

func doExpr(text string) {
  s := parsec.NewScanner([]byte(text))
  v, _ := expr.Y(s)
  fmt.Println(v)
}

func getText(filename string) string {
  if _, err := os.Stat(filename); err != nil {
    return filename
  }
  if b, err := ioutil.ReadFile(filename); err != nil {
    panic(err)
  } else {
    return string(b)
  }
}
