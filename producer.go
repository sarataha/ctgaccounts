package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* let the program behave as a producer when executed */
func mainProducer() {
	var err error
	reader := bufio.NewReader(os.Stdin)
	kafka := newKafkaSyncProducer()

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		args := strings.Split(text, "###")
		cmd := args[0]

		switch cmd {
		case "create":
			if len(args) == 2 {
				accName := args[1]
				event := NewCreateAccountEvent(accName)
				sendMsg(kafka, event)
			} else {
				fmt.Println("Only specify create###Account Name")
			}
		case "payment":
			if len(args) == 3 {
				accId := args[1]
				if amount, err := strconv.Atoi(args[2]); err == nil {
					event := NewDepositEvent(accId, amount)
					sendMsg(kafka, event)
				}
			} else {
				fmt.Println("Only specify payment###Account ID###amount")
			}
		default:
			fmt.Printf("Unknown command %s, only: create, payment\n", cmd)
		}

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			err = nil
		}
	}
}
