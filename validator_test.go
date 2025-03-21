package jsonlogic_test

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qoala-platform/jsonlogic/v3"
)

func TestJSONLogicValidator(t *testing.T) {
	jsonlogic.AddOperator("customOperator", func(values, data any) any {
		return values
	})

	scenarios := map[string]struct {
		IsValid bool
		Rule    io.Reader
	}{
		"invalid rule": {
			IsValid: false,
			Rule:    strings.NewReader(`{"a", "b"}`),
		},
		"invalid operator": {
			IsValid: false,
			Rule:    strings.NewReader(`{"filt":[[10, 1, 100], {">=":[{"var":""},2]}]}`),
		},
		"invalid condition inside a filter": {
			IsValid: false,
			Rule:    strings.NewReader(`{"filter":[{"var":"integers"}, {"=": [{"var":""}, [10]]}]}`),
		},
		"primitive is a valid rule": {
			IsValid: true,
			Rule:    strings.NewReader(`10`),
		},
		"set must be valid": {
			IsValid: true,
			Rule: strings.NewReader(`{
				"map": [
					{"var": "objects"},
					{"set": [
						{"var": ""},
						"age",
						{"+": [{"var": ".age"}, 2]}
					]},
					{"customOperator": [1, 2, 3]}
				]
			}`),
		},
	}

	for name, scenario := range scenarios {
		t.Run(fmt.Sprintf("SCENARIO:%s", name), func(t *testing.T) {
			assert.Equal(t, scenario.IsValid, jsonlogic.IsValid(scenario.Rule))
		})
	}
}
