package goscore

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type truePredicate struct{}
type dummyMiningSchema struct{}

// Node - PMML tree node
type Node struct {
	XMLName            xml.Name
	ID                 string              `xml:"id,attr"`
	Score              string              `xml:"score,attr"`
	RecordCount        string              `xml:"recordCount,attr"`
	Content            []byte              `xml:",innerxml"`
	Nodes              []Node              `xml:",any"`
	True               truePredicate       `xml:"True"`
	DummyMiningSchema  dummyMiningSchema   `xml:"MiningSchema"`
	SimplePredicate    *SimplePredicate    `xml:"SimplePredicate"`
	SimpleSetPredicate *SimpleSetPredicate `xml:"SimpleSetPredicate"`
	ScoreDistribution  *ScoreDistribution  `xml:"ScoreDistribution"`
}

// TraverseTree - traverses Node predicates with features and returns score by terminal node
func (n Node) TraverseTree(features map[string]any) (score float64, err error) {
	curr := n.Nodes[0]
	for len(curr.Nodes) > 0 {
		prevID := curr.ID
		curr = step(curr, features)
		if prevID == curr.ID {
			break
		}
	}

	v, err := strconv.ParseFloat(curr.Score, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid score '%s' in node %s, parse err %v", curr.Score, curr.ID, err)
	}

	return v, nil
}

func step(curr Node, features map[string]any) Node {
	for _, node := range curr.Nodes {
		if node.XMLName.Local == "True" ||
			(node.SimplePredicate != nil && node.SimplePredicate.True(features)) ||
			(node.SimpleSetPredicate != nil && node.SimpleSetPredicate.True(features)) {
			curr = node
			break
		}
	}
	return curr
}
