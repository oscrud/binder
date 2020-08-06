package binder

import (
	"testing"
	"time"
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

	if f1 != 10 {
		t.Errorf("Expect f1 to be 10 after bind w/o error. received %v", f1)
	}

	err = binder.Bind(&f2, "20.3")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing float64 with 20.3. %v", err)
	}

	if f2 != 20.3 {
		t.Errorf("Expect f2 to be 20.3 after bind w/o error. received %v", f2)
	}

	err = binder.Bind(&f2, "abc")
	if err == nil {
		t.Errorf("Should have error when parsing float64 with abc")
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
		t.Errorf("Shouldn't have error when parsing int with 10. received %v", err)
	}

	if i1 != 10 {
		t.Errorf("Expect i1 to be 10 after bind w/o error. received %v", i1)
	}

	err = binder.Bind(&i1, "20.3")
	if err == nil {
		t.Errorf("Should have error when parsing int with 20.3")
	}

	err = binder.Bind(&i1, "abc")
	if err == nil {
		t.Errorf("Should have error when parsing int with abc")
	}

	err = binder.Bind(&i2, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. %v", err)
	}

	if i2 != 10 {
		t.Errorf("Expect i2 to be 10 after bind w/o error. received %v", i2)
	}

	err = binder.Bind(&i3, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. %v", err)
	}

	if i3 != 10 {
		t.Errorf("Expect i3 to be 10 after bind w/o error. received %v", i3)
	}

	err = binder.Bind(&i4, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. %v", err)
	}

	if i4 != 10 {
		t.Errorf("Expect i4 to be 10 after bind w/o error. received %v", i4)
	}

	err = binder.Bind(&i5, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. %v", err)
	}

	if i5 != 10 {
		t.Errorf("Expect i5 to be 10 after bind w/o error. received %v", i5)
	}
}

func TestUInt(t *testing.T) {
	var (
		i1  uint
		i2  uint8
		i3  uint16
		i4  uint32
		i5  uint64
		err error
	)

	binder := NewBinder()
	err = binder.Bind(&i1, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. received %v", err)
	}

	if i1 != 10 {
		t.Errorf("Expect i1 to be 10 after bind w/o error. received %v", i1)
	}

	err = binder.Bind(&i1, "20.3")
	if err == nil {
		t.Errorf("Should have error when parsing int with 20.3")
	}

	err = binder.Bind(&i1, "abc")
	if err == nil {
		t.Errorf("Should have error when parsing int with abc")
	}

	err = binder.Bind(&i2, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. %v", err)
	}

	if i2 != 10 {
		t.Errorf("Expect i2 to be 10 after bind w/o error. received %v", i2)
	}

	err = binder.Bind(&i3, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. %v", err)
	}

	if i3 != 10 {
		t.Errorf("Expect i3 to be 10 after bind w/o error. received %v", i3)
	}

	err = binder.Bind(&i4, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. %v", err)
	}

	if i4 != 10 {
		t.Errorf("Expect i4 to be 10 after bind w/o error. received %v", i4)
	}

	err = binder.Bind(&i5, "10")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing int with 10. %v", err)
	}

	if i5 != 10 {
		t.Errorf("Expect i5 to be 10 after bind w/o error. received %v", i5)
	}
}

func TestString(t *testing.T) {
	var (
		str string
		err error
	)

	binder := NewBinder()
	err = binder.Bind(&str, "string_data")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing string with string_data. received %v", err)
	}

	if str != "string_data" {
		t.Errorf("Expect str to be string_data after bind w/o error. received %v", str)
	}
}

func TestBoolean(t *testing.T) {
	var (
		data bool
		err  error
	)

	binder := NewBinder()
	err = binder.Bind(&data, "true")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing boolean with true. received %v", err)
	}

	if data != true {
		t.Errorf("Expect data to be true after bind w/o error. received %v", data)
	}

	err = binder.Bind(&data, "false")
	if err != nil {
		t.Errorf("Shouldn't have error when parsing boolean with false. received %v", err)
	}

	if data != false {
		t.Errorf("Expect data to be false after bind w/o error. received %v", data)
	}

	err = binder.Bind(&data, 0)
	if err != nil {
		t.Errorf("Shouldn't have error when parsing boolean with 0. received %v", err)
	}

	if data != false {
		t.Errorf("Expect data to be false after bind w/o error. received %v", data)
	}

	err = binder.Bind(&data, 1)
	if err != nil {
		t.Errorf("Shouldn't have error when parsing boolean with 1. received %v", err)
	}

	if data != true {
		t.Errorf("Expect data to be true after bind w/o error. received %v", data)
	}

	err = binder.Bind(&data, "ansadjasod")
	if err == nil {
		t.Errorf("Should have error when parsing invalid bool value")
	}
}

func TestSlice(t *testing.T) {
	var (
		data []string
		err  error
	)

	other := []string{"1", "2"}
	otherInts := []int{1, 2}

	binder := NewBinder()
	err = binder.Bind(&data, other)
	if err != nil {
		t.Errorf("Shouldn't have error when parsing same slice type. received %v", err)
	}

	if data[0] != "1" && data[1] != "2" {
		t.Errorf("Expect data to be same after bind w/o error. received %v", data)
	}

	err = binder.Bind(&data, otherInts)
	if err == nil {
		t.Errorf("Should have error when parsing different slice type")
	}
}

