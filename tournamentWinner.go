package main

import "fmt"

/// 給兩個陣列 competitions 與 result, competitions 裡的陣列順序與 result 的 index 相對應
/// 如 result[0] 代表 competitions[0] 的陣列
/// competitions 裡的每一個陣列代表有兩組隊伍 [home, away] 且用 0 代表 away 得分 1 代表 home 得分
/// result[0] = 1 代表 home 得分
/// 紀錄誰得分並返回最多分的組別

func main() {
	fmt.Println(TournamentWinner([][]string{
		{"HTML", "Java"},
		{"Java", "Python"},
		{"Python", "HTML"},
		{"C#", "Python"},
		{"Java", "C#"},
		{"C#", "HTML"},
		{"SQL", "C#"},
		{"HTML", "SQL"},
		{"SQL", "Python"},
		{"SQL", "Java"},
	}, []int{0, 1, 1, 1, 0, 1, 0, 1, 1, 0}))
}

func TournamentWinner(competitions [][]string, results []int) string {
	board := map[string]int{}
	var bestTeam string
	for i, pair := range competitions {
		var winTeam string
		if results[i] == 0 { // away win
			winTeam = pair[1]
		} else {
			winTeam = pair[0]
		}
		board[winTeam] += 3
		if len(bestTeam) == 0 || board[winTeam] > board[bestTeam] {
			bestTeam = winTeam
		}
	}
	return bestTeam
}