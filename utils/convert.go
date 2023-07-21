package utils

import (
	"errors"
	"reflect"
)

func ToInt(it any) (int, error) {
	itrv, err := toValue(it)
	if err == nil {
		if itrv.CanInt() {
			return int(itrv.Int()), nil
		}
		err = errors.New("not an 'int', but was a: " + itrv.Kind().String())
	}
	return 0, err
}

//goland:noinspection GoUnusedExportedFunction
func ToUint(it any) (uint, error) {
	itrv, err := toValue(it)
	if err == nil {
		if itrv.CanUint() {
			return uint(itrv.Uint()), nil
		}
		err = errors.New("not an 'uint', but was a: " + itrv.Kind().String())
	}
	return 0, err
}

func toValue(it any) (itrv reflect.Value, err error) {
	if it == nil {
		err = errors.New("was nil")
	} else {
		itrv = reflect.ValueOf(it)
		if !itrv.IsValid() {
			err = errors.New("was IsValid")
		}
	}
	return
}
