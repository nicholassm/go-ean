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
	assertChecksum(t, ChecksumEan8, "0", -1, errors.New("incorrect ean 0 to compute a checksum"))
	assertChecksum(t, ChecksumEan8, "0x111111", -1, errors.New("contains non-digit: 'x'"))
	assertChecksum(t, ChecksumEan8, "96385074", 4, nil) // Wikipedia EAN-8 example
	assertChecksum(t, ChecksumEan8, "73513537", 7, nil)
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
	assertChecksum(t, ChecksumEan13, "0", -1, errors.New("incorrect ean 0 to compute a checksum"))
	assertChecksum(t, ChecksumEan13, "00000000000000", -1, errors.New("incorrect ean 00000000000000 to compute a checksum"))
	assertChecksum(t, ChecksumEan13, "0x11111111111", -1, errors.New("contains non-digit: 'x'"))
	assertChecksum(t, ChecksumEan13, "9781934356739", 9, nil)
	assertChecksum(t, ChecksumEan13, "1111111111116", 6, nil)
	assertChecksum(t, ChecksumEan13, "6291041500213", 3, nil) // GS1 example.
	assertChecksum(t, ChecksumEan13, "9780306406157", 7, nil) // Wikipedia ISBN-13 example
	assertChecksum(t, ChecksumEan13, "5711489018800", 0, nil)
	assertChecksum(t, ChecksumEan13, "5711489018824", 4, nil)
}

func TestChecksumUPC(t *testing.T) {
	assertChecksum(t, ChecksumUpc, "012345678905", 5, nil)
}
