passgen
=======

Passgen is a simple random password generator.

Usage
=====

Once compiled, you can run commands like the following examples:

passgen 20
- This will generate a random string of length 20 from lowercase, uppercase, digits and symbols (no whitespace)

passgen 15 a1
- This will generate a random string of length 15 from lowercase, uppercase and digits. The a1 can be any combination of letter and digit.

passgen 30 a%
- This will generate a random string of length 30 from lowercase, uppercase and symbols. Again the a% could have just as easily been r*.

passgen 1 a
- This will generate a random string of length 1 from lowercase and uppercase.