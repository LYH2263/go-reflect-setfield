package reflx

import "testing"

type User struct {
	Name string
}

func TestZeroName(t *testing.T) {
	u := &User{Name: "alice"}
	if err := ZeroField(u, "Name"); err != nil {
		t.Fatal(err)
	}
	if u.Name != "" {
		t.Fatalf("got %q", u.Name)
	}
	var v User
	if err := ZeroField(v, "Name"); err == nil {
		t.Fatal("want error for non-pointer")
	}
}
