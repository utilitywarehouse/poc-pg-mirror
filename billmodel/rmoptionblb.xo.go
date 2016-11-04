// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// RmoptionBlb represents a row from 'equinox.rmoption_blb'.
type RmoptionBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the RmoptionBlb exists in the database.
func (rb *RmoptionBlb) Exists() bool {
	return rb._exists
}

// Deleted provides information if the RmoptionBlb has been deleted from the database.
func (rb *RmoptionBlb) Deleted() bool {
	return rb._deleted
}

// Insert inserts the RmoptionBlb to the database.
func (rb *RmoptionBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if rb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.rmoption_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, rb.BlbData, rb.BlbText)
	err = db.QueryRow(sqlstr, rb.BlbData, rb.BlbText).Scan(&rb.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	rb._exists = true

	return nil
}

// Update updates the RmoptionBlb in the database.
func (rb *RmoptionBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !rb._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if rb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.rmoption_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, rb.BlbData, rb.BlbText, rb.BlbLrn)
	_, err = db.Exec(sqlstr, rb.BlbData, rb.BlbText, rb.BlbLrn)
	return err
}

// Save saves the RmoptionBlb to the database.
func (rb *RmoptionBlb) Save(db XODB) error {
	if rb.Exists() {
		return rb.Update(db)
	}

	return rb.Insert(db)
}

// Upsert performs an upsert for RmoptionBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (rb *RmoptionBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if rb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.rmoption_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, rb.BlbLrn, rb.BlbData, rb.BlbText)
	_, err = db.Exec(sqlstr, rb.BlbLrn, rb.BlbData, rb.BlbText)
	if err != nil {
		return err
	}

	// set existence
	rb._exists = true

	return nil
}

// Delete deletes the RmoptionBlb from the database.
func (rb *RmoptionBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !rb._exists {
		return nil
	}

	// if deleted, bail
	if rb._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.rmoption_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, rb.BlbLrn)
	_, err = db.Exec(sqlstr, rb.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	rb._deleted = true

	return nil
}

// RmoptionBlbByBlbLrn retrieves a row from 'equinox.rmoption_blb' as a RmoptionBlb.
//
// Generated from index 'rmoption_blb_pkey'.
func RmoptionBlbByBlbLrn(db XODB, blbLrn int64) (*RmoptionBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.rmoption_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	rb := RmoptionBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&rb.BlbLrn, &rb.BlbData, &rb.BlbText)
	if err != nil {
		return nil, err
	}

	return &rb, nil
}