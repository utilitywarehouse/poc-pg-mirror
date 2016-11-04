// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Ps72rec represents a row from 'equinox.ps72recs'.
type Ps72rec struct {
	Ps72rejreason sql.NullString `json:"ps72rejreason"` // ps72rejreason
	EquinoxPrn    sql.NullInt64  `json:"equinox_prn"`   // equinox_prn
	EquinoxLrn    int64          `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec    sql.NullInt64  `json:"equinox_sec"`   // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Ps72rec exists in the database.
func (p *Ps72rec) Exists() bool {
	return p._exists
}

// Deleted provides information if the Ps72rec has been deleted from the database.
func (p *Ps72rec) Deleted() bool {
	return p._deleted
}

// Insert inserts the Ps72rec to the database.
func (p *Ps72rec) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.ps72recs (` +
		`ps72rejreason, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, p.Ps72rejreason, p.EquinoxPrn, p.EquinoxSec)
	err = db.QueryRow(sqlstr, p.Ps72rejreason, p.EquinoxPrn, p.EquinoxSec).Scan(&p.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Update updates the Ps72rec in the database.
func (p *Ps72rec) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if p._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.ps72recs SET (` +
		`ps72rejreason, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE equinox_lrn = $4`

	// run query
	XOLog(sqlstr, p.Ps72rejreason, p.EquinoxPrn, p.EquinoxSec, p.EquinoxLrn)
	_, err = db.Exec(sqlstr, p.Ps72rejreason, p.EquinoxPrn, p.EquinoxSec, p.EquinoxLrn)
	return err
}

// Save saves the Ps72rec to the database.
func (p *Ps72rec) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Upsert performs an upsert for Ps72rec.
//
// NOTE: PostgreSQL 9.5+ only
func (p *Ps72rec) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.ps72recs (` +
		`ps72rejreason, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`ps72rejreason, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.ps72rejreason, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, p.Ps72rejreason, p.EquinoxPrn, p.EquinoxLrn, p.EquinoxSec)
	_, err = db.Exec(sqlstr, p.Ps72rejreason, p.EquinoxPrn, p.EquinoxLrn, p.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Delete deletes the Ps72rec from the database.
func (p *Ps72rec) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return nil
	}

	// if deleted, bail
	if p._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.ps72recs WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, p.EquinoxLrn)
	_, err = db.Exec(sqlstr, p.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	p._deleted = true

	return nil
}

// Ps72recByEquinoxLrn retrieves a row from 'equinox.ps72recs' as a Ps72rec.
//
// Generated from index 'ps72recs_pkey'.
func Ps72recByEquinoxLrn(db XODB, equinoxLrn int64) (*Ps72rec, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ps72rejreason, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.ps72recs ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Ps72rec{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Ps72rejreason, &p.EquinoxPrn, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
