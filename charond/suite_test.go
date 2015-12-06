package main

import (
	"database/sql"
	"testing"
	"time"
)

type postgresSuite struct {
	postgres   *sql.DB
	user       UserRepository
	permission PermissionRepository
	group      GroupRepository
}

func (ps *postgresSuite) setup(t *testing.T) {
	config.parse()
	var err error

	ps.postgres, err = sql.Open("postgres", config.postgres.connectionString)
	if err != nil {
		t.Errorf("connection to postgres failed with error: %s", err.Error())
		t.FailNow()
	}

	setupDatabase(ps.postgres)

	ps.user = newUserRepository(ps.postgres)
}

func (ps *postgresSuite) teardown(t *testing.T) {
	if err := tearDownDatabase(ps.postgres); err != nil {
		t.Errorf("unexpected error during database teardown: %s", err.Error())
	}

	ps.postgres.Close()
}

func assertf(t *testing.T, is bool, msg string, args ...interface{}) bool {
	if !is {
		t.Errorf(msg, args...)
	}

	return is
}

func assert(t *testing.T, is bool, msg string) bool {
	if !is {
		t.Errorf(msg)
	}

	return is
}

func assertfTime(t *testing.T, tm *time.Time, msg string, args ...interface{}) bool {
	if tm == nil || tm.IsZero() {
		t.Errorf(msg, args...)
		return false
	}

	return true
}
