package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, fmt.Errorf("%s", "This is an error!"))
}
