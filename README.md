# Hangul-toolkit
[Korean](https://github.com/LoperLee/golang-hangul-toolkit/blob/master/README-kr.md) | [English](https://github.com/LoperLee/golang-hangul-toolkit/blob/master/README.md)

This project is an open source project and can be used by anyone.

It provides the functions of separating, combining, and attaching research to the elementary/middle/final of Hangul.

## Install

```
$ go get github.com/LoperLee/golang-hangul-toolkit
```

## Sample

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

## Document

[Document](https://godoc.org/github.com/LoperLee/golang-hangul-toolkit)

## Example

[Example](https://github.com/LoperLee/golang-hangul-toolkit/blob/master/example/main.go)

## License

[MIT License](https://github.com/LoperLee/golang-hangul-toolkit/blob/master/LICENSE)