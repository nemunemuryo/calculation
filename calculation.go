package main

import (
  "fmt"
  "io"
  //"strconv"
  "flag"
  "io/ioutil"
  "os"
  //"strings"
  "github.com/chzyer/readline"
  "github.com/prataprc/goparsec"
  "github.com/prataprc/goparsec/expr"
)

// goparsec/exprがあるし使えば関係はず...--------------------------------------------------------
// var _ = fmt.Sprintf("dummp print")
//
// // Y is root Parser, usually called as `s` in CFG theory.
// var Y parsec.Parser
// var prod, sum, value parsec.Parser // circular rats
//
// // Terminal rats
// var openparan = parsec.Token(`\(`, "OPENPARAN")
// var closeparan = parsec.Token(`\)`, "CLOSEPARAN")
// var addop = parsec.Token(`\+`, "ADD")
// var subop = parsec.Token(`-`, "SUB")
// var multop = parsec.Token(`\*`, "MULT")
// var divop = parsec.Token(`/`, "DIV")
//
//
// func one2one(ns []parsec.ParsecNode) parsec.ParsecNode {
// 	if ns == nil || len(ns) == 0 {
// 		return nil
// 	}
// 	return ns[0]
// }
//
// func many2many(ns []parsec.ParsecNode) parsec.ParsecNode {
// 	if ns == nil || len(ns) == 0 {
// 		return nil
// 	}
// 	return ns
// }
//
// // NonTerminal rats
// // addop -> "+" |  "-"
// var sumOp = parsec.OrdChoice(one2one, addop, subop)
//
// // mulop -> "*" |  "/"
// var prodOp = parsec.OrdChoice(one2one, multop, divop)
//
// // value -> "(" expr ")"
// var groupExpr = parsec.And(exprNode, openparan, &sum, closeparan)
//
// // (addop prod)*
// var prodK = parsec.Kleene(nil, parsec.And(many2many, sumOp, &prod), nil)
//
// // (mulop value)*
// var valueK = parsec.Kleene(nil, parsec.And(many2many, prodOp, &value), nil)
//
// func init() {
// 	// Circular rats come to life
// 	// sum -> prod (addop prod)*
// 	sum = parsec.And(sumNode, &prod, prodK)
// 	// prod-> value (mulop value)*
// 	prod = parsec.And(prodNode, &value, valueK)
// 	// value -> num | "(" expr ")"
// 	value = parsec.OrdChoice(exprValueNode, intWS(), groupExpr)
// 	// expr  -> sum
// 	Y = parsec.OrdChoice(one2one, sum)
// }
//
// func intWS() parsec.Parser {
// 	return func(s parsec.Scanner) (parsec.ParsecNode, parsec.Scanner) {
// 		_, s = s.SkipAny(`^[  \n\t]+`)
// 		p := parsec.Int()
// 		return p(s)
// 	}
// }
//
// //----------
// // Nodifiers
// //----------
//
// func sumNode(ns []parsec.ParsecNode) parsec.ParsecNode {
// 	if len(ns) > 0 {
// 		val := ns[0].(int)
// 		for _, x := range ns[1].([]parsec.ParsecNode) {
// 			y := x.([]parsec.ParsecNode)
// 			n := y[1].(int)
// 			switch y[0].(*parsec.Terminal).Name {
// 			case "ADD":
// 				val += n
// 			case "SUB":
// 				val -= n
// 			}
// 		}
// 		return val
// 	}
// 	return nil
// }
//
// func prodNode(ns []parsec.ParsecNode) parsec.ParsecNode {
// 	if len(ns) > 0 {
// 		val := ns[0].(int)
// 		for _, x := range ns[1].([]parsec.ParsecNode) {
// 			y := x.([]parsec.ParsecNode)
// 			n := y[1].(int)
// 			switch y[0].(*parsec.Terminal).Name {
// 			case "MULT":
// 				val *= n
// 			case "DIV":
// 				val /= n
// 			}
// 		}
// 		return val
// 	}
// 	return nil
// }
//
// func exprNode(ns []parsec.ParsecNode) parsec.ParsecNode {
// 	if len(ns) == 0 {
// 		return nil
// 	}
// 	return ns[1]
// }
//
// func exprValueNode(ns []parsec.ParsecNode) parsec.ParsecNode {
// 	if len(ns) == 0 {
// 		return nil
// 	} else if term, ok := ns[0].(*parsec.Terminal); ok {
// 		val, _ := strconv.Atoi(term.Value)
// 		return val
// 	}
// 	return ns[0]
// }
//----------------------------------------------------------------------------------------------

var options struct {
  expr string
}

func argParse() {
  // flag.StringVar(&options.expr, "expr", "",
  // "Specify input file or arithmetic expression string")
  // flag.Parse()

  // 引数の間に余計なものが合ってもフラグ解析が切れることはない実装
  f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
  f.StringVar(&options.expr, "expr", "",
  "Specify input file or arithmetic expression string")

	f.Parse(os.Args[1:])
	for 0 < f.NArg() {
		f.Parse(f.Args()[1:])
  }
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

    options.expr = line
    println(options.expr)
    argParse()
    if options.expr != "" {
      doExpr(getText(options.expr))
    }

    //splitと数値変換
    // word := strings.Split(line, " ")
    // var i int
    // var j int
    // i, _ = strconv.Atoi(word[0])
    // j, _ = strconv.Atoi(word[2])

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
