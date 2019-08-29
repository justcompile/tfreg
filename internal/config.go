package internal

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//Config contains configuration values required for tfreg to function
type Config struct {
	Version string `tfreg:"SHA"`
}

// NewConfig returns a Config instance or error
func NewConfig() (*Config, error) {
	config := &Config{}
	if _, err := initStructFromEnvironment(config, "tfreg"); err != nil {
		return nil, err
	}

	return config, nil
}

func initStructFromEnvironment(dest interface{}, envVarPrefix string) (interface{}, error) {
	value := reflect.Indirect(reflect.ValueOf(dest))

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		tag, err := parseTag(value.Type().Field(i))
		if err != nil {
			return nil, err
		}

		envKey := envVarPrefix + "_" + tag.key

		var fieldValue reflect.Value

		if field.Type().Kind() == reflect.Struct {
			nestedStruct := reflect.New(field.Type())

			nestedStructValue, err := initStructFromEnvironment(nestedStruct.Interface(), envKey)
			if err != nil {
				return nil, err
			}
			fieldValue = reflect.ValueOf(nestedStructValue)

		} else {
			envVar := os.Getenv(envKey)

			if tag.required && envVar == "" {
				return nil, fmt.Errorf("Environment variable %s is not set", envKey)
			}

			if envVar == "" && (field.Type().Kind() == reflect.Bool || field.Type().Kind() == reflect.Int) {
				continue
			}

			switch field.Type().Kind() {
			case reflect.Bool:
				valueAsBool, err := strconv.ParseBool(envVar)
				if err != nil {
					return nil, fmt.Errorf("Environment variable %s value is not a valid boolean: %s", envKey, envVar)
				}

				fieldValue = reflect.ValueOf(valueAsBool)
			case reflect.Int:
				valueAsInt, err := strconv.ParseInt(envVar, 10, 0)
				if err != nil {
					return nil, fmt.Errorf("Environment variable %s value is not a valid number: %s", envKey, envVar)
				}

				fieldValue = reflect.ValueOf(valueAsInt)
			default:
				fieldValue = reflect.ValueOf(envVar)
			}
		}

		field.Set(reflect.Indirect(fieldValue))
	}

	return dest, nil
}

type configTag struct {
	key      string
	required bool
}

func parseTag(field reflect.StructField) (*configTag, error) {
	tagValue, ok := field.Tag.Lookup("tfreg")
	if !ok {
		return nil, fmt.Errorf("tfreg tag cannot be found on field: %s", field.Name)
	}

	splitValue := strings.Split(tagValue, ",")
	if len(splitValue) == 0 || splitValue[0] == "" {
		return nil, fmt.Errorf("tfreg tag of field %s is invalid: has an empty value", field.Name)
	}

	tag := &configTag{
		key:      splitValue[0],
		required: len(splitValue) > 1 && splitValue[1] == "required",
	}

	return tag, nil
}
