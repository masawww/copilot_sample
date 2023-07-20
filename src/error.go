package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("test.txt")
	// この行を消してerrの場合の処理をここに書く。
	fmt.Println("Contents of file:", string(data))
}
