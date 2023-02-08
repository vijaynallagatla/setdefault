package setdefault

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	StringField     string  `default:"default string value"`
	IntField        int     `default:"123"`
	Int8Field       int8    `default:"123"`
	Int16Field      int16   `default:"123"`
	Int32Field      int32   `default:"123"`
	Int64Field      int64   `default:"123"`
	UintField       uint    `default:"123"`
	Uint8Field      uint8   `default:"123"`
	Uint16Field     uint16  `default:"123"`
	Uint32Field     uint32  `default:"123"`
	Uint64Field     uint64  `default:"123"`
	Float32Field    float32 `default:"123.456"`
	Float64Field    float64 `default:"123.456"`
	BoolField       bool    `default:"true"`
	InvalidType     int     `default:"invalid value"`
	UnsettableField int     `default:"123"`
}

func TestBindDefault(t *testing.T) {
	testStruct := TestStruct{UnsettableField: 123}
	reflect.ValueOf(&testStruct).Elem().FieldByName("UnsettableField").Set(reflect.ValueOf(0).Convert(reflect.TypeOf(testStruct.UnsettableField)))

	Bind(&testStruct)

	if testStruct.StringField != "default string value" {
		t.Errorf("Expected default value of 'default string value', but got '%s'", testStruct.StringField)
	}
	if testStruct.IntField != 123 {
		t.Errorf("Expected default value of 123, but got '%d'", testStruct.IntField)
	}
	if testStruct.Int8Field != 123 {
		t.Errorf("Expected default value of 123, but got '%d'", testStruct.Int8Field)
	}
	if testStruct.Int16Field != 123 {
		t.Errorf("Expected default value of 123, but got '%d'", testStruct.Int16Field)
	}
	if testStruct.Int32Field != 123 {
		t.Errorf("Expected default value of 123, but got '%d'", testStruct.Int32Field)
	}
	if testStruct.Int64Field != 123 {
		t.Errorf("Expected default value of 123, but got '%d'", testStruct.Int64Field)
	}
	if testStruct.UintField != 123 {
		t.Errorf("Expected default value of 123, but got '%d'", testStruct.UintField)
	}
	if testStruct.Uint8Field != 123 {
		t.Errorf("Expected default value of 123, but got '%d'", testStruct.Uint8Field)
	}
	if testStruct.Uint16Field != 123 {
		t.Errorf("Expected Uint16Field to be 123, but got %v", testStruct.Uint16Field)
	}
	if testStruct.Uint32Field != 123 {
		t.Errorf("Expected Uint32Field to be 123, but got %v", testStruct.Uint32Field)
	}
	if testStruct.Uint64Field != 123 {
		t.Errorf("Expected Uint64Field to be 123, but got %v", testStruct.Uint64Field)
	}
	if testStruct.Float32Field != 123.456 {
		t.Errorf("Expected Float32Field to be 123.456, but got %v", testStruct.Float32Field)
	}
	if testStruct.Float64Field != 123.456 {
		t.Errorf("Expected Float32Field to be 123.456, but got %v", testStruct.Float64Field)
	}
	if testStruct.BoolField != true {
		t.Errorf("Expected BoolField != to be true, but got %v", testStruct.BoolField)
	}
	if testStruct.InvalidType != 0 {
		t.Errorf("Expected BoolField != to be 0, but got %v", testStruct.InvalidType)
	}
	if testStruct.UnsettableField != 123 {
		t.Errorf("Expected UnsettableField != to be 123, but got %v", testStruct.UnsettableField)
	}
}
