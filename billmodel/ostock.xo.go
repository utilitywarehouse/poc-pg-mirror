// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Ostock represents a row from 'equinox.ostock'.
type Ostock struct {
	Ostockimei      sql.NullString `json:"ostockimei"`      // ostockimei
	Ostockmaccode   sql.NullString `json:"ostockmaccode"`   // ostockmaccode
	Ostockdatein    pq.NullTime    `json:"ostockdatein"`    // ostockdatein
	Ostocktimein    pq.NullTime    `json:"ostocktimein"`    // ostocktimein
	Ostockpackageno sql.NullString `json:"ostockpackageno"` // ostockpackageno
	Ostockstatus    sql.NullString `json:"ostockstatus"`    // ostockstatus
	Ostocklocation  sql.NullString `json:"ostocklocation"`  // ostocklocation
	Ostockdateout   pq.NullTime    `json:"ostockdateout"`   // ostockdateout
	Ostocktimeout   pq.NullTime    `json:"ostocktimeout"`   // ostocktimeout
	Ostockaccountno sql.NullString `json:"ostockaccountno"` // ostockaccountno
	EquinoxPrn      sql.NullInt64  `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn      int64          `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`     // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Ostock exists in the database.
func (o *Ostock) Exists() bool {
	return o._exists
}

// Deleted provides information if the Ostock has been deleted from the database.
func (o *Ostock) Deleted() bool {
	return o._deleted
}

// Insert inserts the Ostock to the database.
func (o *Ostock) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if o._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.ostock (` +
		`ostockimei, ostockmaccode, ostockdatein, ostocktimein, ostockpackageno, ostockstatus, ostocklocation, ostockdateout, ostocktimeout, ostockaccountno, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, o.Ostockimei, o.Ostockmaccode, o.Ostockdatein, o.Ostocktimein, o.Ostockpackageno, o.Ostockstatus, o.Ostocklocation, o.Ostockdateout, o.Ostocktimeout, o.Ostockaccountno, o.EquinoxPrn, o.EquinoxSec)
	err = db.QueryRow(sqlstr, o.Ostockimei, o.Ostockmaccode, o.Ostockdatein, o.Ostocktimein, o.Ostockpackageno, o.Ostockstatus, o.Ostocklocation, o.Ostockdateout, o.Ostocktimeout, o.Ostockaccountno, o.EquinoxPrn, o.EquinoxSec).Scan(&o.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	o._exists = true

	return nil
}

// Update updates the Ostock in the database.
func (o *Ostock) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !o._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if o._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.ostock SET (` +
		`ostockimei, ostockmaccode, ostockdatein, ostocktimein, ostockpackageno, ostockstatus, ostocklocation, ostockdateout, ostocktimeout, ostockaccountno, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) WHERE equinox_lrn = $13`

	// run query
	XOLog(sqlstr, o.Ostockimei, o.Ostockmaccode, o.Ostockdatein, o.Ostocktimein, o.Ostockpackageno, o.Ostockstatus, o.Ostocklocation, o.Ostockdateout, o.Ostocktimeout, o.Ostockaccountno, o.EquinoxPrn, o.EquinoxSec, o.EquinoxLrn)
	_, err = db.Exec(sqlstr, o.Ostockimei, o.Ostockmaccode, o.Ostockdatein, o.Ostocktimein, o.Ostockpackageno, o.Ostockstatus, o.Ostocklocation, o.Ostockdateout, o.Ostocktimeout, o.Ostockaccountno, o.EquinoxPrn, o.EquinoxSec, o.EquinoxLrn)
	return err
}

// Save saves the Ostock to the database.
func (o *Ostock) Save(db XODB) error {
	if o.Exists() {
		return o.Update(db)
	}

	return o.Insert(db)
}

// Upsert performs an upsert for Ostock.
//
// NOTE: PostgreSQL 9.5+ only
func (o *Ostock) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if o._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.ostock (` +
		`ostockimei, ostockmaccode, ostockdatein, ostocktimein, ostockpackageno, ostockstatus, ostocklocation, ostockdateout, ostocktimeout, ostockaccountno, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`ostockimei, ostockmaccode, ostockdatein, ostocktimein, ostockpackageno, ostockstatus, ostocklocation, ostockdateout, ostocktimeout, ostockaccountno, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.ostockimei, EXCLUDED.ostockmaccode, EXCLUDED.ostockdatein, EXCLUDED.ostocktimein, EXCLUDED.ostockpackageno, EXCLUDED.ostockstatus, EXCLUDED.ostocklocation, EXCLUDED.ostockdateout, EXCLUDED.ostocktimeout, EXCLUDED.ostockaccountno, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, o.Ostockimei, o.Ostockmaccode, o.Ostockdatein, o.Ostocktimein, o.Ostockpackageno, o.Ostockstatus, o.Ostocklocation, o.Ostockdateout, o.Ostocktimeout, o.Ostockaccountno, o.EquinoxPrn, o.EquinoxLrn, o.EquinoxSec)
	_, err = db.Exec(sqlstr, o.Ostockimei, o.Ostockmaccode, o.Ostockdatein, o.Ostocktimein, o.Ostockpackageno, o.Ostockstatus, o.Ostocklocation, o.Ostockdateout, o.Ostocktimeout, o.Ostockaccountno, o.EquinoxPrn, o.EquinoxLrn, o.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	o._exists = true

	return nil
}

// Delete deletes the Ostock from the database.
func (o *Ostock) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !o._exists {
		return nil
	}

	// if deleted, bail
	if o._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.ostock WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, o.EquinoxLrn)
	_, err = db.Exec(sqlstr, o.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	o._deleted = true

	return nil
}

// OstockByEquinoxLrn retrieves a row from 'equinox.ostock' as a Ostock.
//
// Generated from index 'ostock_pkey'.
func OstockByEquinoxLrn(db XODB, equinoxLrn int64) (*Ostock, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ostockimei, ostockmaccode, ostockdatein, ostocktimein, ostockpackageno, ostockstatus, ostocklocation, ostockdateout, ostocktimeout, ostockaccountno, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.ostock ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	o := Ostock{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&o.Ostockimei, &o.Ostockmaccode, &o.Ostockdatein, &o.Ostocktimein, &o.Ostockpackageno, &o.Ostockstatus, &o.Ostocklocation, &o.Ostockdateout, &o.Ostocktimeout, &o.Ostockaccountno, &o.EquinoxPrn, &o.EquinoxLrn, &o.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &o, nil
}