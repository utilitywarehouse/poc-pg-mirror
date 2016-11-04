// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// OVenue represents a row from 'equinox.o_venues'.
type OVenue struct {
	OvVenueid   sql.NullString `json:"ov_venueid"`   // ov_venueid
	OvVenuename sql.NullString `json:"ov_venuename"` // ov_venuename
	OvAddress1  sql.NullString `json:"ov_address1"`  // ov_address1
	OvAddress2  sql.NullString `json:"ov_address2"`  // ov_address2
	OvAddress3  sql.NullString `json:"ov_address3"`  // ov_address3
	OvAddress4  sql.NullString `json:"ov_address4"`  // ov_address4
	OvAddress5  sql.NullString `json:"ov_address5"`  // ov_address5
	OvPostcode  sql.NullString `json:"ov_postcode"`  // ov_postcode
	OvVancode   sql.NullString `json:"ov_vancode"`   // ov_vancode
	EquinoxLrn  int64          `json:"equinox_lrn"`  // equinox_lrn
	EquinoxSec  sql.NullInt64  `json:"equinox_sec"`  // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the OVenue exists in the database.
func (ov *OVenue) Exists() bool {
	return ov._exists
}

// Deleted provides information if the OVenue has been deleted from the database.
func (ov *OVenue) Deleted() bool {
	return ov._deleted
}

// Insert inserts the OVenue to the database.
func (ov *OVenue) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ov._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.o_venues (` +
		`ov_venueid, ov_venuename, ov_address1, ov_address2, ov_address3, ov_address4, ov_address5, ov_postcode, ov_vancode, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, ov.OvVenueid, ov.OvVenuename, ov.OvAddress1, ov.OvAddress2, ov.OvAddress3, ov.OvAddress4, ov.OvAddress5, ov.OvPostcode, ov.OvVancode, ov.EquinoxSec)
	err = db.QueryRow(sqlstr, ov.OvVenueid, ov.OvVenuename, ov.OvAddress1, ov.OvAddress2, ov.OvAddress3, ov.OvAddress4, ov.OvAddress5, ov.OvPostcode, ov.OvVancode, ov.EquinoxSec).Scan(&ov.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	ov._exists = true

	return nil
}

// Update updates the OVenue in the database.
func (ov *OVenue) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ov._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ov._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.o_venues SET (` +
		`ov_venueid, ov_venuename, ov_address1, ov_address2, ov_address3, ov_address4, ov_address5, ov_postcode, ov_vancode, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`) WHERE equinox_lrn = $11`

	// run query
	XOLog(sqlstr, ov.OvVenueid, ov.OvVenuename, ov.OvAddress1, ov.OvAddress2, ov.OvAddress3, ov.OvAddress4, ov.OvAddress5, ov.OvPostcode, ov.OvVancode, ov.EquinoxSec, ov.EquinoxLrn)
	_, err = db.Exec(sqlstr, ov.OvVenueid, ov.OvVenuename, ov.OvAddress1, ov.OvAddress2, ov.OvAddress3, ov.OvAddress4, ov.OvAddress5, ov.OvPostcode, ov.OvVancode, ov.EquinoxSec, ov.EquinoxLrn)
	return err
}

// Save saves the OVenue to the database.
func (ov *OVenue) Save(db XODB) error {
	if ov.Exists() {
		return ov.Update(db)
	}

	return ov.Insert(db)
}

// Upsert performs an upsert for OVenue.
//
// NOTE: PostgreSQL 9.5+ only
func (ov *OVenue) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if ov._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.o_venues (` +
		`ov_venueid, ov_venuename, ov_address1, ov_address2, ov_address3, ov_address4, ov_address5, ov_postcode, ov_vancode, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`ov_venueid, ov_venuename, ov_address1, ov_address2, ov_address3, ov_address4, ov_address5, ov_postcode, ov_vancode, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.ov_venueid, EXCLUDED.ov_venuename, EXCLUDED.ov_address1, EXCLUDED.ov_address2, EXCLUDED.ov_address3, EXCLUDED.ov_address4, EXCLUDED.ov_address5, EXCLUDED.ov_postcode, EXCLUDED.ov_vancode, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, ov.OvVenueid, ov.OvVenuename, ov.OvAddress1, ov.OvAddress2, ov.OvAddress3, ov.OvAddress4, ov.OvAddress5, ov.OvPostcode, ov.OvVancode, ov.EquinoxLrn, ov.EquinoxSec)
	_, err = db.Exec(sqlstr, ov.OvVenueid, ov.OvVenuename, ov.OvAddress1, ov.OvAddress2, ov.OvAddress3, ov.OvAddress4, ov.OvAddress5, ov.OvPostcode, ov.OvVancode, ov.EquinoxLrn, ov.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	ov._exists = true

	return nil
}

// Delete deletes the OVenue from the database.
func (ov *OVenue) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ov._exists {
		return nil
	}

	// if deleted, bail
	if ov._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.o_venues WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, ov.EquinoxLrn)
	_, err = db.Exec(sqlstr, ov.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	ov._deleted = true

	return nil
}

// OVenueByEquinoxLrn retrieves a row from 'equinox.o_venues' as a OVenue.
//
// Generated from index 'o_venues_pkey'.
func OVenueByEquinoxLrn(db XODB, equinoxLrn int64) (*OVenue, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ov_venueid, ov_venuename, ov_address1, ov_address2, ov_address3, ov_address4, ov_address5, ov_postcode, ov_vancode, equinox_lrn, equinox_sec ` +
		`FROM equinox.o_venues ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	ov := OVenue{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&ov.OvVenueid, &ov.OvVenuename, &ov.OvAddress1, &ov.OvAddress2, &ov.OvAddress3, &ov.OvAddress4, &ov.OvAddress5, &ov.OvPostcode, &ov.OvVancode, &ov.EquinoxLrn, &ov.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &ov, nil
}
