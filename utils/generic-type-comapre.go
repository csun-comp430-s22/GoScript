package utils

import (
	"reflect"
	"strings"
)

func GenericTypeCompare(first any, second any) bool {
	firstType := reflect.TypeOf(first)
	secondType := reflect.TypeOf(second)

	cleanFirstTypeString := strings.Split(firstType.String(), "[")[0]
	cleanSecondTypeString := strings.Split(secondType.String(), "[")[0]

	return cleanFirstTypeString == cleanSecondTypeString
}
