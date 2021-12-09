package decoder

import (
	"errors"
	"fmt"

	"day-8/internal/util"
)

// DigitCount number of different digits.
const DigitCount = 10

// DigitDecoder contains digit encoding table and is used to decode encoded digits.
type DigitDecoder struct {
	encodingTable [DigitCount]string
}

// NewDigitDecoder deduces encoding from the encoded digits and creates a DigitDecoder with deduced encoding table.
func NewDigitDecoder(encodedDigits []string) (*DigitDecoder, error) {
	if len(encodedDigits) != DigitCount {
		return nil, errors.New(fmt.Sprintf("expected %d encoded digits, but got %d", DigitCount, len(encodedDigits)))
	}

	dd := DigitDecoder{}

	// Handle cases digits that can be identified by length
	for _, encodedDigit := range encodedDigits {
		switch len(encodedDigit) {
		case 2:
			dd.encodingTable[1] = util.StringSort(encodedDigit)
		case 3:
			dd.encodingTable[7] = util.StringSort(encodedDigit)
		case 4:
			dd.encodingTable[4] = util.StringSort(encodedDigit)
		case 7:
			dd.encodingTable[8] = util.StringSort(encodedDigit)
		}
	}

	// Handle remaining cases.
	for _, encodedDigit := range encodedDigits {
		switch len(encodedDigit) {
		case 5:
			// Only number 2, 3, 5 are encoded with 5 signals.
			normEncodedString := util.StringSort(encodedDigit)

			// If encoded string contains all signals of 1 it can only be 3.
			if util.StringContainsAll(normEncodedString, dd.encodingTable[1]) {
				dd.encodingTable[3] = normEncodedString
				break
			}

			// We can only consider 2 and 5 from here on.
			// If encoded string contains exactly 3 of 4 signals of number 4 it can only be 5.
			if util.StringCountContainedChars(normEncodedString, dd.encodingTable[4]) == 3 {
				dd.encodingTable[5] = normEncodedString
				break
			}

			// Only 2 remains.
			dd.encodingTable[2] = normEncodedString
		case 6:
			// Only number 0, 6, 9 are encoded with 6 signals.
			normEncodedString := util.StringSort(encodedDigit)

			// If encoded string does not contain all signals of 1 it can only be 6.
			if !util.StringContainsAll(normEncodedString, dd.encodingTable[1]) {
				dd.encodingTable[6] = normEncodedString
				break
			}

			// We can only consider 0 and 9 from here on.
			// If encoded string contain all signals of 4 it can only be 9.
			if util.StringContainsAll(normEncodedString, dd.encodingTable[4]) {
				dd.encodingTable[9] = normEncodedString
				break
			}

			dd.encodingTable[0] = normEncodedString
		}
	}

	// Validate that we composed complete encoding table.
	for _, entry := range dd.encodingTable {
		if len(entry) == 0 {
			return nil, errors.New("failed to compose complete encoding table from input encoded digits")
		}
	}

	return &dd, nil
}

// Decode decodes provided encoded digit.
func (dd *DigitDecoder) Decode(encodedDigit string) (int, error) {
	normEncodedString := util.StringSort(encodedDigit)

	for i, encoding := range dd.encodingTable {
		if normEncodedString == encoding {
			return i, nil
		}
	}

	return -1, errors.New("could not decode digit")
}
