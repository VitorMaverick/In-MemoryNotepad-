package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readCommandAndValidateFields() {
	var list []string
	var command string = ""
	var text string
	var tamanho int

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the maximum number of notes: ")
	fmt.Scan(&tamanho)

	for command != "exit" {
		fmt.Println("Enter command and data:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fields := strings.Fields(input)
		command = fields[0]
		text = strings.Join(fields[1:], " ")

		switch command {
		case "exit":
			fmt.Println("[Info] Bye!")
			continue
		case "create":
			if len(text) <= 0 {
				fmt.Println("[Error] Missing note argument")
				continue
			}
			createPassandoListaPorReferencia(&list, text, tamanho)
		case "list":
			listar(list)
		case "clear":
			list = clear(list)
		case "update":
			if len(text) <= 0 {
				fmt.Println("[Error] Missing position argument")
				continue
			}
			position, err := strconv.ParseInt(fields[1], 10, 0)
			if err != nil {
				fmt.Printf("[Error] Invalid position: %v\n", fields[1])
				continue
			} else if len(fields) == 2 {
				fmt.Println("[Error] Missing note argument")
				continue
			}
			if int64(tamanho) < position-1 {
				fmt.Println("[Error] There is nothing to update")
				continue
			}

			argument := strings.Join(fields[2:], " ")
			update(&list, argument, position-1)
		case "delete":
			if len(fields) < 2 {
				fmt.Println("[Error] Missing position argument")
				continue
			}
			position, err := strconv.ParseInt(fields[1], 10, 0)
			if err != nil {
				fmt.Printf("[Error] Invalid position: %v\n", fields[1])
				continue
			}

			list = delete(list, position-1, tamanho)
		default:
			fmt.Println("[Error] Unknown command")
		}

	}
}

func createSimples(list []string, text string, tamanho int) []string {
	if len(list) >= tamanho {
		fmt.Println("[Error] Notepad is full")
		return list
	}

	fmt.Println("[OK] The note was successfully created")

	return append(list, text)
}

func createPassandoListaPorReferencia(list *[]string, text string, tamanho int) {
	if len(*list) >= tamanho {
		fmt.Println("[Error] Notepad is full")
		return
	}

	fmt.Println("[OK] The note was successfully created")

	*list = append((*list), text)
}

func update(list *[]string, argument string, position int64) {

	for i, _ := range *list {
		if int64(i) == position {
			(*list)[i] = argument
		}
	}

	fmt.Printf("[OK] The note at position %v was successfully updated\n", position+1)
}

func delete(list []string, position int64, tamanho int) []string {
	if position > int64(tamanho) {
		fmt.Printf("[Error] Position %v is out of the boundaries [1, %v]\n", position+1, tamanho)
		return list
	}

	fmt.Printf("[OK] The note at position %v was successfully deleted\n", position+1)

	return append(list[:position], list[position+1:]...)
}

func listar(list []string) {
	if len(list) <= 0 {
		fmt.Println("[Info] Notepad is empty")
	} else {
		for i, s := range list {
			fmt.Printf("[Info] %v: %v\n", i+1, s)
		}
	}
}

func clear(list []string) []string {
	list = nil
	fmt.Println("[OK] All notes were successfully deleted")
	return list
}
func main() {
	readCommandAndValidateFields()

}
