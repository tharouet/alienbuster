package score

type Score struct {
	EnemyCounter int
	Score        int
	Level        int
}

var Board = Score{
	EnemyCounter: 0,
	Score:        0,
	Level:        1,
}
