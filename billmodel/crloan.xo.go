// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Crloan represents a row from 'equinox.crloans'.
type Crloan struct {
	Crloanid     sql.NullInt64   `json:"crloanid"`     // crloanid
	Crlexid      sql.NullString  `json:"crlexid"`      // crlexid
	Crlcomcodeid sql.NullInt64   `json:"crlcomcodeid"` // crlcomcodeid
	Crltype      sql.NullString  `json:"crltype"`      // crltype
	Crlamount    sql.NullFloat64 `json:"crlamount"`    // crlamount
	Crlamount2   sql.NullFloat64 `json:"crlamount2"`   // crlamount2
	Crlfrom      sql.NullString  `json:"crlfrom"`      // crlfrom
	Crlto        sql.NullString  `json:"crlto"`        // crlto
	Crlcommitted pq.NullTime     `json:"crlcommitted"` // crlcommitted
	Crlsparenum1 sql.NullFloat64 `json:"crlsparenum1"` // crlsparenum1
	Crlsparenum2 sql.NullFloat64 `json:"crlsparenum2"` // crlsparenum2
	Crlsparec1   sql.NullString  `json:"crlsparec1"`   // crlsparec1
	Crlsparec2   sql.NullString  `json:"crlsparec2"`   // crlsparec2
	EquinoxPrn   sql.NullInt64   `json:"equinox_prn"`  // equinox_prn
	EquinoxLrn   int64           `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64   `json:"equinox_sec"`  // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Crloan exists in the database.
func (c *Crloan) Exists() bool {
	return c._exists
}

// Deleted provides information if the Crloan has been deleted from the database.
func (c *Crloan) Deleted() bool {
	return c._deleted
}

// Insert inserts the Crloan to the database.
func (c *Crloan) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.crloans (` +
		`crloanid, crlexid, crlcomcodeid, crltype, crlamount, crlamount2, crlfrom, crlto, crlcommitted, crlsparenum1, crlsparenum2, crlsparec1, crlsparec2, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Crloanid, c.Crlexid, c.Crlcomcodeid, c.Crltype, c.Crlamount, c.Crlamount2, c.Crlfrom, c.Crlto, c.Crlcommitted, c.Crlsparenum1, c.Crlsparenum2, c.Crlsparec1, c.Crlsparec2, c.EquinoxPrn, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Crloanid, c.Crlexid, c.Crlcomcodeid, c.Crltype, c.Crlamount, c.Crlamount2, c.Crlfrom, c.Crlto, c.Crlcommitted, c.Crlsparenum1, c.Crlsparenum2, c.Crlsparec1, c.Crlsparec2, c.EquinoxPrn, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Crloan in the database.
func (c *Crloan) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.crloans SET (` +
		`crloanid, crlexid, crlcomcodeid, crltype, crlamount, crlamount2, crlfrom, crlto, crlcommitted, crlsparenum1, crlsparenum2, crlsparec1, crlsparec2, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15` +
		`) WHERE equinox_lrn = $16`

	// run query
	XOLog(sqlstr, c.Crloanid, c.Crlexid, c.Crlcomcodeid, c.Crltype, c.Crlamount, c.Crlamount2, c.Crlfrom, c.Crlto, c.Crlcommitted, c.Crlsparenum1, c.Crlsparenum2, c.Crlsparec1, c.Crlsparec2, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Crloanid, c.Crlexid, c.Crlcomcodeid, c.Crltype, c.Crlamount, c.Crlamount2, c.Crlfrom, c.Crlto, c.Crlcommitted, c.Crlsparenum1, c.Crlsparenum2, c.Crlsparec1, c.Crlsparec2, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Crloan to the database.
func (c *Crloan) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Crloan.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Crloan) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.crloans (` +
		`crloanid, crlexid, crlcomcodeid, crltype, crlamount, crlamount2, crlfrom, crlto, crlcommitted, crlsparenum1, crlsparenum2, crlsparec1, crlsparec2, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`crloanid, crlexid, crlcomcodeid, crltype, crlamount, crlamount2, crlfrom, crlto, crlcommitted, crlsparenum1, crlsparenum2, crlsparec1, crlsparec2, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.crloanid, EXCLUDED.crlexid, EXCLUDED.crlcomcodeid, EXCLUDED.crltype, EXCLUDED.crlamount, EXCLUDED.crlamount2, EXCLUDED.crlfrom, EXCLUDED.crlto, EXCLUDED.crlcommitted, EXCLUDED.crlsparenum1, EXCLUDED.crlsparenum2, EXCLUDED.crlsparec1, EXCLUDED.crlsparec2, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Crloanid, c.Crlexid, c.Crlcomcodeid, c.Crltype, c.Crlamount, c.Crlamount2, c.Crlfrom, c.Crlto, c.Crlcommitted, c.Crlsparenum1, c.Crlsparenum2, c.Crlsparec1, c.Crlsparec2, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Crloanid, c.Crlexid, c.Crlcomcodeid, c.Crltype, c.Crlamount, c.Crlamount2, c.Crlfrom, c.Crlto, c.Crlcommitted, c.Crlsparenum1, c.Crlsparenum2, c.Crlsparec1, c.Crlsparec2, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Crloan from the database.
func (c *Crloan) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.crloans WHERE equinox_lrn = $1`

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

// CrloanByEquinoxLrn retrieves a row from 'equinox.crloans' as a Crloan.
//
// Generated from index 'crloans_pkey'.
func CrloanByEquinoxLrn(db XODB, equinoxLrn int64) (*Crloan, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`crloanid, crlexid, crlcomcodeid, crltype, crlamount, crlamount2, crlfrom, crlto, crlcommitted, crlsparenum1, crlsparenum2, crlsparec1, crlsparec2, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.crloans ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Crloan{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Crloanid, &c.Crlexid, &c.Crlcomcodeid, &c.Crltype, &c.Crlamount, &c.Crlamount2, &c.Crlfrom, &c.Crlto, &c.Crlcommitted, &c.Crlsparenum1, &c.Crlsparenum2, &c.Crlsparec1, &c.Crlsparec2, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}