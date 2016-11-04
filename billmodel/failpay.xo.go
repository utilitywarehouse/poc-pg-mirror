// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Failpay represents a row from 'equinox.failpay'.
type Failpay struct {
	Fpaccount       sql.NullString  `json:"fpaccount"`       // fpaccount
	Fpadd1          sql.NullString  `json:"fpadd1"`          // fpadd1
	Fpadd2          sql.NullString  `json:"fpadd2"`          // fpadd2
	Fpadd3          sql.NullString  `json:"fpadd3"`          // fpadd3
	Fpadd4          sql.NullString  `json:"fpadd4"`          // fpadd4
	Fpcounty        sql.NullString  `json:"fpcounty"`        // fpcounty
	Fppostcode      sql.NullString  `json:"fppostcode"`      // fppostcode
	Fpsqlid         sql.NullString  `json:"fpsqlid"`         // fpsqlid
	Fpccnumber      sql.NullString  `json:"fpccnumber"`      // fpccnumber
	Fpccstart       sql.NullString  `json:"fpccstart"`       // fpccstart
	Fpccend         sql.NullString  `json:"fpccend"`         // fpccend
	Fpccissue       sql.NullString  `json:"fpccissue"`       // fpccissue
	Fpccname        sql.NullString  `json:"fpccname"`        // fpccname
	Fpccpostcode    sql.NullString  `json:"fpccpostcode"`    // fpccpostcode
	Fpccpciuniqueno sql.NullString  `json:"fpccpciuniqueno"` // fpccpciuniqueno
	Fpcctotal       sql.NullFloat64 `json:"fpcctotal"`       // fpcctotal
	Fpcccbc         sql.NullFloat64 `json:"fpcccbc"`         // fpcccbc
	Fpccreddeposit  sql.NullFloat64 `json:"fpccreddeposit"`  // fpccreddeposit
	Fpcchandsetdep  sql.NullFloat64 `json:"fpcchandsetdep"`  // fpcchandsetdep
	Fpcchsetoneoff  sql.NullFloat64 `json:"fpcchsetoneoff"`  // fpcchsetoneoff
	Fpcconlyverify  sql.NullFloat64 `json:"fpcconlyverify"`  // fpcconlyverify
	Fpcctranserrcde sql.NullString  `json:"fpcctranserrcde"` // fpcctranserrcde
	Fpccprepayhdset sql.NullFloat64 `json:"fpccprepayhdset"` // fpccprepayhdset
	Fpccauthcode    sql.NullString  `json:"fpccauthcode"`    // fpccauthcode
	Fpcctype        sql.NullString  `json:"fpcctype"`        // fpcctype
	Fpreexported    sql.NullInt64   `json:"fpreexported"`    // fpreexported
	Fpdate          pq.NullTime     `json:"fpdate"`          // fpdate
	Fpcctranserrmsg sql.NullString  `json:"fpcctranserrmsg"` // fpcctranserrmsg
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Failpay exists in the database.
func (f *Failpay) Exists() bool {
	return f._exists
}

// Deleted provides information if the Failpay has been deleted from the database.
func (f *Failpay) Deleted() bool {
	return f._deleted
}

// Insert inserts the Failpay to the database.
func (f *Failpay) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if f._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.failpay (` +
		`fpaccount, fpadd1, fpadd2, fpadd3, fpadd4, fpcounty, fppostcode, fpsqlid, fpccnumber, fpccstart, fpccend, fpccissue, fpccname, fpccpostcode, fpccpciuniqueno, fpcctotal, fpcccbc, fpccreddeposit, fpcchandsetdep, fpcchsetoneoff, fpcconlyverify, fpcctranserrcde, fpccprepayhdset, fpccauthcode, fpcctype, fpreexported, fpdate, fpcctranserrmsg, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, f.Fpaccount, f.Fpadd1, f.Fpadd2, f.Fpadd3, f.Fpadd4, f.Fpcounty, f.Fppostcode, f.Fpsqlid, f.Fpccnumber, f.Fpccstart, f.Fpccend, f.Fpccissue, f.Fpccname, f.Fpccpostcode, f.Fpccpciuniqueno, f.Fpcctotal, f.Fpcccbc, f.Fpccreddeposit, f.Fpcchandsetdep, f.Fpcchsetoneoff, f.Fpcconlyverify, f.Fpcctranserrcde, f.Fpccprepayhdset, f.Fpccauthcode, f.Fpcctype, f.Fpreexported, f.Fpdate, f.Fpcctranserrmsg, f.EquinoxSec)
	err = db.QueryRow(sqlstr, f.Fpaccount, f.Fpadd1, f.Fpadd2, f.Fpadd3, f.Fpadd4, f.Fpcounty, f.Fppostcode, f.Fpsqlid, f.Fpccnumber, f.Fpccstart, f.Fpccend, f.Fpccissue, f.Fpccname, f.Fpccpostcode, f.Fpccpciuniqueno, f.Fpcctotal, f.Fpcccbc, f.Fpccreddeposit, f.Fpcchandsetdep, f.Fpcchsetoneoff, f.Fpcconlyverify, f.Fpcctranserrcde, f.Fpccprepayhdset, f.Fpccauthcode, f.Fpcctype, f.Fpreexported, f.Fpdate, f.Fpcctranserrmsg, f.EquinoxSec).Scan(&f.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	f._exists = true

	return nil
}

// Update updates the Failpay in the database.
func (f *Failpay) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !f._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if f._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.failpay SET (` +
		`fpaccount, fpadd1, fpadd2, fpadd3, fpadd4, fpcounty, fppostcode, fpsqlid, fpccnumber, fpccstart, fpccend, fpccissue, fpccname, fpccpostcode, fpccpciuniqueno, fpcctotal, fpcccbc, fpccreddeposit, fpcchandsetdep, fpcchsetoneoff, fpcconlyverify, fpcctranserrcde, fpccprepayhdset, fpccauthcode, fpcctype, fpreexported, fpdate, fpcctranserrmsg, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29` +
		`) WHERE equinox_lrn = $30`

	// run query
	XOLog(sqlstr, f.Fpaccount, f.Fpadd1, f.Fpadd2, f.Fpadd3, f.Fpadd4, f.Fpcounty, f.Fppostcode, f.Fpsqlid, f.Fpccnumber, f.Fpccstart, f.Fpccend, f.Fpccissue, f.Fpccname, f.Fpccpostcode, f.Fpccpciuniqueno, f.Fpcctotal, f.Fpcccbc, f.Fpccreddeposit, f.Fpcchandsetdep, f.Fpcchsetoneoff, f.Fpcconlyverify, f.Fpcctranserrcde, f.Fpccprepayhdset, f.Fpccauthcode, f.Fpcctype, f.Fpreexported, f.Fpdate, f.Fpcctranserrmsg, f.EquinoxSec, f.EquinoxLrn)
	_, err = db.Exec(sqlstr, f.Fpaccount, f.Fpadd1, f.Fpadd2, f.Fpadd3, f.Fpadd4, f.Fpcounty, f.Fppostcode, f.Fpsqlid, f.Fpccnumber, f.Fpccstart, f.Fpccend, f.Fpccissue, f.Fpccname, f.Fpccpostcode, f.Fpccpciuniqueno, f.Fpcctotal, f.Fpcccbc, f.Fpccreddeposit, f.Fpcchandsetdep, f.Fpcchsetoneoff, f.Fpcconlyverify, f.Fpcctranserrcde, f.Fpccprepayhdset, f.Fpccauthcode, f.Fpcctype, f.Fpreexported, f.Fpdate, f.Fpcctranserrmsg, f.EquinoxSec, f.EquinoxLrn)
	return err
}

// Save saves the Failpay to the database.
func (f *Failpay) Save(db XODB) error {
	if f.Exists() {
		return f.Update(db)
	}

	return f.Insert(db)
}

// Upsert performs an upsert for Failpay.
//
// NOTE: PostgreSQL 9.5+ only
func (f *Failpay) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if f._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.failpay (` +
		`fpaccount, fpadd1, fpadd2, fpadd3, fpadd4, fpcounty, fppostcode, fpsqlid, fpccnumber, fpccstart, fpccend, fpccissue, fpccname, fpccpostcode, fpccpciuniqueno, fpcctotal, fpcccbc, fpccreddeposit, fpcchandsetdep, fpcchsetoneoff, fpcconlyverify, fpcctranserrcde, fpccprepayhdset, fpccauthcode, fpcctype, fpreexported, fpdate, fpcctranserrmsg, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`fpaccount, fpadd1, fpadd2, fpadd3, fpadd4, fpcounty, fppostcode, fpsqlid, fpccnumber, fpccstart, fpccend, fpccissue, fpccname, fpccpostcode, fpccpciuniqueno, fpcctotal, fpcccbc, fpccreddeposit, fpcchandsetdep, fpcchsetoneoff, fpcconlyverify, fpcctranserrcde, fpccprepayhdset, fpccauthcode, fpcctype, fpreexported, fpdate, fpcctranserrmsg, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.fpaccount, EXCLUDED.fpadd1, EXCLUDED.fpadd2, EXCLUDED.fpadd3, EXCLUDED.fpadd4, EXCLUDED.fpcounty, EXCLUDED.fppostcode, EXCLUDED.fpsqlid, EXCLUDED.fpccnumber, EXCLUDED.fpccstart, EXCLUDED.fpccend, EXCLUDED.fpccissue, EXCLUDED.fpccname, EXCLUDED.fpccpostcode, EXCLUDED.fpccpciuniqueno, EXCLUDED.fpcctotal, EXCLUDED.fpcccbc, EXCLUDED.fpccreddeposit, EXCLUDED.fpcchandsetdep, EXCLUDED.fpcchsetoneoff, EXCLUDED.fpcconlyverify, EXCLUDED.fpcctranserrcde, EXCLUDED.fpccprepayhdset, EXCLUDED.fpccauthcode, EXCLUDED.fpcctype, EXCLUDED.fpreexported, EXCLUDED.fpdate, EXCLUDED.fpcctranserrmsg, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, f.Fpaccount, f.Fpadd1, f.Fpadd2, f.Fpadd3, f.Fpadd4, f.Fpcounty, f.Fppostcode, f.Fpsqlid, f.Fpccnumber, f.Fpccstart, f.Fpccend, f.Fpccissue, f.Fpccname, f.Fpccpostcode, f.Fpccpciuniqueno, f.Fpcctotal, f.Fpcccbc, f.Fpccreddeposit, f.Fpcchandsetdep, f.Fpcchsetoneoff, f.Fpcconlyverify, f.Fpcctranserrcde, f.Fpccprepayhdset, f.Fpccauthcode, f.Fpcctype, f.Fpreexported, f.Fpdate, f.Fpcctranserrmsg, f.EquinoxLrn, f.EquinoxSec)
	_, err = db.Exec(sqlstr, f.Fpaccount, f.Fpadd1, f.Fpadd2, f.Fpadd3, f.Fpadd4, f.Fpcounty, f.Fppostcode, f.Fpsqlid, f.Fpccnumber, f.Fpccstart, f.Fpccend, f.Fpccissue, f.Fpccname, f.Fpccpostcode, f.Fpccpciuniqueno, f.Fpcctotal, f.Fpcccbc, f.Fpccreddeposit, f.Fpcchandsetdep, f.Fpcchsetoneoff, f.Fpcconlyverify, f.Fpcctranserrcde, f.Fpccprepayhdset, f.Fpccauthcode, f.Fpcctype, f.Fpreexported, f.Fpdate, f.Fpcctranserrmsg, f.EquinoxLrn, f.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	f._exists = true

	return nil
}

