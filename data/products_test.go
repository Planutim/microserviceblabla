package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestChecksValidation(t *testing.T) {
// 	p := &Product{
// 		Name:  "test",
// 		Price: 1.10,
// 		SKU:   "abc-abc-abc",
// 	}
// 	err := p.Validate()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
func TestProductMissingNameReturnsErr(t *testing.T) {
	p := Product{
		Price: 1.22,
	}
	v := NewValidation()
	err := v.validate(p)
	assert.Len(t, err, 1)
}
