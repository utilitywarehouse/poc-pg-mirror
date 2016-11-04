// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Exregcus represents a row from 'equinox.exregcus'.
type Exregcus struct {
	ErcExid      sql.NullString `json:"erc_exid"`      // erc_exid
	ErcCustaccno sql.NullString `json:"erc_custaccno"` // erc_custaccno
	ErcRegdate   pq.NullTime    `json:"erc_regdate"`   // erc_regdate
	ErcLivedate  pq.NullTime    `json:"erc_livedate"`  // erc_livedate
	EquinoxLrn   int64          `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec   sql.NullInt64  `json:"equinox_sec"`   // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Exregcus exists in the database.
func (e *Exregcus) Exists() bool {
	return e._exists
}

// Deleted provides information if the Exregcus has been deleted from the database.
func (e *Exregcus) Deleted() bool {
	return e._deleted
}

// Insert inserts the Exregcus to the database.
func (e *Exregcus) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.exregcus (` +
		`erc_exid, erc_custaccno, erc_regdate, erc_livedate, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, e.ErcExid, e.ErcCustaccno, e.ErcRegdate, e.ErcLivedate, e.EquinoxSec)
	err = db.QueryRow(sqlstr, e.ErcExid, e.ErcCustaccno, e.ErcRegdate, e.ErcLivedate, e.EquinoxSec).Scan(&e.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Update updates the Exregcus in the database.
func (e *Exregcus) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if e._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.exregcus SET (` +
		`erc_exid, erc_custaccno, erc_regdate, erc_livedate, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5` +
		`) WHERE equinox_lrn = $6`

	// run query
	XOLog(sqlstr, e.ErcExid, e.ErcCustaccno, e.ErcRegdate, e.ErcLivedate, e.EquinoxSec, e.EquinoxLrn)
	_, err = db.Exec(sqlstr, e.ErcExid, e.ErcCustaccno, e.ErcRegdate, e.ErcLivedate, e.EquinoxSec, e.EquinoxLrn)
	return err
}

// Save saves the Exregcus to the database.
func (e *Exregcus) Save(db XODB) error {
	if e.Exists() {
		return e.Update(db)
	}

	return e.Insert(db)
}

// Upsert performs an upsert for Exregcus.
//
// NOTE: PostgreSQL 9.5+ only
func (e *Exregcus) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.exregcus (` +
		`erc_exid, erc_custaccno, erc_regdate, erc_livedate, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`erc_exid, erc_custaccno, erc_regdate, erc_livedate, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.erc_exid, EXCLUDED.erc_custaccno, EXCLUDED.erc_regdate, EXCLUDED.erc_livedate, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, e.ErcExid, e.ErcCustaccno, e.ErcRegdate, e.ErcLivedate, e.EquinoxLrn, e.EquinoxSec)
	_, err = db.Exec(sqlstr, e.ErcExid, e.ErcCustaccno, e.ErcRegdate, e.ErcLivedate, e.EquinoxLrn, e.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Delete deletes the Exregcus from the database.
func (e *Exregcus) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return nil
	}

	// if deleted, bail
	if e._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.exregcus WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, e.EquinoxLrn)
	_, err = db.Exec(sqlstr, e.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	e._deleted = true

	return nil
}

// ExregcusByEquinoxLrn retrieves a row from 'equinox.exregcus' as a Exregcus.
//
// Generated from index 'exregcus_pkey'.
func ExregcusByEquinoxLrn(db XODB, equinoxLrn int64) (*Exregcus, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`erc_exid, erc_custaccno, erc_regdate, erc_livedate, equinox_lrn, equinox_sec ` +
		`FROM equinox.exregcus ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Exregcus{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.ErcExid, &e.ErcCustaccno, &e.ErcRegdate, &e.ErcLivedate, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}