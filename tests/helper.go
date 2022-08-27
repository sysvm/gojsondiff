package tests

import (
	"encoding/json"
	"os"

	"github.com/onsi/ginkgo"
)

func LoadFixture(file string) map[string]interface{} {
	content, err := os.ReadFile(file)
	if err != nil {
		ginkgo.Fail("Fixture file '" + file + "' not found.")
	}
	var result map[string]interface{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		ginkgo.Fail("Unmarshalling JSON of '" + file + "' failed: " + err.Error())
	}
	return result
}

func LoadFixtureAsArray(file string) []interface{} {
	content, err := os.ReadFile(file)
	if err != nil {
		ginkgo.Fail("Fixture file '" + file + "' not found.")
	}
	var result []interface{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		ginkgo.Fail("Unmarshalling JSON of '" + file + "' failed: " + err.Error())
	}
	return result
}
