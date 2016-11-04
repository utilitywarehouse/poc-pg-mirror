// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Edsmain represents a row from 'equinox.edsmain'.
type Edsmain struct {
	Edscustaccountno sql.NullString  `json:"edscustaccountno"` // edscustaccountno
	Edsimportbalance sql.NullFloat64 `json:"edsimportbalance"` // edsimportbalance
	Edscustname      sql.NullString  `json:"edscustname"`      // edscustname
	Edsbillingadd1   sql.NullString  `json:"edsbillingadd1"`   // edsbillingadd1
	Edsbillingadd2   sql.NullString  `json:"edsbillingadd2"`   // edsbillingadd2
	Edsbillingadd3   sql.NullString  `json:"edsbillingadd3"`   // edsbillingadd3
	Edsbillingadd4   sql.NullString  `json:"edsbillingadd4"`   // edsbillingadd4
	Edsbillingcounty sql.NullString  `json:"edsbillingcounty"` // edsbillingcounty
	Edsbillingpostco sql.NullString  `json:"edsbillingpostco"` // edsbillingpostco
	Edsdefault       pq.NullTime     `json:"edsdefault"`       // edsdefault
	Edsdefault2      pq.NullTime     `json:"edsdefault2"`      // edsdefault2
	Edsfueldirect    sql.NullInt64   `json:"edsfueldirect"`    // edsfueldirect
	Edsfdweekpay     sql.NullFloat64 `json:"edsfdweekpay"`     // edsfdweekpay
	Edsagency        sql.NullString  `json:"edsagency"`        // edsagency
	Edsppagreed      pq.NullTime     `json:"edsppagreed"`      // edsppagreed
	Edsppamount      sql.NullFloat64 `json:"edsppamount"`      // edsppamount
	Edsppagreed2     pq.NullTime     `json:"edsppagreed2"`     // edsppagreed2
	Edsppamount2     sql.NullFloat64 `json:"edsppamount2"`     // edsppamount2
	Edsdeposit       sql.NullFloat64 `json:"edsdeposit"`       // edsdeposit
	Edswalkaway      sql.NullString  `json:"edswalkaway"`      // edswalkaway
	Edsbalancepaid   sql.NullFloat64 `json:"edsbalancepaid"`   // edsbalancepaid
	Edscontactphone  sql.NullString  `json:"edscontactphone"`  // edscontactphone
	Edscontactemail  sql.NullString  `json:"edscontactemail"`  // edscontactemail
	Edscontactfax    sql.NullString  `json:"edscontactfax"`    // edscontactfax
	Edscontactmobile sql.NullString  `json:"edscontactmobile"` // edscontactmobile
	Edsinvtitle      sql.NullString  `json:"edsinvtitle"`      // edsinvtitle
	Edsinvfirstname  sql.NullString  `json:"edsinvfirstname"`  // edsinvfirstname
	Edsinvsurname    sql.NullString  `json:"edsinvsurname"`    // edsinvsurname
	Edsinvinitials   sql.NullString  `json:"edsinvinitials"`   // edsinvinitials
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Edsmain exists in the database.
func (e *Edsmain) Exists() bool {
	return e._exists
}

// Deleted provides information if the Edsmain has been deleted from the database.
func (e *Edsmain) Deleted() bool {
	return e._deleted
}

// Insert inserts the Edsmain to the database.
func (e *Edsmain) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.edsmain (` +
		`edscustaccountno, edsimportbalance, edscustname, edsbillingadd1, edsbillingadd2, edsbillingadd3, edsbillingadd4, edsbillingcounty, edsbillingpostco, edsdefault, edsdefault2, edsfueldirect, edsfdweekpay, edsagency, edsppagreed, edsppamount, edsppagreed2, edsppamount2, edsdeposit, edswalkaway, edsbalancepaid, edscontactphone, edscontactemail, edscontactfax, edscontactmobile, edsinvtitle, edsinvfirstname, edsinvsurname, edsinvinitials, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, e.Edscustaccountno, e.Edsimportbalance, e.Edscustname, e.Edsbillingadd1, e.Edsbillingadd2, e.Edsbillingadd3, e.Edsbillingadd4, e.Edsbillingcounty, e.Edsbillingpostco, e.Edsdefault, e.Edsdefault2, e.Edsfueldirect, e.Edsfdweekpay, e.Edsagency, e.Edsppagreed, e.Edsppamount, e.Edsppagreed2, e.Edsppamount2, e.Edsdeposit, e.Edswalkaway, e.Edsbalancepaid, e.Edscontactphone, e.Edscontactemail, e.Edscontactfax, e.Edscontactmobile, e.Edsinvtitle, e.Edsinvfirstname, e.Edsinvsurname, e.Edsinvinitials, e.EquinoxSec)
	err = db.QueryRow(sqlstr, e.Edscustaccountno, e.Edsimportbalance, e.Edscustname, e.Edsbillingadd1, e.Edsbillingadd2, e.Edsbillingadd3, e.Edsbillingadd4, e.Edsbillingcounty, e.Edsbillingpostco, e.Edsdefault, e.Edsdefault2, e.Edsfueldirect, e.Edsfdweekpay, e.Edsagency, e.Edsppagreed, e.Edsppamount, e.Edsppagreed2, e.Edsppamount2, e.Edsdeposit, e.Edswalkaway, e.Edsbalancepaid, e.Edscontactphone, e.Edscontactemail, e.Edscontactfax, e.Edscontactmobile, e.Edsinvtitle, e.Edsinvfirstname, e.Edsinvsurname, e.Edsinvinitials, e.EquinoxSec).Scan(&e.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Update updates the Edsmain in the database.
func (e *Edsmain) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if e._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.edsmain SET (` +
		`edscustaccountno, edsimportbalance, edscustname, edsbillingadd1, edsbillingadd2, edsbillingadd3, edsbillingadd4, edsbillingcounty, edsbillingpostco, edsdefault, edsdefault2, edsfueldirect, edsfdweekpay, edsagency, edsppagreed, edsppamount, edsppagreed2, edsppamount2, edsdeposit, edswalkaway, edsbalancepaid, edscontactphone, edscontactemail, edscontactfax, edscontactmobile, edsinvtitle, edsinvfirstname, edsinvsurname, edsinvinitials, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30` +
		`) WHERE equinox_lrn = $31`

	// run query
	XOLog(sqlstr, e.Edscustaccountno, e.Edsimportbalance, e.Edscustname, e.Edsbillingadd1, e.Edsbillingadd2, e.Edsbillingadd3, e.Edsbillingadd4, e.Edsbillingcounty, e.Edsbillingpostco, e.Edsdefault, e.Edsdefault2, e.Edsfueldirect, e.Edsfdweekpay, e.Edsagency, e.Edsppagreed, e.Edsppamount, e.Edsppagreed2, e.Edsppamount2, e.Edsdeposit, e.Edswalkaway, e.Edsbalancepaid, e.Edscontactphone, e.Edscontactemail, e.Edscontactfax, e.Edscontactmobile, e.Edsinvtitle, e.Edsinvfirstname, e.Edsinvsurname, e.Edsinvinitials, e.EquinoxSec, e.EquinoxLrn)
	_, err = db.Exec(sqlstr, e.Edscustaccountno, e.Edsimportbalance, e.Edscustname, e.Edsbillingadd1, e.Edsbillingadd2, e.Edsbillingadd3, e.Edsbillingadd4, e.Edsbillingcounty, e.Edsbillingpostco, e.Edsdefault, e.Edsdefault2, e.Edsfueldirect, e.Edsfdweekpay, e.Edsagency, e.Edsppagreed, e.Edsppamount, e.Edsppagreed2, e.Edsppamount2, e.Edsdeposit, e.Edswalkaway, e.Edsbalancepaid, e.Edscontactphone, e.Edscontactemail, e.Edscontactfax, e.Edscontactmobile, e.Edsinvtitle, e.Edsinvfirstname, e.Edsinvsurname, e.Edsinvinitials, e.EquinoxSec, e.EquinoxLrn)
	return err
}

// Save saves the Edsmain to the database.
func (e *Edsmain) Save(db XODB) error {
	if e.Exists() {
		return e.Update(db)
	}

	return e.Insert(db)
}

// Upsert performs an upsert for Edsmain.
//
// NOTE: PostgreSQL 9.5+ only
func (e *Edsmain) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.edsmain (` +
		`edscustaccountno, edsimportbalance, edscustname, edsbillingadd1, edsbillingadd2, edsbillingadd3, edsbillingadd4, edsbillingcounty, edsbillingpostco, edsdefault, edsdefault2, edsfueldirect, edsfdweekpay, edsagency, edsppagreed, edsppamount, edsppagreed2, edsppamount2, edsdeposit, edswalkaway, edsbalancepaid, edscontactphone, edscontactemail, edscontactfax, edscontactmobile, edsinvtitle, edsinvfirstname, edsinvsurname, edsinvinitials, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`edscustaccountno, edsimportbalance, edscustname, edsbillingadd1, edsbillingadd2, edsbillingadd3, edsbillingadd4, edsbillingcounty, edsbillingpostco, edsdefault, edsdefault2, edsfueldirect, edsfdweekpay, edsagency, edsppagreed, edsppamount, edsppagreed2, edsppamount2, edsdeposit, edswalkaway, edsbalancepaid, edscontactphone, edscontactemail, edscontactfax, edscontactmobile, edsinvtitle, edsinvfirstname, edsinvsurname, edsinvinitials, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.edscustaccountno, EXCLUDED.edsimportbalance, EXCLUDED.edscustname, EXCLUDED.edsbillingadd1, EXCLUDED.edsbillingadd2, EXCLUDED.edsbillingadd3, EXCLUDED.edsbillingadd4, EXCLUDED.edsbillingcounty, EXCLUDED.edsbillingpostco, EXCLUDED.edsdefault, EXCLUDED.edsdefault2, EXCLUDED.edsfueldirect, EXCLUDED.edsfdweekpay, EXCLUDED.edsagency, EXCLUDED.edsppagreed, EXCLUDED.edsppamount, EXCLUDED.edsppagreed2, EXCLUDED.edsppamount2, EXCLUDED.edsdeposit, EXCLUDED.edswalkaway, EXCLUDED.edsbalancepaid, EXCLUDED.edscontactphone, EXCLUDED.edscontactemail, EXCLUDED.edscontactfax, EXCLUDED.edscontactmobile, EXCLUDED.edsinvtitle, EXCLUDED.edsinvfirstname, EXCLUDED.edsinvsurname, EXCLUDED.edsinvinitials, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, e.Edscustaccountno, e.Edsimportbalance, e.Edscustname, e.Edsbillingadd1, e.Edsbillingadd2, e.Edsbillingadd3, e.Edsbillingadd4, e.Edsbillingcounty, e.Edsbillingpostco, e.Edsdefault, e.Edsdefault2, e.Edsfueldirect, e.Edsfdweekpay, e.Edsagency, e.Edsppagreed, e.Edsppamount, e.Edsppagreed2, e.Edsppamount2, e.Edsdeposit, e.Edswalkaway, e.Edsbalancepaid, e.Edscontactphone, e.Edscontactemail, e.Edscontactfax, e.Edscontactmobile, e.Edsinvtitle, e.Edsinvfirstname, e.Edsinvsurname, e.Edsinvinitials, e.EquinoxLrn, e.EquinoxSec)
	_, err = db.Exec(sqlstr, e.Edscustaccountno, e.Edsimportbalance, e.Edscustname, e.Edsbillingadd1, e.Edsbillingadd2, e.Edsbillingadd3, e.Edsbillingadd4, e.Edsbillingcounty, e.Edsbillingpostco, e.Edsdefault, e.Edsdefault2, e.Edsfueldirect, e.Edsfdweekpay, e.Edsagency, e.Edsppagreed, e.Edsppamount, e.Edsppagreed2, e.Edsppamount2, e.Edsdeposit, e.Edswalkaway, e.Edsbalancepaid, e.Edscontactphone, e.Edscontactemail, e.Edscontactfax, e.Edscontactmobile, e.Edsinvtitle, e.Edsinvfirstname, e.Edsinvsurname, e.Edsinvinitials, e.EquinoxLrn, e.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Delete deletes the Edsmain from the database.
func (e *Edsmain) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return nil
	}

	// if deleted, bail
	if e._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.edsmain WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, e.EquinoxLrn)
	_, err = db.Exec(sqlstr, e.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	e._deleted = true

	return nil
}

