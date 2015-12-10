passgen [![Build Status](https://travis-ci.org/voutasaurus/passgen.svg?branch=master)](https://travis-ci.org/voutasaurus/passgen)
=======

passgen is a simple random password generator.

Usage
=====

Once passgen is compiled and added to PATH, you can run commands like the following examples:

passgen
- By default passgen will generate a random string of length 20 from lowercase, uppercase, digits and symbols (no whitespace)

passgen -l 10
- This will generate a random string of length 10 from lowercase, uppercase, digits and symbols (no whitespace)

passgen -l 15 -c a1
- This will generate a random string of length 15 from lowercase, uppercase and digits. The a1 can be any combination of letter and digit.

passgen -l 30 -c a%
- This will generate a random string of length 30 from lowercase, uppercase and symbols. Again the a% could have just as easily been r*.

passgen -l 1 -c a
- This will generate a random string of length 1 from lowercase and uppercase.

passgen -h
- This will print the help output.


Build Instructions
==================

First install and set up Go:
Install the latest version of Go (https://golang.org/dl/)

Set $GOPATH and create three directories in $GOPATH called bin, src and pkg.

Install the app (method 1):
Add $GOPATH/bin to $PATH

Run the following commands:
- go get github.com/voutasaurus/passgen
- cd $GOPATH/src/github.com/voutasaurus/passgen
- go install

Install the app (alternative method):

Run the following commands:
- go get github.com/voutasaurus/passgen
- cd $GOPATH/src/github.com/voutasaurus/passgen
- go build

Copy the resultant executable (passgen) to a path already in $PATH.
