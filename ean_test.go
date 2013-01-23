package ean

import "testing"
import "errors"

func TestValid(t *testing.T) {
	if ean := ""; Valid(ean) {
		t.Errorf("Valid(%v) should be false, was true", ean)
	}

	if ean := "abc"; Valid(ean) {
		t.Errorf("Valid(%v) should be false, was true", ean)
	}

	if ean := "9781934356739"; !Valid(ean) {
		t.Errorf("Valid(%v) should be true, was false", ean)
	}

	if ean := "1111111111111"; Valid(ean) {
		t.Errorf("Valid(%v) should be false, was true", ean)
	}

	if ean := "978193435673912123123"; Valid(ean) {
		t.Errorf("Valid(%v) should be false, was true", ean)
	}
}

func TestChecksumEan8(t *testing.T) {
	assertChecksum8(t, "0", -1, errors.New("Ean 0 is too short to compute a checksum."))
	assertChecksum8(t, "0x111111", -1, errors.New("Contains non-digit: 'x'."))
	assertChecksum8(t, "96385074", 4, nil) // Wikipedia EAN-8 example
	assertChecksum8(t, "73513537", 7, nil)
}

func assertChecksum8(t *testing.T, ean string, expectedChecksum int, err error) {
	x, e := ChecksumEan8(ean)

	if e != nil && err != nil && e.Error() != err.Error() {
		t.Errorf("Checksum(%v) returned error %v, want %v", ean, e, err)
	}

	if x != expectedChecksum {
		t.Errorf("Checksum(%v) = %v, want %v", ean, x, expectedChecksum)
	}
}

func TestChecksumEan13(t *testing.T) {
	assertChecksum13(t, "0", -1, errors.New("Ean 0 is too short to compute a checksum."))
	assertChecksum13(t, "0x11111111111", -1, errors.New("Contains non-digit: 'x'."))
	assertChecksum13(t, "9781934356739", 9, nil)
	assertChecksum13(t, "1111111111116", 6, nil)
	assertChecksum13(t, "6291041500213", 3, nil) // GS1 example.
	assertChecksum13(t, "9780306406157", 7, nil) // Wikipedia ISBN-13 example
	assertChecksum13(t, "5711489018800", 0, nil)
	assertChecksum13(t, "5711489018824", 4, nil)
}

func assertChecksum13(t *testing.T, ean string, expectedChecksum int, err error) {
	x, e := ChecksumEan13(ean)

	if e != nil && err != nil && e.Error() != err.Error() {
		t.Errorf("Checksum(%v) returned error %v, want %v", ean, e, err)
	}

	if x != expectedChecksum {
		t.Errorf("Checksum(%v) = %v, want %v", ean, x, expectedChecksum)
	}
}
