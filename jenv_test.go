package jenv

import (
	"encoding/json"
	"os"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	j := DefaultConfig()
	filePathCorrect := j.FilePath == "env.json"
	initialDataCorrect := j.Data == nil
	if !filePathCorrect || !initialDataCorrect {
		t.Errorf("Default Config Produces Bad Output, filepath: %v, initdata: %v", filePathCorrect, initialDataCorrect)
	}
}

func TestLoad(t *testing.T) {
	j := DefaultConfig()
	err := j.Load()
	if err != nil && err.Error() != "File Not Found" {
		t.Errorf("It should produces error because no file is found")
	}
}

func TestLoadWithFile(t *testing.T) {
	j := DefaultConfig()

	f, _ := os.Create(j.FilePath)
	_, _ = f.Write([]byte(`
		{
			"data": ["1", "2", "3"]
		}
	`))
	f.Close()

	err := j.Load()
	if err != nil && err.Error() == "File Not Found" {
		t.Errorf("Error in file path lookup")
	} else if err != nil && err.Error() == "Decode Error" {
		t.Errorf("Error when decoding")
	}

	_ = os.Remove(j.FilePath)
}
func TestLoadFromString(t *testing.T) {
	j := Jenv{}
	err := j.LoadFromString(`
		{
			"data": ["1", "2", "3"]
		}
	`)
	if err != nil {
		t.Errorf("Decode should not produce error")
	}
}

func TestGetEmptyData(t *testing.T) {
	j := Jenv{}
	data, ok := j.Get("data")
	if data != nil || ok != false {
		t.Errorf("Get should return nil data and false boolean")
	}
}

func TestGetNullData(t *testing.T) {
	j := Jenv{}
	j.LoadFromString(`
		{
			"data": null
		}
	`)
	data, ok := j.Get("data")
	if data != nil || ok != true {
		t.Errorf("Get should return nil data and true boolean")
	}
}

func TestGetFilledData(t *testing.T) {
	j := Jenv{}
	_ = j.LoadFromString(`
		{
			"data": ["1", "2", "3"]
		}
	`)
	data, ok := j.Get("data")
	if ok == false || data == nil {
		t.Errorf("Data not loaded")
	}
}

func TestBoolean(t *testing.T) {
	j := Jenv{}
	_ = j.LoadFromString(`
		{
			"data": true
		}
	`)
	data, _ := j.Get("data")
	arr, err := Boolean(data)
	if err != nil {
		t.Errorf("It should not produce any error")
	}
	if arr != true {
		t.Errorf("Data parsed uncorrectly")
	}
}

func TestFloat64(t *testing.T) {
	j := Jenv{}
	_ = j.LoadFromString(`
		{
			"data": 3.14
		}
	`)
	data, _ := j.Get("data")
	arr, err := Float64(data)
	if err != nil {
		t.Errorf("It should not produce any error")
	}
	if arr != 3.14 {
		t.Errorf("Data parsed uncorrectly")
	}
}

func TestString(t *testing.T) {
	j := Jenv{}
	_ = j.LoadFromString(`
		{
			"data": "testing"
		}
	`)
	data, _ := j.Get("data")
	arr, err := String(data)
	if err != nil {
		t.Errorf("It should not produce any error")
	}
	if arr != "testing" {
		t.Errorf("Data parsed uncorrectly")
	}
}

func TestObject(t *testing.T) {
	j := Jenv{}
	_ = j.LoadFromString(`
		{
			"data": {
				"inside": "the value inside"
			}
		}
	`)
	obj := struct {
		Inside string `json:"inside"`
	}{}
	data, _ := j.Get("data")
	err := Object(data, &obj)
	if err != nil {
		t.Errorf("It should not produce any error")
	}
	if obj.Inside != "the value inside" {
		t.Errorf("Data parsed uncorrectly")
	}
}

func TestMap(t *testing.T) {
	j := Jenv{}
	_ = j.LoadFromString(`
		{
			"data": {
				"inside": "the value inside"
			}
		}
	`)
	data, _ := j.Get("data")
	m, err := Map(data)
	if err != nil {
		t.Errorf("It should not produce any error")
	}
	str, _ := String(m["inside"])
	if str != "the value inside" {
		t.Errorf("Data parsed uncorrectly")
	}
}

func TestStringArray(t *testing.T) {
	j := Jenv{}
	_ = j.LoadFromString(`
		{
			"data": ["1", "2", "3"]
		}
	`)
	data, _ := j.Get("data")
	arr, err := StringArray(data)
	if err != nil {
		t.Errorf("It should not produce any error")
	}
	if len(arr) != 3 || arr[0] != "1" || arr[1] != "2" || arr[2] != "3" {
		t.Errorf("Data parsed uncorrectly")
	}
}

func TestFloat64Array(t *testing.T) {
	j := Jenv{}
	_ = j.LoadFromString(`
		{
			"data": [1, 2.22, 3.14]
		}
	`)
	data, _ := j.Get("data")
	arr, err := Float64Array(data)
	if err != nil {
		t.Errorf("It should not produce any error")
	}
	if len(arr) != 3 || arr[0] != 1 || arr[1] != 2.22 || arr[2] != 3.14 {
		t.Errorf("Data parsed uncorrectly")
	}
}

func TestBooleanArray(t *testing.T) {
	j := Jenv{}
	_ = j.LoadFromString(`
		{
			"data": [true, false, true]
		}
	`)
	data, _ := j.Get("data")
	arr, err := BooleanArray(data)
	if err != nil {
		t.Errorf("It should not produce any error")
	}
	if len(arr) != 3 || arr[0] != true || arr[1] != false || arr[2] != true {
		t.Errorf("Data parsed uncorrectly")
	}
}

func TestObjectArray(t *testing.T) {
	j := Jenv{}
	_ = j.LoadFromString(`
		{
			"data": [
				{
					"number": "1"
				},
				{
					"number": "2"
				}
			]
		}
	`)
	data, _ := j.Get("data")
	arr, err := ObjectArray(data)
	if err != nil {
		t.Errorf("It should not produce any error")
	}
	if len(arr) != 2 {
		t.Errorf("Data parsed uncorrectly")
	} else {
		val1 := struct {
			Number string `json:"number"`
		}{}
		val2 := struct {
			Number string `json:"number"`
		}{}
		err = json.Unmarshal(arr[0], &val1)
		if err != nil {
			t.Errorf("Value of first object parsed uncorrectly")
		}
		err = json.Unmarshal(arr[1], &val2)
		if err != nil {
			t.Errorf("Value of second object parsed uncorrectly")
		}
		if val1.Number != "1" {
			t.Errorf("Value of number in the first object parsed uncorrectly")
		}
		if val2.Number != "2" {
			t.Errorf("Value of number in the second object parsed uncorrectly")
		}
	}
}
