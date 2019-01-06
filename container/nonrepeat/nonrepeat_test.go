package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct{
		s string
		ans int
	} {
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},
		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbb",1},
		{"abcabcabcd", 4},
		// chinese support
		{"这里是慕课网",6},
		{"一二三三二", 3},
		{"寻找最长不含有重复字符的子串", 14},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubString(tt.s)
		if actual != tt.ans{
			t.Errorf("got %d for input %s; " + "expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发会黑会飞花"
	ans := 9

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubString(s)
		if actual != ans {
			b.Errorf("got %d for input %s; " + "expected %d", actual, s, ans)
		}
	}
}
