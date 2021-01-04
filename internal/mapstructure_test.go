package internal

import (
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	ID       int `mapstructure:"id"`
	Name     string
	Embedded struct {
		ID   int `mapstructure:"id"`
		Name string
	}
}

var _ = Describe("MapStructure", func() {
	t := GinkgoT()

	Context("data type is map[string]interface{}", func() {
		It("ok", func() {
			var m = map[string]interface{}{
				"id":   1234,
				"name": "test",
				"embedded": map[string]interface{}{
					"id":   1234,
					"name": "embedded",
				},
			}
			var out TestStruct

			decoder, err := NewMapStructureDecoder(&out)
			assert.Nil(t, err)

			assert.Nil(t, decoder.Decode(m))

			assert.Equal(t, 1234, out.ID)
			assert.Equal(t, "test", out.Name)
			assert.Equal(t, 1234, out.Embedded.ID)
			assert.Equal(t, "embedded", out.Embedded.Name)
		})
	})

	Context("data type is map[interface{}]interface{}", func() {
		It("ok", func() {
			var m = map[interface{}]interface{}{
				"id":   1234,
				"name": "test",
				"embedded": map[interface{}]interface{}{
					"id":   1234,
					"name": "embedded",
				},
			}
			var out TestStruct

			decoder, err := NewMapStructureDecoder(&out)
			assert.Nil(t, err)

			assert.Nil(t, decoder.Decode(m))

			assert.Equal(t, 1234, out.ID)
			assert.Equal(t, "test", out.Name)
			assert.Equal(t, 1234, out.Embedded.ID)
			assert.Equal(t, "embedded", out.Embedded.Name)
		})
	})

	Context("data type is not map", func() {
		It("fail", func() {
			var m = map[interface{}]interface{}{
				"id":       1234,
				"name":     "test",
				"embedded": []string{},
			}
			var out TestStruct

			decoder, err := NewMapStructureDecoder(&out)
			assert.Nil(t, err)

			assert.NotNil(t, decoder.Decode(m))
		})
	})
})
