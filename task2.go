package main

import "fmt"

type Employee struct {
	Name     string // имя
	Age      int    // возраст
	Position string // позиция
	Salary   int    // зарплата
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

func task2() {
	const size = 512
	empls := [size]*Employee{}
	for {
		cmd := 0
		fmt.Print(commands)
		fmt.Scanf("%d", &cmd)

		var discard string
		fmt.Scanln(&discard)

		switch cmd {
		case 1:
			// Добавляем нового сотрудника
			empl := new(Employee)
			fmt.Println("\nИмя:")
			fmt.Scanf("%s", &empl.Name)
			fmt.Scanln(&discard)

			fmt.Println("Возраст:")
			fmt.Scanf("%d", &empl.Age)
			fmt.Scanln(&discard)

			fmt.Println("Позиция:")
			fmt.Scanf("%s", &empl.Position)
			fmt.Scanln(&discard)

			fmt.Println("Зарплата:")
			fmt.Scanf("%d", &empl.Salary)
			fmt.Scanln(&discard)

			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					break
				}
			}
		case 2:
			fmt.Println("Введите имя сотрудника для удаления:")
			var name string
			fmt.Scanf("%s", &name)

			for i := 0; i < size; i++ {
				if empls[i] != nil && empls[i].Name == name {
					empls[i] = nil
					fmt.Printf("Сотрудник %s успешно удален!\n", name)
					break
				}
			}
		case 3:
			fmt.Printf("\n")
			count := 0
			for i := 0; i < size; i++ {
				if empls[i] != nil {
					count++
					fmt.Printf("%d %s %d %s %d\n",
						count,
						empls[i].Name,
						empls[i].Age,
						empls[i].Position,
						empls[i].Salary)
				}
			}
		case 4:
			return
		}
	}
}
