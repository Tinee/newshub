package main

import (
	"fmt"
	"os"

	"github.com/Tinee/newshub/theskimm"
)

func main() {

	//"https://www.theskimm.com/news/daily-skimm/2020-03-05"
	parser := theskimm.NewParser()
	news, err := parser.Parse("https://www.theskimm.com/news/daily-skimm/2020-03-05")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
