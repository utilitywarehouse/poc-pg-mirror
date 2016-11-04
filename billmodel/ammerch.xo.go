// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Ammerch represents a row from 'equinox.ammerch'.
type Ammerch struct {
	Ammercid       sql.NullString `json:"ammercid"`       // ammercid
	Ammercname     sql.NullString `json:"ammercname"`     // ammercname
	Ammercdatefrom pq.NullTime    `json:"ammercdatefrom"` // ammercdatefrom
	Ammercdateto   pq.NullTime    `json:"ammercdateto"`   // ammercdateto
	EquinoxPrn     sql.NullInt64  `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64          `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64  `json:"equinox_sec"`    // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Ammerch exists in the database.
func (a *Ammerch) Exists() bool {
	return a._exists
}

// Deleted provides information if the Ammerch has been deleted from the database.
func (a *Ammerch) Deleted() bool {
	return a._deleted
}

// Insert inserts the Ammerch to the database.
func (a *Ammerch) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.ammerch (` +
		`ammercid, ammercname, ammercdatefrom, ammercdateto, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, a.Ammercid, a.Ammercname, a.Ammercdatefrom, a.Ammercdateto, a.EquinoxPrn, a.EquinoxSec)
	err = db.QueryRow(sqlstr, a.Ammercid, a.Ammercname, a.Ammercdatefrom, a.Ammercdateto, a.EquinoxPrn, a.EquinoxSec).Scan(&a.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Update updates the Ammerch in the database.
func (a *Ammerch) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.ammerch SET (` +
		`ammercid, ammercname, ammercdatefrom, ammercdateto, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6` +
		`) WHERE equinox_lrn = $7`

	// run query
	XOLog(sqlstr, a.Ammercid, a.Ammercname, a.Ammercdatefrom, a.Ammercdateto, a.EquinoxPrn, a.EquinoxSec, a.EquinoxLrn)
	_, err = db.Exec(sqlstr, a.Ammercid, a.Ammercname, a.Ammercdatefrom, a.Ammercdateto, a.EquinoxPrn, a.EquinoxSec, a.EquinoxLrn)
	return err
}

// Save saves the Ammerch to the database.
func (a *Ammerch) Save(db XODB) error {
	if a.Exists() {
		return a.Update(db)
	}

	return a.Insert(db)
}

// Upsert performs an upsert for Ammerch.
//
// NOTE: PostgreSQL 9.5+ only
func (a *Ammerch) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.ammerch (` +
		`ammercid, ammercname, ammercdatefrom, ammercdateto, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`ammercid, ammercname, ammercdatefrom, ammercdateto, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.ammercid, EXCLUDED.ammercname, EXCLUDED.ammercdatefrom, EXCLUDED.ammercdateto, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, a.Ammercid, a.Ammercname, a.Ammercdatefrom, a.Ammercdateto, a.EquinoxPrn, a.EquinoxLrn, a.EquinoxSec)
	_, err = db.Exec(sqlstr, a.Ammercid, a.Ammercname, a.Ammercdatefrom, a.Ammercdateto, a.EquinoxPrn, a.EquinoxLrn, a.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Delete deletes the Ammerch from the database.
func (a *Ammerch) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.ammerch WHERE equinox_lrn = $1`

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

// AmmerchByEquinoxLrn retrieves a row from 'equinox.ammerch' as a Ammerch.
//
// Generated from index 'ammerch_pkey'.
func AmmerchByEquinoxLrn(db XODB, equinoxLrn int64) (*Ammerch, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ammercid, ammercname, ammercdatefrom, ammercdateto, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.ammerch ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	a := Ammerch{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&a.Ammercid, &a.Ammercname, &a.Ammercdatefrom, &a.Ammercdateto, &a.EquinoxPrn, &a.EquinoxLrn, &a.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
