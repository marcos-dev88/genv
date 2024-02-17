package genv

import (
	"log"
	"testing"
)

// go test -v -count=1 -failfast -run ^Test_New$
func Test_New(t *testing.T) {

	t.Run("new_without_env_defined", func(*testing.T) {
		if err := New(); err != nil {
			log.Print(err)
		}
	})

	t.Run("new_with_env_defined", func(*testing.T) {
		if err := New(".env.example"); err != nil {
			log.Print(err)
		}
	})

	t.Run("new_with_env_defined_error", func(*testing.T) {
		if err := New(".envsh"); err != nil {
			log.Print(err)
		}
	})
}

// go test -v -count=1 -failfast -run ^Test_NewFast$
func Test_NewFast(t *testing.T) {

	t.Run("new_without_env_defined", func(*testing.T) {
		if err := NewFast(); err != nil {
			log.Print(err)
		}
	})

	t.Run("new_with_env_defined", func(*testing.T) {
		if err := NewFast(".env.example"); err != nil {
			log.Print(err)
		}
	})

	t.Run("new_with_env_defined_error", func(*testing.T) {
		if err := NewFast(".envsh"); err != nil {
			log.Print(err)
		}
	})
}
