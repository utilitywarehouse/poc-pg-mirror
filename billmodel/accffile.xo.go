// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Accffile represents a row from 'equinox.accffile'.
type Accffile struct {
	Gaccftranstype   sql.NullString  `json:"gaccftranstype"`   // gaccftranstype
	Gaccfackcode     sql.NullString  `json:"gaccfackcode"`     // gaccfackcode
	Gaccffiletype    sql.NullString  `json:"gaccffiletype"`    // gaccffiletype
	Gaccfcontactref  sql.NullFloat64 `json:"gaccfcontactref"`  // gaccfcontactref
	Gaccfreadref     sql.NullString  `json:"gaccfreadref"`     // gaccfreadref
	Gaccfreasoncode  sql.NullString  `json:"gaccfreasoncode"`  // gaccfreasoncode
	Gaccfapptextval  sql.NullInt64   `json:"gaccfapptextval"`  // gaccfapptextval
	Gaccfshippercode sql.NullString  `json:"gaccfshippercode"` // gaccfshippercode
	Gaccfmprn        sql.NullFloat64 `json:"gaccfmprn"`        // gaccfmprn
	Gaccfadjstartdte pq.NullTime     `json:"gaccfadjstartdte"` // gaccfadjstartdte
	Gaccfadjenddte   pq.NullTime     `json:"gaccfadjenddte"`   // gaccfadjenddte
	Gaccfadjtype     sql.NullString  `json:"gaccfadjtype"`     // gaccfadjtype
	Gaccfdataitemchg sql.NullString  `json:"gaccfdataitemchg"` // gaccfdataitemchg
	Gaccfrvsdvou     sql.NullInt64   `json:"gaccfrvsdvou"`     // gaccfrvsdvou
	Gaccfrsnremarks  sql.NullInt64   `json:"gaccfrsnremarks"`  // gaccfrsnremarks
	Gaccffileread    sql.NullString  `json:"gaccffileread"`    // gaccffileread
	Gaccfdateread    pq.NullTime     `json:"gaccfdateread"`    // gaccfdateread
	Gaccftimeread    pq.NullTime     `json:"gaccftimeread"`    // gaccftimeread
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Accffile exists in the database.
func (a *Accffile) Exists() bool {
	return a._exists
}

// Deleted provides information if the Accffile has been deleted from the database.
func (a *Accffile) Deleted() bool {
	return a._deleted
}

// Insert inserts the Accffile to the database.
func (a *Accffile) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.accffile (` +
		`gaccftranstype, gaccfackcode, gaccffiletype, gaccfcontactref, gaccfreadref, gaccfreasoncode, gaccfapptextval, gaccfshippercode, gaccfmprn, gaccfadjstartdte, gaccfadjenddte, gaccfadjtype, gaccfdataitemchg, gaccfrvsdvou, gaccfrsnremarks, gaccffileread, gaccfdateread, gaccftimeread, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, a.Gaccftranstype, a.Gaccfackcode, a.Gaccffiletype, a.Gaccfcontactref, a.Gaccfreadref, a.Gaccfreasoncode, a.Gaccfapptextval, a.Gaccfshippercode, a.Gaccfmprn, a.Gaccfadjstartdte, a.Gaccfadjenddte, a.Gaccfadjtype, a.Gaccfdataitemchg, a.Gaccfrvsdvou, a.Gaccfrsnremarks, a.Gaccffileread, a.Gaccfdateread, a.Gaccftimeread, a.EquinoxPrn, a.EquinoxSec)
	err = db.QueryRow(sqlstr, a.Gaccftranstype, a.Gaccfackcode, a.Gaccffiletype, a.Gaccfcontactref, a.Gaccfreadref, a.Gaccfreasoncode, a.Gaccfapptextval, a.Gaccfshippercode, a.Gaccfmprn, a.Gaccfadjstartdte, a.Gaccfadjenddte, a.Gaccfadjtype, a.Gaccfdataitemchg, a.Gaccfrvsdvou, a.Gaccfrsnremarks, a.Gaccffileread, a.Gaccfdateread, a.Gaccftimeread, a.EquinoxPrn, a.EquinoxSec).Scan(&a.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Update updates the Accffile in the database.
func (a *Accffile) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if a._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.accffile SET (` +
		`gaccftranstype, gaccfackcode, gaccffiletype, gaccfcontactref, gaccfreadref, gaccfreasoncode, gaccfapptextval, gaccfshippercode, gaccfmprn, gaccfadjstartdte, gaccfadjenddte, gaccfadjtype, gaccfdataitemchg, gaccfrvsdvou, gaccfrsnremarks, gaccffileread, gaccfdateread, gaccftimeread, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20` +
		`) WHERE equinox_lrn = $21`

	// run query
	XOLog(sqlstr, a.Gaccftranstype, a.Gaccfackcode, a.Gaccffiletype, a.Gaccfcontactref, a.Gaccfreadref, a.Gaccfreasoncode, a.Gaccfapptextval, a.Gaccfshippercode, a.Gaccfmprn, a.Gaccfadjstartdte, a.Gaccfadjenddte, a.Gaccfadjtype, a.Gaccfdataitemchg, a.Gaccfrvsdvou, a.Gaccfrsnremarks, a.Gaccffileread, a.Gaccfdateread, a.Gaccftimeread, a.EquinoxPrn, a.EquinoxSec, a.EquinoxLrn)
	_, err = db.Exec(sqlstr, a.Gaccftranstype, a.Gaccfackcode, a.Gaccffiletype, a.Gaccfcontactref, a.Gaccfreadref, a.Gaccfreasoncode, a.Gaccfapptextval, a.Gaccfshippercode, a.Gaccfmprn, a.Gaccfadjstartdte, a.Gaccfadjenddte, a.Gaccfadjtype, a.Gaccfdataitemchg, a.Gaccfrvsdvou, a.Gaccfrsnremarks, a.Gaccffileread, a.Gaccfdateread, a.Gaccftimeread, a.EquinoxPrn, a.EquinoxSec, a.EquinoxLrn)
	return err
}

// Save saves the Accffile to the database.
func (a *Accffile) Save(db XODB) error {
	if a.Exists() {
		return a.Update(db)
	}

	return a.Insert(db)
}

// Upsert performs an upsert for Accffile.
//
// NOTE: PostgreSQL 9.5+ only
func (a *Accffile) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.accffile (` +
		`gaccftranstype, gaccfackcode, gaccffiletype, gaccfcontactref, gaccfreadref, gaccfreasoncode, gaccfapptextval, gaccfshippercode, gaccfmprn, gaccfadjstartdte, gaccfadjenddte, gaccfadjtype, gaccfdataitemchg, gaccfrvsdvou, gaccfrsnremarks, gaccffileread, gaccfdateread, gaccftimeread, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`gaccftranstype, gaccfackcode, gaccffiletype, gaccfcontactref, gaccfreadref, gaccfreasoncode, gaccfapptextval, gaccfshippercode, gaccfmprn, gaccfadjstartdte, gaccfadjenddte, gaccfadjtype, gaccfdataitemchg, gaccfrvsdvou, gaccfrsnremarks, gaccffileread, gaccfdateread, gaccftimeread, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.gaccftranstype, EXCLUDED.gaccfackcode, EXCLUDED.gaccffiletype, EXCLUDED.gaccfcontactref, EXCLUDED.gaccfreadref, EXCLUDED.gaccfreasoncode, EXCLUDED.gaccfapptextval, EXCLUDED.gaccfshippercode, EXCLUDED.gaccfmprn, EXCLUDED.gaccfadjstartdte, EXCLUDED.gaccfadjenddte, EXCLUDED.gaccfadjtype, EXCLUDED.gaccfdataitemchg, EXCLUDED.gaccfrvsdvou, EXCLUDED.gaccfrsnremarks, EXCLUDED.gaccffileread, EXCLUDED.gaccfdateread, EXCLUDED.gaccftimeread, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, a.Gaccftranstype, a.Gaccfackcode, a.Gaccffiletype, a.Gaccfcontactref, a.Gaccfreadref, a.Gaccfreasoncode, a.Gaccfapptextval, a.Gaccfshippercode, a.Gaccfmprn, a.Gaccfadjstartdte, a.Gaccfadjenddte, a.Gaccfadjtype, a.Gaccfdataitemchg, a.Gaccfrvsdvou, a.Gaccfrsnremarks, a.Gaccffileread, a.Gaccfdateread, a.Gaccftimeread, a.EquinoxPrn, a.EquinoxLrn, a.EquinoxSec)
	_, err = db.Exec(sqlstr, a.Gaccftranstype, a.Gaccfackcode, a.Gaccffiletype, a.Gaccfcontactref, a.Gaccfreadref, a.Gaccfreasoncode, a.Gaccfapptextval, a.Gaccfshippercode, a.Gaccfmprn, a.Gaccfadjstartdte, a.Gaccfadjenddte, a.Gaccfadjtype, a.Gaccfdataitemchg, a.Gaccfrvsdvou, a.Gaccfrsnremarks, a.Gaccffileread, a.Gaccfdateread, a.Gaccftimeread, a.EquinoxPrn, a.EquinoxLrn, a.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Delete deletes the Accffile from the database.
func (a *Accffile) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return nil
	}

	// if deleted, bail
	if a._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.accffile WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, a.EquinoxLrn)
	_, err = db.Exec(sqlstr, a.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	a._deleted = true

	return nil
}

