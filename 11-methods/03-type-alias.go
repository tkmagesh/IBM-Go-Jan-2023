package main

import "fmt"

type MyString string

func (str MyString) Length() int {
	return len(str)
}

func main() {
	s := MyString("Elit quis tempor ut duis quis reprehenderit nostrud ex cupidatat dolor ea aute occaecat. Minim excepteur proident laboris qui irure ad nisi dolor anim officia enim nisi sit. Do quis labore esse magna est laborum. Veniam enim ea magna enim nostrud aliqua culpa quis aute in consequat amet.")
	fmt.Println(s.Length())
}
