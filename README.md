# Hangul-toolkit

이 프로젝트는 오픈소스 프로젝트로 누구나 이용할 수 있습니다.

한글의 자모음 분리와 조합, 한글여부를 판단해주는 기능을 제공합니다.

## 사용방법

```
$ go get github.com/LoperLee/golang-hangul-toolkit
```

## 문서

[Document](https://godoc.org/github.com/LoperLee/golang-hangul-toolkit)

## 예제

```
package main

import (
	"fmt"
	hangul "github.com/LoperLee/golang-hangul-toolkit"
)

func main() {
	s := "안녕하세요"

	han := hangul.ExtractHangul(s)
	gul := hangul.Hangul{Chosung: "ㅇ", Jungsung: "ㅏ", Jongsung: "ㄴ"}
	err := hangul.CombineHangul(gul)
	if err != nil {
		panic(err)
	}
	fmt.Println(han[0].Chosung, han[0].Jungsung, han[0].Jongsung) // ㅇ ㅏ ㄴ
	fmt.Println(gul.Word)                                         // 안
	fmt.Println(hangul.IsHangul(s))                               // true
}
```

## 라이선스

[MIT License](https://github.com/LoperLee/golang-hangul-toolkit/blob/master/LICENSE)