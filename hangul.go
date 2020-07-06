package hangul

import "fmt"

const (
	HANGUL_BASE = 0xAC00
	HANGUL_END  = 0xD7AF

	HANGUL_CHO = 0x1100
	HANGUL_JUN = 0x1161
	HANGUL_JON = 0x11A8 - 1

	HANGUL_JA = 0x3131
	HANGUL_MO = 0x314F
)

func Extract(s string) (ret []string) {
	for _, c := range s {
		choInt := (c - HANGUL_BASE) / 28 / 21
		jungInt := ((c - HANGUL_BASE) / 28) % 21
		jongInt := (c - HANGUL_BASE) % 28
		cho := (choInt + HANGUL_CHO)
		jung := (jungInt + HANGUL_JUN)
		ret = append(ret, string(choInt+HANGUL_CHO))
		ret = append(ret, string(jungInt+HANGUL_JUN))
		fmt.Println(string(cho), string(jung))
		if jongInt != 0 {
			jong := (jongInt + HANGUL_JON)
			ret = append(ret, string(jong))
			fmt.Println(string(jong))
		}
	}
	return
}
