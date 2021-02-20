package data

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-hclog"
)

func TestNewRates(t *testing.T) {
	tr, err := NewRates(hclog.Default())

	if err != nil {
		t.Fatal("COULD NOT CREATE STRUCT!", err)
	}
	fmt.Printf("%+v", tr.rates)
}
