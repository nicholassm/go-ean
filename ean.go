package ean

import (
	"fmt"
	"strconv"
)

func Valid(ean string) bool {
	return ValidEan8(ean) || ValidEan13(ean)
}

func ValidEan8(ean string) bool {
	if len(ean) != 8 {
		return false
	}

	checksum, err := ChecksumEan8(ean)

	return err == nil && strconv.Itoa(checksum) == ean[7:8]
}

func ValidEan13(ean string) bool {
	if len(ean) != 13 {
		return false
	}

	checksum, err := ChecksumEan13(ean)

	return err == nil && strconv.Itoa(checksum) == ean[12:13]
}

func ChecksumEan8(ean string) (int, error) {
	if len(ean) < 7 {
		return -1, fmt.Errorf("Ean %v is too short to compute a checksum.", ean)
	}

	return checksum(ean[:7], true)
}

func ChecksumEan13(ean string) (int, error) {
	if len(ean) < 12 {
		return -1, fmt.Errorf("Ean %v is too short to compute a checksum.", ean)
	}

	return checksum(ean[:12], false)
}

func checksum(ean string, multiplyWhenEven bool) (int, error) {
	sum := 0

	for i, v := range ean {
		value, err := strconv.Atoi(string(v))

		if err != nil {
			return -1, fmt.Errorf("Contains non-digit: %q.", v)
		}

		if (i%2 == 0) == multiplyWhenEven {
			sum += 3 * value
		} else {
			sum += value
		}
	}

	return (10 - sum%10) % 10, nil
}
