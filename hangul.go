package hangul

var (
	HANGUL_CHO = []string{"ㄱ", "ㄲ", "ㄴ", "ㄷ", "ㄸ", "ㄹ", "ㅁ", "ㅂ", "ㅃ", "ㅅ", "ㅆ", "ㅇ", "ㅈ", "ㅉ", "ㅊ", "ㅋ", "ㅌ", "ㅍ", "ㅎ"}
	HANGUL_JUN = []string{"ㅏ", "ㅐ", "ㅑ", "ㅒ", "ㅓ", "ㅔ", "ㅕ", "ㅖ", "ㅗ", "ㅘ", "ㅙ", "ㅚ", "ㅛ", "ㅜ", "ㅝ", "ㅞ", "ㅟ", "ㅠ", "ㅡ", "ㅢ", "ㅣ"}
	HANGUL_JON = []string{"", "ㄱ", "ㄲ", "ㄳ", "ㄴ", "ㄵ", "ㄶ", "ㄷ", "ㄹ", "ㄺ", "ㄻ", "ㄼ", "ㄽ", "ㄾ", "ㄿ", "ㅀ", "ㅁ", "ㅂ", "ㅄ", "ㅅ", "ㅆ", "ㅇ", "ㅈ", "ㅊ", "ㅋ", "ㅌ", "ㅍ", "ㅎ"}
)

const (
	// 한글의 시작은 '가' 끝은 '힣'
	HANGUL_BASE = rune('가')
	HANGUL_END  = rune('힣')

	HANGUL_JA = rune('ㄱ')
	HANGUL_MO = rune('ㅏ')
)

// 한글의 초성, 중성, 종성을 추출하는 함수
// input : string
// output : []string
func ExtractHangul(s string) []string {
	ret := []string{}
	for _, c := range []rune(s) {
		if c >= HANGUL_JA && c <= (HANGUL_JA+29) {
			// 자음 단독 입력
			ret = append(ret, string(c))
		} else if c >= HANGUL_MO && c <= (HANGUL_MO+20) {
			// 모음 단독 입력
			ret = append(ret, string(c))
		} else if c >= HANGUL_BASE && c <= HANGUL_END {
			// 이미 배열에 모든 문자열이 있기때문에 별도의 유니코드 계산이 불필요 '가'의 유니코드를 빼면 가의경우 0이 됨
			temp := c - HANGUL_BASE
			cho := (temp / 588)
			jung := (temp - (cho * 588)) / 28
			ret = append(ret, HANGUL_CHO[cho])
			ret = append(ret, HANGUL_JUN[jung])
			jong := temp % 28
			if jong != 0 {
				ret = append(ret, HANGUL_JON[jong])
			}
		} else {
			// 한글이 아닌경우 그냥 바로 array에 넣음
			ret = append(ret, string(c))
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
		if c >= HANGUL_JA && c <= (HANGUL_JA+29) {
			// 자음 단독 입력
			continue
		} else if c >= HANGUL_MO && c <= (HANGUL_MO+20) {
			// 모음 단독 입력
			continue
		} else if c >= HANGUL_BASE && c <= HANGUL_END {
			// 일반 한글 입력
			continue
		} else {
			return false
		}
	}
	return true
}
