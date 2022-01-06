package main

import "fmt"

type player struct {
	position int
	score    int
}

type game struct {
	player_one player
	player_two player
}

func roll_die(i *int) int {
	*i++
	result := *i % 100
	if result == 0 {
		result = 100
	}
	return result
}

func play_turn(p player, roll int) player {
	new_player := player{p.position, p.score}
	new_player.position = (p.position + roll) % 10
	if new_player.position == 0 {
		new_player.position = 10
	}
	new_player.score += new_player.position
	return new_player
}

func play_practice_game(practice_game game) {
	i := 0
	for practice_game.player_one.score < 1000 && practice_game.player_two.score < 1000 {
		player := i%2 + 1
		roll := roll_die(&i) + roll_die(&i) + roll_die(&i)
		if player == 1 {
			practice_game.player_one = play_turn(practice_game.player_one, roll)
		} else if player == 2 {
			practice_game.player_two = play_turn(practice_game.player_two, roll)
		}
	}

	if practice_game.player_one.score < 1000 {
		fmt.Println("Pratice Game Score:", practice_game.player_one.score*i)

	} else if practice_game.player_two.score < 1000 {
		fmt.Println("Practice Game Score:", practice_game.player_two.score*i)
	}
}

func generate_roll_map() map[int]int {
	roll_map := make(map[int]int)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				roll_map[i+j+k]++
			}
		}
	}
	return roll_map
}

func play_quantum_turn(results_map map[game]int64, player int) map[game]int64 {
	new_results := make(map[game]int64)
	roll_map := generate_roll_map()
	for g, g_count := range results_map {
		new_game := g
		updated := false
		if g.player_one.score < 21 && g.player_two.score < 21 {
			for roll, roll_count := range roll_map {
				if player == 1 {
					new_player_one := play_turn(g.player_one, roll)
					new_game = game{player_one: new_player_one, player_two: g.player_two}
					new_results[new_game] += int64(roll_count) * g_count

				} else if player == 2 {
					new_player_two := play_turn(g.player_two, roll)
					new_game = game{player_one: g.player_one, player_two: new_player_two}
					new_results[new_game] += int64(roll_count) * g_count
				}
			}
			updated = true
		}
		if !updated {
			new_results[new_game] += g_count
		}
	}
	return new_results
}

func check_results(results_map map[game]int64) bool {
	for g := range results_map {
		if g.player_one.score < 21 && g.player_two.score < 21 {
			return true
		}
	}
	return false
}

func play_game(main_game game) {
	results_map := make(map[game]int64)
	results_map[main_game] = 1
	for check_results(results_map) {
		results_map = play_quantum_turn(results_map, 1)
		results_map = play_quantum_turn(results_map, 2)
	}
	player_one_win := int64(0)
	player_two_win := int64(0)
	for g, g_count := range results_map {
		if g.player_one.score > g.player_two.score {
			player_one_win += g_count
		} else if g.player_one.score < g.player_two.score {
			player_two_win += g_count
		}
	}
	fmt.Println("Player One Wins:", player_one_win, "Player Two Wins:", player_two_win)
}

func main() {
	play_practice_game(game{player_one: player{position: 4, score: 0}, player_two: player{position: 8, score: 0}})
	play_practice_game(game{player_one: player{position: 6, score: 0}, player_two: player{position: 3, score: 0}})
	play_game(game{player_one: player{position: 4, score: 0}, player_two: player{position: 8, score: 0}})
	play_game(game{player_one: player{position: 6, score: 0}, player_two: player{position: 3, score: 0}})
}
