# dotenv

**A simple .env file loader for go(golang)**

![Test](https://github.com/radulucut/dotenv/actions/workflows/test.yml/badge.svg)

## Install

`go get github.com/radulucut/dotenv`

## Usage

.env

```
MY_VAR=some value
SECRET_KEY="s3cret"
SINGLE_QUOTE='some "value"'
```

main.go

```
package main

import (
	"log"
	"os"

	"github.com/radulucut/dotenv"
)

func main() {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	myVar := os.Getenv("MY_VAR")
	secretKey := os.Getenv("SECRET_KEY")
	singleQuote := os.Getenv("SINGLE_QUOTE")

	// ...
}
```

### Notes:

- Overrides existing variables
- Does not trim spaces
- Does not support comments
