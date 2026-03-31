package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	if len(dataset) == 0 {
		log.Println("dataset is empty")
		return
	}

	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Printf("error parsing item %s: %v", v, err)
			continue
		}
	}

	msg, err := dp.ActionInfo()
	if err != nil {
		log.Printf("error getting action info: %v", err)
	}
	fmt.Println(msg)
}
