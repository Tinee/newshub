package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Tinee/newshub/theskimm"
)

func main() {

	//"https://www.theskimm.com/news/daily-skimm/2020-03-05"
	parser := theskimm.NewParser("https://www.theskimm.com/news/daily-skimm/2020-03-05")
	t := time.Now()

	fmt.Println(t.Format("06-01-02"))
	_, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
