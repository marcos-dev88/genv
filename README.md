# genv ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/marcos-dev88/genv?color=%2329D6D8&logo=Go&logoColor=%2329D6D8)
This repository contains a library to define environment varibales by file.

## Clone the project
```bash
$ git clone https://github.com/marcos-dev88/genv
$ cd example
```

## Example of usage
1. Defining env files:
```go
import (
	"log"
	"os"
	"github.com/marcos-dev88/genv"
)

func init() {
	if err := genv.New(".env", ".env.example"); err != nil {
		log.Printf("error -> %v", err)
	}
}

func main() {
	log.Printf("defined env: %s\ndefined example from other .env file: %s", os.Getenv("TEST_KEY"), os.Getenv("TEST_EXAMPLE_KEY"))
}
```

2. Not files defined in function:
```go
import (
	"log"
	"os"
	"github.com/marcos-dev88/genv"
)

func init() {
	if err := genv.New(); err != nil {
		log.Printf("error -> %v", err)
	}
}

func main() {
	log.Printf("defined env: %s\ndefined example from other .env file: %s", os.Getenv("TEST_KEY"), os.Getenv("TEST_EXAMPLE_KEY"))
}
```
