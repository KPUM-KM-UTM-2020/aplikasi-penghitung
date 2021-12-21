package helper

import (
	"testing"
)

func TestNimToEmail(t *testing.T) {
	param := "123456789012"
	expected := "123456789012@student.trunojoyo.ac.id"
	found, err := NimToEmail(param)
	if err != nil {
		t.Errorf("Should not be error on params '%s': %s", param, err)
	} else {
		if found != expected {
			t.Errorf("On input '%s', found '%s', expected '%s'", param, found, expected)
		}
	}
}

func TestJenisEmail(t *testing.T) {
	params := []string{
		"123456789012@trunojoyo.ac.id",
		"123456789012@student.trunojoyo.ac.id",
		"saya@trunojoyo.ac.id",
		"saya@student.trunojoyo.ac.id",
		"mereka@gmail.com",
		"mereka@",
	}
	expecteds := []string{
		"bukan student",
		"student",
		"bukan student",
		"student",
		"bukan student",
		"bukan student",
	}
	for i := 0; i < len(params); i++ {
		v := params[i]
		expected := expecteds[i]

		found, err := JenisEmail(v)
		if err != nil {
			t.Errorf("Should not be error on params '%s': %s", v, err)
		} else {
			if found != expected {
				t.Errorf("On input '%s', found '%s', expected '%s'", v, found, expected)
			}
		}
	}
}

func TestJenisEmailError(t *testing.T) {
	params := []string{
		"",
		"A",
		"Ab",
		"Ab.com",
	}
	for _, v := range params {
		_, err := JenisEmail(v)
		if err == nil {
			t.Errorf("Should be returning error on params '%s'", v)
		}
	}
}

func TestSingkatanFakultas(t *testing.T) {
	params := []string{
		"Fakultas Hukum",
		"Fakultas Ekonomi dan Bisnis",
		"Fakultas Pertanian",
		"Fakultas Teknik",
		"Fakultas Ilmu Sosial dan Ilmu Budaya",
		"Fakultas Ilmu Pendidikan",
		"Fakultas Keislaman",
	}
	expecteds := []string{
		"FH",
		"FEB",
		"FP",
		"FT",
		"FISIB",
		"FIP",
		"FKis",
	}

	for i := 0; i < len(params); i++ {
		v := params[i]
		expected := expecteds[i]

		found, err := SingkatanFakultas(v)
		if err != nil {
			t.Errorf("Should not be error on params '%s': %s", v, err)
		} else {
			if found != expected {
				t.Errorf("On input '%s', found '%s', expected '%s'", v, found, expected)
			}
		}
	}
}

func TestSingkatanFakultasError(t *testing.T) {
	params := []string{
		"Any fakultas",
	}
	for _, v := range params {
		_, err := SingkatanFakultas(v)
		if err == nil {
			t.Errorf("Should be returning error on params '%s'", v)
		}
	}
}
