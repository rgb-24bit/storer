package storer

import (
	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

func NewMapStructureDecoder(out interface{}, hooks ...mapstructure.DecodeHookFunc) (*mapstructure.Decoder, error) {
	hooks = append(hooks, MapKeyToCamelHookFunc())
	config := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(hooks...),
		Metadata: nil,
		Result:   out,
	}
	return mapstructure.NewDecoder(config)
}

func MapKeyToCamelHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if t.Kind() != reflect.Struct {
			return data, nil
		}

		nm := make(map[string]interface{})

		switch data := data.(type) {
		case map[string]interface{}:
			for k, v := range data {
				nm[strcase.ToCamel(k)] = v
			}
		case map[interface{}]interface{}:
			for k, v := range data {
				nm[strcase.ToCamel(k.(string))] = v
			}
		default:
			return data, nil
		}

		return nm, nil
	}
}

