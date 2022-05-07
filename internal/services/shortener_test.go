package services

import "testing"

func TestGenerateCode(t *testing.T) {
	got, err := generateCode()
	if err != nil {
		t.Fatalf("Error while generating code: %q", err)
	}
	if len(got) != 10 {
		t.Logf("Generated code length is not 10: %q", got)
	}
	allowed := make(map[rune]struct{})
	for _, r := range alphabet {
		allowed[r] = struct{}{}
	}
	for _, r := range got {
		if _, ok := allowed[r]; !ok {
			t.Logf("Generated code contains not allowed chars: %q", r)
		}
	}
}
