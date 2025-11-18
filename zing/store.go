package zing

import (
	"errors"
	"strings"

	"github.com/dgraph-io/badger/v4"
)

type Store struct {
	db *badger.DB
}

func NewStore(path string) (*Store, error) {
	opts := badger.DefaultOptions(path)

	opts.Logger = nil

	instance, err := badger.Open(opts)

	if err != nil {
		return nil, err
	}

	return &Store{db: instance}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) Exist(key string) (bool, error) {
	exist := false

	err := s.db.View(
		func(txn *badger.Txn) error {
			item, err := txn.Get([]byte(key))
			if err != nil {
				return err
			}
			if item != nil {
				exist = true
				return nil
			}
			return nil
		})

	if errors.Is(err, badger.ErrKeyNotFound) {
		err = nil
	}

	return exist, err
}

func (s *Store) Get(key string) (string, error) {
	value := ""

	err := s.db.View(
		func(txn *badger.Txn) error {
			item, err := txn.Get([]byte(key))
			if err != nil {
				return err
			}
			valueByte, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			value = string(valueByte)
			return nil
		},
	)

	return value, err
}

func (s *Store) Set(key, data string) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(data))
	})
}

func (s *Store) Remove(key string) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

func (s *Store) List() (string, error) {
	var b strings.Builder

	err := s.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.KeyCopy(nil)
			v, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			b.WriteString(string(k))
			b.WriteString(": ")
			b.WriteString(string(v))
			b.WriteByte('\n')
		}

		return nil
	})

	return b.String(), err
}
