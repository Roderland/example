package parser_example

import "log"

type token struct {
	name    string
	literal string
}

var nilLtr string
var tokens []token
var curIdx uint
var curTkn token

func next() {
	curTkn = tokens[curIdx]
	curIdx++
}

func error(msg string) {
	log.Println(msg)
}

func accept(names ...string) bool {
	for _, name := range names {
		if curTkn.name == name {
			next()
			return true
		}
	}
	return false
}

func expect(name string) {
	if !accept(name) {
		error("expect: unexpected symbol")
	}
}
