package storer

import (
	"reflect"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"github.com/rgb-24bit/storer/internal"
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

		if !internal.IsMapType(data) {
			return data, nil
		}

		sm := internal.ToMapStringInterface(data)
		nm := make(map[string]interface{})

		for k, v := range sm {
			nm[strcase.ToCamel(k)] = v
		}

		return nm, nil
	}
}

