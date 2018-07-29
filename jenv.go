// Simple Package to load JSON-based environment variables. The environment variables could be any type of valid JSON object containing arrays, strings, numbers, boolean, etc.
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

// Get returns the value stored in Data field. It will return json.RawMessage and boolean indicating value exists or not.
func (env *Jenv) Get(key string) (json.RawMessage, bool) {
	val, ok := env.Data[key]
	return val, ok
}

// StringArray will convert the json.RawMessage value into array of string. It will return error if the json.RawMessage format is invalid or if the data is not an array of string.
func StringArray(data json.RawMessage) ([]string, error) {
	u := []string{}
	err := json.Unmarshal([]byte(data), &u)
	if err != nil {
		return nil, errors.New("Invalid Format")
	}
	return u, nil
}

// ObjectArray will convert the json.RawMessage value into array of json object. It will return error if the json.RawMessage format is invalid or if the data is not an array of json object.
func ObjectArray(data json.RawMessage) ([]json.RawMessage, error) {
	u := []json.RawMessage{}
	err := json.Unmarshal([]byte(data), &u)
	if err != nil {
		return nil, errors.New("Invalid Format")
	}
	return u, nil
}
