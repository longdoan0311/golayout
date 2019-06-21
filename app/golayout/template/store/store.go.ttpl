package store

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"{{.ModName}}"
)

type Store struct {
	db    *sqlx.DB
	user {{.ProjName}}.UserStore
}

type StoreConfig struct {
	Use            string
	DBCfg          SQLConfig
}

type SQLConfig struct {
	Host           string
	Port           int
	DB             string
	User           string
	Pass           string
	SSLMode        bool
	MaxOpenCnn     int
	MaxCnnLifeTime time.Duration
}

var scfg StoreConfig

func New(cfg StoreConfig) ({{.ProjName}}.Store, error) {
	var s Store
	scfg = cfg

	switch cfg.Use {
	case "postgres":
		connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", cfg.DBCfg.Host, cfg.DBCfg.Port, cfg.DBCfg.User, cfg.DBCfg.DB, cfg.DBCfg.Pass)
		if !cfg.DBCfg.SSLMode {
			connStr += " sslmode=disable"
		}

		db, err := sqlx.Connect("postgres", connStr)
		if err != nil {
			return nil, errors.Wrap(err, "Open Postgresql DB fail")
		}

		s.db = db
		s.db.SetMaxOpenConns(cfg.DBCfg.MaxOpenCnn)
		s.db.SetConnMaxLifetime(cfg.DBCfg.MaxCnnLifeTime)

		s.user = &sqlUser{
			db: s.db,
		}

	case "imdb":
		i := newImdbUser()
		s.user = &i
	default:
		return nil, errors.New(fmt.Sprintf("Unsupported store type %s", cfg.Use))
	}

	return &s, nil
}

func (s *Store) Initialize() error {
	// TODO: fill your code here
	return nil
}

func (s *Store) User() {{.ProjName}}.UserStore {
	return s.user
}

func (s *Store) Close() error {
	return nil
}
