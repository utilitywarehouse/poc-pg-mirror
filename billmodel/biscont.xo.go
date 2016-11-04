// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Biscont represents a row from 'equinox.biscont'.
type Biscont struct {
	Biscontcompany sql.NullString `json:"biscontcompany"` // biscontcompany
	Biscontid      sql.NullString `json:"biscontid"`      // biscontid
	Biscontdeptid  sql.NullString `json:"biscontdeptid"`  // biscontdeptid
	Biscontname    sql.NullString `json:"biscontname"`    // biscontname
	Biscontphone   sql.NullString `json:"biscontphone"`   // biscontphone
	Biscontfax     sql.NullString `json:"biscontfax"`     // biscontfax
	Biscontemail   sql.NullString `json:"biscontemail"`   // biscontemail
	Biscontprocess sql.NullString `json:"biscontprocess"` // biscontprocess
	EquinoxLrn     int64          `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64  `json:"equinox_sec"`    // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Biscont exists in the database.
func (b *Biscont) Exists() bool {
	return b._exists
}

// Deleted provides information if the Biscont has been deleted from the database.
func (b *Biscont) Deleted() bool {
	return b._deleted
}

// Insert inserts the Biscont to the database.
func (b *Biscont) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.biscont (` +
		`biscontcompany, biscontid, biscontdeptid, biscontname, biscontphone, biscontfax, biscontemail, biscontprocess, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, b.Biscontcompany, b.Biscontid, b.Biscontdeptid, b.Biscontname, b.Biscontphone, b.Biscontfax, b.Biscontemail, b.Biscontprocess, b.EquinoxSec)
	err = db.QueryRow(sqlstr, b.Biscontcompany, b.Biscontid, b.Biscontdeptid, b.Biscontname, b.Biscontphone, b.Biscontfax, b.Biscontemail, b.Biscontprocess, b.EquinoxSec).Scan(&b.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	b._exists = true

	return nil
}

// Update updates the Biscont in the database.
func (b *Biscont) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !b._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if b._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.biscont SET (` +
		`biscontcompany, biscontid, biscontdeptid, biscontname, biscontphone, biscontfax, biscontemail, biscontprocess, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) WHERE equinox_lrn = $10`

	// run query
	XOLog(sqlstr, b.Biscontcompany, b.Biscontid, b.Biscontdeptid, b.Biscontname, b.Biscontphone, b.Biscontfax, b.Biscontemail, b.Biscontprocess, b.EquinoxSec, b.EquinoxLrn)
	_, err = db.Exec(sqlstr, b.Biscontcompany, b.Biscontid, b.Biscontdeptid, b.Biscontname, b.Biscontphone, b.Biscontfax, b.Biscontemail, b.Biscontprocess, b.EquinoxSec, b.EquinoxLrn)
	return err
}

// Save saves the Biscont to the database.
func (b *Biscont) Save(db XODB) error {
	if b.Exists() {
		return b.Update(db)
	}

	return b.Insert(db)
}

// Upsert performs an upsert for Biscont.
//
// NOTE: PostgreSQL 9.5+ only
func (b *Biscont) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.biscont (` +
		`biscontcompany, biscontid, biscontdeptid, biscontname, biscontphone, biscontfax, biscontemail, biscontprocess, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`biscontcompany, biscontid, biscontdeptid, biscontname, biscontphone, biscontfax, biscontemail, biscontprocess, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.biscontcompany, EXCLUDED.biscontid, EXCLUDED.biscontdeptid, EXCLUDED.biscontname, EXCLUDED.biscontphone, EXCLUDED.biscontfax, EXCLUDED.biscontemail, EXCLUDED.biscontprocess, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, b.Biscontcompany, b.Biscontid, b.Biscontdeptid, b.Biscontname, b.Biscontphone, b.Biscontfax, b.Biscontemail, b.Biscontprocess, b.EquinoxLrn, b.EquinoxSec)
	_, err = db.Exec(sqlstr, b.Biscontcompany, b.Biscontid, b.Biscontdeptid, b.Biscontname, b.Biscontphone, b.Biscontfax, b.Biscontemail, b.Biscontprocess, b.EquinoxLrn, b.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	b._exists = true

	return nil
}

// Delete deletes the Biscont from the database.
func (b *Biscont) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !b._exists {
		return nil
	}

	// if deleted, bail
	if b._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.biscont WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, b.EquinoxLrn)
	_, err = db.Exec(sqlstr, b.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	b._deleted = true

	return nil
}

// BiscontByEquinoxLrn retrieves a row from 'equinox.biscont' as a Biscont.
//
// Generated from index 'biscont_pkey'.
func BiscontByEquinoxLrn(db XODB, equinoxLrn int64) (*Biscont, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`biscontcompany, biscontid, biscontdeptid, biscontname, biscontphone, biscontfax, biscontemail, biscontprocess, equinox_lrn, equinox_sec ` +
		`FROM equinox.biscont ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	b := Biscont{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&b.Biscontcompany, &b.Biscontid, &b.Biscontdeptid, &b.Biscontname, &b.Biscontphone, &b.Biscontfax, &b.Biscontemail, &b.Biscontprocess, &b.EquinoxLrn, &b.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
