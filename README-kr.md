# Hangul-toolkit
[한국어](https://github.com/LoperLee/golang-hangul-toolkit/blob/master/README-kr.md) | [영어](https://github.com/LoperLee/golang-hangul-toolkit/blob/master/README.md)

이 프로젝트는 오픈소스 프로젝트로 누구나 이용할 수 있습니다.

한글의 초/중/종성 분리, 조합, 조사 붙히기 기능을 제공합니다.

## 사용방법

```
$ go get github.com/LoperLee/golang-hangul-toolkit
```

## 샘플

### Extract

```
han := hangul.ExtractHangul("안녕하세요")
fmt.Println(han[0].Chosung, han[0].Jungsung, han[0].Jongsung)

>> ㅇ ㅏ ㄴ
```

### Combine

```
han := hangul.Hangul{Chosung: "ㅇ", Jungsung: "ㅏ", Jongsung: "ㄴ"}
err := hangul.CombineHangul(&han)
if err != nil {
	panic(err)
}
fmt.Println(han.Word)

>> 안
```

### Check

```
fmt.Println(hangul.IsHangul("안녕하세요"))
fmt.Println(hangul.IsHangul("안s녕하세요"))
fmt.Println(hangul.IsHangul("Hello"))

>> true
>> false
>> false
```

### Josa

| 조사 | const name |
|------|---------|
|은/는 | EUN_NEUN |
|이/가 | I_GA |
|을/를 | EUL_REUL |
|과/와 | GWA_WA |
|이다/다 | IDA_DA |
|으로/로 | EURO_RO |
|률/율 | RYUL_YUL |

```
fmt.Println(hangul.GetJosa("커피", hangul.EUN_NEUN))
>> 커피를
```

## 문서

[Document](https://godoc.org/github.com/LoperLee/golang-hangul-toolkit)

## 예제

[Example](https://github.com/LoperLee/golang-hangul-toolkit/blob/master/example/main.go)

## 라이선스

[MIT License](https://github.com/LoperLee/golang-hangul-toolkit/blob/master/LICENSE)