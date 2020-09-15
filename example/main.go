package main

import (
	"fmt"

	"github.com/LoperLee/golang-hangul-toolkit/hangul"
)

func main() {
	s := "안녕하세요"

	han := hangul.ExtractHangul(s)
	gul := hangul.Hangul{Chosung: "ㅇ", Jungsung: "ㅏ", Jongsung: "ㄴ"}
	err := hangul.CombineHangul(&gul)
	if err != nil {
		panic(err)
	}
	fmt.Println(han[0].Chosung, han[0].Jungsung, han[0].Jongsung) // ㅇ ㅏ ㄴ
	fmt.Println(gul.Word)                                         // 안
	fmt.Println(hangul.IsHangul(s))                               // true
}
