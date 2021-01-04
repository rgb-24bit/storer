package internal

import (
	"errors"
	"reflect"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
)

// NewMapStructureDecoder create MapStructureDecoder with MapKeyToCamelHook
func NewMapStructureDecoder(out interface{}, hooks ...mapstructure.DecodeHookFunc) (*mapstructure.Decoder, error) {
	hooks = append(hooks, MapKeyToCamelHookFunc())
	config := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(hooks...),
		Metadata:   nil,
		Result:     out,
		Squash:     true, // support embedded struct, ref https://github.com/mitchellh/mapstructure/pull/168
	}
	return mapstructure.NewDecoder(config)
}

// MapKeyToCamelHookFunc convert key of map[string]interface{} or map[interface{}(string)]interface{} to Camel style
func MapKeyToCamelHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if t.Kind() != reflect.Struct {
			return data, nil
		}

		if !isMapType(data) {
			return nil, errors.New("convert data to struct, but data is not a map")
		}

		nm := make(map[string]interface{})
		switch om := data.(type) {
		case map[string]interface{}:
			for k, v := range om {
				nm[strcase.ToCamel(k)] = v
				nm[k] = v
			}
		case map[interface{}]interface{}:
			for k, v := range om {
				nm[strcase.ToCamel(k.(string))] = v
				nm[k.(string)] = v
			}
		}

		return nm, nil
	}
}

func isMapType(i interface{}) bool {
	_, i2ip := i.(map[interface{}]interface{})
	_, s2ip := i.(map[string]interface{})
	return i2ip || s2ip
}
