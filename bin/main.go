package main

import (
	"fmt"

	peg "github.com/ridho9/pegchamp"
)

func main() {
	weatherString := peg.String("Weather")

	timeString := peg.SequenceOf(
		peg.String("("),
		peg.Choice(
			peg.String("today"),
			peg.String("yesterday"),
		),
		peg.String(")"),
	).Map(func(ps peg.ParserState) (interface{}, error) {
		return ps.Result().([]interface{})[1], nil
	})

	weatherType := peg.Choice(
		peg.String("Sunny"),
		peg.String("Rainy"),
		peg.String("Rainy"),
	)

	fullParser := peg.SequenceOf(
		weatherString,
		peg.TakeSecond(
			peg.String(" "),
			timeString,
		),
		peg.TakeSecond(
			peg.String(": "),
			weatherType,
		),
	)

	run := fullParser.Run("Weather (today): Sunny")
	fmt.Printf("%e\n", run.Error())
	fmt.Printf("%#v\n", run.Result())
}
