package formatter

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/sysvm/gojsondiff/diff"
	"github.com/sysvm/gojsondiff/tests"
)

var _ = ginkgo.Describe("Ascii", func() {
	ginkgo.Describe("AsciiPrinter", func() {
		var (
			a, b map[string]interface{}
		)

		ginkgo.It("Prints the given diff", func() {
			a = tests.LoadFixture("../testdata/base.json")
			b = tests.LoadFixture("../testdata/base_changed.json")

			diff := diff.New().CompareObjects(a, b)
			gomega.Expect(diff.Modified()).To(gomega.BeTrue())

			f := NewAsciiFormatter(a, AsciiFormatterDefaultConfig)
			deltaJson, err := f.Format(diff)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(deltaJson).To(gomega.Equal(
				` {
   "arr": [
     "arr0",
     21,
     {
       "num": 1,
-      "str": "pek3f"
+      "str": "changed"
     },
     [
       0,
-      "1"
+      "changed"
     ]
   ],
   "bool": true,
-  "null": null,
   "num_float": 39.39,
   "num_int": 13,
   "obj": {
     "arr": [
       17,
       "str",
       {
-        "str": "eafeb"
+        "str": "changed"
       }
     ],
-    "num": 19,
     "obj": {
-      "num": 14,
+      "num": 9999,
-      "str": "efj3"
+      "str": "changed"
     },
     "str": "bcded"
+    "new": "added"
   },
   "str": "abcde"
 }
`,
			),
			)
		})

		ginkgo.It("Prints the given diff", func() {
			a = tests.LoadFixture("../testdata/add_delete_from.json")
			b = tests.LoadFixture("../testdata/add_delete_to.json")

			diff := diff.New().CompareObjects(a, b)
			gomega.Expect(diff.Modified()).To(gomega.BeTrue())

			f := NewAsciiFormatter(a, AsciiFormatterDefaultConfig)
			deltaJson, err := f.Format(diff)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(deltaJson).To(gomega.Equal(
				` {
-  "delete": {
-    "l0a": [
-      "abcd",
-      [
-        "efcg"
-      ]
-    ],
-    "l0o": {
-      "l1o": {
-        "l2s": "efed"
-      },
-      "l1s": "abcd"
-    }
-  }
+  "add": {
+    "l0a": [
+      "abcd",
+      [
+        "efcg"
+      ]
+    ],
+    "l0o": {
+      "l1o": {
+        "l2s": "efed"
+      },
+      "l1s": "abcd"
+    }
+  }
 }
`,
			),
			)
		})
	})

})
