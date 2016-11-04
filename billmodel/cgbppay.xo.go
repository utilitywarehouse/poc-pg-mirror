// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Cgbppay represents a row from 'equinox.cgbppay'.
type Cgbppay struct {
	Cgbpcrnumber sql.NullInt64   `json:"cgbpcrnumber"` // cgbpcrnumber
	Cgbpamount   sql.NullFloat64 `json:"cgbpamount"`   // cgbpamount
	EquinoxPrn   sql.NullInt64   `json:"equinox_prn"`  // equinox_prn
	EquinoxLrn   int64           `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64   `json:"equinox_sec"`  // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Cgbppay exists in the database.
func (c *Cgbppay) Exists() bool {
	return c._exists
}

// Deleted provides information if the Cgbppay has been deleted from the database.
func (c *Cgbppay) Deleted() bool {
	return c._deleted
}

// Insert inserts the Cgbppay to the database.
func (c *Cgbppay) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cgbppay (` +
		`cgbpcrnumber, cgbpamount, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Cgbpcrnumber, c.Cgbpamount, c.EquinoxPrn, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Cgbpcrnumber, c.Cgbpamount, c.EquinoxPrn, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Cgbppay in the database.
func (c *Cgbppay) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if c._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.cgbppay SET (` +
		`cgbpcrnumber, cgbpamount, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE equinox_lrn = $5`

	// run query
	XOLog(sqlstr, c.Cgbpcrnumber, c.Cgbpamount, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Cgbpcrnumber, c.Cgbpamount, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Cgbppay to the database.
func (c *Cgbppay) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Cgbppay.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Cgbppay) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cgbppay (` +
		`cgbpcrnumber, cgbpamount, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`cgbpcrnumber, cgbpamount, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.cgbpcrnumber, EXCLUDED.cgbpamount, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Cgbpcrnumber, c.Cgbpamount, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Cgbpcrnumber, c.Cgbpamount, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Cgbppay from the database.
func (c *Cgbppay) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return nil
	}

	// if deleted, bail
	if c._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.cgbppay WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	c._deleted = true

	return nil
}

// CgbppayByEquinoxLrn retrieves a row from 'equinox.cgbppay' as a Cgbppay.
//
// Generated from index 'cgbppay_pkey'.
func CgbppayByEquinoxLrn(db XODB, equinoxLrn int64) (*Cgbppay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cgbpcrnumber, cgbpamount, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cgbppay ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cgbppay{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cgbpcrnumber, &c.Cgbpamount, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
