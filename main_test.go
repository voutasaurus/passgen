package main

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

var testWriter io.Writer = os.Stdout

// RandomPasswordApp([]string{}) should print an error about not enough args
// codegangsta's cli currently throws a runtime error if run is called with an empty slice
// Not sure if this scenario is even possible in practice
/*
func TestRandomPasswordNoArgs(t *testing.T) {
	buf := &bytes.Buffer{}
	testWriter = buf

	app := randomPasswordApp(testWriter)
	app.Run([]string{})

	message := "Called with no executable name or other arguments. This function needs an executable name."
	messageln := message + "\n"

	if buf.String() != messageln {
		t.Error("randomPassword([]string{}) printed '", buf.String(), "' expected '", messageln, "'")
	}
}
*/

// RandomPasswordApp([]string{passgen}) should exhibit default app behavior
// That is, chars=a1! and length=20
func TestRandomPasswordOneArg(t *testing.T) {
	premessage := "Random password: "
	postmessage1 := " - "
	postmessage2 := " characters long\n"

	buf := &bytes.Buffer{}
	testWriter = buf
	app := randomPasswordApp(testWriter)

	input := []string{"passgen"}
	app.Run(input)

	n := 20
	pre := buf.String()[:len(premessage)]
	pass := buf.String()[len(premessage) : len(premessage)+n]
	post1 := buf.String()[len(premessage)+n : len(premessage)+n+len(postmessage1)]
	_ = buf.String()[len(premessage)+n+len(postmessage1) : len(premessage)+n+len(postmessage1)+len(strconv.Itoa(n))]
	post2 := buf.String()[len(premessage)+n+len(postmessage1)+len(strconv.Itoa(n)):]

	if pre != premessage {
		t.Error("Expected premessage of: '" + premessage + "' Got: '" + pre + "'")
	}
	if post1 != postmessage1 {
		t.Error("Expected postmessage of: '" + postmessage1 + "' Got: '" + post1 + "'")
	}
	if post2 != postmessage2 {
		t.Error("Expected postmessage of: '" + postmessage2 + "' Got: '" + post2 + "'")
	}

	// Check char set membership
	all := union(alphabet, digit, special)
	for _, x := range []rune(pass) {
		if !all[x] {
			t.Error("pass =", pass, "contains", string(x), "but this is not in", all)
		}
	}

}

// RandomPasswordApp([]string{passgen, n}) n is not a number should print error about atoi or invalid length arg
func TestRandomPasswordNAN(t *testing.T) {
	buf := &bytes.Buffer{}
	testWriter = buf
	n := "stuff"

	app := randomPasswordApp(testWriter)
	app.Run([]string{"passgen", "-l", n})

	expected := `Incorrect Usage.

NAME:
   passgen - generate a random password

USAGE:
   /tmp/go-build496363236/github.com/voutasaurus/passgen/_test/passgen.test [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --chars, -c "a1*"	specify which character sets to include in the generated password
   --length, -l "20"	specify the length of the generated password
   --help, -h		show help
   --version, -v	print the version

`
	expectWords := strings.Fields(expected)
	actualWords := strings.Fields(buf.String())

	if len(actualWords) != len(expectWords) {
		t.Error("expected", len(expectWords), "words. got", len(actualWords), "words.")
		t.Error("passgen -l", n, "printed '", buf.String(), "' expected", expected)
		t.FailNow()
	}

	for i := range actualWords {
		if i != 10 && actualWords[i] != expectWords[i] { // i == 10 contains the build number for the tests
			t.Error("Output at word", i, "expected:", expectWords[i], "got:", actualWords[i])
			t.Error("passgen -l", n, "printed '", buf.String(), "' expected", expected)
		}
	}

}

// RandomPasswordApp([]string{passgen, -1})
func TestRandomPasswordNegLength(t *testing.T) {
	buf := &bytes.Buffer{}
	testWriter = buf
	n := "-1"

	app := randomPasswordApp(testWriter)
	app.Run([]string{"passgen", "-l", n})

	if !strings.Contains(buf.String(), "is not a valid password length. Must be greater than zero.") {
		t.Error("passgen -l", n, "printed '", buf.String(), "' expected invalid password length message")
	}
}

// RandomPasswordApp([]string{passgen, n}) where n>0 should print n random characters with a message
func TestRandomPasswordJustN(t *testing.T) {

	premessage := "Random password: "
	postmessage1 := " - "
	postmessage2 := " characters long\n"

	buf := &bytes.Buffer{}
	testWriter = buf
	app := randomPasswordApp(testWriter)

	for n := 0; n < 100; n++ {
		input := []string{"passgen", "-l", strconv.Itoa(n)}
		app.Run(input)

		pre := buf.String()[:len(premessage)]
		pass := buf.String()[len(premessage) : len(premessage)+n]
		post1 := buf.String()[len(premessage)+n : len(premessage)+n+len(postmessage1)]
		_ = buf.String()[len(premessage)+n+len(postmessage1) : len(premessage)+n+len(postmessage1)+len(strconv.Itoa(n))]
		post2 := buf.String()[len(premessage)+n+len(postmessage1)+len(strconv.Itoa(n)):]

		if pre != premessage {
			t.Error("Expected premessage of: '" + premessage + "' Got: '" + pre + "'")
		}
		if pass == "" && n != 0 {
			t.Error("passgen 0 should return empty password.")
		}
		if post1 != postmessage1 {
			t.Error("Expected postmessage of: '" + postmessage1 + "' Got: '" + post1 + "'")
		}
		if post2 != postmessage2 {
			t.Error("Expected postmessage of: '" + postmessage2 + "' Got: '" + post2 + "'")
		}

		// reset buffer for next test case
		buf.Reset()
	}
}

// RandomPasswordApp([]string{"passgen", n, set}) should print characters from set
func TestRandomPasswordSetMembership(t *testing.T) {

	premessage := "Random password: "

	buf := &bytes.Buffer{}
	testWriter = buf

	app := randomPasswordApp(testWriter)

	all := union(alphabet, digit, special)

	input := []string{"passgen", "-l", strconv.Itoa(0)}
	for n := 0; n < 100; n++ {
		input[2] = strconv.Itoa(n)
		app.Run(input)

		pass := buf.String()[len(premessage) : len(premessage)+n]

		for _, x := range []rune(pass) {
			if !all[x] {
				t.Error("pass =", pass, "contains", string(x), "but this is not in", all)
			}
		}

		// reset buffer for next test case
		buf.Reset()
	}
}
