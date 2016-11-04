// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Allowed represents a row from 'equinox.allowed'.
type Allowed struct {
	Allowedmeter  sql.NullString `json:"allowedmeter"`  // allowedmeter
	Allowedexecid sql.NullString `json:"allowedexecid"` // allowedexecid
	EquinoxLrn    int64          `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec    sql.NullInt64  `json:"equinox_sec"`   // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Allowed exists in the database.
func (a *Allowed) Exists() bool {
	return a._exists
}

// Deleted provides information if the Allowed has been deleted from the database.
func (a *Allowed) Deleted() bool {
	return a._deleted
}

// Insert inserts the Allowed to the database.
func (a *Allowed) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.allowed (` +
		`allowedmeter, allowedexecid, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, a.Allowedmeter, a.Allowedexecid, a.EquinoxSec)
	err = db.QueryRow(sqlstr, a.Allowedmeter, a.Allowedexecid, a.EquinoxSec).Scan(&a.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Update updates the Allowed in the database.
func (a *Allowed) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if a._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.allowed SET (` +
		`allowedmeter, allowedexecid, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE equinox_lrn = $4`

	// run query
	XOLog(sqlstr, a.Allowedmeter, a.Allowedexecid, a.EquinoxSec, a.EquinoxLrn)
	_, err = db.Exec(sqlstr, a.Allowedmeter, a.Allowedexecid, a.EquinoxSec, a.EquinoxLrn)
	return err
}

// Save saves the Allowed to the database.
func (a *Allowed) Save(db XODB) error {
	if a.Exists() {
		return a.Update(db)
	}

	return a.Insert(db)
}

// Upsert performs an upsert for Allowed.
//
// NOTE: PostgreSQL 9.5+ only
func (a *Allowed) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.allowed (` +
		`allowedmeter, allowedexecid, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`allowedmeter, allowedexecid, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.allowedmeter, EXCLUDED.allowedexecid, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, a.Allowedmeter, a.Allowedexecid, a.EquinoxLrn, a.EquinoxSec)
	_, err = db.Exec(sqlstr, a.Allowedmeter, a.Allowedexecid, a.EquinoxLrn, a.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Delete deletes the Allowed from the database.
func (a *Allowed) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return nil
	}

	// if deleted, bail
	if a._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.allowed WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, a.EquinoxLrn)
	_, err = db.Exec(sqlstr, a.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	a._deleted = true

	return nil
}

// AllowedByEquinoxLrn retrieves a row from 'equinox.allowed' as a Allowed.
//
// Generated from index 'allowed_pkey'.
func AllowedByEquinoxLrn(db XODB, equinoxLrn int64) (*Allowed, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`allowedmeter, allowedexecid, equinox_lrn, equinox_sec ` +
		`FROM equinox.allowed ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	a := Allowed{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&a.Allowedmeter, &a.Allowedexecid, &a.EquinoxLrn, &a.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
