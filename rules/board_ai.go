package rules

func (b *Board) GetBestMove(depth int) Move {
	var (
		bestMove  Move
		bestScore = -1000000
		color     = b.color()
	)

	moves := b.validMoves()
	for _, move := range moves {
		cloneBoard := b.Clone()
		cloneBoard.move(move.from.X, move.from.Y, move.to.X, move.to.Y, true)

		score := minimax(cloneBoard, color, depth, -1000000, 1000000, false)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}

	return bestMove
}

// minimax tries to select the best policy to move. It simulates two
// players to move in turn, and evaluate the score of each move.
// Refer to
//
//	-https://cs.stanford.edu/people/eroberts/courses/soco/projects/2003-04/intelligent-search/minimax.html.
//	-https://cs.stanford.edu/people/eroberts/courses/soco/projects/2003-04/intelligent-search/alphabeta.html
//
// Parameters:
//   - b:     the board of the game
//   - color: the color of the player against which the evaluation is performed
//   - depth: the steps that we are trying to explore
//   - alpha: the score of the best choice we have found so far at any choice point along the path for the maximizer (usually me)
//   - beta:  the score of the best choice we have found so far at any choice point along the path for the minimizer (usually my opponent)
//   - isMaximizing: which side to evaluate, the maximizer (true) or the minimizer (false)?
func minimax(b *Board, color PieceColor, depth int, alpha, beta int, isMaximizing bool) int {
	if depth == 0 || b.isGameOver() {
		return evaluate(b, color)
	}

	moves := b.validMoves()

	if isMaximizing {
		// I am trying to maximize my score.
		maxEval := -1000000
		for _, m := range moves {
			cloneBoard := b.Clone()
			cloneBoard.move(m.from.X, m.from.Y, m.to.X, m.to.Y, true)

			eval := minimax(cloneBoard, color, depth-1, alpha, beta, false)
			maxEval = max(maxEval, eval)
			alpha = max(alpha, eval)

			// purge: The minimizer at the parent or at any choice point
			// further up won't reach this node as it's greater than the
			// best score that he/she is aware of.
			if beta <= alpha {
				break
			}
		}
		return maxEval
	} else {
		// My opponent is trying to minimize my score.
		minEval := 1000000
		for _, m := range moves {
			cloneBoard := b.Clone()
			cloneBoard.move(m.from.X, m.from.Y, m.to.X, m.to.Y, true)

			eval := minimax(cloneBoard, color, depth-1, alpha, beta, true)
			minEval = min(minEval, eval)
			beta = min(beta, eval)

			// purge: The maximizer at the parent or at any choice point
			// further up definitely won't reach this node as it's less
			// than the best score that he/she is aware of.
			if beta <= alpha {
				break
			}
		}
		return minEval
	}
}

func evaluate(b *Board, color PieceColor) int {
	score := 0
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 8; j++ {
			p := b.pieceMatrix[i][j]
			if p == nil {
				continue
			}
			if p.color == color {
				score += pieceValueMap[p.role]
			} else {
				score -= pieceValueMap[p.role]
			}
		}
	}

	return score
}
