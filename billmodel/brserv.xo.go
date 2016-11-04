// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Brserv represents a row from 'equinox.brserv'.
type Brserv struct {
	Brstype      sql.NullString  `json:"brstype"`      // brstype
	Brscallcnt   sql.NullFloat64 `json:"brscallcnt"`   // brscallcnt
	Brscallsecs  sql.NullFloat64 `json:"brscallsecs"`  // brscallsecs
	Brsretail    sql.NullFloat64 `json:"brsretail"`    // brsretail
	Brsmonnet    sql.NullFloat64 `json:"brsmonnet"`    // brsmonnet
	Brsmontot    sql.NullFloat64 `json:"brsmontot"`    // brsmontot
	Brsmonvatstd sql.NullFloat64 `json:"brsmonvatstd"` // brsmonvatstd
	Brsmonvatnrg sql.NullFloat64 `json:"brsmonvatnrg"` // brsmonvatnrg
	Brswhole     sql.NullFloat64 `json:"brswhole"`     // brswhole
	Brsvatstd    sql.NullFloat64 `json:"brsvatstd"`    // brsvatstd
	Brsvatnrg    sql.NullFloat64 `json:"brsvatnrg"`    // brsvatnrg
	EquinoxPrn   sql.NullInt64   `json:"equinox_prn"`  // equinox_prn
	EquinoxLrn   int64           `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec   sql.NullInt64   `json:"equinox_sec"`  // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Brserv exists in the database.
func (b *Brserv) Exists() bool {
	return b._exists
}

// Deleted provides information if the Brserv has been deleted from the database.
func (b *Brserv) Deleted() bool {
	return b._deleted
}

// Insert inserts the Brserv to the database.
func (b *Brserv) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.brserv (` +
		`brstype, brscallcnt, brscallsecs, brsretail, brsmonnet, brsmontot, brsmonvatstd, brsmonvatnrg, brswhole, brsvatstd, brsvatnrg, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, b.Brstype, b.Brscallcnt, b.Brscallsecs, b.Brsretail, b.Brsmonnet, b.Brsmontot, b.Brsmonvatstd, b.Brsmonvatnrg, b.Brswhole, b.Brsvatstd, b.Brsvatnrg, b.EquinoxPrn, b.EquinoxSec)
	err = db.QueryRow(sqlstr, b.Brstype, b.Brscallcnt, b.Brscallsecs, b.Brsretail, b.Brsmonnet, b.Brsmontot, b.Brsmonvatstd, b.Brsmonvatnrg, b.Brswhole, b.Brsvatstd, b.Brsvatnrg, b.EquinoxPrn, b.EquinoxSec).Scan(&b.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	b._exists = true

	return nil
}

// Update updates the Brserv in the database.
func (b *Brserv) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.brserv SET (` +
		`brstype, brscallcnt, brscallsecs, brsretail, brsmonnet, brsmontot, brsmonvatstd, brsmonvatnrg, brswhole, brsvatstd, brsvatnrg, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`) WHERE equinox_lrn = $14`

	// run query
	XOLog(sqlstr, b.Brstype, b.Brscallcnt, b.Brscallsecs, b.Brsretail, b.Brsmonnet, b.Brsmontot, b.Brsmonvatstd, b.Brsmonvatnrg, b.Brswhole, b.Brsvatstd, b.Brsvatnrg, b.EquinoxPrn, b.EquinoxSec, b.EquinoxLrn)
	_, err = db.Exec(sqlstr, b.Brstype, b.Brscallcnt, b.Brscallsecs, b.Brsretail, b.Brsmonnet, b.Brsmontot, b.Brsmonvatstd, b.Brsmonvatnrg, b.Brswhole, b.Brsvatstd, b.Brsvatnrg, b.EquinoxPrn, b.EquinoxSec, b.EquinoxLrn)
	return err
}

// Save saves the Brserv to the database.
func (b *Brserv) Save(db XODB) error {
	if b.Exists() {
		return b.Update(db)
	}

	return b.Insert(db)
}

// Upsert performs an upsert for Brserv.
//
// NOTE: PostgreSQL 9.5+ only
func (b *Brserv) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.brserv (` +
		`brstype, brscallcnt, brscallsecs, brsretail, brsmonnet, brsmontot, brsmonvatstd, brsmonvatnrg, brswhole, brsvatstd, brsvatnrg, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`brstype, brscallcnt, brscallsecs, brsretail, brsmonnet, brsmontot, brsmonvatstd, brsmonvatnrg, brswhole, brsvatstd, brsvatnrg, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.brstype, EXCLUDED.brscallcnt, EXCLUDED.brscallsecs, EXCLUDED.brsretail, EXCLUDED.brsmonnet, EXCLUDED.brsmontot, EXCLUDED.brsmonvatstd, EXCLUDED.brsmonvatnrg, EXCLUDED.brswhole, EXCLUDED.brsvatstd, EXCLUDED.brsvatnrg, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, b.Brstype, b.Brscallcnt, b.Brscallsecs, b.Brsretail, b.Brsmonnet, b.Brsmontot, b.Brsmonvatstd, b.Brsmonvatnrg, b.Brswhole, b.Brsvatstd, b.Brsvatnrg, b.EquinoxPrn, b.EquinoxLrn, b.EquinoxSec)
	_, err = db.Exec(sqlstr, b.Brstype, b.Brscallcnt, b.Brscallsecs, b.Brsretail, b.Brsmonnet, b.Brsmontot, b.Brsmonvatstd, b.Brsmonvatnrg, b.Brswhole, b.Brsvatstd, b.Brsvatnrg, b.EquinoxPrn, b.EquinoxLrn, b.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	b._exists = true

	return nil
}

// Delete deletes the Brserv from the database.
func (b *Brserv) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.brserv WHERE equinox_lrn = $1`

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

// BrservByEquinoxLrn retrieves a row from 'equinox.brserv' as a Brserv.
//
// Generated from index 'brserv_pkey'.
func BrservByEquinoxLrn(db XODB, equinoxLrn int64) (*Brserv, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`brstype, brscallcnt, brscallsecs, brsretail, brsmonnet, brsmontot, brsmonvatstd, brsmonvatnrg, brswhole, brsvatstd, brsvatnrg, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.brserv ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	b := Brserv{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&b.Brstype, &b.Brscallcnt, &b.Brscallsecs, &b.Brsretail, &b.Brsmonnet, &b.Brsmontot, &b.Brsmonvatstd, &b.Brsmonvatnrg, &b.Brswhole, &b.Brsvatstd, &b.Brsvatnrg, &b.EquinoxPrn, &b.EquinoxLrn, &b.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
