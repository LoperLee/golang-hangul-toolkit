package hangul

import (
	"testing"
)

func TestExtractHangul(t *testing.T) {
	s := ExtractHangul("안녕하세요")
	if len(s) != 5 {
		t.Error("IsHangul correct is wrong!")
	}
}

func TestIsHangul(t *testing.T) {
	if !IsHangul("안녕하세요") {
		t.Error("IsHangul correct is wrong!")
	}
	if IsHangul("Hello") {
		t.Error("IsHangul error case 1 is wrong!")
	}
	if IsHangul("Hello안녕") {
		t.Error("IsHangul error case 2 is wrong!")
	}
}

func TestCombineHangul(t *testing.T) {
	correct := Hangul{Chosung: "ㅇ", Jungsung: "ㅏ", Jongsung: "ㄴ"}
	error1 := Hangul{Chosung: "ㅇ", Jungsung: "a"}
	error2 := Hangul{Chosung: "a"}
	error3 := Hangul{Chosung: "ㅇㅏ", Jungsung: "ㅏㅏ"}
	error4 := Hangul{Chosung: "안", Jungsung: "녕"}

	// 정상동작
	err := CombineHangul(&correct)
	if err != nil {
		t.Error("Combine correct is wrong!")
	}

	// 의도적 오류 발생
	err = CombineHangul(&error1)
	if err == nil {
		t.Error("Combine error case 1 is wrong!")
	}
	err = CombineHangul(&error2)
	if err == nil {
		t.Error("Combine error case 2 is wrong!")
	}
	err = CombineHangul(&error3)
	if err == nil {
		t.Error("Combine error case 3 is wrong!")
	}
	err = CombineHangul(&error4)
	if err == nil {
		t.Error("Combine error case 4 is wrong!")
	}
}
