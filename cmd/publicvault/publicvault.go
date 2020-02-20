package main

import (
	"fmt"

	"github.com/lpinter/publicvault"
	"github.com/lpinter/publicvault/internal/auth"
	"github.com/lpinter/publicvault/secret"
)

func main() {
	fmt.Println(publicvault.Config())
	fmt.Println(secret.GetSecret())
	fmt.Println(auth.GetAuth())
}
