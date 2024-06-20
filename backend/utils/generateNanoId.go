package utils

import "github.com/aidarkhanov/nanoid"

func GenerateNanoId() string {
	id, err := nanoid.Generate(nanoid.DefaultAlphabet, 5)

	if err != nil {
		panic(err)
	}

	return id
}
