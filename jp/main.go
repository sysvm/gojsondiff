package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sysvm/gojsondiff/diff"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "jd"
	app.Usage = "JSON Diff"
	app.Version = "0.0.2"

	app.Flags = []cli.Flag{}

	app.Action = func(c *cli.Context) {
		if len(c.Args()) < 2 {
			fmt.Println("Not enough arguments")
			fmt.Printf("Usage: %s diff json_file\n", app.Name)
			os.Exit(1)
		}

		diffFilePath := c.Args()[0]
		jsonFilePath := c.Args()[1]

		// Diff file
		diffFile, err := os.ReadFile(diffFilePath)
		if err != nil {
			fmt.Printf("Failed to open file '%s': %s\n", diffFilePath, err.Error())
			os.Exit(2)
		}

		// Load Diff file
		um := diff.NewUnmarshaler()
		diffObject, err := um.UnmarshalBytes(diffFile)
		if err != nil {
			fmt.Printf("Failed to load diff file '%s': %s\n", diffFilePath, err.Error())
			os.Exit(2)
		}

		// JSON file
		jsonFile, err := os.ReadFile(jsonFilePath)
		if err != nil {
			fmt.Printf("Failed to open file '%s': %s\n", jsonFilePath, err.Error())
			os.Exit(2)
		}

		// Load JSON
		var jsonObject map[string]interface{}
		if err = json.Unmarshal(jsonFile, &jsonObject); err != nil {
			fmt.Printf("json Unmarshal error %s\n", err.Error())
			os.Exit(2)
		}

		// Apply
		differ := diff.New()
		differ.ApplyPatch(jsonObject, diffObject)

		patchedJson, _ := json.MarshalIndent(jsonObject, "", "  ")
		fmt.Println(string(patchedJson))
	}

	app.Run(os.Args)
}
