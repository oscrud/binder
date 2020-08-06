package binder

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// Error Message
var (
	ErrSourceNotAddressable     = errors.New("binder: source must be addressable")
	ErrSourceCantBeInterface    = errors.New("binder: source can't be interface")
	ErrSourceAndAssignCantBeNil = errors.New("binder: source and assigned can't be nil")
)

// Bind :
type Bind func(interface{}) (interface{}, error)

// Binder :
type Binder struct {
	binder map[string]Bind
}

// NewBinder :
func NewBinder() *Binder {
	return &Binder{
		binder: make(map[string]Bind, 0),
	}
}

func getTypeName(rtype interface{}) string {
	typ := reflect.TypeOf(rtype)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Array {
		arrayType := typ.Elem()
		len := fmt.Sprintf("%v", typ.Len())
		if arrayType.PkgPath() != "" {
			return "array" + len + "$" + arrayType.PkgPath() + "." + arrayType.Name()
		}
		return "array" + len + "$" + arrayType.Name()
	} else if typ.Kind() == reflect.Slice {
		sliceType := typ.Elem()
		if sliceType.PkgPath() != "" {
			return "slice$" + sliceType.PkgPath() + "." + sliceType.Name()
		}
		return "slice$" + sliceType.Name()
	}

	if typ.PkgPath() != "" {
		return typ.PkgPath() + "." + typ.Name()
	}
	return typ.Name()
}

// Register :
func (b *Binder) Register(from interface{}, to interface{}, bind Bind) *Binder {
	fromStr := getTypeName(from)
	toStr := getTypeName(to)
	b.binder[fromStr+","+toStr] = bind
	return b
}

// Bind :
func (b Binder) Bind(assign interface{}, value interface{}) error {
	if value == nil || assign == nil {
		return ErrSourceAndAssignCantBeNil
	}

	typ := reflect.TypeOf(assign)
	if typ.Kind() != reflect.Ptr {
		return ErrSourceNotAddressable
	}

	if typ.Elem().Kind() == reflect.Interface {
		return ErrSourceCantBeInterface
	}

	field := reflect.ValueOf(assign).Elem()
	sederName := getTypeName(value) + "," + getTypeName(field.Interface())
	if seder, ok := b.binder[sederName]; ok {
		val, err := seder(value)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(val))
		return nil
	}

	switch field.Type().Kind() {
	case reflect.Float32, reflect.Float64:
		bit := field.Type().Bits()
		str := fmt.Sprintf("%v", value)
		result, err := strconv.ParseFloat(str, bit)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to float%d", value, bit)
		}
		field.SetFloat(result)
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		bit := field.Type().Bits()
		str := fmt.Sprintf("%v", value)
		result, err := strconv.ParseUint(str, 10, bit)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to uint%d", value, bit)
		}
		field.SetUint(result)
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		bit := field.Type().Bits()
		str := fmt.Sprintf("%v", value)
		result, err := strconv.ParseInt(str, 10, bit)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to int%d", value, bit)
		}
		field.SetInt(result)
		break
	case reflect.String:
		result := fmt.Sprintf("%v", value)
		field.SetString(result)
		break
	case reflect.Bool:
		str := fmt.Sprintf("%v", value)
		result, err := strconv.ParseBool(str)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to bool", value)
		}
		field.SetBool(result)
		break
	case reflect.Slice:
		qt := reflect.TypeOf(value)
		if !field.Type().AssignableTo(qt) {
			return fmt.Errorf("Trying to convert %v to slice %v", value, field.Type().Elem().Name())
		}
		field.Set(reflect.ValueOf(value))
		break
	case reflect.Array:
		qt := reflect.TypeOf(value)
		if !field.Type().AssignableTo(qt) {
			return fmt.Errorf("Trying to convert %v to array %v", value, field.Type().Elem().Name())
		}
		field.Set(reflect.ValueOf(value))
		break
	case reflect.Struct:
		qt := reflect.TypeOf(value)
		if !field.Type().AssignableTo(qt) {
			return fmt.Errorf("Trying to convert %v to struct %v", value, field.Type().Name())
		}
		field.Set(reflect.ValueOf(value))
		break
	case reflect.Map:
		qt := reflect.TypeOf(value)
		if !field.Type().AssignableTo(qt) {
			return fmt.Errorf("Trying to convert %v to map %v", value, field.Type().Name())
		}
		field.Set(reflect.ValueOf(value))
		break
	}
	return nil
}
