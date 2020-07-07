package hangul

import (
	"errors"
)

var (
	HangulCHO = []string{"ㄱ", "ㄲ", "ㄴ", "ㄷ", "ㄸ", "ㄹ", "ㅁ", "ㅂ", "ㅃ", "ㅅ", "ㅆ", "ㅇ", "ㅈ", "ㅉ", "ㅊ", "ㅋ", "ㅌ", "ㅍ", "ㅎ"}
	HangulJUN = []string{"ㅏ", "ㅐ", "ㅑ", "ㅒ", "ㅓ", "ㅔ", "ㅕ", "ㅖ", "ㅗ", "ㅘ", "ㅙ", "ㅚ", "ㅛ", "ㅜ", "ㅝ", "ㅞ", "ㅟ", "ㅠ", "ㅡ", "ㅢ", "ㅣ"}
	HangulJON = []string{"", "ㄱ", "ㄲ", "ㄳ", "ㄴ", "ㄵ", "ㄶ", "ㄷ", "ㄹ", "ㄺ", "ㄻ", "ㄼ", "ㄽ", "ㄾ", "ㄿ", "ㅀ", "ㅁ", "ㅂ", "ㅄ", "ㅅ", "ㅆ", "ㅇ", "ㅈ", "ㅊ", "ㅋ", "ㅌ", "ㅍ", "ㅎ"}
)

const (
	// 한글의 시작은 '가' 끝은 '힣'
	HangulBASE = rune('가')
	HangulEND  = rune('힣')

	// 자음은 29개, 모음은 20개
	HangulJA = rune('ㄱ')
	HangulMO = rune('ㅏ')
)

type Hangul struct {
	Word     string
	Chosung  string
	Jungsung string
	Jongsung string
}

// 한글의 초성, 중성, 종성을 추출하는 함수
// input : string
// output : []Hangul
func ExtractHangul(s string) []Hangul {
	ret := []Hangul{}
	for _, c := range []rune(s) {
		item := Hangul{Word: string(c)}
		if c >= HangulJA && c <= (HangulJA+29) {
			// 자음 단독 입력
			ret = append(ret, item)
		} else if c >= HangulMO && c <= (HangulMO+20) {
			// 모음 단독 입력
			ret = append(ret, item)
		} else if c >= HangulBASE && c <= HangulEND {
			// 이미 배열에 모든 문자열이 있기때문에 별도의 유니코드 계산이 불필요 '가'의 유니코드를 빼면 가의경우 0이 됨
			temp := c - HangulBASE
			cho := (temp / 588)
			jung := (temp - (cho * 588)) / 28
			item.Chosung = HangulCHO[cho]
			item.Jungsung = HangulJUN[jung]
			if jong := temp % 28; jong != 0 {
				item.Jongsung = HangulJON[jong]
			}
			ret = append(ret, item)
		} else {
			// 한글이 아닌경우 그냥 바로 array에 넣음
			ret = append(ret, item)
		}
	}
	return ret
}

// 문자열의 한글 여부를 판단해주는 함수
// input : string
// output : bool
func IsHangul(s string) bool {
	// 한글 판단 로직
	for _, c := range []rune(s) {
		if c >= HangulJA && c <= (HangulJA+29) {
			// 자음 단독 입력
			continue
		} else if c >= HangulMO && c <= (HangulMO+20) {
			// 모음 단독 입력
			continue
		} else if c >= HangulBASE && c <= HangulEND {
			// 일반 한글 입력
			continue
		} else {
			return false
		}
	}
	return true
}

func CombineHangul(word *Hangul) error {
	// 몇가지 오류상황을 가정한다.
	// 1. 초성이나 중성이 없을 경우
	// 2. 초성이나 중성에 문자열이 1글자 이상일 경우
	// 3. 초성이나 중성에 한글 이외의 문자가 있을 경우
	if (len(word.Chosung) <= 0 || len(word.Jungsung) <= 0) || (len(word.Chosung) > 3 || len(word.Jungsung) > 3) {
		// 한글은 3바이트 len에서도 3으로 출력됨
		return errors.New("Invalid input value")
	}
	cho := findIndex(HangulCHO, word.Chosung)
	jun := findIndex(HangulJUN, word.Jungsung)
	jon := findIndex(HangulJON, word.Jongsung)

	// 초, 중, 종성의 유효성 검사
	if cho < 0 {
		return errors.New("Chosung isn't correct string")
	} else if jun < 0 {
		return errors.New("Jungsung isn't correct string")
	} else if jon < 0 {
		return errors.New("Jongsung isn't correct string")
	}

	// (588 * cho) + (28 * jun) + jon + HangulBASE
	ret := (588 * rune(cho)) + (28 * rune(jun)) + rune(jon) + HangulBASE
	word.Word = string(ret)

	return nil
}

func findIndex(array []string, target string) int {
	ret := -1
	for idx, value := range array {
		if value == target {
			ret = idx
			break
		}
	}
	return ret
}
