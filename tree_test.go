package goscore_test

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/asafschers/goscore"
)

var TreeTests = []struct {
	features map[string]any
	score    float64
	err      error
}{
	{map[string]any{},
		4.3463944950723456e-4,
		nil,
	},
	{
		map[string]any{"f2": "f2v1"},
		-1.8361380219689046e-4,
		nil,
	},
	{
		map[string]any{"f2": "f2v1", "f1": "f1v3"},
		-6.237581139073701e-4,
		nil,
	},
	{
		map[string]any{"f2": "f2v1", "f1": "f1v3", "f4": 0.08},
		0.0021968294712358194,
		nil,
	},
	{
		map[string]any{"f2": "f2v1", "f1": "f1v3", "f4": 0.09},
		-9.198573460887271e-4,
		nil,
	},
	{
		map[string]any{"f2": "f2v1", "f1": "f1v3", "f4": 0.09, "f3": "f3v2"},
		-0.0021187239505556523,
		nil,
	},
	{
		map[string]any{"f2": "f2v1", "f1": "f1v3", "f4": 0.09, "f3": "f3v4"},
		-3.3516227414227926e-4,
		nil,
	},
	{
		map[string]any{"f2": "f2v1", "f1": "f1v4"},
		0.0011015286521365208,
		nil,
	},
	{
		map[string]any{"f2": "f2v4"},
		0.0022726641744997256,
		nil,
	},
	{
		map[string]any{"f1": "f1v3", "f2": "f2v1", "f3": "f3v7", "f4": 0.09},
		-1,
		errors.New("Terminal node without score, Node id: 5"),
	},
}

// TODO: test score distribution

func TestTree(t *testing.T) {
	treeXml, err := ioutil.ReadFile("fixtures/tree.pmml")
	if err != nil {
		panic(err)
	}

	tree := []byte(treeXml)
	var n goscore.Node
	xml.Unmarshal(tree, &n)

	for _, tt := range TreeTests {
		actual, err := n.TraverseTree(tt.features)

		if err != nil {
			if tt.err == nil {
				t.Errorf("expected no error, actual: %s",
					err)
			} else if tt.err.Error() != err.Error() {
				t.Errorf("expected error %s, actual: %s",
					tt.err.Error(),
					err)
			}
		}

		if actual != tt.score {
			t.Errorf("expected %f, actual %f",
				tt.score,
				actual)
		}
	}
}
