// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Oshist represents a row from 'equinox.oshist'.
type Oshist struct {
	Oshdateout   pq.NullTime    `json:"oshdateout"`   // oshdateout
	Oshtimeout   pq.NullTime    `json:"oshtimeout"`   // oshtimeout
	Oshaccountno sql.NullString `json:"oshaccountno"` // oshaccountno
	Oshlocation  sql.NullString `json:"oshlocation"`  // oshlocation
	EquinoxPrn   sql.NullInt64  `json:"equinox_prn"`  // equinox_prn
	EquinoxLrn   int64          `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64  `json:"equinox_sec"`  // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Oshist exists in the database.
func (o *Oshist) Exists() bool {
	return o._exists
}

// Deleted provides information if the Oshist has been deleted from the database.
func (o *Oshist) Deleted() bool {
	return o._deleted
}

// Insert inserts the Oshist to the database.
func (o *Oshist) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if o._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.oshist (` +
		`oshdateout, oshtimeout, oshaccountno, oshlocation, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, o.Oshdateout, o.Oshtimeout, o.Oshaccountno, o.Oshlocation, o.EquinoxPrn, o.EquinoxSec)
	err = db.QueryRow(sqlstr, o.Oshdateout, o.Oshtimeout, o.Oshaccountno, o.Oshlocation, o.EquinoxPrn, o.EquinoxSec).Scan(&o.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	o._exists = true

	return nil
}

// Update updates the Oshist in the database.
func (o *Oshist) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !o._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if o._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.oshist SET (` +
		`oshdateout, oshtimeout, oshaccountno, oshlocation, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6` +
		`) WHERE equinox_lrn = $7`

	// run query
	XOLog(sqlstr, o.Oshdateout, o.Oshtimeout, o.Oshaccountno, o.Oshlocation, o.EquinoxPrn, o.EquinoxSec, o.EquinoxLrn)
	_, err = db.Exec(sqlstr, o.Oshdateout, o.Oshtimeout, o.Oshaccountno, o.Oshlocation, o.EquinoxPrn, o.EquinoxSec, o.EquinoxLrn)
	return err
}

// Save saves the Oshist to the database.
func (o *Oshist) Save(db XODB) error {
	if o.Exists() {
		return o.Update(db)
	}

	return o.Insert(db)
}

// Upsert performs an upsert for Oshist.
//
// NOTE: PostgreSQL 9.5+ only
func (o *Oshist) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if o._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.oshist (` +
		`oshdateout, oshtimeout, oshaccountno, oshlocation, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`oshdateout, oshtimeout, oshaccountno, oshlocation, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.oshdateout, EXCLUDED.oshtimeout, EXCLUDED.oshaccountno, EXCLUDED.oshlocation, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, o.Oshdateout, o.Oshtimeout, o.Oshaccountno, o.Oshlocation, o.EquinoxPrn, o.EquinoxLrn, o.EquinoxSec)
	_, err = db.Exec(sqlstr, o.Oshdateout, o.Oshtimeout, o.Oshaccountno, o.Oshlocation, o.EquinoxPrn, o.EquinoxLrn, o.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	o._exists = true

	return nil
}

// Delete deletes the Oshist from the database.
func (o *Oshist) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !o._exists {
		return nil
	}

	// if deleted, bail
	if o._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.oshist WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, o.EquinoxLrn)
	_, err = db.Exec(sqlstr, o.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	o._deleted = true

	return nil
}

// OshistByEquinoxLrn retrieves a row from 'equinox.oshist' as a Oshist.
//
// Generated from index 'oshist_pkey'.
func OshistByEquinoxLrn(db XODB, equinoxLrn int64) (*Oshist, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`oshdateout, oshtimeout, oshaccountno, oshlocation, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.oshist ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	o := Oshist{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&o.Oshdateout, &o.Oshtimeout, &o.Oshaccountno, &o.Oshlocation, &o.EquinoxPrn, &o.EquinoxLrn, &o.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &o, nil
}