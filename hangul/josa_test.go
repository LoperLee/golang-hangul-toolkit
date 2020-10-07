package hangul

import (
	"testing"
)

func TestJosa(t *testing.T) {
	for _, keyword := range []string{"네이버", "구글", "카카오톡", "하늘", "과거", "커피"} {
		t.Log(GetJosa(keyword, EUN_NEUN))
		t.Log(GetJosa(keyword, I_GA))
		t.Log(GetJosa(keyword, EUL_REUL))
		t.Log(GetJosa(keyword, GWA_WA))
		t.Log(GetJosa(keyword, IDA_DA))
		t.Log(GetJosa(keyword, EURO_RO))
	}
	for _, keyword := range []string{"방어", "공격", "반환"} {
		t.Log(GetJosa(keyword, RYUL_YUL))
	}
}
