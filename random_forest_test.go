package goscore_test

import (
	"testing"

	"github.com/asafschers/goscore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RandomForest", func() {
	var (
		lowScoreFeatures, highScoreFeatures map[string]any
		lowScore, highScore                 float64
		errorFeatures                       map[string]any
		rf                                  goscore.RandomForest
		err                                 error
		first                               bool = true
	)

	BeforeEach(func() {
		lowScoreFeatures = map[string]any{
			"Sex":      "male",
			"Parch":    0,
			"Age":      30,
			"Fare":     9.6875,
			"Pclass":   2,
			"SibSp":    0,
			"Embarked": "Q",
		}
		lowScore = 2.0 / 15.0

		highScoreFeatures = map[string]any{
			"Sex":      "female",
			"Parch":    0,
			"Age":      38,
			"Fare":     71.2833,
			"Pclass":   2,
			"SibSp":    1,
			"Embarked": "C",
		}
		highScore = 14.0 / 15.0

		errorFeatures = map[string]any{
			"Sex":      "female",
			"Parch":    0,
			"Age":      38,
			"Fare":     71.2833,
			"Pclass":   2,
			"SibSp":    1,
			"Embarked": "UnknownCategory",
		}

		if first {
			rf, err = goscore.LoadRandomForest("fixtures/random_forest.pmml")
			if err != nil {
				panic(err)
			}
			first = false
		}
	})

	Describe("Loads model", func() {
		It("Loads all trees", func() {
			Expect(len(rf.Trees)).To(Equal(15))
		})
	})

	Describe("Scores", func() {
		It("Scores low", func() {
			Expect(rf.Score(lowScoreFeatures, "1")).To(Equal(lowScore))
		})

		It("Scores low concurrently", func() {
			Expect(rf.ScoreConcurrently(lowScoreFeatures, "1")).To(Equal(lowScore))
		})

		It("Scores high", func() {
			Expect(rf.Score(highScoreFeatures, "1")).To(Equal(highScore))
		})

		It("Scores high concurrently", func() {
			Expect(rf.ScoreConcurrently(highScoreFeatures, "1")).To(Equal(highScore))
		})
	})

	Describe("Errors", func() {
		It("Errors", func() {
			_, err = rf.Score(errorFeatures, "1")
			Expect(err).To(MatchError(HavePrefix("Terminal node without score")))
			_, err = rf.ScoreConcurrently(errorFeatures, "1")
			Expect(err).To(MatchError(HavePrefix("Terminal node without score")))
		})
	})
})

func TestRandomForest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RandomForest Suite")
}
