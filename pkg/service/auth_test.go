package service

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"mementor-back/pkg/parser"
	"testing"
)

func TestToken(t *testing.T) {
	data := []struct {
		testName string
		input    string
		expected string
	}{{
		"int",
		"1",
		"1",
	},
		{
			"uuid",
			"b36af69e-76d2-4df6-9b1d-27f6415a700c",
			"b36af69e-76d2-4df6-9b1d-27f6415a700c",
		},
	}

	for _, test := range data {
		t.Run(test.testName, func(t *testing.T) {
			token, err := generateToken(test.input)
			if err != nil {
				assert.Error(t, err, "generateToken error")
			}
			id, err := parser.ParseToken(token, []byte(viper.GetString("signing_key")))
			if err != nil {
				assert.Error(t, err, "ParseToken error")
			}
			assert.Equal(t, test.expected, id)
		})
	}
}
