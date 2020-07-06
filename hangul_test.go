package hangul

import (
	"testing"
)

func TestHangul(t *testing.T) {
	s := Extract("안녕하세요")
	// for _, ss := range s {
	// 	fmt.Println(ss)
	// }
	if len(s) != 6 {
		t.Error("Wrong result")
	}
}
