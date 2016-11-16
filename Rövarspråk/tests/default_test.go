package testing

import (
	"regexp"
	"strings"
	"testing"
)

var (
	reg  = regexp.MustCompile("([B-D|b-d|F-H|f-h|J-N|j-n|P-T|p-t|V-X|v-x|Z|z])")
	reg2 = regexp.MustCompile("([B-Db-dF-Hf-hJ-Nj-nP-Tp-tV-Xv-xZz])")
	reg3 = regexp.MustCompile("((?i)[b-df-gj-np-tv-xz])")
)

type rövare struct {
	phrase   string
	expected string
}

var tests = []rövare{
	{"Dennis Planting", "Dodenonnonisos Poplolanontotinongog"},
	{"Detta är ett test", "Dodetottota äror etottot totesostot"},
	{"Programmering är väldigt roligt", "Poprorogogroramommomerorinongog äror voväloldodigogtot rorololigogtot"},
}

func stringInSlice(str string, list []string) bool {
	str = strings.ToUpper(str)
	for _, tecken := range list {
		if str == tecken {
			return true
		}
	}
	return false
}

func RövarspråkLoop(input string) string {
	var rövarspråk string
	konsonanter := []string{"B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "X", "Z"}
	inteRövarspråkSlice := strings.Split(input, "")
	for _, tecken := range inteRövarspråkSlice {
		if tecken == "\n" || tecken == "\r" {
			continue
		}
		if stringInSlice(tecken, konsonanter) {
			rövarspråk += tecken + "o" + strings.ToLower(tecken)
		} else {
			rövarspråk += tecken
		}
	}
	return rövarspråk
}

func TestLoop(t *testing.T) {
	for _, test := range tests {
		v := RövarspråkLoop(test.phrase)
		if v != test.expected {
			t.Error("Phrase", test.phrase, "expected", test.expected, "got", v)
		}
	}
}
func TestReg3(t *testing.T) {
	for _, test := range tests {
		v := reg3.ReplaceAllStringFunc(test.phrase, func(s string) string {
			return s + "o" + strings.ToLower(s)
		})
		if v != test.expected {
			t.Error("Phrase", test.phrase, "expected", test.expected, "got", v)
		}
	}
}

func BenchmarkLoop1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RövarspråkLoop(tests[0].phrase)
	}
}

func BenchmarkRegex1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reg.ReplaceAllString(tests[0].phrase, "${1}o${1}")
	}
}

func BenchmarkRegex2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reg2.ReplaceAllString(tests[0].phrase, "${1}o${1}")
	}
}

func BenchmarkRegex3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reg3.ReplaceAllStringFunc(tests[0].phrase, func(s string) string {
			return s + "o" + strings.ToLower(s)
		})
	}
}
