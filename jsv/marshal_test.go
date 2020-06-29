package jsv

import (
	"testing"

	"github.com/petuhovskiy/cian-search/nometa"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	obj := nometa.Building{
		BuildYear:   228,
		FloorsCount: 1337,
	}

	expected := map[string]string{
		"BuildYear":   "228",
		"FloorsCount": "1337",
	}

	actual, err := Marshal(obj)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

type myStruct struct {
	A myA
	B string
}

type myA struct {
	DarkKeks string
}

func TestNestedMarshal(t *testing.T) {
	obj := myStruct{
		A: myA{
			DarkKeks: "flyce",
		},
		B: "Baliuk",
	}

	expected := map[string]string{
		"A.DarkKeks": "flyce",
		"B":          "Baliuk",
	}

	actual, err := Marshal(obj)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
