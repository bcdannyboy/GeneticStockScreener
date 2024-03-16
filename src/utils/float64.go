package utils

import (
	"math/rand"
	"reflect"
	"time"
)

func SetRandomFloats(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		rand.Seed(time.Now().UnixNano())
		field := v.Field(i)

		switch field.Kind() {
		case reflect.Struct:
			SetRandomFloats(field)
		case reflect.Float64:
			field.SetFloat(rand.Float64()*2 - 1) // Random float64 between -1 and 1
		}
	}
}
