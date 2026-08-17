package reflx

import (
	"fmt"
	"reflect"
)

func ZeroField(ptr any, name string) error {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Pointer {
		return fmt.Errorf("reflx: ZeroField requires a pointer, got %s", v.Kind())
	}
	if v.IsNil() {
		return fmt.Errorf("reflx: ZeroField requires a non-nil pointer")
	}
	// v.Elem() is the value behind the pointer: it is addressable and
	// settable, so mutating its fields writes back to the caller's object.
	// Do NOT copy into a temporary value — that would discard the changes.
	v = v.Elem()
	f := v.FieldByName(name)
	if !f.IsValid() {
		return fmt.Errorf("reflx: field %q not found in %s", name, v.Type())
	}
	if !f.CanSet() {
		return fmt.Errorf("reflx: field %q of %s is not settable", name, v.Type())
	}
	f.Set(reflect.Zero(f.Type()))
	return nil
}
