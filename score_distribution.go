package goscore

// ScoreDistribution - PMML score distribution
type ScoreDistribution struct {
	Value       string `xml:"value,attr"`
	RecordCount string `xml:"recordCount,attr"`
}
