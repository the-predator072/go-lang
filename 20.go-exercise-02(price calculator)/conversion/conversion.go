package conversion

import (
	"strconv"
)

func StrigToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for index, strVal := range strings {
		floatPrice, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return nil, err
		}
		floats[index] = floatPrice
	}
	return floats, nil
}
