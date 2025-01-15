package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	uid := uuid.New()
	fmt.Println(fmt.Sprintf("_%s", uid.String()))
}
