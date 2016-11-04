// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Sysletl represents a row from 'equinox.sysletl'.
type Sysletl struct {
	Sllettercode     sql.NullString `json:"sllettercode"`     // sllettercode
	Slcountquarterly sql.NullInt64  `json:"slcountquarterly"` // slcountquarterly
	Slcountannual    sql.NullInt64  `json:"slcountannual"`    // slcountannual
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Sysletl exists in the database.
func (s *Sysletl) Exists() bool {
	return s._exists
}

// Deleted provides information if the Sysletl has been deleted from the database.
func (s *Sysletl) Deleted() bool {
	return s._deleted
}

// Insert inserts the Sysletl to the database.
func (s *Sysletl) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.sysletl (` +
		`sllettercode, slcountquarterly, slcountannual, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, s.Sllettercode, s.Slcountquarterly, s.Slcountannual, s.EquinoxSec)
	err = db.QueryRow(sqlstr, s.Sllettercode, s.Slcountquarterly, s.Slcountannual, s.EquinoxSec).Scan(&s.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Update updates the Sysletl in the database.
func (s *Sysletl) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.sysletl SET (` +
		`sllettercode, slcountquarterly, slcountannual, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE equinox_lrn = $5`

	// run query
	XOLog(sqlstr, s.Sllettercode, s.Slcountquarterly, s.Slcountannual, s.EquinoxSec, s.EquinoxLrn)
	_, err = db.Exec(sqlstr, s.Sllettercode, s.Slcountquarterly, s.Slcountannual, s.EquinoxSec, s.EquinoxLrn)
	return err
}

// Save saves the Sysletl to the database.
func (s *Sysletl) Save(db XODB) error {
	if s.Exists() {
		return s.Update(db)
	}

	return s.Insert(db)
}

// Upsert performs an upsert for Sysletl.
//
// NOTE: PostgreSQL 9.5+ only
func (s *Sysletl) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.sysletl (` +
		`sllettercode, slcountquarterly, slcountannual, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`sllettercode, slcountquarterly, slcountannual, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.sllettercode, EXCLUDED.slcountquarterly, EXCLUDED.slcountannual, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, s.Sllettercode, s.Slcountquarterly, s.Slcountannual, s.EquinoxLrn, s.EquinoxSec)
	_, err = db.Exec(sqlstr, s.Sllettercode, s.Slcountquarterly, s.Slcountannual, s.EquinoxLrn, s.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	s._exists = true

	return nil
}

// Delete deletes the Sysletl from the database.
func (s *Sysletl) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.sysletl WHERE equinox_lrn = $1`

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

// SysletlByEquinoxLrn retrieves a row from 'equinox.sysletl' as a Sysletl.
//
// Generated from index 'sysletl_pkey'.
func SysletlByEquinoxLrn(db XODB, equinoxLrn int64) (*Sysletl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`sllettercode, slcountquarterly, slcountannual, equinox_lrn, equinox_sec ` +
		`FROM equinox.sysletl ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	s := Sysletl{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&s.Sllettercode, &s.Slcountquarterly, &s.Slcountannual, &s.EquinoxLrn, &s.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &s, nil
}