// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"fmt"
	"math"
	"text/scanner"
	"strings"
	"strconv"
)

// Expr は算術式
type Expr interface{
	Eval(env Env) float64
}

type Env map[Var]float64

// Var は変数を特定します。
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

// literal は数値定数。
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// unary は単項演算式を表します。
type unary struct {
	op rune  // +, -
	x Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

// binary は二項演算式を表します。
type binary struct {
	op rune // +, -, *, /
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

// call は関数呼び出しを表します。
type call struct {
	fn string  // "pow", "sin", "sqrt"
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %q", c.fn))
}

type lexer struct {
	scan scanner.Scanner
	token rune
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token))
}

func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

func Parse(input string) (_ Expr, err error) {
	defer func() {
		switch x := recover().(type){
		case nil:
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			panic(x)
		}
	}()
	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	lex.next()
	e := parseExpr(lex)
	if lex.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %s", lex.describe())
	}
	return e, nil
}

func parseExpr(lex *lexer) Expr { return parseBinary(lex, 1) }

func parseBinary(lex *lexer, prec1 int) Expr {
	lhs := parseUnary(lex)
	for prec := precedence(lex.token); prec >= prec1; prec-- {
		for precedence(lex.token) == prec {
			op := lex.token
			lex.next()
			rhs := parseBinary(lex, prec+1)
			lhs = binary{op, lhs, rhs}
		}
	}
	return lhs
}

func parseUnary(lex *lexer) Expr {
	if lex.token == '+' || lex.token == '-' {
		op := lex.token
		lex.next()
		return unary{op, parseUnary(lex)}
	}
	return parsePrimary(lex)
}

func parsePrimary(lex *lexer) Expr {
	switch lex.token {
	case scanner.Ident:
		id := lex.text()
		lex.next()
		if lex.token != '(' {
			return Var(id)
		}
		lex.next()
		var args []Expr
		if lex.token != ')' {
			for {
				args = append(args, parseExpr(lex))
				if lex.token != ',' {
					break
				}
				lex.next()
			}
			if lex.token != ')' {
				msg := fmt.Sprintf("got %q, want ')'", lex.token)
				panic(lexPanic(msg))
			}
		}
		lex.next()
		return call{id, args}

	case scanner.Int, scanner.Float:
		f, err := strconv.ParseFloat(lex.text(), 64)
		if err != nil {
			panic(lexPanic(err.Error()))
		}
		lex.next()
		return literal(f)

	case '(':
		lex.next()
		e := parseExpr(lex)
		if lex.token != ')' {
			msg := fmt.Sprintf("got %s, want ')'", lex.describe)
			panic(lexPanic(msg))
		}
		lex.next()
		return e
	}
  msg := fmt.Sprintf("unexpected %s", lex.describe())
  panic(lexPanic(msg))
}