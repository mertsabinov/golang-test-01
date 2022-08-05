package iterasyon_01

import "testing"

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 5)
	expected := "aaaaa"

	if repeat != expected {
		t.Errorf("repeat = %s expected = %s ", repeat, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestReplaceString(t *testing.T) {
	replace := ReplaceString("test", "t", "b")
	expected := "besb"

	if replace != expected {
		t.Errorf("replace = %s expected = %s", replace, expected)
	}
}

func BenchmarkReplaceString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReplaceString("test", "t", "b")
	}
}
