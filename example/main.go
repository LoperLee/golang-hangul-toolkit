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

	for _, keyword := range []string{"커피", "구글"} {
		fmt.Println(hangul.GetJosa(keyword, hangul.EUN_NEUN))
		fmt.Println(hangul.GetJosa(keyword, hangul.I_GA))
		fmt.Println(hangul.GetJosa(keyword, hangul.EUL_REUL))
		fmt.Println(hangul.GetJosa(keyword, hangul.GWA_WA))
		fmt.Println(hangul.GetJosa(keyword, hangul.IDA_DA))
		fmt.Println(hangul.GetJosa(keyword, hangul.EURO_RO))
	}
	for _, keyword := range []string{"방어", "공격"} {
		fmt.Println(hangul.GetJosa(keyword, hangul.RYUL_YUL))
	}
}
