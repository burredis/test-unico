package feiralivre

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateWithError(t *testing.T) {
	entity := FeiraLivre{
		ID: 1,
	}

	err := entity.Validate()
	assert.NotNil(t, err)
}

func TestValidateWithoutError(t *testing.T) {
	entity := FeiraLivre{
		ID:   1,
		Nome: "Pinheiros",
	}

	err := entity.Validate()
	assert.Nil(t, err)
}
