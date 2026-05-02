// GRID-SAST-007: data race — concurrent map access without sync.
package cache

import "time"

type Store struct {
	data map[string]string
}

func New() *Store { return &Store{data: map[string]string{}} }

func (s *Store) Set(k, v string)           { s.data[k] = v }   // unsafe concurrent write
func (s *Store) Get(k string) (string, bool) { v, ok := s.data[k]; return v, ok } // unsafe concurrent read

func (s *Store) ExpireSweeper() {
	for range time.Tick(time.Minute) {
		for k := range s.data {
			delete(s.data, k)   // concurrent map iteration + write
		}
	}
}
