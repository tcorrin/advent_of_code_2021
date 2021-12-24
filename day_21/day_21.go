package main

import "fmt"

func roll_die(i *int) int {
	*i++
	result := *i % 100
	if result == 0 {
		result = 100
	}
	return result
}

func play_turn(pos *int, score *int, i *int) {
	roll := roll_die(i) + roll_die(i) + roll_die(i)
	*pos = (*pos + roll) % 10
	if *pos == 0 {
		*pos = 10
	}
	*score += *pos
}

func play_practice_game(player_one_pos int, player_two_pos int) {
	player_one_score := 0
	player_two_score := 0
	i := 0
	for player_one_score < 1000 && player_two_score < 1000 {
		if i%2 == 0 {
			play_turn(&player_one_pos, &player_one_score, &i)
		} else if i%2 == 1 {
			play_turn(&player_two_pos, &player_two_score, &i)
		}
	}

	if player_one_score < 1000 {
		fmt.Println("Pratice Game Score:", player_one_score*i)

	} else if player_two_score < 1000 {
		fmt.Println("Practice Game Score:", player_two_score*i)
	}
}

func play_game(player_one_pos int, player_two_pos int) {
	player_one_score := 0
	player_two_score := 0
	i := 0
	for player_one_score < 21 && player_two_score < 21 {
		if i%2 == 0 {
			play_turn(&player_one_pos, &player_one_score, &i)
		} else if i%2 == 1 {
			play_turn(&player_two_pos, &player_two_score, &i)
		}
	}

	if player_one_score < 21 {
		fmt.Println("Game Score:", player_one_score*i)

	} else if player_two_score < 21 {
		fmt.Println("Game Score:", player_two_score*i)
	}
}

func main() {
	play_practice_game(4, 8)
	play_practice_game(6, 3)
}
