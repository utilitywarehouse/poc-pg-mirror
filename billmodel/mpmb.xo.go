// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Mpmb represents a row from 'equinox.mpmb'.
type Mpmb struct {
	Mpbmpancore    sql.NullString `json:"mpbmpancore"`    // mpbmpancore
	Mpbsupplierid  sql.NullString `json:"mpbsupplierid"`  // mpbsupplierid
	Mpbsettledate  pq.NullTime    `json:"mpbsettledate"`  // mpbsettledate
	Mpbmetertype   sql.NullString `json:"mpbmetertype"`   // mpbmetertype
	Mpbinstalldate pq.NullTime    `json:"mpbinstalldate"` // mpbinstalldate
	Mpbmeterecoes  sql.NullString `json:"mpbmeterecoes"`  // mpbmeterecoes
	EquinoxPrn     sql.NullInt64  `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64          `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64  `json:"equinox_sec"`    // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Mpmb exists in the database.
func (m *Mpmb) Exists() bool {
	return m._exists
}

// Deleted provides information if the Mpmb has been deleted from the database.
func (m *Mpmb) Deleted() bool {
	return m._deleted
}

// Insert inserts the Mpmb to the database.
func (m *Mpmb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if m._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.mpmb (` +
		`mpbmpancore, mpbsupplierid, mpbsettledate, mpbmetertype, mpbinstalldate, mpbmeterecoes, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, m.Mpbmpancore, m.Mpbsupplierid, m.Mpbsettledate, m.Mpbmetertype, m.Mpbinstalldate, m.Mpbmeterecoes, m.EquinoxPrn, m.EquinoxSec)
	err = db.QueryRow(sqlstr, m.Mpbmpancore, m.Mpbsupplierid, m.Mpbsettledate, m.Mpbmetertype, m.Mpbinstalldate, m.Mpbmeterecoes, m.EquinoxPrn, m.EquinoxSec).Scan(&m.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	m._exists = true

	return nil
}

// Update updates the Mpmb in the database.
func (m *Mpmb) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.mpmb SET (` +
		`mpbmpancore, mpbsupplierid, mpbsettledate, mpbmetertype, mpbinstalldate, mpbmeterecoes, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) WHERE equinox_lrn = $9`

	// run query
	XOLog(sqlstr, m.Mpbmpancore, m.Mpbsupplierid, m.Mpbsettledate, m.Mpbmetertype, m.Mpbinstalldate, m.Mpbmeterecoes, m.EquinoxPrn, m.EquinoxSec, m.EquinoxLrn)
	_, err = db.Exec(sqlstr, m.Mpbmpancore, m.Mpbsupplierid, m.Mpbsettledate, m.Mpbmetertype, m.Mpbinstalldate, m.Mpbmeterecoes, m.EquinoxPrn, m.EquinoxSec, m.EquinoxLrn)
	return err
}

// Save saves the Mpmb to the database.
func (m *Mpmb) Save(db XODB) error {
	if m.Exists() {
		return m.Update(db)
	}

	return m.Insert(db)
}

// Upsert performs an upsert for Mpmb.
//
// NOTE: PostgreSQL 9.5+ only
func (m *Mpmb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if m._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.mpmb (` +
		`mpbmpancore, mpbsupplierid, mpbsettledate, mpbmetertype, mpbinstalldate, mpbmeterecoes, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`mpbmpancore, mpbsupplierid, mpbsettledate, mpbmetertype, mpbinstalldate, mpbmeterecoes, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.mpbmpancore, EXCLUDED.mpbsupplierid, EXCLUDED.mpbsettledate, EXCLUDED.mpbmetertype, EXCLUDED.mpbinstalldate, EXCLUDED.mpbmeterecoes, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, m.Mpbmpancore, m.Mpbsupplierid, m.Mpbsettledate, m.Mpbmetertype, m.Mpbinstalldate, m.Mpbmeterecoes, m.EquinoxPrn, m.EquinoxLrn, m.EquinoxSec)
	_, err = db.Exec(sqlstr, m.Mpbmpancore, m.Mpbsupplierid, m.Mpbsettledate, m.Mpbmetertype, m.Mpbinstalldate, m.Mpbmeterecoes, m.EquinoxPrn, m.EquinoxLrn, m.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	m._exists = true

	return nil
}

// Delete deletes the Mpmb from the database.
func (m *Mpmb) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.mpmb WHERE equinox_lrn = $1`

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

// MpmbByEquinoxLrn retrieves a row from 'equinox.mpmb' as a Mpmb.
//
// Generated from index 'mpmb_pkey'.
func MpmbByEquinoxLrn(db XODB, equinoxLrn int64) (*Mpmb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`mpbmpancore, mpbsupplierid, mpbsettledate, mpbmetertype, mpbinstalldate, mpbmeterecoes, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.mpmb ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	m := Mpmb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&m.Mpbmpancore, &m.Mpbsupplierid, &m.Mpbsettledate, &m.Mpbmetertype, &m.Mpbinstalldate, &m.Mpbmeterecoes, &m.EquinoxPrn, &m.EquinoxLrn, &m.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
