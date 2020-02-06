package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// Error Message
var (
	ErrSourceNotAddressable = errors.New("binder source must be addressable")
	ErrSourceNotUpdatable   = errors.New("binder source unable to set new value")
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
		return "array$" + typ.Elem().Name()
	} else if typ.Kind() == reflect.Slice {
		return "slice$" + typ.Elem().Name()
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
	typ := reflect.TypeOf(assign)
	if typ.Kind() != reflect.Ptr {
		return ErrSourceNotAddressable
	}

	field := reflect.ValueOf(assign).Elem()
	if !field.CanSet() {
		return ErrSourceNotUpdatable
	}

	sederName := getTypeName(value) + "," + getTypeName(field.Interface())
	switch field.Type().Kind() {
	case reflect.Float32, reflect.Float64:
		if seder, ok := b.binder[sederName]; ok {
			val, err := seder(value)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(val))
			break
		}
		bit := field.Type().Bits()
		str := fmt.Sprintf("%v", value)
		result, err := strconv.ParseFloat(str, bit)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to float%d", value, bit)
		}
		field.SetFloat(result)
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if seder, ok := b.binder[sederName]; ok {
			val, err := seder(value)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(val))
			break
		}
		bit := field.Type().Bits()
		str := fmt.Sprintf("%v", value)
		result, err := strconv.ParseUint(str, 10, bit)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to uint%d", value, bit)
		}
		field.SetUint(result)
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if seder, ok := b.binder[sederName]; ok {
			val, err := seder(value)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(val))
			break
		}
		bit := field.Type().Bits()
		str := fmt.Sprintf("%v", value)
		result, err := strconv.ParseInt(str, 10, bit)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to int%d", value, bit)
		}
		field.SetInt(result)
		break
	case reflect.String:
		if seder, ok := b.binder[sederName]; ok {
			val, err := seder(value)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(val))
			break
		}
		result := fmt.Sprintf("%v", value)
		field.SetString(result)
		break
	case reflect.Bool:
		if seder, ok := b.binder[sederName]; ok {
			val, err := seder(value)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(val))
			break
		}
		str := fmt.Sprintf("%v", value)
		result, err := strconv.ParseBool(str)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to bool", value)
		}
		field.SetBool(result)
		break
	case reflect.Slice:
		if binder, ok := b.binder[sederName]; ok {
			deserialized, err := binder(value)
			if err != nil {
				return fmt.Errorf("Trying to deserialize %v to %v, %v", value, field.Type().Name(), err)
			}
			field.Set(reflect.ValueOf(deserialized))
		} else {
			qt := reflect.TypeOf(value)
			if !field.Type().AssignableTo(qt) {
				return fmt.Errorf("Trying to convert %v to slice %v", value, field.Type().Elem().Name())
			}
			field.Set(reflect.ValueOf(value))
		}
		break
	case reflect.Array:
		if binder, ok := b.binder[sederName]; ok {
			deserialized, err := binder(value)
			if err != nil {
				return fmt.Errorf("Trying to deserialize %v to %v, %v", value, field.Type().Name(), err)
			}
			field.Set(reflect.ValueOf(deserialized))
		} else {
			qt := reflect.TypeOf(value)
			if !field.Type().AssignableTo(qt) {
				return fmt.Errorf("Trying to convert %v to array %v", value, field.Type().Elem().Name())
			}
			field.Set(reflect.ValueOf(value))
		}
		break
	case reflect.Struct:
		if binder, ok := b.binder[sederName]; ok {
			deserialized, err := binder(value)
			if err != nil {
				return fmt.Errorf("Trying to deserialize %v to %v, %v", value, field.Type().Name(), err)
			}
			field.Set(reflect.ValueOf(deserialized))
		} else {
			qt := reflect.TypeOf(value)
			if !field.Type().AssignableTo(qt) {
				return fmt.Errorf("Trying to convert %v to struct %v", value, field.Type().Name())
			}
			field.Set(reflect.ValueOf(value))
		}
		break
	default:
		qt := reflect.TypeOf(value)
		if !field.Type().AssignableTo(qt) {
			return fmt.Errorf("Trying to convert %v to %v", value, field.Addr().Type())
		}
		field.Set(reflect.ValueOf(value))
		break
	}

	return nil
}
