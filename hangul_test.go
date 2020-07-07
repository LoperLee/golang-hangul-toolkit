package hangul

import (
	"testing"
)

func TestExtractHangul(t *testing.T) {
	s := ExtractHangul("안녕하세요")
	if len(s) != 12 {
		t.Error("Wrong result")
	}
}

func TestIsHangul(t *testing.T) {
	if !IsHangul("안녕하세요") {
		t.Error("Wrong result")
	}
	if IsHangul("Hello") {
		t.Error("Wrong result")
	}
	if IsHangul("Hello안녕") {
		t.Error("Wrong result")
	}
}
