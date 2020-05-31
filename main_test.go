package binder

import (
	"testing"
)

func TestFloat(t *testing.T) {
	var (
		f1  float32
		f2  float64
		err error
	)

	binder := NewBinder()
	err = binder.Bind(&f1, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing float32 with 10. %v", err)
	}

	err = binder.Bind(&f2, "20.3")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing float64 with 20.3. %v", err)
	}

	err = binder.Bind(&f2, "abc")
	if err == nil {
		t.Errorf("Should have error when parsing float64 with abc. %v", err)
	}
}

func TestInt(t *testing.T) {

	var (
		i1  int
		i2  int8
		i3  int16
		i4  int32
		i5  int64
		err error
	)

	binder := NewBinder()
	err = binder.Bind(&i1, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. %v", err)
	}

	err = binder.Bind(&i1, "20.3")
	if err == nil {
		t.Errorf("Should have error when parsing int with 20.3. %v", err)
	}

	err = binder.Bind(&i1, "abc")
	if err == nil {
		t.Errorf("Should have error when parsing int with abc. %v", err)
	}

	err = binder.Bind(&i2, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. %v", err)
	}

	err = binder.Bind(&i3, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. %v", err)
	}

	err = binder.Bind(&i4, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. %v", err)
	}

	err = binder.Bind(&i5, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. %v", err)
	}
}
