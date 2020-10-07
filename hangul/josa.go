package hangul

var josaList = [][]string{{"은", "는"}, {"이", "가"}, {"을", "를"}, {"과", "와"}, {"이다", "다"}, {"으로", "로"}, {"률", "율"}}

// 조사 타입 지정
type Josa int

// 은/는, 이/가, 을/를, 과/와, 이다/다, 으로/로, 률/율
const (
	EUN_NEUN Josa = iota
	I_GA
	EUL_REUL
	GWA_WA
	IDA_DA
	EURO_RO
	RYUL_YUL
)

func GetJosa(keyword string, josa Josa) string {
	ret := keyword
	pic := josaList[josa]
	han := ExtractHangul(keyword)
	if han[len(han)-1].Jongsung != "" {
		if josa == EURO_RO && han[len(han)-1].Jongsung == "ㄹ" {
			ret += pic[1]
		} else {
			ret += pic[0]
		}
	} else {
		ret += pic[1]
	}
	return ret
}