// AccffileByEquinoxLrn retrieves a row from 'equinox.accffile' as a Accffile.
//
// Generated from index 'accffile_pkey'.
func AccffileByEquinoxLrn(db XODB, equinoxLrn int64) (*Accffile, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`gaccftranstype, gaccfackcode, gaccffiletype, gaccfcontactref, gaccfreadref, gaccfreasoncode, gaccfapptextval, gaccfshippercode, gaccfmprn, gaccfadjstartdte, gaccfadjenddte, gaccfadjtype, gaccfdataitemchg, gaccfrvsdvou, gaccfrsnremarks, gaccffileread, gaccfdateread, gaccftimeread, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.accffile ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	a := Accffile{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&a.Gaccftranstype, &a.Gaccfackcode, &a.Gaccffiletype, &a.Gaccfcontactref, &a.Gaccfreadref, &a.Gaccfreasoncode, &a.Gaccfapptextval, &a.Gaccfshippercode, &a.Gaccfmprn, &a.Gaccfadjstartdte, &a.Gaccfadjenddte, &a.Gaccfadjtype, &a.Gaccfdataitemchg, &a.Gaccfrvsdvou, &a.Gaccfrsnremarks, &a.Gaccffileread, &a.Gaccfdateread, &a.Gaccftimeread, &a.EquinoxPrn, &a.EquinoxLrn, &a.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &a, nil
}