package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sysvm/gojsondiff/diff"
	"github.com/sysvm/gojsondiff/formatter"
	"github.com/urfave/cli"
)

const (
	fileSuffix = ".json"
	diffStr    = "diff"
)

func main() {
	app := cli.NewApp()
	app.Name = "jd"
	app.Usage = "JSON Diff"
	app.Version = "0.0.2"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "format, f",
			Value:  "ascii",
			Usage:  "Diff Output Format (ascii, delta)",
			EnvVar: "DIFF_FORMAT",
		},
		cli.BoolFlag{
			Name:   "coloring, c",
			Usage:  "Enable coloring in the ASCII mode (not available in the delta mode)",
			EnvVar: "COLORING",
		},
		cli.BoolFlag{
			Name:   "quiet, q",
			Usage:  "Suppress output, if no differences are found",
			EnvVar: "QUIET",
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(c.Args()) < 2 {
			fmt.Println("Not enough arguments.")
			fmt.Printf("Usage: %s json_file another_json_file\n", app.Name)
			os.Exit(1)
		}

		aFilePath := c.Args()[0]
		bFilePath := c.Args()[1]

		// Prepare your JSON string as `[]byte`, not `string`
		aString, err := os.ReadFile(aFilePath)
		if err != nil {
			fmt.Printf("Failed to open file '%s': %s\n", aFilePath, err.Error())
			os.Exit(2)
		}

		// Another JSON string
		bString, err := os.ReadFile(bFilePath)
		if err != nil {
			fmt.Printf("Failed to open file '%s': %s\n", bFilePath, err.Error())
			os.Exit(2)
		}

		path, fileName := filepath.Split(aFilePath)
		fmt.Printf("path: %s, fileName: %s\n", path, fileName)
		file := strings.TrimSuffix(fileName, fileSuffix)

		// Then, compare them
		differ := diff.New()
		d, err := differ.Compare(aString, bString)
		if err != nil {
			fmt.Printf("Failed to unmarshal file: %s\n", err.Error())
			os.Exit(3)
		}

		// Output the result
		if d.Modified() || !c.Bool("quiet") {
			format := c.String("format")
			var diffString string
			if format == "ascii" {
				var aJson map[string]interface{}
				err = json.Unmarshal(aString, &aJson)
				if err != nil {
					fmt.Printf("")
				}

				config := formatter.AsciiFormatterConfig{
					ShowArrayIndex: true,
					Coloring:       c.Bool("coloring"),
				}

				asciiFormatter := formatter.NewAsciiFormatter(aJson, config)
				diffString, err = asciiFormatter.Format(d)
				if err != nil {
					// No error can occur
				}
			} else if format == "delta" {
				deltaFormatter := formatter.NewDeltaFormatter()
				diffString, err = deltaFormatter.Format(d)
				if err != nil {
					// No error can occur
				}
			} else {
				fmt.Printf("Unknown Foramt %s\n", format)
				os.Exit(4)
			}

			writeToFile(file, diffString)
			//fmt.Print(diffString)
			return cli.NewExitError("", 1)
		}
		return nil
	}

	app.Run(os.Args)
}

func writeToFile(fileName, diffString string) {
	f, err := os.Create(fileName + diffStr + fileSuffix)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	defer f.Close()

	l, err := f.WriteString(diffString)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	fmt.Println(l, "bytes written successfully")
}
