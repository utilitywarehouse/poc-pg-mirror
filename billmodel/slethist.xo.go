// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Slethist represents a row from 'equinox.slethist'.
type Slethist struct {
	Slhmonth   sql.NullString `json:"slhmonth"`    // slhmonth
	Slhcount   sql.NullInt64  `json:"slhcount"`    // slhcount
	EquinoxPrn sql.NullInt64  `json:"equinox_prn"` // equinox_prn
	EquinoxLrn int64          `json:"equinox_lrn"` // equinox_lrn
	EquinoxSec sql.NullInt64  `json:"equinox_sec"` // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Slethist exists in the database.
func (s *Slethist) Exists() bool {
	return s._exists
}

// Deleted provides information if the Slethist has been deleted from the database.
func (s *Slethist) Deleted() bool {
	return s._deleted
}

// Insert inserts the Slethist to the database.
func (s *Slethist) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.slethist (` +
		`slhmonth, slhcount, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, s.Slhmonth, s.Slhcount, s.EquinoxPrn, s.EquinoxSec)
	err = db.QueryRow(sqlstr, s.Slhmonth, s.Slhcount, s.EquinoxPrn, s.EquinoxSec).Scan(&s.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Update updates the Slethist in the database.
func (s *Slethist) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if s._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.slethist SET (` +
		`slhmonth, slhcount, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE equinox_lrn = $5`

	// run query
	XOLog(sqlstr, s.Slhmonth, s.Slhcount, s.EquinoxPrn, s.EquinoxSec, s.EquinoxLrn)
	_, err = db.Exec(sqlstr, s.Slhmonth, s.Slhcount, s.EquinoxPrn, s.EquinoxSec, s.EquinoxLrn)
	return err
}

// Save saves the Slethist to the database.
func (s *Slethist) Save(db XODB) error {
	if s.Exists() {
		return s.Update(db)
	}

	return s.Insert(db)
}

// Upsert performs an upsert for Slethist.
//
// NOTE: PostgreSQL 9.5+ only
func (s *Slethist) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.slethist (` +
		`slhmonth, slhcount, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`slhmonth, slhcount, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.slhmonth, EXCLUDED.slhcount, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, s.Slhmonth, s.Slhcount, s.EquinoxPrn, s.EquinoxLrn, s.EquinoxSec)
	_, err = db.Exec(sqlstr, s.Slhmonth, s.Slhcount, s.EquinoxPrn, s.EquinoxLrn, s.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Delete deletes the Slethist from the database.
func (s *Slethist) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return nil
	}

	// if deleted, bail
	if s._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.slethist WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, s.EquinoxLrn)
	_, err = db.Exec(sqlstr, s.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	s._deleted = true

	return nil
}

// SlethistByEquinoxLrn retrieves a row from 'equinox.slethist' as a Slethist.
//
// Generated from index 'slethist_pkey'.
func SlethistByEquinoxLrn(db XODB, equinoxLrn int64) (*Slethist, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`slhmonth, slhcount, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.slethist ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	s := Slethist{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&s.Slhmonth, &s.Slhcount, &s.EquinoxPrn, &s.EquinoxLrn, &s.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
