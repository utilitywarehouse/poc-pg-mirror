// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Mpma represents a row from 'equinox.mpma'.
type Mpma struct {
	Mpasupplierid  sql.NullString `json:"mpasupplierid"`  // mpasupplierid
	Mpaeffsettdate pq.NullTime    `json:"mpaeffsettdate"` // mpaeffsettdate
	Mpameterid     sql.NullString `json:"mpameterid"`     // mpameterid
	Mpametertype   sql.NullString `json:"mpametertype"`   // mpametertype
	Mpainstalldate pq.NullTime    `json:"mpainstalldate"` // mpainstalldate
	Mpampanecoes   sql.NullString `json:"mpampanecoes"`   // mpampanecoes
	EquinoxPrn     sql.NullInt64  `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64          `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64  `json:"equinox_sec"`    // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Mpma exists in the database.
func (m *Mpma) Exists() bool {
	return m._exists
}

// Deleted provides information if the Mpma has been deleted from the database.
func (m *Mpma) Deleted() bool {
	return m._deleted
}

// Insert inserts the Mpma to the database.
func (m *Mpma) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if m._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.mpma (` +
		`mpasupplierid, mpaeffsettdate, mpameterid, mpametertype, mpainstalldate, mpampanecoes, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, m.Mpasupplierid, m.Mpaeffsettdate, m.Mpameterid, m.Mpametertype, m.Mpainstalldate, m.Mpampanecoes, m.EquinoxPrn, m.EquinoxSec)
	err = db.QueryRow(sqlstr, m.Mpasupplierid, m.Mpaeffsettdate, m.Mpameterid, m.Mpametertype, m.Mpainstalldate, m.Mpampanecoes, m.EquinoxPrn, m.EquinoxSec).Scan(&m.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	m._exists = true

	return nil
}

// Update updates the Mpma in the database.
func (m *Mpma) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.mpma SET (` +
		`mpasupplierid, mpaeffsettdate, mpameterid, mpametertype, mpainstalldate, mpampanecoes, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) WHERE equinox_lrn = $9`

	// run query
	XOLog(sqlstr, m.Mpasupplierid, m.Mpaeffsettdate, m.Mpameterid, m.Mpametertype, m.Mpainstalldate, m.Mpampanecoes, m.EquinoxPrn, m.EquinoxSec, m.EquinoxLrn)
	_, err = db.Exec(sqlstr, m.Mpasupplierid, m.Mpaeffsettdate, m.Mpameterid, m.Mpametertype, m.Mpainstalldate, m.Mpampanecoes, m.EquinoxPrn, m.EquinoxSec, m.EquinoxLrn)
	return err
}

// Save saves the Mpma to the database.
func (m *Mpma) Save(db XODB) error {
	if m.Exists() {
		return m.Update(db)
	}

	return m.Insert(db)
}

// Upsert performs an upsert for Mpma.
//
// NOTE: PostgreSQL 9.5+ only
func (m *Mpma) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if m._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.mpma (` +
		`mpasupplierid, mpaeffsettdate, mpameterid, mpametertype, mpainstalldate, mpampanecoes, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`mpasupplierid, mpaeffsettdate, mpameterid, mpametertype, mpainstalldate, mpampanecoes, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.mpasupplierid, EXCLUDED.mpaeffsettdate, EXCLUDED.mpameterid, EXCLUDED.mpametertype, EXCLUDED.mpainstalldate, EXCLUDED.mpampanecoes, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, m.Mpasupplierid, m.Mpaeffsettdate, m.Mpameterid, m.Mpametertype, m.Mpainstalldate, m.Mpampanecoes, m.EquinoxPrn, m.EquinoxLrn, m.EquinoxSec)
	_, err = db.Exec(sqlstr, m.Mpasupplierid, m.Mpaeffsettdate, m.Mpameterid, m.Mpametertype, m.Mpainstalldate, m.Mpampanecoes, m.EquinoxPrn, m.EquinoxLrn, m.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	m._exists = true

	return nil
}

// Delete deletes the Mpma from the database.
func (m *Mpma) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.mpma WHERE equinox_lrn = $1`

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

// MpmaByEquinoxLrn retrieves a row from 'equinox.mpma' as a Mpma.
//
// Generated from index 'mpma_pkey'.
func MpmaByEquinoxLrn(db XODB, equinoxLrn int64) (*Mpma, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`mpasupplierid, mpaeffsettdate, mpameterid, mpametertype, mpainstalldate, mpampanecoes, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.mpma ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	m := Mpma{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&m.Mpasupplierid, &m.Mpaeffsettdate, &m.Mpameterid, &m.Mpametertype, &m.Mpainstalldate, &m.Mpampanecoes, &m.EquinoxPrn, &m.EquinoxLrn, &m.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &m, nil
}