// Delete deletes the Failpay from the database.
func (f *Failpay) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !f._exists {
		return nil
	}

	// if deleted, bail
	if f._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.failpay WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, f.EquinoxLrn)
	_, err = db.Exec(sqlstr, f.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	f._deleted = true

	return nil
}

// FailpayByEquinoxLrn retrieves a row from 'equinox.failpay' as a Failpay.
//
// Generated from index 'failpay_pkey'.
func FailpayByEquinoxLrn(db XODB, equinoxLrn int64) (*Failpay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`fpaccount, fpadd1, fpadd2, fpadd3, fpadd4, fpcounty, fppostcode, fpsqlid, fpccnumber, fpccstart, fpccend, fpccissue, fpccname, fpccpostcode, fpccpciuniqueno, fpcctotal, fpcccbc, fpccreddeposit, fpcchandsetdep, fpcchsetoneoff, fpcconlyverify, fpcctranserrcde, fpccprepayhdset, fpccauthcode, fpcctype, fpreexported, fpdate, fpcctranserrmsg, equinox_lrn, equinox_sec ` +
		`FROM equinox.failpay ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	f := Failpay{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&f.Fpaccount, &f.Fpadd1, &f.Fpadd2, &f.Fpadd3, &f.Fpadd4, &f.Fpcounty, &f.Fppostcode, &f.Fpsqlid, &f.Fpccnumber, &f.Fpccstart, &f.Fpccend, &f.Fpccissue, &f.Fpccname, &f.Fpccpostcode, &f.Fpccpciuniqueno, &f.Fpcctotal, &f.Fpcccbc, &f.Fpccreddeposit, &f.Fpcchandsetdep, &f.Fpcchsetoneoff, &f.Fpcconlyverify, &f.Fpcctranserrcde, &f.Fpccprepayhdset, &f.Fpccauthcode, &f.Fpcctype, &f.Fpreexported, &f.Fpdate, &f.Fpcctranserrmsg, &f.EquinoxLrn, &f.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &f, nil
}
