// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Cb represents a row from 'equinox.cb'.
type Cb struct {
	Cbuniquesys      sql.NullInt64  `json:"cbuniquesys"`      // cbuniquesys
	Cblocation       sql.NullString `json:"cblocation"`       // cblocation
	Cbdescription    sql.NullString `json:"cbdescription"`    // cbdescription
	Cbretailband     sql.NullString `json:"cbretailband"`     // cbretailband
	Cbwholebandbt    sql.NullString `json:"cbwholebandbt"`    // cbwholebandbt
	Cbwholebandcw    sql.NullString `json:"cbwholebandcw"`    // cbwholebandcw
	Cbwholebandcwin  sql.NullString `json:"cbwholebandcwin"`  // cbwholebandcwin
	Cbwholebandenerg sql.NullString `json:"cbwholebandenerg"` // cbwholebandenerg
	Cbwholebandgamma sql.NullString `json:"cbwholebandgamma"` // cbwholebandgamma
	Cbwholebandmpc   sql.NullString `json:"cbwholebandmpc"`   // cbwholebandmpc
	Cbwholebandmpcp  sql.NullString `json:"cbwholebandmpcp"`  // cbwholebandmpcp
	Cbwholebandntl   sql.NullString `json:"cbwholebandntl"`   // cbwholebandntl
	Cbwholebando2    sql.NullString `json:"cbwholebando2"`    // cbwholebando2
	Cbwholebandopal  sql.NullString `json:"cbwholebandopal"`  // cbwholebandopal
	Cbtype           sql.NullInt64  `json:"cbtype"`           // cbtype
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Cb exists in the database.
func (c *Cb) Exists() bool {
	return c._exists
}

// Deleted provides information if the Cb has been deleted from the database.
func (c *Cb) Deleted() bool {
	return c._deleted
}

// Insert inserts the Cb to the database.
func (c *Cb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cb (` +
		`cbuniquesys, cblocation, cbdescription, cbretailband, cbwholebandbt, cbwholebandcw, cbwholebandcwin, cbwholebandenerg, cbwholebandgamma, cbwholebandmpc, cbwholebandmpcp, cbwholebandntl, cbwholebando2, cbwholebandopal, cbtype, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Cbuniquesys, c.Cblocation, c.Cbdescription, c.Cbretailband, c.Cbwholebandbt, c.Cbwholebandcw, c.Cbwholebandcwin, c.Cbwholebandenerg, c.Cbwholebandgamma, c.Cbwholebandmpc, c.Cbwholebandmpcp, c.Cbwholebandntl, c.Cbwholebando2, c.Cbwholebandopal, c.Cbtype, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Cbuniquesys, c.Cblocation, c.Cbdescription, c.Cbretailband, c.Cbwholebandbt, c.Cbwholebandcw, c.Cbwholebandcwin, c.Cbwholebandenerg, c.Cbwholebandgamma, c.Cbwholebandmpc, c.Cbwholebandmpcp, c.Cbwholebandntl, c.Cbwholebando2, c.Cbwholebandopal, c.Cbtype, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Cb in the database.
func (c *Cb) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.cb SET (` +
		`cbuniquesys, cblocation, cbdescription, cbretailband, cbwholebandbt, cbwholebandcw, cbwholebandcwin, cbwholebandenerg, cbwholebandgamma, cbwholebandmpc, cbwholebandmpcp, cbwholebandntl, cbwholebando2, cbwholebandopal, cbtype, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16` +
		`) WHERE equinox_lrn = $17`

	// run query
	XOLog(sqlstr, c.Cbuniquesys, c.Cblocation, c.Cbdescription, c.Cbretailband, c.Cbwholebandbt, c.Cbwholebandcw, c.Cbwholebandcwin, c.Cbwholebandenerg, c.Cbwholebandgamma, c.Cbwholebandmpc, c.Cbwholebandmpcp, c.Cbwholebandntl, c.Cbwholebando2, c.Cbwholebandopal, c.Cbtype, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Cbuniquesys, c.Cblocation, c.Cbdescription, c.Cbretailband, c.Cbwholebandbt, c.Cbwholebandcw, c.Cbwholebandcwin, c.Cbwholebandenerg, c.Cbwholebandgamma, c.Cbwholebandmpc, c.Cbwholebandmpcp, c.Cbwholebandntl, c.Cbwholebando2, c.Cbwholebandopal, c.Cbtype, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Cb to the database.
func (c *Cb) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Cb.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Cb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cb (` +
		`cbuniquesys, cblocation, cbdescription, cbretailband, cbwholebandbt, cbwholebandcw, cbwholebandcwin, cbwholebandenerg, cbwholebandgamma, cbwholebandmpc, cbwholebandmpcp, cbwholebandntl, cbwholebando2, cbwholebandopal, cbtype, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`cbuniquesys, cblocation, cbdescription, cbretailband, cbwholebandbt, cbwholebandcw, cbwholebandcwin, cbwholebandenerg, cbwholebandgamma, cbwholebandmpc, cbwholebandmpcp, cbwholebandntl, cbwholebando2, cbwholebandopal, cbtype, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.cbuniquesys, EXCLUDED.cblocation, EXCLUDED.cbdescription, EXCLUDED.cbretailband, EXCLUDED.cbwholebandbt, EXCLUDED.cbwholebandcw, EXCLUDED.cbwholebandcwin, EXCLUDED.cbwholebandenerg, EXCLUDED.cbwholebandgamma, EXCLUDED.cbwholebandmpc, EXCLUDED.cbwholebandmpcp, EXCLUDED.cbwholebandntl, EXCLUDED.cbwholebando2, EXCLUDED.cbwholebandopal, EXCLUDED.cbtype, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Cbuniquesys, c.Cblocation, c.Cbdescription, c.Cbretailband, c.Cbwholebandbt, c.Cbwholebandcw, c.Cbwholebandcwin, c.Cbwholebandenerg, c.Cbwholebandgamma, c.Cbwholebandmpc, c.Cbwholebandmpcp, c.Cbwholebandntl, c.Cbwholebando2, c.Cbwholebandopal, c.Cbtype, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Cbuniquesys, c.Cblocation, c.Cbdescription, c.Cbretailband, c.Cbwholebandbt, c.Cbwholebandcw, c.Cbwholebandcwin, c.Cbwholebandenerg, c.Cbwholebandgamma, c.Cbwholebandmpc, c.Cbwholebandmpcp, c.Cbwholebandntl, c.Cbwholebando2, c.Cbwholebandopal, c.Cbtype, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Cb from the database.
func (c *Cb) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.cb WHERE equinox_lrn = $1`

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

// CbByEquinoxLrn retrieves a row from 'equinox.cb' as a Cb.
//
// Generated from index 'cb_pkey'.
func CbByEquinoxLrn(db XODB, equinoxLrn int64) (*Cb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cbuniquesys, cblocation, cbdescription, cbretailband, cbwholebandbt, cbwholebandcw, cbwholebandcwin, cbwholebandenerg, cbwholebandgamma, cbwholebandmpc, cbwholebandmpcp, cbwholebandntl, cbwholebando2, cbwholebandopal, cbtype, equinox_lrn, equinox_sec ` +
		`FROM equinox.cb ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cbuniquesys, &c.Cblocation, &c.Cbdescription, &c.Cbretailband, &c.Cbwholebandbt, &c.Cbwholebandcw, &c.Cbwholebandcwin, &c.Cbwholebandenerg, &c.Cbwholebandgamma, &c.Cbwholebandmpc, &c.Cbwholebandmpcp, &c.Cbwholebandntl, &c.Cbwholebando2, &c.Cbwholebandopal, &c.Cbtype, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
