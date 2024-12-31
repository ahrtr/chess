package rules

var pieceValueMap map[PieceRole]int = map[PieceRole]int{
	RoleKing:   100,
	RoleRook:   10,
	RoleHorse:  5,
	RoleCannon: 5,
	RoleBishop: 2,
	RoleGuard:  2,
	RoleSolder: 1,
}
