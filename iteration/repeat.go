package main

// Although I am not a fan of this hard coding of the repeated
// value since it exsits in two differnt files, this is ok for now
// becuase this is just for pratice iterations
const repeatCount = 5

func Repeat(character string) string {

	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
