package main

import (
	"fmt"
	"time"

	"github.com/lpinter/publicvault"
	"github.com/lpinter/publicvault/internal/auth"
	"github.com/lpinter/publicvault/secret"
)

// Main function here
func main() {
	fmt.Println(publicvault.Config())
	fmt.Println(secret.GetSecret())
	fmt.Println(auth.GetAuth())
	seconds := 60 * time.Minute
	fmt.Printf("Sleeping %v \n", seconds)
	time.Sleep(seconds)
	fmt.Println("Done")
}
