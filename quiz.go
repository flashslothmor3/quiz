package quiz

import (
	"encoding/csv"
	"fmt"
	"os"
)

type item struct {
	question string
	answer   string
}

type List []item

func (l *List) Load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	for _, line := range records {
		*l = append(*l, item{
			question: line[0],
			answer:   line[1],
		})
	}
	return nil
}

func (l *List) Play() {
	var correct int = 0
	var wrong int = 0
	for _, line := range *l {
		fmt.Printf("Problem: %s\n", line.question)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == line.answer {
			correct++
		} else {
			wrong++
		}
	}
	fmt.Printf("%d correct answer out of %d\n", correct, len(*l))
}
