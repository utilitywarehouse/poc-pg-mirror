// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Excustac represents a row from 'equinox.excustac'.
type Excustac struct {
	Excustacno      sql.NullString `json:"excustacno"`      // excustacno
	Excustacdate    pq.NullTime    `json:"excustacdate"`    // excustacdate
	Excustacquality sql.NullInt64  `json:"excustacquality"` // excustacquality
	EquinoxPrn      sql.NullInt64  `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn      int64          `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`     // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Excustac exists in the database.
func (e *Excustac) Exists() bool {
	return e._exists
}

// Deleted provides information if the Excustac has been deleted from the database.
func (e *Excustac) Deleted() bool {
	return e._deleted
}

// Insert inserts the Excustac to the database.
func (e *Excustac) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.excustac (` +
		`excustacno, excustacdate, excustacquality, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, e.Excustacno, e.Excustacdate, e.Excustacquality, e.EquinoxPrn, e.EquinoxSec)
	err = db.QueryRow(sqlstr, e.Excustacno, e.Excustacdate, e.Excustacquality, e.EquinoxPrn, e.EquinoxSec).Scan(&e.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Update updates the Excustac in the database.
func (e *Excustac) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.excustac SET (` +
		`excustacno, excustacdate, excustacquality, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5` +
		`) WHERE equinox_lrn = $6`

	// run query
	XOLog(sqlstr, e.Excustacno, e.Excustacdate, e.Excustacquality, e.EquinoxPrn, e.EquinoxSec, e.EquinoxLrn)
	_, err = db.Exec(sqlstr, e.Excustacno, e.Excustacdate, e.Excustacquality, e.EquinoxPrn, e.EquinoxSec, e.EquinoxLrn)
	return err
}

// Save saves the Excustac to the database.
func (e *Excustac) Save(db XODB) error {
	if e.Exists() {
		return e.Update(db)
	}

	return e.Insert(db)
}

// Upsert performs an upsert for Excustac.
//
// NOTE: PostgreSQL 9.5+ only
func (e *Excustac) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.excustac (` +
		`excustacno, excustacdate, excustacquality, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`excustacno, excustacdate, excustacquality, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.excustacno, EXCLUDED.excustacdate, EXCLUDED.excustacquality, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, e.Excustacno, e.Excustacdate, e.Excustacquality, e.EquinoxPrn, e.EquinoxLrn, e.EquinoxSec)
	_, err = db.Exec(sqlstr, e.Excustacno, e.Excustacdate, e.Excustacquality, e.EquinoxPrn, e.EquinoxLrn, e.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Delete deletes the Excustac from the database.
func (e *Excustac) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.excustac WHERE equinox_lrn = $1`

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

// ExcustacByEquinoxLrn retrieves a row from 'equinox.excustac' as a Excustac.
//
// Generated from index 'excustac_pkey'.
func ExcustacByEquinoxLrn(db XODB, equinoxLrn int64) (*Excustac, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`excustacno, excustacdate, excustacquality, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.excustac ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Excustac{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Excustacno, &e.Excustacdate, &e.Excustacquality, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}