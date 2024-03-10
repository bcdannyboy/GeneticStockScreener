package utils

import (
	"fmt"
	"reflect"
)

// IsZeroValue checks if the provided struct is in its zero value state.
// This function uses reflection to work with any arbitrary struct.
func IsZeroValue(s interface{}) bool {
	val := reflect.ValueOf(s)

	// Check if the passed interface is a slice
	if val.Kind() == reflect.Slice {
		// Handle slice of interfaces
		if val.Len() == 0 {
			// Slice is empty
			return true
		}

		// Check if all elements in the slice are zero values
		allElementsZero := true
		for i := 0; i < val.Len(); i++ {
			element := val.Index(i)
			if !reflect.DeepEqual(element.Interface(), reflect.Zero(element.Type()).Interface()) {
				allElementsZero = false
				break // Found an element that is not its zero value
			}
		}
		return allElementsZero
	} else if val.Kind() == reflect.Struct {
		// Handle structs
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if !reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
				return false // Found a field that is not its zero value
			}
		}
		// All fields are their zero value
		return true
	} else {
		fmt.Printf("error, expected struct or slice, got %v\n", val.Kind())
		return false
	}
}
