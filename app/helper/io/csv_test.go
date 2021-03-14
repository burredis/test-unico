package io

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSV(t *testing.T) {
	records := ReadCSV("../../../DEINFO_AB_FEIRASLIVRES_2014.csv")
	assert.GreaterOrEqual(t, len(records), 1)
}
