package quiz

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
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

func (l *List) Play(limit int) {
	var correct int = 0
	timer := time.NewTimer(time.Duration(limit) * time.Second)

	func() {
		for _, line := range *l {
			fmt.Printf("Problem: %s\n", line.question)
			ansChan := make(chan string)

			go func() {
				var answer string
				fmt.Scanf("%s", &answer)
				ansChan <- answer
			}()

			select {
			case <-timer.C:
				return
			case ans := <-ansChan:
				if ans == line.answer {
					correct++
				}
			}
		}
	}()
	fmt.Printf("%d correct answer out of %d\n", correct, len(*l))
}