// EdsmainByEquinoxLrn retrieves a row from 'equinox.edsmain' as a Edsmain.
//
// Generated from index 'edsmain_pkey'.
func EdsmainByEquinoxLrn(db XODB, equinoxLrn int64) (*Edsmain, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`edscustaccountno, edsimportbalance, edscustname, edsbillingadd1, edsbillingadd2, edsbillingadd3, edsbillingadd4, edsbillingcounty, edsbillingpostco, edsdefault, edsdefault2, edsfueldirect, edsfdweekpay, edsagency, edsppagreed, edsppamount, edsppagreed2, edsppamount2, edsdeposit, edswalkaway, edsbalancepaid, edscontactphone, edscontactemail, edscontactfax, edscontactmobile, edsinvtitle, edsinvfirstname, edsinvsurname, edsinvinitials, equinox_lrn, equinox_sec ` +
		`FROM equinox.edsmain ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Edsmain{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Edscustaccountno, &e.Edsimportbalance, &e.Edscustname, &e.Edsbillingadd1, &e.Edsbillingadd2, &e.Edsbillingadd3, &e.Edsbillingadd4, &e.Edsbillingcounty, &e.Edsbillingpostco, &e.Edsdefault, &e.Edsdefault2, &e.Edsfueldirect, &e.Edsfdweekpay, &e.Edsagency, &e.Edsppagreed, &e.Edsppamount, &e.Edsppagreed2, &e.Edsppamount2, &e.Edsdeposit, &e.Edswalkaway, &e.Edsbalancepaid, &e.Edscontactphone, &e.Edscontactemail, &e.Edscontactfax, &e.Edscontactmobile, &e.Edsinvtitle, &e.Edsinvfirstname, &e.Edsinvsurname, &e.Edsinvinitials, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