func TestArray(t *testing.T) {
	var (
		data [2]string
		err  error
	)

	other := [2]string{"1", "2"}
	otherInts := [2]int{1, 2}

	binder := NewBinder()
	err = binder.Bind(&data, other)
	if err != nil {
		t.Errorf("Shouldn't have error when parsing same array type. received %v", err)
	}

	if data[0] != "1" && data[1] != "2" {
		t.Errorf("Expect data to be same after bind w/o error. received %v", data)
	}

	err = binder.Bind(&data, otherInts)
	if err == nil {
		t.Errorf("Should have error when parsing different array type")
	}
}
func TestMap(t *testing.T) {
	var (
		data map[string]string
		err  error
	)

	other := map[string]string{"1": "1", "2": "2"}
	otherMaps := map[int]string{1: "1", 2: "2"}

	binder := NewBinder()
	err = binder.Bind(&data, other)
	if err != nil {
		t.Errorf("Shouldn't have error when parsing same map type. received %v", err)
	}

	if data["1"] != "1" && data["2"] != "2" {
		t.Errorf("Expect data to be same after bind w/o error. received %v", data)
	}

	err = binder.Bind(&data, otherMaps)
	if err == nil {
		t.Errorf("Should have error when parsing different map type")
	}
}

func TestStruct(t *testing.T) {

	type ts struct {
		Data string
	}

	type ts2 struct {
		Data string
	}

	var (
		data ts
		err  error
	)

	other := ts{Data: "1"}
	other2 := ts2{Data: "1"}

	binder := NewBinder()
	err = binder.Bind(&data, other)
	if err != nil {
		t.Errorf("Shouldn't have error when parsing same struct type. received %v", err)
	}

	if data.Data != "1" {
		t.Errorf("Expect data to be same after bind w/o error. received %v", data)
	}

	err = binder.Bind(&data, other2)
	if err == nil {
		t.Errorf("Should have error when parsing different struct type")
	}
}

func TestTypeName(t *testing.T) {
	t1 := string("")
	t1Name := getTypeName(t1)
	if t1Name != "string" {
		t.Errorf("Should type name be string. received %v", t1Name)
	}

	t2 := time.Time{}
	t2Name := getTypeName(t2)
	if t2Name != "time.Time" {
		t.Errorf("Should type name be time.Time. received %v", t2Name)
	}

	t3 := []time.Time{}
	t3Name := getTypeName(t3)
	if t3Name != "slice$time.Time" {
		t.Errorf("Should type name be slice$time.Time. received %v", t2Name)
	}

	t4 := []string{}
	t4Name := getTypeName(t4)
	if t4Name != "slice$string" {
		t.Errorf("Should type name be slice$string. received %v", t4Name)
	}

	t5 := [2]string{}
	t5Name := getTypeName(t5)
	if t5Name != "array2$string" {
		t.Errorf("Should type name be array2$string. received %v", t5Name)
	}

	t6 := [2]time.Time{}
	t6Name := getTypeName(t6)
	if t6Name != "array2$time.Time" {
		t.Errorf("Should type name be array2$string. received %v", t6Name)
	}

	t7 := new([2]time.Time)
	t7Name := getTypeName(t7)
	if t7Name != "array2$time.Time" {
		t.Errorf("Should type name be array2$string. received %v", t6Name)
	}
}

func TestRegister(t *testing.T) {
	binder := NewBinder()

	fromType := string("")
	fromTypeName := getTypeName(fromType)
	toType := time.Time{}
	toTypeName := getTypeName(toType)

	binder.Register(
		fromType, toType,
		func(raw interface{}) (interface{}, error) {
			text := raw.(string)
			return time.Parse("2006-01-02", text)
		},
	)

	if _, ok := binder.binder[fromTypeName+","+toTypeName]; !ok {
		t.Errorf("Should have type registered in binder instance")
	}

	var (
		data time.Time
		err  error
	)

	err = binder.Bind(&data, "2020-08-01")
	if err != nil {
		t.Errorf("Shouldn't have error when bind custom correct value. received %v", err)
	}

	if data.Year() != 2020 || data.Month() != 8 || data.Day() != 01 {
		t.Errorf("Expect data to be same after bind w/o error. received %v", data)
	}

	err = binder.Bind(&data, "asdoaskdok")
	if err == nil {
		t.Errorf("Should return error when bind custom invalid value")
	}
}

func TestError(t *testing.T) {
	binder := NewBinder()

	var (
		data  int
		idata interface{}
		err   error
	)

	err = binder.Bind(data, 1)
	if err != ErrSourceNotAddressable {
		t.Errorf("Should return error if assignee is not address")
	}

	err = binder.Bind(&idata, 1)
	if err != ErrSourceCantBeInterface {
		t.Errorf("Should return error if assignee is private")
	}

	err = binder.Bind(nil, nil)
	if err != ErrSourceAndAssignCantBeNil {
		t.Errorf("Should return error if assignee or source is nil")
	}
}
