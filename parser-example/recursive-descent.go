package parser_example

func program() {
	next()
	block()
	expect(".")
}

func block() {
	// const
	if accept("const") {
		expect("ident")
		expect("=")
		expect("number")
		for accept(",") {
			expect("ident")
			expect("=")
			expect("number")
		}
		expect(";")
	}

	// var
	if accept("var") {
		expect("ident")
		for accept(",") {
			expect("ident")
		}
		expect(";")
	}

	// procedure
	for accept("procedure") {
		expect("ident")
		expect(";")
		block()
		expect(";")
	}

	// statement
	statement()
}

func statement() {
	if accept(".") {
		return
	}

	if accept("ident") {
		expect(":=")
		expression()
	} else if accept("call") {
		expect("ident")
	} else if accept("begin") {
		statement()
		for accept(";") {
			statement()
		}
		expect("end")
	} else if accept("if") {
		condition()
		expect("then")
		statement()
	} else if accept("while") {
		condition()
		expect("do")
		statement()
	} else if accept("!") {
		expect("ident")
	} else {
		error("statement: syntax error")
		next()
	}
}

func condition() {
	if accept("odd") {
		expression()
	} else {
		expression()
		if accept("=", "#", "<", "<=", ">", ">=") {
			expression()
		} else {
			error("condition: invalid operator")
			next()
		}
	}
}

func expression() {
	accept("+", "-")
	term()
	for accept("+", "-") {
		term()
	}
}

func term() {
	factor()
	for accept("*", "/") {
		factor()
	}
}

func factor() {
	if accept("ident") {

	} else if accept("number") {

	} else if accept("(") {
		expression()
		expect(")")
	} else {
		error("factor: syntax error")
		next()
	}
}
