package goscore_test

import (
	"encoding/xml"
	"testing"

	"github.com/userpro/goscore"
)

var simplePredicateTests = []struct {
	predicate []byte
	features  map[string]any
	expected  bool
}{
	{[]byte(`<SimplePredicate field="f33" operator="equal" value="18.85"/>`),
		map[string]any{"f33": 18.850},
		true},
	{[]byte(`<SimplePredicate field="f33" operator="lessOrEqual" value="18.85"/>`),
		map[string]any{"f33": 18.84},
		true},
	{[]byte(`<SimplePredicate field="f33" operator="lessOrEqual" value="18.85"/>`),
		map[string]any{"f33": 18.86},
		false},
	{[]byte(`<SimplePredicate field="f33" operator="lessOrEqual" value="18.85"/>`),
		map[string]any{"f33": "18.84"},
		true},
	{[]byte(`<SimplePredicate field="f33" operator="isMissing" value="18.85"/>`),
		map[string]any{"f33": 18.86},
		false},
	{[]byte(`<SimplePredicate field="f33" operator="isMissing" value="18.85"/>`),
		map[string]any{},
		true},
	{[]byte(`<SimplePredicate field="f33" operator="isMissing" value="18.85"/>`),
		map[string]any{"f33": nil},
		true},
	{[]byte(`<SimplePredicate field="f33" operator="isMissing" value="18.85"/>`),
		map[string]any{"f33": ""},
		true},
}

func TestSimplePredicate(t *testing.T) {

	for _, tt := range simplePredicateTests {
		var predicate goscore.SimplePredicate
		xml.Unmarshal(tt.predicate, &predicate)

		actual := predicate.True(tt.features)
		if actual != tt.expected {
			t.Errorf("Predicate: %s %s %s, Feature value : %s, expected %t, actual %t",
				predicate.Field,
				predicate.Operator,
				predicate.Value,
				tt.features[predicate.Field],
				tt.expected,
				actual)
		}
	}
}
