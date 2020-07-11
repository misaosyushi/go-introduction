package main

import (
	"fmt"
	"os"
)
import sample "./sample"

func main() {
	fmt.Println("Hello World")
	sample.Hello()
	sample.Pointa()
	hoge()
	pointaSample()
}

func hoge() {
	var hoge = "hoge"
	hoge = "fuga"
	fmt.Println(hoge)

	const fuga = "fuga"
	fmt.Println(fuga)

	piyo := ""
	fmt.Println(piyo)
}

func pointaSample() {

	var ptr *int
	var i int = 123

	ptr = &i

	fmt.Println("iのアドレス", &i)
	fmt.Println("ptrの値(変数iのアドレス)", ptr)

	*ptr = 999

	fmt.Println("ポインタ経由で変更したiの値", i)
}

func errorHandlingSample()  {
	file, err := os.Open("test.txt")

	// エラーチェック
	if err != nil {
		fmt.Println(err.Error()) // エラーの詳細を出力
		os.Exit(1) // 終了
	}

	file.Close()
}
