package actioninfo

import "fmt"

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, v := range dataset {
		if err := dp.Parse(v); err != nil {
			fmt.Printf("error parsing: %v\n", err)
			continue

		}
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Error formating: %v\n", err)
			continue
		}
		fmt.Println(info)
	}

}
