// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Chrgcall represents a row from 'equinox.chrgcall'.
type Chrgcall struct {
	Chrgcode         sql.NullString  `json:"chrgcode"`         // chrgcode
	Chrgmincostpence sql.NullFloat64 `json:"chrgmincostpence"` // chrgmincostpence
	Chrgcallunits    sql.NullInt64   `json:"chrgcallunits"`    // chrgcallunits
	Chrgminduration  sql.NullInt64   `json:"chrgminduration"`  // chrgminduration
	Chrgroundingpenc sql.NullInt64   `json:"chrgroundingpenc"` // chrgroundingpenc
	Chrgcallsetup    sql.NullFloat64 `json:"chrgcallsetup"`    // chrgcallsetup
	Chrgaccesspence  sql.NullFloat64 `json:"chrgaccesspence"`  // chrgaccesspence
	Chrgroundup      sql.NullInt64   `json:"chrgroundup"`      // chrgroundup
	Chrgdescrip      sql.NullString  `json:"chrgdescrip"`      // chrgdescrip
	Chrgmaxcostpence sql.NullFloat64 `json:"chrgmaxcostpence"` // chrgmaxcostpence
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Chrgcall exists in the database.
func (c *Chrgcall) Exists() bool {
	return c._exists
}

// Deleted provides information if the Chrgcall has been deleted from the database.
func (c *Chrgcall) Deleted() bool {
	return c._deleted
}

// Insert inserts the Chrgcall to the database.
func (c *Chrgcall) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.chrgcall (` +
		`chrgcode, chrgmincostpence, chrgcallunits, chrgminduration, chrgroundingpenc, chrgcallsetup, chrgaccesspence, chrgroundup, chrgdescrip, chrgmaxcostpence, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Chrgcode, c.Chrgmincostpence, c.Chrgcallunits, c.Chrgminduration, c.Chrgroundingpenc, c.Chrgcallsetup, c.Chrgaccesspence, c.Chrgroundup, c.Chrgdescrip, c.Chrgmaxcostpence, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Chrgcode, c.Chrgmincostpence, c.Chrgcallunits, c.Chrgminduration, c.Chrgroundingpenc, c.Chrgcallsetup, c.Chrgaccesspence, c.Chrgroundup, c.Chrgdescrip, c.Chrgmaxcostpence, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Chrgcall in the database.
func (c *Chrgcall) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.chrgcall SET (` +
		`chrgcode, chrgmincostpence, chrgcallunits, chrgminduration, chrgroundingpenc, chrgcallsetup, chrgaccesspence, chrgroundup, chrgdescrip, chrgmaxcostpence, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) WHERE equinox_lrn = $12`

	// run query
	XOLog(sqlstr, c.Chrgcode, c.Chrgmincostpence, c.Chrgcallunits, c.Chrgminduration, c.Chrgroundingpenc, c.Chrgcallsetup, c.Chrgaccesspence, c.Chrgroundup, c.Chrgdescrip, c.Chrgmaxcostpence, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Chrgcode, c.Chrgmincostpence, c.Chrgcallunits, c.Chrgminduration, c.Chrgroundingpenc, c.Chrgcallsetup, c.Chrgaccesspence, c.Chrgroundup, c.Chrgdescrip, c.Chrgmaxcostpence, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Chrgcall to the database.
func (c *Chrgcall) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Chrgcall.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Chrgcall) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.chrgcall (` +
		`chrgcode, chrgmincostpence, chrgcallunits, chrgminduration, chrgroundingpenc, chrgcallsetup, chrgaccesspence, chrgroundup, chrgdescrip, chrgmaxcostpence, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`chrgcode, chrgmincostpence, chrgcallunits, chrgminduration, chrgroundingpenc, chrgcallsetup, chrgaccesspence, chrgroundup, chrgdescrip, chrgmaxcostpence, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.chrgcode, EXCLUDED.chrgmincostpence, EXCLUDED.chrgcallunits, EXCLUDED.chrgminduration, EXCLUDED.chrgroundingpenc, EXCLUDED.chrgcallsetup, EXCLUDED.chrgaccesspence, EXCLUDED.chrgroundup, EXCLUDED.chrgdescrip, EXCLUDED.chrgmaxcostpence, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Chrgcode, c.Chrgmincostpence, c.Chrgcallunits, c.Chrgminduration, c.Chrgroundingpenc, c.Chrgcallsetup, c.Chrgaccesspence, c.Chrgroundup, c.Chrgdescrip, c.Chrgmaxcostpence, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Chrgcode, c.Chrgmincostpence, c.Chrgcallunits, c.Chrgminduration, c.Chrgroundingpenc, c.Chrgcallsetup, c.Chrgaccesspence, c.Chrgroundup, c.Chrgdescrip, c.Chrgmaxcostpence, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Chrgcall from the database.
func (c *Chrgcall) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.chrgcall WHERE equinox_lrn = $1`

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

// ChrgcallByEquinoxLrn retrieves a row from 'equinox.chrgcall' as a Chrgcall.
//
// Generated from index 'chrgcall_pkey'.
func ChrgcallByEquinoxLrn(db XODB, equinoxLrn int64) (*Chrgcall, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`chrgcode, chrgmincostpence, chrgcallunits, chrgminduration, chrgroundingpenc, chrgcallsetup, chrgaccesspence, chrgroundup, chrgdescrip, chrgmaxcostpence, equinox_lrn, equinox_sec ` +
		`FROM equinox.chrgcall ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Chrgcall{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Chrgcode, &c.Chrgmincostpence, &c.Chrgcallunits, &c.Chrgminduration, &c.Chrgroundingpenc, &c.Chrgcallsetup, &c.Chrgaccesspence, &c.Chrgroundup, &c.Chrgdescrip, &c.Chrgmaxcostpence, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
