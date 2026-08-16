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
	// BUG: mutate a copy, never write back
	tmp := reflect.New(v.Type()).Elem()
	tmp.Set(v)
	f := tmp.FieldByName(name)
	if !f.IsValid() || !f.CanSet() {
		return fmt.Errorf("field %s", name)
	}
	f.Set(reflect.Zero(f.Type()))
	return nil
}
