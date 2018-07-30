// Package jenv is a simple package to load JSON-based environment variables. The environment variables could be any type of valid JSON object containing arrays, strings, numbers, boolean, etc.
package jenv

import (
	"encoding/json"
	"errors"
	"os"
)

// Jenv is a type containing FilePath and Data for the environment variables. It is needed by the Load and LoadFromString method. To create one simply just create the object or call DefaultConfig method.
type Jenv struct {
	FilePath string
	Data     map[string]json.RawMessage
}

// DefaultConfig returns a Jenv object with empty data and "env.json" filename.
func DefaultConfig() Jenv {
	return Jenv{
		FilePath: "env.json",
	}
}

// Load is used to decode the environment variables into Data field in Jenv struct. It will return error if the decoding process has an error.
func (env *Jenv) Load() error {
	envFile, err := os.Open(env.FilePath)
	if err != nil {
		return errors.New("File Not Found")
	}
	defer envFile.Close()

	decoder := json.NewDecoder(envFile)
	err = decoder.Decode(&env.Data)
	if err != nil {
		return errors.New("Decode Error")
	}
	return nil
}

// LoadFromString is used to decode the environment variables provided in the string parameter into Data field in Jenv struct. It will return error if the decoding process has an error.
func (env *Jenv) LoadFromString(s string) error {
	err := json.Unmarshal([]byte(s), &env.Data)
	if err != nil {
		return errors.New("Decode Error")
	}
	return nil
}

// Get returns the value stored in Data field. It will return []byte and boolean indicating value exists or not.
func (env *Jenv) Get(key string) ([]byte, bool) {
	val, ok := env.Data[key]
	if string(val) == "null" {
		return nil, ok
	}
	return []byte(val), ok
}

// Boolean will convert the []byte value into float64. It will return error if the []byte format is invalid or if the data is not a boolean.
func Boolean(data []byte) (bool, error) {
	var u bool
	err := json.Unmarshal(data, &u)
	if err != nil {
		return false, errors.New("Invalid Format")
	}
	return u, nil
}

// Float64 will convert the []byte value into float64. It will return error if the []byte format is invalid or if the data is not a number.
func Float64(data []byte) (float64, error) {
	var u float64
	err := json.Unmarshal(data, &u)
	if err != nil {
		return 0, errors.New("Invalid Format")
	}
	return u, nil
}

// String will convert the []byte value into string. It will return error if the []byte format is invalid or if the data is not a string.
func String(data []byte) (string, error) {
	var u string
	err := json.Unmarshal(data, &u)
	if err != nil {
		return "", errors.New("Invalid Format")
	}
	return u, nil
}

// Object will convert the []byte value into struct object passed into the parameter. It will return error if the []byte format is invalid or if the data is not the appropriate type.
func Object(data []byte, dest interface{}) error {
	err := json.Unmarshal(data, dest)
	if err != nil {
		return errors.New("Invalid Format")
	}
	return nil
}

// Map will convert the []byte value into a string map of json.RawMessage. It will return error if the []byte format is invalid or if the data is not the appropriate type.
func Map(data []byte) (map[string]json.RawMessage, error) {
	u := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &u)
	if err != nil {
		return nil, errors.New("Invalid Format")
	}
	return u, nil
}

// StringArray will convert the []byte value into array of string. It will return error if the []byte format is invalid or if the data is not an array of string.
func StringArray(data []byte) ([]string, error) {
	u := []string{}
	err := json.Unmarshal(data, &u)
	if err != nil {
		return nil, errors.New("Invalid Format")
	}
	return u, nil
}

// Float64Array will convert the []byte value into array of float64. It will return error if the []byte format is invalid or if the data is not an array of number.
func Float64Array(data []byte) ([]float64, error) {
	u := []float64{}
	err := json.Unmarshal(data, &u)
	if err != nil {
		return nil, errors.New("Invalid Format")
	}
	return u, nil
}

// BooleanArray will convert the []byte value into array of boolean. It will return error if the []byte format is invalid or if the data is not an array of boolean.
func BooleanArray(data []byte) ([]bool, error) {
	u := []bool{}
	err := json.Unmarshal(data, &u)
	if err != nil {
		return nil, errors.New("Invalid Format")
	}
	return u, nil
}

// ObjectArray will convert the []byte value into array of json object. It will return error if the []byte format is invalid or if the data is not an array of json object.
func ObjectArray(data []byte) ([][]byte, error) {
	u := []json.RawMessage{}
	err := json.Unmarshal(data, &u)
	if err != nil {
		return nil, errors.New("Invalid Format")
	}
	ret := [][]byte{}
	for _, val := range u {
		ret = append(ret, []byte(val))
	}
	return ret, nil
}
