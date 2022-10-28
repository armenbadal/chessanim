package chess

import "testing"

func TestGame(t *testing.T) {
	g := make(Game, 5)
	g[0] = Step{from: "d2", to: "d4"}
	g[1] = Step{from: "e7", to: "e5"}
	g[2] = Step{from: "g1", to: "f3"}
	g[3] = Step{from: "b8", to: "c6"}
	g[4] = Step{from: "c1", to: "h6"}
	g.Animation("a.gif")
}
