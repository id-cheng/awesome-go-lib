package main

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
)

var str = "我来到北京清华大学"

func main() {
	x := gojieba.NewJieba()
	defer x.Free()
	words := x.CutAll(str)
	fmt.Println("words1:", words) // 全模式
	words = x.Cut(str, true)
	fmt.Println("words2:", words) // 精确模式

	str = "比特币"
	words = x.Cut(str, true)
	fmt.Println("words3:", words) // 精确模式
	x.AddWord(str)                // 添加词典
	words = x.Cut(str, true)
	fmt.Println("words4:", words) // 精确模式

	str = "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
	words = x.CutForSearch(str, true)
	fmt.Println("搜索引擎模式:", words)

	str = "广东省潮汕药店"
	words = x.Tag(str)
	fmt.Println("词性标注:", words)

	str = "区块链"
	words = x.Tag(str)
	fmt.Println("词性标注:", words)

	str = "长江大桥"
	words = x.CutForSearch(str, false)
	fmt.Println("搜索引擎模式:", words)

	tokenize := x.Tokenize(str, gojieba.SearchMode, false)
	fmt.Println("Tokenize:(搜索引擎模式)", tokenize)

	tokenize = x.Tokenize(str, gojieba.DefaultMode, false)
	fmt.Println("Tokenize:(默认模式)", tokenize)

	keywords := x.ExtractWithWeight(str, 5)
	fmt.Println("Extract:", keywords)
}
