// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Eventlog represents a row from 'equinox.eventlog'.
type Eventlog struct {
	Elid          sql.NullInt64  `json:"elid"`          // elid
	Eleventid     sql.NullInt64  `json:"eleventid"`     // eleventid
	Eldate        pq.NullTime    `json:"eldate"`        // eldate
	Eltime        pq.NullTime    `json:"eltime"`        // eltime
	Elstartdate   pq.NullTime    `json:"elstartdate"`   // elstartdate
	Elstarttime   pq.NullTime    `json:"elstarttime"`   // elstarttime
	Elenddate     pq.NullTime    `json:"elenddate"`     // elenddate
	Elendtime     pq.NullTime    `json:"elendtime"`     // elendtime
	Elmessage     sql.NullString `json:"elmessage"`     // elmessage
	Eldetails     sql.NullInt64  `json:"eldetails"`     // eldetails
	Eltype        sql.NullString `json:"eltype"`        // eltype
	Elcategory    sql.NullString `json:"elcategory"`    // elcategory
	Elraisedby    sql.NullString `json:"elraisedby"`    // elraisedby
	Eladditional1 sql.NullString `json:"eladditional1"` // eladditional1
	Eladditional2 sql.NullString `json:"eladditional2"` // eladditional2
	Eladditional3 sql.NullString `json:"eladditional3"` // eladditional3
	Eladditional4 sql.NullString `json:"eladditional4"` // eladditional4
	EquinoxLrn    int64          `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec    sql.NullInt64  `json:"equinox_sec"`   // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Eventlog exists in the database.
func (e *Eventlog) Exists() bool {
	return e._exists
}

// Deleted provides information if the Eventlog has been deleted from the database.
func (e *Eventlog) Deleted() bool {
	return e._deleted
}

// Insert inserts the Eventlog to the database.
func (e *Eventlog) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.eventlog (` +
		`elid, eleventid, eldate, eltime, elstartdate, elstarttime, elenddate, elendtime, elmessage, eldetails, eltype, elcategory, elraisedby, eladditional1, eladditional2, eladditional3, eladditional4, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, e.Elid, e.Eleventid, e.Eldate, e.Eltime, e.Elstartdate, e.Elstarttime, e.Elenddate, e.Elendtime, e.Elmessage, e.Eldetails, e.Eltype, e.Elcategory, e.Elraisedby, e.Eladditional1, e.Eladditional2, e.Eladditional3, e.Eladditional4, e.EquinoxSec)
	err = db.QueryRow(sqlstr, e.Elid, e.Eleventid, e.Eldate, e.Eltime, e.Elstartdate, e.Elstarttime, e.Elenddate, e.Elendtime, e.Elmessage, e.Eldetails, e.Eltype, e.Elcategory, e.Elraisedby, e.Eladditional1, e.Eladditional2, e.Eladditional3, e.Eladditional4, e.EquinoxSec).Scan(&e.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Update updates the Eventlog in the database.
func (e *Eventlog) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.eventlog SET (` +
		`elid, eleventid, eldate, eltime, elstartdate, elstarttime, elenddate, elendtime, elmessage, eldetails, eltype, elcategory, elraisedby, eladditional1, eladditional2, eladditional3, eladditional4, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18` +
		`) WHERE equinox_lrn = $19`

	// run query
	XOLog(sqlstr, e.Elid, e.Eleventid, e.Eldate, e.Eltime, e.Elstartdate, e.Elstarttime, e.Elenddate, e.Elendtime, e.Elmessage, e.Eldetails, e.Eltype, e.Elcategory, e.Elraisedby, e.Eladditional1, e.Eladditional2, e.Eladditional3, e.Eladditional4, e.EquinoxSec, e.EquinoxLrn)
	_, err = db.Exec(sqlstr, e.Elid, e.Eleventid, e.Eldate, e.Eltime, e.Elstartdate, e.Elstarttime, e.Elenddate, e.Elendtime, e.Elmessage, e.Eldetails, e.Eltype, e.Elcategory, e.Elraisedby, e.Eladditional1, e.Eladditional2, e.Eladditional3, e.Eladditional4, e.EquinoxSec, e.EquinoxLrn)
	return err
}

// Save saves the Eventlog to the database.
func (e *Eventlog) Save(db XODB) error {
	if e.Exists() {
		return e.Update(db)
	}

	return e.Insert(db)
}

// Upsert performs an upsert for Eventlog.
//
// NOTE: PostgreSQL 9.5+ only
func (e *Eventlog) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.eventlog (` +
		`elid, eleventid, eldate, eltime, elstartdate, elstarttime, elenddate, elendtime, elmessage, eldetails, eltype, elcategory, elraisedby, eladditional1, eladditional2, eladditional3, eladditional4, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`elid, eleventid, eldate, eltime, elstartdate, elstarttime, elenddate, elendtime, elmessage, eldetails, eltype, elcategory, elraisedby, eladditional1, eladditional2, eladditional3, eladditional4, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.elid, EXCLUDED.eleventid, EXCLUDED.eldate, EXCLUDED.eltime, EXCLUDED.elstartdate, EXCLUDED.elstarttime, EXCLUDED.elenddate, EXCLUDED.elendtime, EXCLUDED.elmessage, EXCLUDED.eldetails, EXCLUDED.eltype, EXCLUDED.elcategory, EXCLUDED.elraisedby, EXCLUDED.eladditional1, EXCLUDED.eladditional2, EXCLUDED.eladditional3, EXCLUDED.eladditional4, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, e.Elid, e.Eleventid, e.Eldate, e.Eltime, e.Elstartdate, e.Elstarttime, e.Elenddate, e.Elendtime, e.Elmessage, e.Eldetails, e.Eltype, e.Elcategory, e.Elraisedby, e.Eladditional1, e.Eladditional2, e.Eladditional3, e.Eladditional4, e.EquinoxLrn, e.EquinoxSec)
	_, err = db.Exec(sqlstr, e.Elid, e.Eleventid, e.Eldate, e.Eltime, e.Elstartdate, e.Elstarttime, e.Elenddate, e.Elendtime, e.Elmessage, e.Eldetails, e.Eltype, e.Elcategory, e.Elraisedby, e.Eladditional1, e.Eladditional2, e.Eladditional3, e.Eladditional4, e.EquinoxLrn, e.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Delete deletes the Eventlog from the database.
func (e *Eventlog) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.eventlog WHERE equinox_lrn = $1`

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

// EventlogByEquinoxLrn retrieves a row from 'equinox.eventlog' as a Eventlog.
//
// Generated from index 'eventlog_pkey'.
func EventlogByEquinoxLrn(db XODB, equinoxLrn int64) (*Eventlog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`elid, eleventid, eldate, eltime, elstartdate, elstarttime, elenddate, elendtime, elmessage, eldetails, eltype, elcategory, elraisedby, eladditional1, eladditional2, eladditional3, eladditional4, equinox_lrn, equinox_sec ` +
		`FROM equinox.eventlog ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Eventlog{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Elid, &e.Eleventid, &e.Eldate, &e.Eltime, &e.Elstartdate, &e.Elstarttime, &e.Elenddate, &e.Elendtime, &e.Elmessage, &e.Eldetails, &e.Eltype, &e.Elcategory, &e.Elraisedby, &e.Eladditional1, &e.Eladditional2, &e.Eladditional3, &e.Eladditional4, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
