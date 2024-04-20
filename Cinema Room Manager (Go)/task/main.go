package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	//"strconv"
)

var (
	rows          int
	seats         int
	userRow       int
	userSeat      int
	userInput     int
	currentIncome int
	cinema        [][]rune
)

func main() {
	fmt.Println("Enter the number of rows:")
	askUserInput("Enter the number of rows:", &rows)
	askUserInput("Enter the number of seats in each row:", &seats)
	initCinema()
	for {
		printMenu()
		askUserInput("", &userInput)
		switch userInput {
		case 1:
			printCinema()
		case 2:
			for {
				askUserInput("Enter a row number:", &userRow)
				askUserInput("Enter a seat number in that row:", &userSeat)
				if userRow < 1 || userRow > rows || userSeat < 1 || userSeat > seats {
					fmt.Println("Wrong input!")
					continue
				}
				if cinema[userRow-1][userSeat-1] == 'B' {
					fmt.Println("That ticket has already been purchased!")
				} else {
					printTicketPrice()
					break
				}
			}
		case 3:
			printStat()
		case 0:
			os.Exit(0)
		}
	}
}

func initCinema() {
	for row := 0; row < rows; row++ {
		cinema = append(cinema, []rune{})
		for seat := 0; seat < seats; seat++ {
			cinema[row] = append(cinema[row], 'S')
		}
	}
}

func printStat() {
	fmt.Printf("Number of purchased tickets: %v\n", calculatePurchasedTickets())
	fmt.Printf("Percentage: %.2f%%\n", (float64(calculatePurchasedTickets())/float64(rows*seats))*100)
	fmt.Printf("Current income: $%v\n", currentIncome)
	fmt.Printf("Total income: $%v\n", calculateTotalIncome())
}

func calculatePurchasedTickets() int {
	tickets := 0
	for _, row := range cinema {
		for _, seat := range row {
			if seat == 'B' {
				tickets++
			}
		}
	}
	return tickets
}

func printMenu() {
	fmt.Println("1. Show the seats")
	fmt.Println("2. Buy a ticket")
	fmt.Println("3. Statistics")
	fmt.Println("0. Exit")
}

func calculateTotalIncome() int {
	if rows*seats <= 60 {
		return rows * seats * 10
	}
	return (rows/2)*seats*10 + (rows-rows/2)*seats*8
}

func printTicketPrice() {
	price := 0
	if rows*seats <= 60 {
		price = 10
	} else {
		if rows/2 >= userRow {
			price = 10
		} else {
			price = 8
		}
	}
	cinema[userRow-1][userSeat-1] = 'B'
	fmt.Printf("Ticket price: $%v\n", price)
	currentIncome += price
	printCinema()
}

func askUserInput(str string, inputVar *int) {
	fmt.Println(str)
	_, err := fmt.Scan(inputVar)
	if err != nil {
		log.Fatal(err)
	}
}

func printCinema() {
	fmt.Println("Cinema:")
	rowsToPrint := "  "
	for i := 0; i <= rows; i++ {
		rowsToPrint += strconv.Itoa(i+1) + " "
	}
	fmt.Println(rowsToPrint)
	for row, _ := range cinema {
		str := fmt.Sprintf("%v ", row+1)
		for i, _ := range cinema[row] {
			str += string(cinema[row][i]) + " "
		}
		fmt.Println(str)
	}
	//fmt.Println("Total income:")
	//fmt.Printf("$%v\n", calculateTotalIncome())
}
