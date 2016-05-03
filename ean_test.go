package ean

import "testing"
import "errors"

func TestValid(t *testing.T) {
	invalids := []string{"", "abc", "1111111111111", "978193435673912123123"}
	for _, ean := range invalids {
		if Valid(ean) {
			t.Errorf("Valid(%v) should be false, was true", ean)
		}
	}

	valids := []string{"9781934356739", "0012345678905"}
	for _, ean := range valids {
		if !Valid(ean) {
			t.Errorf("Valid(%v) should be true, was false", ean)
		}
	}
}

func TestChecksumEan8(t *testing.T) {
	tests := []struct {
		ean      string
		expected int
		err      error
	}{
		{"0", -1, errors.New("incorrect ean 0 to compute a checksum")},
		{"0x111111", -1, errors.New("contains non-digit: 'x'")},
		{"96385074", 4, nil}, // Wikipedia EAN-8 example
		{"73513537", 7, nil},
	}
	for _, v := range tests {
		assertChecksum(t, ChecksumEan8, v.ean, v.expected, v.err)
	}
}

func assertChecksum(t *testing.T, f func(string) (int, error), ean string, expectedChecksum int, err error) {
	x, e := f(ean)

	if e != nil && err != nil && e.Error() != err.Error() {
		t.Errorf("Checksum(%v) returned error %v, want %v", ean, e, err)
	}

	if x != expectedChecksum {
		t.Errorf("Checksum(%v) = %v, want %v", ean, x, expectedChecksum)
	}
}

func TestChecksumEan13(t *testing.T) {
	tests := []struct {
		ean      string
		expected int
		err      error
	}{
		{"0", -1, errors.New("incorrect ean 0 to compute a checksum")},
		{"00000000000000", -1, errors.New("incorrect ean 00000000000000 to compute a checksum")},
		{"0x11111111111", -1, errors.New("contains non-digit: 'x'")},
		{"9781934356739", 9, nil},
		{"1111111111116", 6, nil},
		{"6291041500213", 3, nil}, // GS1 example.
		{"9780306406157", 7, nil}, // Wikipedia ISBN-13 example
		{"5711489018800", 0, nil},
		{"5711489018824", 4, nil},
	}
	for _, v := range tests {
		assertChecksum(t, ChecksumEan13, v.ean, v.expected, v.err)
	}
}

func TestChecksumUPC(t *testing.T) {
	assertChecksum(t, ChecksumUpc, "012345678905", 5, nil)
}
