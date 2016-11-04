// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Cvcconf represents a row from 'equinox.cvcconf'.
type Cvcconf struct {
	Cvcconfigid      sql.NullInt64   `json:"cvcconfigid"`      // cvcconfigid
	Cvcdescription   sql.NullString  `json:"cvcdescription"`   // cvcdescription
	Cvcbillsfield    sql.NullString  `json:"cvcbillsfield"`    // cvcbillsfield
	Cvcentrycode     sql.NullInt64   `json:"cvcentrycode"`     // cvcentrycode
	Cvcexectotal     sql.NullInt64   `json:"cvcexectotal"`     // cvcexectotal
	Cvcdownlinetotal sql.NullInt64   `json:"cvcdownlinetotal"` // cvcdownlinetotal
	Cvcinvoicetotal  sql.NullInt64   `json:"cvcinvoicetotal"`  // cvcinvoicetotal
	Cvcdownline      sql.NullInt64   `json:"cvcdownline"`      // cvcdownline
	Cvcratingid      sql.NullInt64   `json:"cvcratingid"`      // cvcratingid
	Cvcservicecode   sql.NullString  `json:"cvcservicecode"`   // cvcservicecode
	Cvcrateposition  sql.NullInt64   `json:"cvcrateposition"`  // cvcrateposition
	Cvcservposition  sql.NullInt64   `json:"cvcservposition"`  // cvcservposition
	Cvcsparen2       sql.NullFloat64 `json:"cvcsparen2"`       // cvcsparen2
	Cvcsparec1       sql.NullString  `json:"cvcsparec1"`       // cvcsparec1
	Cvcsparec2       sql.NullString  `json:"cvcsparec2"`       // cvcsparec2
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Cvcconf exists in the database.
func (c *Cvcconf) Exists() bool {
	return c._exists
}

// Deleted provides information if the Cvcconf has been deleted from the database.
func (c *Cvcconf) Deleted() bool {
	return c._deleted
}

// Insert inserts the Cvcconf to the database.
func (c *Cvcconf) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cvcconf (` +
		`cvcconfigid, cvcdescription, cvcbillsfield, cvcentrycode, cvcexectotal, cvcdownlinetotal, cvcinvoicetotal, cvcdownline, cvcratingid, cvcservicecode, cvcrateposition, cvcservposition, cvcsparen2, cvcsparec1, cvcsparec2, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Cvcconfigid, c.Cvcdescription, c.Cvcbillsfield, c.Cvcentrycode, c.Cvcexectotal, c.Cvcdownlinetotal, c.Cvcinvoicetotal, c.Cvcdownline, c.Cvcratingid, c.Cvcservicecode, c.Cvcrateposition, c.Cvcservposition, c.Cvcsparen2, c.Cvcsparec1, c.Cvcsparec2, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Cvcconfigid, c.Cvcdescription, c.Cvcbillsfield, c.Cvcentrycode, c.Cvcexectotal, c.Cvcdownlinetotal, c.Cvcinvoicetotal, c.Cvcdownline, c.Cvcratingid, c.Cvcservicecode, c.Cvcrateposition, c.Cvcservposition, c.Cvcsparen2, c.Cvcsparec1, c.Cvcsparec2, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Cvcconf in the database.
func (c *Cvcconf) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.cvcconf SET (` +
		`cvcconfigid, cvcdescription, cvcbillsfield, cvcentrycode, cvcexectotal, cvcdownlinetotal, cvcinvoicetotal, cvcdownline, cvcratingid, cvcservicecode, cvcrateposition, cvcservposition, cvcsparen2, cvcsparec1, cvcsparec2, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16` +
		`) WHERE equinox_lrn = $17`

	// run query
	XOLog(sqlstr, c.Cvcconfigid, c.Cvcdescription, c.Cvcbillsfield, c.Cvcentrycode, c.Cvcexectotal, c.Cvcdownlinetotal, c.Cvcinvoicetotal, c.Cvcdownline, c.Cvcratingid, c.Cvcservicecode, c.Cvcrateposition, c.Cvcservposition, c.Cvcsparen2, c.Cvcsparec1, c.Cvcsparec2, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Cvcconfigid, c.Cvcdescription, c.Cvcbillsfield, c.Cvcentrycode, c.Cvcexectotal, c.Cvcdownlinetotal, c.Cvcinvoicetotal, c.Cvcdownline, c.Cvcratingid, c.Cvcservicecode, c.Cvcrateposition, c.Cvcservposition, c.Cvcsparen2, c.Cvcsparec1, c.Cvcsparec2, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Cvcconf to the database.
func (c *Cvcconf) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Cvcconf.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Cvcconf) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cvcconf (` +
		`cvcconfigid, cvcdescription, cvcbillsfield, cvcentrycode, cvcexectotal, cvcdownlinetotal, cvcinvoicetotal, cvcdownline, cvcratingid, cvcservicecode, cvcrateposition, cvcservposition, cvcsparen2, cvcsparec1, cvcsparec2, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`cvcconfigid, cvcdescription, cvcbillsfield, cvcentrycode, cvcexectotal, cvcdownlinetotal, cvcinvoicetotal, cvcdownline, cvcratingid, cvcservicecode, cvcrateposition, cvcservposition, cvcsparen2, cvcsparec1, cvcsparec2, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.cvcconfigid, EXCLUDED.cvcdescription, EXCLUDED.cvcbillsfield, EXCLUDED.cvcentrycode, EXCLUDED.cvcexectotal, EXCLUDED.cvcdownlinetotal, EXCLUDED.cvcinvoicetotal, EXCLUDED.cvcdownline, EXCLUDED.cvcratingid, EXCLUDED.cvcservicecode, EXCLUDED.cvcrateposition, EXCLUDED.cvcservposition, EXCLUDED.cvcsparen2, EXCLUDED.cvcsparec1, EXCLUDED.cvcsparec2, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Cvcconfigid, c.Cvcdescription, c.Cvcbillsfield, c.Cvcentrycode, c.Cvcexectotal, c.Cvcdownlinetotal, c.Cvcinvoicetotal, c.Cvcdownline, c.Cvcratingid, c.Cvcservicecode, c.Cvcrateposition, c.Cvcservposition, c.Cvcsparen2, c.Cvcsparec1, c.Cvcsparec2, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Cvcconfigid, c.Cvcdescription, c.Cvcbillsfield, c.Cvcentrycode, c.Cvcexectotal, c.Cvcdownlinetotal, c.Cvcinvoicetotal, c.Cvcdownline, c.Cvcratingid, c.Cvcservicecode, c.Cvcrateposition, c.Cvcservposition, c.Cvcsparen2, c.Cvcsparec1, c.Cvcsparec2, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Cvcconf from the database.
func (c *Cvcconf) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.cvcconf WHERE equinox_lrn = $1`

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

// CvcconfByEquinoxLrn retrieves a row from 'equinox.cvcconf' as a Cvcconf.
//
// Generated from index 'cvcconf_pkey'.
func CvcconfByEquinoxLrn(db XODB, equinoxLrn int64) (*Cvcconf, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cvcconfigid, cvcdescription, cvcbillsfield, cvcentrycode, cvcexectotal, cvcdownlinetotal, cvcinvoicetotal, cvcdownline, cvcratingid, cvcservicecode, cvcrateposition, cvcservposition, cvcsparen2, cvcsparec1, cvcsparec2, equinox_lrn, equinox_sec ` +
		`FROM equinox.cvcconf ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cvcconf{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cvcconfigid, &c.Cvcdescription, &c.Cvcbillsfield, &c.Cvcentrycode, &c.Cvcexectotal, &c.Cvcdownlinetotal, &c.Cvcinvoicetotal, &c.Cvcdownline, &c.Cvcratingid, &c.Cvcservicecode, &c.Cvcrateposition, &c.Cvcservposition, &c.Cvcsparen2, &c.Cvcsparec1, &c.Cvcsparec2, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
