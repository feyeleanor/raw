package raw

import (
	"reflect"
	"testing"
)

func TestDereferenceAll(t *testing.T) {
	ConfirmDereference := func(l, r interface{}) {
		if v := DereferenceAll(l); v != r {
			t.Fatalf("%v: DereferenceAll(%v) != %v", v, l, r)
		}
	}

	i := -1
	ConfirmDereference(i, i)

	ConfirmDissimilar := func(i, v interface{}) {
		if i == v {
			t.Fatalf("%v == %v", i, v)
		}
	}

	ip1 := &i
	ConfirmDissimilar(ip1, i)
	ConfirmDereference(ip1, i)

	ip2 := &ip1
	ConfirmDissimilar(ip2, i)
	ConfirmDissimilar(ip2, ip1)
	ConfirmDereference(ip2, i)
	ConfirmDereference(ip2, DereferenceAll(ip1))

	ip3 := &ip2
	ConfirmDissimilar(ip3, i)
	ConfirmDissimilar(ip3, ip1)
	ConfirmDissimilar(ip3, ip2)
	ConfirmDereference(ip3, i)
	ConfirmDereference(ip3, DereferenceAll(ip1))
	ConfirmDereference(ip3, DereferenceAll(ip2))
}

func Test_shallowCopy(t *testing.T) {
	i := -1
	v := shallowCopy(reflect.ValueOf(i)).Interface()
	switch p, ok := v.(*int); {
	case !ok:
		t.Fatalf("reflect.TypeOf(%v) = %v", v, reflect.TypeOf(v))
	case *p != i:
		t.Fatalf("*(%v) != %v", p, i)
	}
}

func TestShallowCopy(t *testing.T) {
	i := -1
	v := ShallowCopy(i)
	switch p, ok := v.(*int); {
	case !ok:
		t.Fatalf("reflect.TypeOf(%v) = %v", v, reflect.TypeOf(v))
	case *p != i:
		t.Fatalf("*(%v) != %v", p, i)
	}
}

func TestMakeAddressable(t *testing.T) {
	AssertAddressable := func(v reflect.Value, r bool) {
		if v.CanAddr() != r {
			t.Fatalf("reflect.ValueOf(%v).CanAddr() should be %v", v, r)
		}
	}

	ConfirmMadeAddressable := func(v reflect.Value) {
		if a := MakeAddressable(v); a.Kind() != reflect.Ptr {
			t.Fatalf("MakeAddressable(%v) should return a pointer not a %v", v.Interface(), a.Kind())
		}
	}

	v := reflect.ValueOf(-1)
	AssertAddressable(v, false)
	ConfirmMadeAddressable(v)

	var i interface{} = -1
	v = reflect.ValueOf(i)
	AssertAddressable(v, false)
	ConfirmMadeAddressable(v)

	var p int
	v = reflect.ValueOf(p)
	AssertAddressable(v, false)
	ConfirmMadeAddressable(v)

	v = reflect.ValueOf(&p)
	AssertAddressable(v, false)
	AssertAddressable(reflect.Indirect(v), true)
	ConfirmMadeAddressable(v)
}

func TestDuplicate(t *testing.T) {
	t.Log("Test not yet implemented")
}