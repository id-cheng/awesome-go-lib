package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	hans := "中国人"

	// 默认
	a := pinyin.NewArgs()
	fmt.Println(pinyin.Pinyin(hans, a))

	// 包含声调
	a.Style = pinyin.Tone
	fmt.Println(pinyin.Pinyin(hans, a))

	// 声调用数字表示
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))

	// 开启多音字模式
	a = pinyin.NewArgs()
	a.Heteronym = true
	fmt.Println(pinyin.Pinyin(hans, a))

	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))

	fmt.Println(pinyin.LazyPinyin(hans, pinyin.NewArgs()))

	fmt.Println(pinyin.Convert(hans, nil))

	fmt.Println(pinyin.LazyConvert(hans, nil))
}
