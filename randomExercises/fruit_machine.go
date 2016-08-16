package main


import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"strconv"
)

const price = 2

func main() {
	player_money := 50
	slot_machine_money := 10000

	//bla bla if above 0
	game := true

	for game {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Play more?: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		if text == "n" {
			game = false
		}
		player_money -= price
		playRound(&player_money, &slot_machine_money)
	}

}

func playRound(player_money, slot_machine_money *int) {
	slot_windows := [4]int {}

	jackpot := true
	two_in_a_row := false
	all_different := true

	for index := range slot_windows{
		color := (rand.Intn(4)+1)
		slot_windows[index] = color
		if index > 0 {
			if slot_windows[index-1] == color{
				two_in_a_row = true
			}else{
				jackpot = false
			}
			for i := 0; i < index; i++ {
				if slot_windows[i] == color {
					all_different = false
				}
			}
		}
	}

	win_sum := 0
	if jackpot {
		win_sum = *slot_machine_money
		*slot_machine_money = 0
	} else if two_in_a_row {
		win_sum = price*5
		*slot_machine_money -= win_sum
	} else if all_different {
		win_sum = (*slot_machine_money/2)
		*slot_machine_money /= 2
	}

	*player_money += win_sum
	fmt.Printf("%v", slot_windows)
	fmt.Println("\nYou won: " + strconv.Itoa(win_sum))
	fmt.Println("Your new total: " + strconv.Itoa(*player_money))
}
