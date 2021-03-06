package main

import (
	"fmt"
	"io"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := randomPasswordApp(os.Stdout)
	app.Run(os.Args)
}

// randomPasswordApp handles the length argument and calls the generate function
// printing the results
func randomPasswordApp(out io.Writer) *cli.App {
	app := cli.NewApp()
	app.Name = "passgen"
	app.Usage = "generate a random password"
	app.Writer = out
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "chars, c",
			Value: "a1*",
			Usage: "specify which character sets to include in the generated password",
		},
		cli.IntFlag{
			Name:  "length, l",
			Value: 20,
			Usage: "specify the length of the generated password",
		},
		cli.BoolFlag{
			Name:  "words, w",
			Usage: "use words instead",
		},
	}
	app.Action = func(c *cli.Context) {
		symbolSubsets, subsetErr := getSymbolSubsets(c.String("chars"))
		if subsetErr != nil {
			fmt.Fprintln(out, "Could not generate a random password successfully.")
			fmt.Fprintln(out, subsetErr)
			return
		}

		// Convert the given argument so it can be used as the length of the password to generate
		length := c.Int("length")
		if length < 0 {
			fmt.Fprintln(out, "Could not generate a random password successfully.")
			fmt.Fprintln(out, length, "is not a valid password length. Must be greater than zero.")
			return
		}

		// generate the password
		var symbolSet alphabet
		if c.Bool("words") {
			symbolSet = englishWords
		} else {
			symbolSet = valid(symbolSubsets)
		}
		password, err := generate(length, symbolSet)
		if err != nil {
			fmt.Fprintln(out, "Could not generate a random password successfully.")
			fmt.Fprintln(out, err)
			return
		}

		// Print the generated password
		fmt.Fprintln(out, "Random password:", password, "-", len(password), "characters long")
		return
	}

	return app
}
