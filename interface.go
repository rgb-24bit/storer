package storer

type Loader interface {
	Load() ([]byte, error)
}

type Decoder interface {
	Decode([]byte) (map[string]interface{}, error)
}

type Storer interface {
	Store(map[string]interface{}) error
}

