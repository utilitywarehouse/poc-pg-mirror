// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Evlog represents a row from 'equinox.evlog'.
type Evlog struct {
	Eventlogid     sql.NullInt64  `json:"eventlogid"`     // eventlogid
	Evlchildid     sql.NullInt64  `json:"evlchildid"`     // evlchildid
	Evldate        pq.NullTime    `json:"evldate"`        // evldate
	Evltime        pq.NullTime    `json:"evltime"`        // evltime
	Evlstartdate   pq.NullTime    `json:"evlstartdate"`   // evlstartdate
	Evlstarttime   pq.NullTime    `json:"evlstarttime"`   // evlstarttime
	Evlenddate     pq.NullTime    `json:"evlenddate"`     // evlenddate
	Evlendtime     pq.NullTime    `json:"evlendtime"`     // evlendtime
	Evlmessage     sql.NullString `json:"evlmessage"`     // evlmessage
	Evldetails     sql.NullInt64  `json:"evldetails"`     // evldetails
	Evltype        sql.NullString `json:"evltype"`        // evltype
	Evlcategory    sql.NullString `json:"evlcategory"`    // evlcategory
	Evlraisedby    sql.NullString `json:"evlraisedby"`    // evlraisedby
	Evlmachine     sql.NullString `json:"evlmachine"`     // evlmachine
	Evladditional1 sql.NullString `json:"evladditional1"` // evladditional1
	Evladditional2 sql.NullString `json:"evladditional2"` // evladditional2
	Evladditional3 sql.NullString `json:"evladditional3"` // evladditional3
	Evladditional4 sql.NullString `json:"evladditional4"` // evladditional4
	EquinoxPrn     sql.NullInt64  `json:"equinox_prn"`    // equinox_prn
	EquinoxLrn     int64          `json:"equinox_lrn"`    // equinox_lrn
	EquinoxSec     sql.NullInt64  `json:"equinox_sec"`    // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Evlog exists in the database.
func (e *Evlog) Exists() bool {
	return e._exists
}

// Deleted provides information if the Evlog has been deleted from the database.
func (e *Evlog) Deleted() bool {
	return e._deleted
}

// Insert inserts the Evlog to the database.
func (e *Evlog) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.evlog (` +
		`eventlogid, evlchildid, evldate, evltime, evlstartdate, evlstarttime, evlenddate, evlendtime, evlmessage, evldetails, evltype, evlcategory, evlraisedby, evlmachine, evladditional1, evladditional2, evladditional3, evladditional4, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, e.Eventlogid, e.Evlchildid, e.Evldate, e.Evltime, e.Evlstartdate, e.Evlstarttime, e.Evlenddate, e.Evlendtime, e.Evlmessage, e.Evldetails, e.Evltype, e.Evlcategory, e.Evlraisedby, e.Evlmachine, e.Evladditional1, e.Evladditional2, e.Evladditional3, e.Evladditional4, e.EquinoxPrn, e.EquinoxSec)
	err = db.QueryRow(sqlstr, e.Eventlogid, e.Evlchildid, e.Evldate, e.Evltime, e.Evlstartdate, e.Evlstarttime, e.Evlenddate, e.Evlendtime, e.Evlmessage, e.Evldetails, e.Evltype, e.Evlcategory, e.Evlraisedby, e.Evlmachine, e.Evladditional1, e.Evladditional2, e.Evladditional3, e.Evladditional4, e.EquinoxPrn, e.EquinoxSec).Scan(&e.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Update updates the Evlog in the database.
func (e *Evlog) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.evlog SET (` +
		`eventlogid, evlchildid, evldate, evltime, evlstartdate, evlstarttime, evlenddate, evlendtime, evlmessage, evldetails, evltype, evlcategory, evlraisedby, evlmachine, evladditional1, evladditional2, evladditional3, evladditional4, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20` +
		`) WHERE equinox_lrn = $21`

	// run query
	XOLog(sqlstr, e.Eventlogid, e.Evlchildid, e.Evldate, e.Evltime, e.Evlstartdate, e.Evlstarttime, e.Evlenddate, e.Evlendtime, e.Evlmessage, e.Evldetails, e.Evltype, e.Evlcategory, e.Evlraisedby, e.Evlmachine, e.Evladditional1, e.Evladditional2, e.Evladditional3, e.Evladditional4, e.EquinoxPrn, e.EquinoxSec, e.EquinoxLrn)
	_, err = db.Exec(sqlstr, e.Eventlogid, e.Evlchildid, e.Evldate, e.Evltime, e.Evlstartdate, e.Evlstarttime, e.Evlenddate, e.Evlendtime, e.Evlmessage, e.Evldetails, e.Evltype, e.Evlcategory, e.Evlraisedby, e.Evlmachine, e.Evladditional1, e.Evladditional2, e.Evladditional3, e.Evladditional4, e.EquinoxPrn, e.EquinoxSec, e.EquinoxLrn)
	return err
}

// Save saves the Evlog to the database.
func (e *Evlog) Save(db XODB) error {
	if e.Exists() {
		return e.Update(db)
	}

	return e.Insert(db)
}

// Upsert performs an upsert for Evlog.
//
// NOTE: PostgreSQL 9.5+ only
func (e *Evlog) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.evlog (` +
		`eventlogid, evlchildid, evldate, evltime, evlstartdate, evlstarttime, evlenddate, evlendtime, evlmessage, evldetails, evltype, evlcategory, evlraisedby, evlmachine, evladditional1, evladditional2, evladditional3, evladditional4, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`eventlogid, evlchildid, evldate, evltime, evlstartdate, evlstarttime, evlenddate, evlendtime, evlmessage, evldetails, evltype, evlcategory, evlraisedby, evlmachine, evladditional1, evladditional2, evladditional3, evladditional4, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.eventlogid, EXCLUDED.evlchildid, EXCLUDED.evldate, EXCLUDED.evltime, EXCLUDED.evlstartdate, EXCLUDED.evlstarttime, EXCLUDED.evlenddate, EXCLUDED.evlendtime, EXCLUDED.evlmessage, EXCLUDED.evldetails, EXCLUDED.evltype, EXCLUDED.evlcategory, EXCLUDED.evlraisedby, EXCLUDED.evlmachine, EXCLUDED.evladditional1, EXCLUDED.evladditional2, EXCLUDED.evladditional3, EXCLUDED.evladditional4, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, e.Eventlogid, e.Evlchildid, e.Evldate, e.Evltime, e.Evlstartdate, e.Evlstarttime, e.Evlenddate, e.Evlendtime, e.Evlmessage, e.Evldetails, e.Evltype, e.Evlcategory, e.Evlraisedby, e.Evlmachine, e.Evladditional1, e.Evladditional2, e.Evladditional3, e.Evladditional4, e.EquinoxPrn, e.EquinoxLrn, e.EquinoxSec)
	_, err = db.Exec(sqlstr, e.Eventlogid, e.Evlchildid, e.Evldate, e.Evltime, e.Evlstartdate, e.Evlstarttime, e.Evlenddate, e.Evlendtime, e.Evlmessage, e.Evldetails, e.Evltype, e.Evlcategory, e.Evlraisedby, e.Evlmachine, e.Evladditional1, e.Evladditional2, e.Evladditional3, e.Evladditional4, e.EquinoxPrn, e.EquinoxLrn, e.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Delete deletes the Evlog from the database.
func (e *Evlog) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.evlog WHERE equinox_lrn = $1`

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

// EvlogByEquinoxLrn retrieves a row from 'equinox.evlog' as a Evlog.
//
// Generated from index 'evlog_pkey'.
func EvlogByEquinoxLrn(db XODB, equinoxLrn int64) (*Evlog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`eventlogid, evlchildid, evldate, evltime, evlstartdate, evlstarttime, evlenddate, evlendtime, evlmessage, evldetails, evltype, evlcategory, evlraisedby, evlmachine, evladditional1, evladditional2, evladditional3, evladditional4, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.evlog ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Evlog{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Eventlogid, &e.Evlchildid, &e.Evldate, &e.Evltime, &e.Evlstartdate, &e.Evlstarttime, &e.Evlenddate, &e.Evlendtime, &e.Evlmessage, &e.Evldetails, &e.Evltype, &e.Evlcategory, &e.Evlraisedby, &e.Evlmachine, &e.Evladditional1, &e.Evladditional2, &e.Evladditional3, &e.Evladditional4, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
