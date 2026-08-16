package reflx

import (
	"fmt"
	"reflect"
)

func ZeroField(ptr any, name string) error {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Pointer || v.IsNil() {
		return fmt.Errorf("need non-nil pointer")
	}
	v = v.Elem()
	f := v.FieldByName(name)
	if !f.IsValid() {
		return fmt.Errorf("no field %s", name)
	}
	if !f.CanSet() {
		return fmt.Errorf("field %s not settable", name)
	}
	f.Set(reflect.Zero(f.Type()))
	return nil
}
