// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Tenpromo represents a row from 'equinox.tenpromo'.
type Tenpromo struct {
	Tencustaccountno sql.NullString `json:"tencustaccountno"` // tencustaccountno
	Tenexidold       sql.NullString `json:"tenexidold"`       // tenexidold
	Tenexidnew       sql.NullString `json:"tenexidnew"`       // tenexidnew
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Tenpromo exists in the database.
func (t *Tenpromo) Exists() bool {
	return t._exists
}

// Deleted provides information if the Tenpromo has been deleted from the database.
func (t *Tenpromo) Deleted() bool {
	return t._deleted
}

// Insert inserts the Tenpromo to the database.
func (t *Tenpromo) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.tenpromo (` +
		`tencustaccountno, tenexidold, tenexidnew, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, t.Tencustaccountno, t.Tenexidold, t.Tenexidnew, t.EquinoxSec)
	err = db.QueryRow(sqlstr, t.Tencustaccountno, t.Tenexidold, t.Tenexidnew, t.EquinoxSec).Scan(&t.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Update updates the Tenpromo in the database.
func (t *Tenpromo) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if t._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.tenpromo SET (` +
		`tencustaccountno, tenexidold, tenexidnew, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE equinox_lrn = $5`

	// run query
	XOLog(sqlstr, t.Tencustaccountno, t.Tenexidold, t.Tenexidnew, t.EquinoxSec, t.EquinoxLrn)
	_, err = db.Exec(sqlstr, t.Tencustaccountno, t.Tenexidold, t.Tenexidnew, t.EquinoxSec, t.EquinoxLrn)
	return err
}

// Save saves the Tenpromo to the database.
func (t *Tenpromo) Save(db XODB) error {
	if t.Exists() {
		return t.Update(db)
	}

	return t.Insert(db)
}

// Upsert performs an upsert for Tenpromo.
//
// NOTE: PostgreSQL 9.5+ only
func (t *Tenpromo) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.tenpromo (` +
		`tencustaccountno, tenexidold, tenexidnew, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`tencustaccountno, tenexidold, tenexidnew, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.tencustaccountno, EXCLUDED.tenexidold, EXCLUDED.tenexidnew, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, t.Tencustaccountno, t.Tenexidold, t.Tenexidnew, t.EquinoxLrn, t.EquinoxSec)
	_, err = db.Exec(sqlstr, t.Tencustaccountno, t.Tenexidold, t.Tenexidnew, t.EquinoxLrn, t.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Delete deletes the Tenpromo from the database.
func (t *Tenpromo) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return nil
	}

	// if deleted, bail
	if t._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.tenpromo WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, t.EquinoxLrn)
	_, err = db.Exec(sqlstr, t.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	t._deleted = true

	return nil
}

// TenpromoByEquinoxLrn retrieves a row from 'equinox.tenpromo' as a Tenpromo.
//
// Generated from index 'tenpromo_pkey'.
func TenpromoByEquinoxLrn(db XODB, equinoxLrn int64) (*Tenpromo, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tencustaccountno, tenexidold, tenexidnew, equinox_lrn, equinox_sec ` +
		`FROM equinox.tenpromo ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	t := Tenpromo{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&t.Tencustaccountno, &t.Tenexidold, &t.Tenexidnew, &t.EquinoxLrn, &t.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &t, nil
}