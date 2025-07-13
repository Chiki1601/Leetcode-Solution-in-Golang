import "sort"

func matchPlayersAndTrainers(players, trainers []int) int {
    sort.Ints(players)
    sort.Ints(trainers)

    i, j, matches := 0, 0, 0
    for i < len(players) && j < len(trainers) {
        if players[i] <= trainers[j] {
            matches++
            i++
        }
        j++
    }
    return matches
}
