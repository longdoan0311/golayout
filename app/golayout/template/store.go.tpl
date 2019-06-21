package {{.ProjName}}

type Store interface {
	Initialize() error
	User() UserStore
	Close() error
}