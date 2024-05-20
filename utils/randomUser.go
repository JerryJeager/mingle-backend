package utils

import (
    "github.com/goombaio/namegenerator"
	"time"
)

func RandomUserName() string {
    seed := time.Now().UTC().UnixNano()
    nameGenerator := namegenerator.NewNameGenerator(seed)

    name := nameGenerator.Generate()

    return name
}