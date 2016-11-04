// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Attendee represents a row from 'equinox.attendee'.
type Attendee struct {
	AttSparec1     sql.NullString `json:"att_sparec1"`     // att_sparec1
	AttExid        sql.NullString `json:"att_exid"`        // att_exid
	AttSpaces      sql.NullInt64  `json:"att_spaces"`      // att_spaces
	AttSparec2     sql.NullString `json:"att_sparec2"`     // att_sparec2
	AttNotifyby    sql.NullString `json:"att_notifyby"`    // att_notifyby
	AttNotified    sql.NullInt64  `json:"att_notified"`    // att_notified
	AttNoshow      sql.NullInt64  `json:"att_noshow"`      // att_noshow
	AttLevel       sql.NullString `json:"att_level"`       // att_level
	AttCompleted   sql.NullInt64  `json:"att_completed"`   // att_completed
	AttCertificate pq.NullTime    `json:"att_certificate"` // att_certificate
	AttOldexec     sql.NullInt64  `json:"att_oldexec"`     // att_oldexec
	AttNewexec     sql.NullInt64  `json:"att_newexec"`     // att_newexec
	AttExname      sql.NullString `json:"att_exname"`      // att_exname
	AttBookedby    sql.NullString `json:"att_bookedby"`    // att_bookedby
	AttBookingno   sql.NullInt64  `json:"att_bookingno"`   // att_bookingno
	EquinoxPrn     sql.NullInt64  `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn     int64          `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec     sql.NullInt64  `json:"equinox_sec"`     // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Attendee exists in the database.
func (a *Attendee) Exists() bool {
	return a._exists
}

// Deleted provides information if the Attendee has been deleted from the database.
func (a *Attendee) Deleted() bool {
	return a._deleted
}

// Insert inserts the Attendee to the database.
func (a *Attendee) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.attendee (` +
		`att_sparec1, att_exid, att_spaces, att_sparec2, att_notifyby, att_notified, att_noshow, att_level, att_completed, att_certificate, att_oldexec, att_newexec, att_exname, att_bookedby, att_bookingno, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, a.AttSparec1, a.AttExid, a.AttSpaces, a.AttSparec2, a.AttNotifyby, a.AttNotified, a.AttNoshow, a.AttLevel, a.AttCompleted, a.AttCertificate, a.AttOldexec, a.AttNewexec, a.AttExname, a.AttBookedby, a.AttBookingno, a.EquinoxPrn, a.EquinoxSec)
	err = db.QueryRow(sqlstr, a.AttSparec1, a.AttExid, a.AttSpaces, a.AttSparec2, a.AttNotifyby, a.AttNotified, a.AttNoshow, a.AttLevel, a.AttCompleted, a.AttCertificate, a.AttOldexec, a.AttNewexec, a.AttExname, a.AttBookedby, a.AttBookingno, a.EquinoxPrn, a.EquinoxSec).Scan(&a.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Update updates the Attendee in the database.
func (a *Attendee) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.attendee SET (` +
		`att_sparec1, att_exid, att_spaces, att_sparec2, att_notifyby, att_notified, att_noshow, att_level, att_completed, att_certificate, att_oldexec, att_newexec, att_exname, att_bookedby, att_bookingno, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17` +
		`) WHERE equinox_lrn = $18`

	// run query
	XOLog(sqlstr, a.AttSparec1, a.AttExid, a.AttSpaces, a.AttSparec2, a.AttNotifyby, a.AttNotified, a.AttNoshow, a.AttLevel, a.AttCompleted, a.AttCertificate, a.AttOldexec, a.AttNewexec, a.AttExname, a.AttBookedby, a.AttBookingno, a.EquinoxPrn, a.EquinoxSec, a.EquinoxLrn)
	_, err = db.Exec(sqlstr, a.AttSparec1, a.AttExid, a.AttSpaces, a.AttSparec2, a.AttNotifyby, a.AttNotified, a.AttNoshow, a.AttLevel, a.AttCompleted, a.AttCertificate, a.AttOldexec, a.AttNewexec, a.AttExname, a.AttBookedby, a.AttBookingno, a.EquinoxPrn, a.EquinoxSec, a.EquinoxLrn)
	return err
}

// Save saves the Attendee to the database.
func (a *Attendee) Save(db XODB) error {
	if a.Exists() {
		return a.Update(db)
	}

	return a.Insert(db)
}

// Upsert performs an upsert for Attendee.
//
// NOTE: PostgreSQL 9.5+ only
func (a *Attendee) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.attendee (` +
		`att_sparec1, att_exid, att_spaces, att_sparec2, att_notifyby, att_notified, att_noshow, att_level, att_completed, att_certificate, att_oldexec, att_newexec, att_exname, att_bookedby, att_bookingno, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`att_sparec1, att_exid, att_spaces, att_sparec2, att_notifyby, att_notified, att_noshow, att_level, att_completed, att_certificate, att_oldexec, att_newexec, att_exname, att_bookedby, att_bookingno, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.att_sparec1, EXCLUDED.att_exid, EXCLUDED.att_spaces, EXCLUDED.att_sparec2, EXCLUDED.att_notifyby, EXCLUDED.att_notified, EXCLUDED.att_noshow, EXCLUDED.att_level, EXCLUDED.att_completed, EXCLUDED.att_certificate, EXCLUDED.att_oldexec, EXCLUDED.att_newexec, EXCLUDED.att_exname, EXCLUDED.att_bookedby, EXCLUDED.att_bookingno, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, a.AttSparec1, a.AttExid, a.AttSpaces, a.AttSparec2, a.AttNotifyby, a.AttNotified, a.AttNoshow, a.AttLevel, a.AttCompleted, a.AttCertificate, a.AttOldexec, a.AttNewexec, a.AttExname, a.AttBookedby, a.AttBookingno, a.EquinoxPrn, a.EquinoxLrn, a.EquinoxSec)
	_, err = db.Exec(sqlstr, a.AttSparec1, a.AttExid, a.AttSpaces, a.AttSparec2, a.AttNotifyby, a.AttNotified, a.AttNoshow, a.AttLevel, a.AttCompleted, a.AttCertificate, a.AttOldexec, a.AttNewexec, a.AttExname, a.AttBookedby, a.AttBookingno, a.EquinoxPrn, a.EquinoxLrn, a.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Delete deletes the Attendee from the database.
func (a *Attendee) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.attendee WHERE equinox_lrn = $1`

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

// AttendeeByEquinoxLrn retrieves a row from 'equinox.attendee' as a Attendee.
//
// Generated from index 'attendee_pkey'.
func AttendeeByEquinoxLrn(db XODB, equinoxLrn int64) (*Attendee, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`att_sparec1, att_exid, att_spaces, att_sparec2, att_notifyby, att_notified, att_noshow, att_level, att_completed, att_certificate, att_oldexec, att_newexec, att_exname, att_bookedby, att_bookingno, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.attendee ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	a := Attendee{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&a.AttSparec1, &a.AttExid, &a.AttSpaces, &a.AttSparec2, &a.AttNotifyby, &a.AttNotified, &a.AttNoshow, &a.AttLevel, &a.AttCompleted, &a.AttCertificate, &a.AttOldexec, &a.AttNewexec, &a.AttExname, &a.AttBookedby, &a.AttBookingno, &a.EquinoxPrn, &a.EquinoxLrn, &a.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &a, nil
}