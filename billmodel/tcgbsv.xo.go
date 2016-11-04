// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Tcgbsv represents a row from 'equinox.tcgbsvs'.
type Tcgbsv struct {
	Tcgbnumbersvs sql.NullInt64   `json:"tcgbnumbersvs"` // tcgbnumbersvs
	Tcgbamount    sql.NullFloat64 `json:"tcgbamount"`    // tcgbamount
	Tcgbstartdate pq.NullTime     `json:"tcgbstartdate"` // tcgbstartdate
	Tcgbenddate   pq.NullTime     `json:"tcgbenddate"`   // tcgbenddate
	EquinoxLrn    int64           `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec    sql.NullInt64   `json:"equinox_sec"`   // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Tcgbsv exists in the database.
func (t *Tcgbsv) Exists() bool {
	return t._exists
}

// Deleted provides information if the Tcgbsv has been deleted from the database.
func (t *Tcgbsv) Deleted() bool {
	return t._deleted
}

// Insert inserts the Tcgbsv to the database.
func (t *Tcgbsv) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.tcgbsvs (` +
		`tcgbnumbersvs, tcgbamount, tcgbstartdate, tcgbenddate, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, t.Tcgbnumbersvs, t.Tcgbamount, t.Tcgbstartdate, t.Tcgbenddate, t.EquinoxSec)
	err = db.QueryRow(sqlstr, t.Tcgbnumbersvs, t.Tcgbamount, t.Tcgbstartdate, t.Tcgbenddate, t.EquinoxSec).Scan(&t.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Update updates the Tcgbsv in the database.
func (t *Tcgbsv) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.tcgbsvs SET (` +
		`tcgbnumbersvs, tcgbamount, tcgbstartdate, tcgbenddate, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5` +
		`) WHERE equinox_lrn = $6`

	// run query
	XOLog(sqlstr, t.Tcgbnumbersvs, t.Tcgbamount, t.Tcgbstartdate, t.Tcgbenddate, t.EquinoxSec, t.EquinoxLrn)
	_, err = db.Exec(sqlstr, t.Tcgbnumbersvs, t.Tcgbamount, t.Tcgbstartdate, t.Tcgbenddate, t.EquinoxSec, t.EquinoxLrn)
	return err
}

// Save saves the Tcgbsv to the database.
func (t *Tcgbsv) Save(db XODB) error {
	if t.Exists() {
		return t.Update(db)
	}

	return t.Insert(db)
}

// Upsert performs an upsert for Tcgbsv.
//
// NOTE: PostgreSQL 9.5+ only
func (t *Tcgbsv) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.tcgbsvs (` +
		`tcgbnumbersvs, tcgbamount, tcgbstartdate, tcgbenddate, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`tcgbnumbersvs, tcgbamount, tcgbstartdate, tcgbenddate, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.tcgbnumbersvs, EXCLUDED.tcgbamount, EXCLUDED.tcgbstartdate, EXCLUDED.tcgbenddate, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, t.Tcgbnumbersvs, t.Tcgbamount, t.Tcgbstartdate, t.Tcgbenddate, t.EquinoxLrn, t.EquinoxSec)
	_, err = db.Exec(sqlstr, t.Tcgbnumbersvs, t.Tcgbamount, t.Tcgbstartdate, t.Tcgbenddate, t.EquinoxLrn, t.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Delete deletes the Tcgbsv from the database.
func (t *Tcgbsv) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.tcgbsvs WHERE equinox_lrn = $1`

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

// TcgbsvByEquinoxLrn retrieves a row from 'equinox.tcgbsvs' as a Tcgbsv.
//
// Generated from index 'tcgbsvs_pkey'.
func TcgbsvByEquinoxLrn(db XODB, equinoxLrn int64) (*Tcgbsv, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tcgbnumbersvs, tcgbamount, tcgbstartdate, tcgbenddate, equinox_lrn, equinox_sec ` +
		`FROM equinox.tcgbsvs ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	t := Tcgbsv{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&t.Tcgbnumbersvs, &t.Tcgbamount, &t.Tcgbstartdate, &t.Tcgbenddate, &t.EquinoxLrn, &t.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
