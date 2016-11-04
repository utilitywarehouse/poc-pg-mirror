// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Udcpay represents a row from 'equinox.udcpay'.
type Udcpay struct {
	Udcpayment      sql.NullFloat64 `json:"udcpayment"`      // udcpayment
	Udcpaymentdate  pq.NullTime     `json:"udcpaymentdate"`  // udcpaymentdate
	Udcpaymenttype  sql.NullString  `json:"udcpaymenttype"`  // udcpaymenttype
	Udcpaymenttime  pq.NullTime     `json:"udcpaymenttime"`  // udcpaymenttime
	Udcpaymentno    sql.NullInt64   `json:"udcpaymentno"`    // udcpaymentno
	Udcpayvspcode   sql.NullString  `json:"udcpayvspcode"`   // udcpayvspcode
	Udcpaybankcode  sql.NullString  `json:"udcpaybankcode"`  // udcpaybankcode
	Udcpaymentvalid sql.NullString  `json:"udcpaymentvalid"` // udcpaymentvalid
	Udcpayledsys    sql.NullInt64   `json:"udcpayledsys"`    // udcpayledsys
	EquinoxPrn      sql.NullInt64   `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Udcpay exists in the database.
func (u *Udcpay) Exists() bool {
	return u._exists
}

// Deleted provides information if the Udcpay has been deleted from the database.
func (u *Udcpay) Deleted() bool {
	return u._deleted
}

// Insert inserts the Udcpay to the database.
func (u *Udcpay) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.udcpay (` +
		`udcpayment, udcpaymentdate, udcpaymenttype, udcpaymenttime, udcpaymentno, udcpayvspcode, udcpaybankcode, udcpaymentvalid, udcpayledsys, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, u.Udcpayment, u.Udcpaymentdate, u.Udcpaymenttype, u.Udcpaymenttime, u.Udcpaymentno, u.Udcpayvspcode, u.Udcpaybankcode, u.Udcpaymentvalid, u.Udcpayledsys, u.EquinoxPrn, u.EquinoxSec)
	err = db.QueryRow(sqlstr, u.Udcpayment, u.Udcpaymentdate, u.Udcpaymenttype, u.Udcpaymenttime, u.Udcpaymentno, u.Udcpayvspcode, u.Udcpaybankcode, u.Udcpaymentvalid, u.Udcpayledsys, u.EquinoxPrn, u.EquinoxSec).Scan(&u.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	u._exists = true

	return nil
}

// Update updates the Udcpay in the database.
func (u *Udcpay) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if u._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.udcpay SET (` +
		`udcpayment, udcpaymentdate, udcpaymenttype, udcpaymenttime, udcpaymentno, udcpayvspcode, udcpaybankcode, udcpaymentvalid, udcpayledsys, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) WHERE equinox_lrn = $12`

	// run query
	XOLog(sqlstr, u.Udcpayment, u.Udcpaymentdate, u.Udcpaymenttype, u.Udcpaymenttime, u.Udcpaymentno, u.Udcpayvspcode, u.Udcpaybankcode, u.Udcpaymentvalid, u.Udcpayledsys, u.EquinoxPrn, u.EquinoxSec, u.EquinoxLrn)
	_, err = db.Exec(sqlstr, u.Udcpayment, u.Udcpaymentdate, u.Udcpaymenttype, u.Udcpaymenttime, u.Udcpaymentno, u.Udcpayvspcode, u.Udcpaybankcode, u.Udcpaymentvalid, u.Udcpayledsys, u.EquinoxPrn, u.EquinoxSec, u.EquinoxLrn)
	return err
}

// Save saves the Udcpay to the database.
func (u *Udcpay) Save(db XODB) error {
	if u.Exists() {
		return u.Update(db)
	}

	return u.Insert(db)
}

// Upsert performs an upsert for Udcpay.
//
// NOTE: PostgreSQL 9.5+ only
func (u *Udcpay) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.udcpay (` +
		`udcpayment, udcpaymentdate, udcpaymenttype, udcpaymenttime, udcpaymentno, udcpayvspcode, udcpaybankcode, udcpaymentvalid, udcpayledsys, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`udcpayment, udcpaymentdate, udcpaymenttype, udcpaymenttime, udcpaymentno, udcpayvspcode, udcpaybankcode, udcpaymentvalid, udcpayledsys, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.udcpayment, EXCLUDED.udcpaymentdate, EXCLUDED.udcpaymenttype, EXCLUDED.udcpaymenttime, EXCLUDED.udcpaymentno, EXCLUDED.udcpayvspcode, EXCLUDED.udcpaybankcode, EXCLUDED.udcpaymentvalid, EXCLUDED.udcpayledsys, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, u.Udcpayment, u.Udcpaymentdate, u.Udcpaymenttype, u.Udcpaymenttime, u.Udcpaymentno, u.Udcpayvspcode, u.Udcpaybankcode, u.Udcpaymentvalid, u.Udcpayledsys, u.EquinoxPrn, u.EquinoxLrn, u.EquinoxSec)
	_, err = db.Exec(sqlstr, u.Udcpayment, u.Udcpaymentdate, u.Udcpaymenttype, u.Udcpaymenttime, u.Udcpaymentno, u.Udcpayvspcode, u.Udcpaybankcode, u.Udcpaymentvalid, u.Udcpayledsys, u.EquinoxPrn, u.EquinoxLrn, u.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	u._exists = true

	return nil
}

// Delete deletes the Udcpay from the database.
func (u *Udcpay) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return nil
	}

	// if deleted, bail
	if u._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.udcpay WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, u.EquinoxLrn)
	_, err = db.Exec(sqlstr, u.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

// UdcpayByEquinoxLrn retrieves a row from 'equinox.udcpay' as a Udcpay.
//
// Generated from index 'udcpay_pkey'.
func UdcpayByEquinoxLrn(db XODB, equinoxLrn int64) (*Udcpay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`udcpayment, udcpaymentdate, udcpaymenttype, udcpaymenttime, udcpaymentno, udcpayvspcode, udcpaybankcode, udcpaymentvalid, udcpayledsys, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.udcpay ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	u := Udcpay{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&u.Udcpayment, &u.Udcpaymentdate, &u.Udcpaymenttype, &u.Udcpaymenttime, &u.Udcpaymentno, &u.Udcpayvspcode, &u.Udcpaybankcode, &u.Udcpaymentvalid, &u.Udcpayledsys, &u.EquinoxPrn, &u.EquinoxLrn, &u.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
