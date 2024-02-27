package main

/*
- nama file harus diakhiri dengan _test.go
- sedangkan nama funcnya harus diawali dengan Test
*/

import (
	"testing"
)

func lulus(nilaiAkhir, absen int) bool {
	if nilaiAkhir > 80 && absen > 80 {
		return true
	} else {
		return false
	}
}

func TestKelulusan(t *testing.T) {
	resultKelulusan := lulus(85, 85)
	expected := true

	if resultKelulusan != expected {
		t.Errorf("resultKelulusan returned %t, expected %t", resultKelulusan, expected)
	}
}
