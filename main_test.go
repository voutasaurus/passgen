package main

import (
	"testing"
	"io"
	"bytes"
	"os"
	"strconv"
)

var testWriter io.Writer = os.Stdout

// RandomPassword([]string{}) should print an error about not enough args
func TestRandomPasswordNoArgs(t *testing.T) {
	buf := &bytes.Buffer{}
	testWriter = buf
	randomPassword([]string{}, testWriter)

	message := "Called with no executable name or other arguments. This function needs an executable name."
	messageln := message + "\n"

	if buf.String() != messageln {
		t.Error("randomPassword([]string{}) printed '", buf.String(), "' expected '", messageln, "'")
	}
}

// RandomPassword([]string{passgen}) should print err about not enough args
func TestRandomPasswordOneArg(t *testing.T) {
	buf := &bytes.Buffer{}
	testWriter = buf
	randomPassword([]string{"passgen"}, testWriter)

	message := "passgen must be called with at least one argument.\nThat is, a length for the password to be generated."
	messageln := message + "\n"

	if buf.String() != messageln {
		t.Error("passgen printed '", buf.String(), "' expected '", messageln, "'")
	}
}

// RandomPassword([]string{passgen, n}) n is not a number should print error about atoi or invalid length arg
func TestRandomPasswordNAN(t *testing.T) {
	buf := &bytes.Buffer{}
	testWriter = buf
	n := "stuff"
	randomPassword([]string{"passgen", n}, testWriter)

	message := string(n)+" is not a valid password length."
	messageln := message + "\n"

	if buf.String() != messageln {
		t.Error("passgen", n, "printed '", buf.String(), "' expected '", messageln, "'")
	}
}

// RandomPassword([]string{passgen, n}) where n>0 should print n random characters with a message
func TestRandomPasswordJustN(t *testing.T) {

	premessage := "Random password: "
	postmessage1 := " - "
	postmessage2 := " characters long\n"

	buf := &bytes.Buffer{}
	testWriter = buf

	for n := 0; n < 100; n++ {
		input := []string{"passgen", strconv.Itoa(n)}
		randomPassword(input, testWriter)

		pre := buf.String()[:len(premessage)]
		pass := buf.String()[len(premessage):len(premessage)+n]
		post1 := buf.String()[len(premessage)+n:len(premessage)+n+len(postmessage1)]
		_ = buf.String()[len(premessage)+n+len(postmessage1):len(premessage)+n+len(postmessage1)+len(strconv.Itoa(n))]
		post2 := buf.String()[len(premessage)+n+len(postmessage1)+len(strconv.Itoa(n)):]

		if pre != premessage {
			t.Error("Expected premessage of: '"+premessage+"' Got: '"+pre+"'")			
		}
		if pass == "" && n != 0 {
			t.Error("passgen 0 should return empty password.")
		}
		if post1 != postmessage1 {
			t.Error("Expected postmessage of: '"+postmessage1+"' Got: '"+post1+"'")			
		}
		if post2 != postmessage2 {
			t.Error("Expected postmessage of: '"+postmessage2+"' Got: '"+post2+"'")			
		}

		// reset buffer for next test case
		buf.Reset()	
	}
}

// RandomPassword([]string{"passgen", n, set}) should print characters from set
func TestRandomPasswordSetMembership(t *testing.T) {

	premessage := "Random password: "

	buf := &bytes.Buffer{}
	testWriter = buf

	all := union(alphabet,digit,special)

	input := []string{"passgen", strconv.Itoa(0)}
	for n := 0; n < 100; n++ {
		input[1] = strconv.Itoa(n)
		randomPassword(input, testWriter)

		pass := buf.String()[len(premessage):len(premessage)+n]

		for _, x := range []rune(pass) {
			if !all[x] {
				t.Error("pass =", pass, "contains", x, "but this is not in", all)
			}
		}

		// reset buffer for next test case
		buf.Reset()	
	}
}

// RandomPassword([]string{"passgen", a,b,c}) should print an error about too many args
func TestRandomPassword(t *testing.T) {
	buf := &bytes.Buffer{}
	testWriter = buf

	executableName := "passgen"

	randomPassword([]string{executableName, "10", "a@3", "omglol"}, testWriter)

	message := executableName+" must be called with at most two arguments.\n"
	message += "That is, a length for the password to be generated, and a combination of characters representing the desired character sets\n"

	if buf.String() != message {
		t.Error("passgen 10 a@3 omglol printed '", buf.String(), "' expected '", message, "'")
	}
}
