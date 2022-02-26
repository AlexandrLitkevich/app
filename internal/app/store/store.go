package store

import (
	"database/sql"

	_ "github.com/lib/pq" //анонимный импорт без методов
)

type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

/*
ф-ия принимает config и возвращает
указатель на хранилище(store)
*/
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

/*
Метод Open используем
при инициализации и подключении к БД
*/

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DataBaseURL)
	if err != nil {
		return err
	}
	// проверяем верный ли конфиг базы данных
	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

/*
Метод Close используем
когда веб сервер заканчиват работу
мы могли отключиться от БД и еще
какие либо действия
*/

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
