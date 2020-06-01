package main

import (
	"bufio"
	"errors"
	"example/tic_tac_toe_app/components"
	"example/tic_tac_toe_app/service"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// select board size
	readobj := bufio.NewReader(os.Stdin)
	var board *components.Board
	fmt.Print("\n\n------------------------TIC TAC TOE------------------------\n\n")
	fmt.Print("select board size\n\n2 for 2 X 2\n3 for 3 X 3\n4 for 4 X 4\n\n")

	for {
		fmt.Println("enter number : ")
		inp, err := readobj.ReadString('\n')
		if checkWarning(err) {
			continue
		}
		size, err := getUint8(inp)
		if checkWarning(err) {
			continue
		}
		if size < 2 || size > 4 {
			checkWarning(errors.New("Enter number in range (2,4)"))
			continue
		}
		board = components.CreateBoard(size)
		break
	}

	//enter player1

	var player1 *components.Player
	fmt.Println("Enter name of first player : ")
	p1Name, err := readobj.ReadString('\n')
	checkerror(err)
	p1Name = strings.TrimSpace(p1Name)
	for {
		fmt.Println("select your mark X or O : ")
		p1Mark, err := readobj.ReadString('\n')
		if checkWarning(err) {
			continue
		}
		p1Mark = strings.TrimSpace(p1Mark)
		if p1Mark != components.OMark && p1Mark != components.XMark {
			checkWarning(errors.New("mark was neither X nor O"))
			continue
		}
		player1 = components.NewPlayer(p1Name, p1Mark)
		break
	}

	//enter player2

	var player2 *components.Player
	p2Name := ""
	for {
		fmt.Println("Enter name of second player : ")
		p2Name, err = readobj.ReadString('\n')
		checkerror(err)
		p2Name = strings.TrimSpace(p2Name)
		if p2Name == player1.Name {
			checkWarning(errors.New("player name is already taken"))
			continue
		}
		break
	}
	p2Mark := ""
	if player1.Mark == components.OMark {
		p2Mark = components.XMark
	} else {
		p2Mark = components.OMark
	}
	player2 = components.NewPlayer(p2Name, p2Mark)

	// print players info

	fmt.Print("\n\n-----------------------Players Info-----------------------\n")
	fmt.Printf("\tPlayer 1 :- Name : %10s\tMark : %s\n", player1.Name, player1.Mark)
	fmt.Printf("\tPlayer 2 :- Name : %10s\tMark : %s\n", player2.Name, player2.Mark)
	fmt.Print("----------------------------------------------------------\n\n")

	//intializing all services

	boardServiceObj := service.NewBoardService(board)
	resultserviceObj := service.NewResultService(boardServiceObj)
	gameserviceObj := service.NewGameService(resultserviceObj, [2]*components.Player{player1, player2})

	//game starts
	fmt.Println("-----------------------Game Starts------------------------")
	fmt.Println(gameserviceObj.PrintBoard())
	for {
		var res service.Result
		for {
			fmt.Print(player1.Name, " enter your pos : ")
			inp, err := readobj.ReadString('\n')
			checkerror(err)
			pos, err := getUint8(inp)
			if checkWarning(err) {
				continue
			}
			res, err = gameserviceObj.Play(pos)
			if checkWarning(err) {
				continue
			}
			break
		}
		fmt.Println(gameserviceObj.PrintBoard())
		if res.Win == true {
			fmt.Print(res.CurrResult)
			//fmt.Printf("--------------------%s Won--------------------", player1.Name)
			break
		} else if res.Draw == true {
			fmt.Print(res.CurrResult)
			//fmt.Println("--------------------Draw--------------------")
			break
		}

		for {
			fmt.Print(player2.Name, " enter your pos : ")
			inp, err := readobj.ReadString('\n')
			checkerror(err)
			pos, err := getUint8(inp)
			if checkWarning(err) {
				continue
			}
			res, err = gameserviceObj.Play(pos)
			if checkWarning(err) {
				continue
			}
			break
		}
		fmt.Println(gameserviceObj.PrintBoard())
		if res.Win == true {
			fmt.Print(res.CurrResult)
			//fmt.Printf("--------------------%s Won--------------------", player2.Name)
			break
		} else if res.Draw == true {
			fmt.Print(res.CurrResult)
			//fmt.Println("--------------------Draw--------------------")
			break
		}

	}

}

func getUint8(numstring string) (uint8, error) {
	numstring = strings.TrimSpace(numstring)
	num, err := strconv.Atoi(numstring)
	if err != nil {
		return 0, errors.New("please enter integer")
	}
	return uint8(num), nil
}

func checkerror(err error) { // checkerror and terminates the program
	if err != nil {
		log.Fatal(err)
	}
}

func checkWarning(err error) bool { // checkerror and print if any
	if err != nil {
		fmt.Println("Warning : ", err)
		return true
	}
	return false
}
