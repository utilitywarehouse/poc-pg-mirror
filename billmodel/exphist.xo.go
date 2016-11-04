// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Exphist represents a row from 'equinox.exphist'.
type Exphist struct {
	EphPromoname   sql.NullString `json:"eph_promoname"`   // eph_promoname
	EphQualifying  sql.NullInt64  `json:"eph_qualifying"`  // eph_qualifying
	EphDate        pq.NullTime    `json:"eph_date"`        // eph_date
	EphNotes       sql.NullInt64  `json:"eph_notes"`       // eph_notes
	EphPeriodcount sql.NullInt64  `json:"eph_periodcount"` // eph_periodcount
	EphPeriodstart pq.NullTime    `json:"eph_periodstart"` // eph_periodstart
	EphPeriodend   pq.NullTime    `json:"eph_periodend"`   // eph_periodend
	EphStatus      sql.NullInt64  `json:"eph_status"`      // eph_status
	EphTime        pq.NullTime    `json:"eph_time"`        // eph_time
	EquinoxPrn     sql.NullInt64  `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn     int64          `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec     sql.NullInt64  `json:"equinox_sec"`     // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Exphist exists in the database.
func (e *Exphist) Exists() bool {
	return e._exists
}

// Deleted provides information if the Exphist has been deleted from the database.
func (e *Exphist) Deleted() bool {
	return e._deleted
}

// Insert inserts the Exphist to the database.
func (e *Exphist) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.exphist (` +
		`eph_promoname, eph_qualifying, eph_date, eph_notes, eph_periodcount, eph_periodstart, eph_periodend, eph_status, eph_time, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, e.EphPromoname, e.EphQualifying, e.EphDate, e.EphNotes, e.EphPeriodcount, e.EphPeriodstart, e.EphPeriodend, e.EphStatus, e.EphTime, e.EquinoxPrn, e.EquinoxSec)
	err = db.QueryRow(sqlstr, e.EphPromoname, e.EphQualifying, e.EphDate, e.EphNotes, e.EphPeriodcount, e.EphPeriodstart, e.EphPeriodend, e.EphStatus, e.EphTime, e.EquinoxPrn, e.EquinoxSec).Scan(&e.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Update updates the Exphist in the database.
func (e *Exphist) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.exphist SET (` +
		`eph_promoname, eph_qualifying, eph_date, eph_notes, eph_periodcount, eph_periodstart, eph_periodend, eph_status, eph_time, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) WHERE equinox_lrn = $12`

	// run query
	XOLog(sqlstr, e.EphPromoname, e.EphQualifying, e.EphDate, e.EphNotes, e.EphPeriodcount, e.EphPeriodstart, e.EphPeriodend, e.EphStatus, e.EphTime, e.EquinoxPrn, e.EquinoxSec, e.EquinoxLrn)
	_, err = db.Exec(sqlstr, e.EphPromoname, e.EphQualifying, e.EphDate, e.EphNotes, e.EphPeriodcount, e.EphPeriodstart, e.EphPeriodend, e.EphStatus, e.EphTime, e.EquinoxPrn, e.EquinoxSec, e.EquinoxLrn)
	return err
}

// Save saves the Exphist to the database.
func (e *Exphist) Save(db XODB) error {
	if e.Exists() {
		return e.Update(db)
	}

	return e.Insert(db)
}

// Upsert performs an upsert for Exphist.
//
// NOTE: PostgreSQL 9.5+ only
func (e *Exphist) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.exphist (` +
		`eph_promoname, eph_qualifying, eph_date, eph_notes, eph_periodcount, eph_periodstart, eph_periodend, eph_status, eph_time, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`eph_promoname, eph_qualifying, eph_date, eph_notes, eph_periodcount, eph_periodstart, eph_periodend, eph_status, eph_time, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.eph_promoname, EXCLUDED.eph_qualifying, EXCLUDED.eph_date, EXCLUDED.eph_notes, EXCLUDED.eph_periodcount, EXCLUDED.eph_periodstart, EXCLUDED.eph_periodend, EXCLUDED.eph_status, EXCLUDED.eph_time, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, e.EphPromoname, e.EphQualifying, e.EphDate, e.EphNotes, e.EphPeriodcount, e.EphPeriodstart, e.EphPeriodend, e.EphStatus, e.EphTime, e.EquinoxPrn, e.EquinoxLrn, e.EquinoxSec)
	_, err = db.Exec(sqlstr, e.EphPromoname, e.EphQualifying, e.EphDate, e.EphNotes, e.EphPeriodcount, e.EphPeriodstart, e.EphPeriodend, e.EphStatus, e.EphTime, e.EquinoxPrn, e.EquinoxLrn, e.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Delete deletes the Exphist from the database.
func (e *Exphist) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.exphist WHERE equinox_lrn = $1`

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

// ExphistByEquinoxLrn retrieves a row from 'equinox.exphist' as a Exphist.
//
// Generated from index 'exphist_pkey'.
func ExphistByEquinoxLrn(db XODB, equinoxLrn int64) (*Exphist, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`eph_promoname, eph_qualifying, eph_date, eph_notes, eph_periodcount, eph_periodstart, eph_periodend, eph_status, eph_time, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.exphist ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Exphist{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.EphPromoname, &e.EphQualifying, &e.EphDate, &e.EphNotes, &e.EphPeriodcount, &e.EphPeriodstart, &e.EphPeriodend, &e.EphStatus, &e.EphTime, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
