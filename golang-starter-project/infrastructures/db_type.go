package infrastructures

type DBType int

const (
	POSTGRES DBType = iota
	MYSQL
	SQLSERVER
)
