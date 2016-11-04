// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Minicrit represents a row from 'equinox.minicrit'.
type Minicrit struct {
	Minicrcriteriano sql.NullInt64 `json:"minicrcriteriano"` // minicrcriteriano
	Minicrcustomers  sql.NullInt64 `json:"minicrcustomers"`  // minicrcustomers
	Minicrids        sql.NullInt64 `json:"minicrids"`        // minicrids
	EquinoxPrn       sql.NullInt64 `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64         `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64 `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Minicrit exists in the database.
func (m *Minicrit) Exists() bool {
	return m._exists
}

// Deleted provides information if the Minicrit has been deleted from the database.
func (m *Minicrit) Deleted() bool {
	return m._deleted
}

// Insert inserts the Minicrit to the database.
func (m *Minicrit) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if m._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.minicrit (` +
		`minicrcriteriano, minicrcustomers, minicrids, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, m.Minicrcriteriano, m.Minicrcustomers, m.Minicrids, m.EquinoxPrn, m.EquinoxSec)
	err = db.QueryRow(sqlstr, m.Minicrcriteriano, m.Minicrcustomers, m.Minicrids, m.EquinoxPrn, m.EquinoxSec).Scan(&m.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	m._exists = true

	return nil
}

// Update updates the Minicrit in the database.
func (m *Minicrit) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !m._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if m._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.minicrit SET (` +
		`minicrcriteriano, minicrcustomers, minicrids, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5` +
		`) WHERE equinox_lrn = $6`

	// run query
	XOLog(sqlstr, m.Minicrcriteriano, m.Minicrcustomers, m.Minicrids, m.EquinoxPrn, m.EquinoxSec, m.EquinoxLrn)
	_, err = db.Exec(sqlstr, m.Minicrcriteriano, m.Minicrcustomers, m.Minicrids, m.EquinoxPrn, m.EquinoxSec, m.EquinoxLrn)
	return err
}

// Save saves the Minicrit to the database.
func (m *Minicrit) Save(db XODB) error {
	if m.Exists() {
		return m.Update(db)
	}

	return m.Insert(db)
}

// Upsert performs an upsert for Minicrit.
//
// NOTE: PostgreSQL 9.5+ only
func (m *Minicrit) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if m._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.minicrit (` +
		`minicrcriteriano, minicrcustomers, minicrids, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`minicrcriteriano, minicrcustomers, minicrids, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.minicrcriteriano, EXCLUDED.minicrcustomers, EXCLUDED.minicrids, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, m.Minicrcriteriano, m.Minicrcustomers, m.Minicrids, m.EquinoxPrn, m.EquinoxLrn, m.EquinoxSec)
	_, err = db.Exec(sqlstr, m.Minicrcriteriano, m.Minicrcustomers, m.Minicrids, m.EquinoxPrn, m.EquinoxLrn, m.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	m._exists = true

	return nil
}

// Delete deletes the Minicrit from the database.
func (m *Minicrit) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !m._exists {
		return nil
	}

	// if deleted, bail
	if m._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.minicrit WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, m.EquinoxLrn)
	_, err = db.Exec(sqlstr, m.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	m._deleted = true

	return nil
}

// MinicritByEquinoxLrn retrieves a row from 'equinox.minicrit' as a Minicrit.
//
// Generated from index 'minicrit_pkey'.
func MinicritByEquinoxLrn(db XODB, equinoxLrn int64) (*Minicrit, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`minicrcriteriano, minicrcustomers, minicrids, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.minicrit ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	m := Minicrit{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&m.Minicrcriteriano, &m.Minicrcustomers, &m.Minicrids, &m.EquinoxPrn, &m.EquinoxLrn, &m.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &m, nil
}