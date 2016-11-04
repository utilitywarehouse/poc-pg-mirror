// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// OptmonthBlb represents a row from 'equinox.optmonth_blb'.
type OptmonthBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the OptmonthBlb exists in the database.
func (ob *OptmonthBlb) Exists() bool {
	return ob._exists
}

// Deleted provides information if the OptmonthBlb has been deleted from the database.
func (ob *OptmonthBlb) Deleted() bool {
	return ob._deleted
}

// Insert inserts the OptmonthBlb to the database.
func (ob *OptmonthBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ob._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.optmonth_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, ob.BlbData, ob.BlbText)
	err = db.QueryRow(sqlstr, ob.BlbData, ob.BlbText).Scan(&ob.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	ob._exists = true

	return nil
}

// Update updates the OptmonthBlb in the database.
func (ob *OptmonthBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ob._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ob._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.optmonth_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, ob.BlbData, ob.BlbText, ob.BlbLrn)
	_, err = db.Exec(sqlstr, ob.BlbData, ob.BlbText, ob.BlbLrn)
	return err
}

// Save saves the OptmonthBlb to the database.
func (ob *OptmonthBlb) Save(db XODB) error {
	if ob.Exists() {
		return ob.Update(db)
	}

	return ob.Insert(db)
}

// Upsert performs an upsert for OptmonthBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (ob *OptmonthBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if ob._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.optmonth_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, ob.BlbLrn, ob.BlbData, ob.BlbText)
	_, err = db.Exec(sqlstr, ob.BlbLrn, ob.BlbData, ob.BlbText)
	if err != nil {
		return err
	}

	// set existence
	ob._exists = true

	return nil
}

// Delete deletes the OptmonthBlb from the database.
func (ob *OptmonthBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ob._exists {
		return nil
	}

	// if deleted, bail
	if ob._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.optmonth_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, ob.BlbLrn)
	_, err = db.Exec(sqlstr, ob.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	ob._deleted = true

	return nil
}

// OptmonthBlbByBlbLrn retrieves a row from 'equinox.optmonth_blb' as a OptmonthBlb.
//
// Generated from index 'optmonth_blb_pkey'.
func OptmonthBlbByBlbLrn(db XODB, blbLrn int64) (*OptmonthBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.optmonth_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	ob := OptmonthBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&ob.BlbLrn, &ob.BlbData, &ob.BlbText)
	if err != nil {
		return nil, err
	}

	return &ob, nil
}