package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	parser "gparser/internal/iantlr/grammar"
	iparser "gparser/internal/iparser"
)

//	func main() {
//		is := antlr.NewInputStream("select id , name from mart_waimai.fact_wm_food_spu where dt = '20240609' and 1=1 ")
//		lexer := parser.NewHiveLexer(is)
//		lexerErrorListener := iparser.NewGparserErrorListener()
//		lexer.AddErrorListener(lexerErrorListener)
//		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
//		hiveParser := parser.NewHiveParser(stream)
//		gparser := iparser.NewGparserParserListener()
//		parserErrorListener := iparser.NewGparserErrorListener()
//		hiveParser.AddErrorListener(parserErrorListener)
//		hiveParser.BuildParseTrees = true
//		//baseVisitor := new(parser.BaseHiveParserVisitor)
//		//baseVisitor.Visit(hiveParser.Statement())
//		////hiveVisitor := iparser.NewGHiveVisitor()
//		////tree := hiveParser.Statement()
//
//		antlr.ParseTreeWalkerDefault.Walk(gparser, hiveParser.Statement())
//		if len(lexerErrorListener.GrammarErrors) > 0 {
//			fmt.Println("<===============lexer error :=============>")
//			fmt.Println(lexerErrorListener.GrammarErrors)
//		}
//		if len(parserErrorListener.GrammarErrors) > 0 {
//			fmt.Println("<===============grammar error :=============>")
//			fmt.Println(parserErrorListener.GrammarErrors)
//		}
//
// }
const (
	sql1 = `select id , name from mart_waimai.fact_wm_food_spu spu where dt = '20240609' and 1=1 `
	sql2 = `insert into mart_waimai.aggr_food_spu_info_dd partition(dt) 
		select spu.id as spu_id , spu.name as spu_name , spu.dt ,
		 sku.id as sku_id, sku.name as sku_name  , dt 
		 from mart_waimai.fact_wm_food_spu spu 
		left join mart_waimai.fact_wm_food_snapshot sku 
		on spu.dt = sku.dt and spu.id = sku.wm_food_spu_id 
		 where spu.dt = '20240609' and 1=1`
)

func main() {
	is := antlr.NewInputStream(sql2)

	lexer := parser.NewHiveLexer(is)
	lexerErrorListener := iparser.NewGparserErrorListener()
	lexer.AddErrorListener(lexerErrorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	hiveParser := parser.NewHiveParser(stream)
	parserErrorListener := iparser.NewGparserErrorListener()
	hiveParser.AddErrorListener(parserErrorListener)
	baseVisitor := iparser.NewGHiveVisitor()
	//hiveVisitor := iparser.NewGHiveVisitor()
	//tree := hiveParser.Statement()
	hiveParser.BuildParseTrees = true
	tree := hiveParser.Statement()
	baseVisitor.Visit(tree)
	if len(lexerErrorListener.GrammarErrors) > 0 {
		fmt.Println("<===============lexer error :=============>")
		fmt.Println(lexerErrorListener.GrammarErrors)
	}
	if len(parserErrorListener.GrammarErrors) > 0 {
		fmt.Println("<===============grammar error :=============>")
		fmt.Println(parserErrorListener.GrammarErrors)
	}
	fmt.Printf("columns :%v , DB :%v , table :%v  alias :%v\n", baseVisitor.Columns, baseVisitor.DB, baseVisitor.Table, baseVisitor.Alias)
}
