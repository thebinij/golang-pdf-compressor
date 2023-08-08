// options-utils.go

package pdfcompresser

import (
	"fmt"
	"strconv"
)

func getDPIValue(dpiStr string) (string, error) {
	if dpiStr != "" {
		dpiValue, err := strconv.Atoi(dpiStr)
		if err != nil || dpiValue < 25 || dpiValue > 600 {
			return "", fmt.Errorf("invalid DPI value: %s", dpiStr)
		}
	}
	return dpiStr, nil
}

func getGrayscaleValue(grayscaleStr string) (bool, error) {
	if grayscaleStr == "" {
		return false, nil // Default to false
	}
	grayscale, err := strconv.ParseBool(grayscaleStr)
	if err != nil {
		return false, err
	}
	return grayscale, nil
}
