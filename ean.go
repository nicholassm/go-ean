package ean

import (
	"fmt"
	"strconv"
)

func Valid(ean string) bool {
	return ValidEan8(ean) || ValidEan13(ean) || ValidUpc(ean)
}

func ValidUpc(upc string) bool {
	return validCode(upc, 12)

}

func ValidEan8(ean string) bool {
	return validCode(ean, 8)
}

func ValidEan13(ean string) bool {
	return validCode(ean, 13)
}

func validCode(ean string, size int) bool {
	checksum, err := checksum(ean, size)

	return err == nil && strconv.Itoa(checksum) == ean[size-1:size]
}

func ChecksumEan8(ean string) (int, error) {
	return checksum(ean, 8)
}

func ChecksumEan13(ean string) (int, error) {
	return checksum(ean, 13)
}

func ChecksumUpc(upc string) (int, error) {
	return checksum(upc, 12)
}

func checksum(ean string, size int) (int, error) {
	if len(ean) != size {
		return -1, fmt.Errorf("incorrect ean %v to compute a checksum", ean)
	}

	code := ean[:size-1]
	multiplyWhenEven := size%2 == 0
	sum := 0

	for i, v := range code {
		value, err := strconv.Atoi(string(v))

		if err != nil {
			return -1, fmt.Errorf("contains non-digit: %q", v)
		}

		if (i%2 == 0) == multiplyWhenEven {
			sum += 3 * value
		} else {
			sum += value
		}
	}

	return (10 - sum%10) % 10, nil
}
