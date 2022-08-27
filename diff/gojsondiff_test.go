package diff

import (
	"os"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/sysvm/gojsondiff/tests"
)

var _ = ginkgo.Describe("GoJSONDiff", func() {
	ginkgo.Describe("Differ", func() {
		ginkgo.Describe("CompareObjects", func() {

			var (
				a, b   map[string]interface{}
				differ *Differ
			)

			ginkgo.BeforeEach(func() {
				differ = New()
			})

			ginkgo.Context("There are no difference between the two JSON strings", func() {
				ginkgo.It("Detects nothing", func() {
					a = tests.LoadFixture("../testdata/base.json")
					b = tests.LoadFixture("../testdata/base.json")

					diff := differ.CompareObjects(a, b)
					gomega.Expect(diff.Modified()).To(gomega.BeFalse())
				})
			})

			ginkgo.Context("There are some values modified", func() {
				ginkgo.It("Detects changes", func() {
					a = tests.LoadFixture("../testdata/base.json")
					b = tests.LoadFixture("../testdata/base_changed.json")

					diff := differ.CompareObjects(a, b)
					gomega.Expect(diff.Modified()).To(gomega.BeTrue())
					differ.ApplyPatch(a, diff)
					gomega.Expect(a).To(gomega.Equal(tests.LoadFixture("../testdata/base_changed.json")))
				})
			})

			ginkgo.Context("There are values only types are changed", func() {
				ginkgo.It("Detects changed types", func() {
					a := tests.LoadFixture("../testdata/changed_types_from.json")
					b := tests.LoadFixture("../testdata/changed_types_to.json")

					diff := differ.CompareObjects(a, b)
					differ.ApplyPatch(a, diff)
					gomega.Expect(a).To(gomega.Equal(tests.LoadFixture("../testdata/changed_types_to.json")))
				})
			})

			ginkgo.Context("There is a moved item in an array", func() {
				ginkgo.It("Detects changed types", func() {
					a := tests.LoadFixture("../testdata/move_from.json")
					b := tests.LoadFixture("../testdata/move_to.json")

					diff := differ.CompareObjects(a, b)
					gomega.Expect(diff.Modified()).To(gomega.BeTrue())
					differ.ApplyPatch(a, diff)
					gomega.Expect(a).To(gomega.Equal(tests.LoadFixture("../testdata/move_to.json")))
				})
			})

			ginkgo.Context("There are long text diff", func() {
				ginkgo.It("Detects changes", func() {
					a = tests.LoadFixture("../testdata/long_text_from.json")
					b = tests.LoadFixture("../testdata/long_text_to.json")

					diff := differ.CompareObjects(a, b)
					gomega.Expect(diff.Modified()).To(gomega.BeTrue())
					differ.ApplyPatch(a, diff)
					gomega.Expect(a).To(gomega.Equal(tests.LoadFixture("../testdata/long_text_to.json")))
				})
			})
		})
		ginkgo.Describe("CompareArrays", func() {

			var (
				a, b   []interface{}
				differ *Differ
			)

			ginkgo.BeforeEach(func() {
				differ = New()
			})

			ginkgo.Context("There are no difference between the two JSON strings", func() {
				ginkgo.It("Detects nothing", func() {
					a = tests.LoadFixtureAsArray("../testdata/array.json")
					b = tests.LoadFixtureAsArray("../testdata/array.json")

					diff := differ.CompareArrays(a, b)
					gomega.Expect(diff.Modified()).To(gomega.BeFalse())
				})
			})

			ginkgo.Context("There are some values modified", func() {
				ginkgo.It("Detects changes", func() {
					a = tests.LoadFixtureAsArray("../testdata/array.json")
					b = tests.LoadFixtureAsArray("../testdata/array_changed.json")

					diff := differ.CompareArrays(a, b)
					gomega.Expect(diff.Modified()).To(gomega.BeTrue())
					gomega.Expect(len(diff.Deltas())).To(gomega.Equal(1))
				})
			})
		})
		ginkgo.Describe("Compare", func() {
			ginkgo.Context("There are some values modified", func() {
				ginkgo.It("Detects changes", func() {
					aFile := "../testdata/base.json"
					bFile := "../testdata/base_changed.json"
					aObj := tests.LoadFixture(aFile)
					bObj := tests.LoadFixture(bFile)

					differ := New()

					diffObj := differ.CompareObjects(aObj, bObj)
					gomega.Expect(diffObj.Modified()).To(gomega.BeTrue())

					aStr, err := os.ReadFile(aFile)
					gomega.Expect(err).To(gomega.BeNil())
					bStr, err := os.ReadFile(bFile)
					gomega.Expect(err).To(gomega.BeNil())

					diffStr, err := differ.Compare(aStr, bStr)
					gomega.Expect(err).To(gomega.BeNil())
					gomega.Expect(diffStr).To(gomega.Equal(diffObj))
				})
			})
		})
	})
})
