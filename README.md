# Storer

> Go configuration reading and deserialization

# Introduction

Now this library is just a simple idea. It reads the deserialization configuration through a general rule and maintains good expansibility.

For configuration reading and deserialization scenarios, usually, we need to first read and load configuration from a source:
```go
type Loader interface {
	Load() ([]byte, error)
}
```

For these configurations, we can usually parse them into the format of `map[string]interface{}`:
```go
type Decoder interface {
	Decode([]byte) (map[string]interface{}, error)
}
```

Finally, these `map[string]interface{}` objects can be processed and saved:
```go
type Storer interface {
	Store(map[string]interface{}) error
}
```

# Unmarshal

For unmarshal scenarios, when converting `map[string]interface{}` into a structure, we can often use [mapstructure](https://github.com/mitchellh/mapstructure) to achieve.

In particular, the fields of the target structure are usually in the form of `Camel`, so I think that the key of the map can be converted to the form of `Camel` through the hook.
```go
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
```

