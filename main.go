package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	parser "gparser/internal/iantlr/grammar"
	iparser "gparser/internal/iparser"
)

func main() {
	is := antlr.NewInputStream("select id , name from mart_waimai.fact_wm_food_spu where dt = '20240609' and 1=1 ")
	lexer := parser.NewHiveLexer(is)
	lexerErrorListener := iparser.NewGparserErrorListener()
	lexer.AddErrorListener(lexerErrorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	hiveParser := parser.NewHiveParser(stream)
	gparser := iparser.NewGparserParserListener()
	parserErrorListener := iparser.NewGparserErrorListener()
	hiveParser.AddErrorListener(parserErrorListener)
	hiveParser.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(gparser, hiveParser.Statement())
	if len(lexerErrorListener.GrammarErrors) > 0 {
		fmt.Println("<===============lexer error :=============>")
		fmt.Println(lexerErrorListener.GrammarErrors)
	}
	if len(parserErrorListener.GrammarErrors) > 0 {
		fmt.Println("<===============grammar error :=============>")
		fmt.Println(parserErrorListener.GrammarErrors)
	}

}
