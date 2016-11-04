// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// HsHistBlb represents a row from 'equinox.hs_hist_blb'.
type HsHistBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the HsHistBlb exists in the database.
func (hhb *HsHistBlb) Exists() bool {
	return hhb._exists
}

// Deleted provides information if the HsHistBlb has been deleted from the database.
func (hhb *HsHistBlb) Deleted() bool {
	return hhb._deleted
}

// Insert inserts the HsHistBlb to the database.
func (hhb *HsHistBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if hhb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.hs_hist_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, hhb.BlbData, hhb.BlbText)
	err = db.QueryRow(sqlstr, hhb.BlbData, hhb.BlbText).Scan(&hhb.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	hhb._exists = true

	return nil
}

// Update updates the HsHistBlb in the database.
func (hhb *HsHistBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !hhb._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if hhb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.hs_hist_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, hhb.BlbData, hhb.BlbText, hhb.BlbLrn)
	_, err = db.Exec(sqlstr, hhb.BlbData, hhb.BlbText, hhb.BlbLrn)
	return err
}

// Save saves the HsHistBlb to the database.
func (hhb *HsHistBlb) Save(db XODB) error {
	if hhb.Exists() {
		return hhb.Update(db)
	}

	return hhb.Insert(db)
}

// Upsert performs an upsert for HsHistBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (hhb *HsHistBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if hhb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.hs_hist_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, hhb.BlbLrn, hhb.BlbData, hhb.BlbText)
	_, err = db.Exec(sqlstr, hhb.BlbLrn, hhb.BlbData, hhb.BlbText)
	if err != nil {
		return err
	}

	// set existence
	hhb._exists = true

	return nil
}

// Delete deletes the HsHistBlb from the database.
func (hhb *HsHistBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !hhb._exists {
		return nil
	}

	// if deleted, bail
	if hhb._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.hs_hist_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, hhb.BlbLrn)
	_, err = db.Exec(sqlstr, hhb.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	hhb._deleted = true

	return nil
}

// HsHistBlbByBlbLrn retrieves a row from 'equinox.hs_hist_blb' as a HsHistBlb.
//
// Generated from index 'hs_hist_blb_pkey'.
func HsHistBlbByBlbLrn(db XODB, blbLrn int64) (*HsHistBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.hs_hist_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	hhb := HsHistBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&hhb.BlbLrn, &hhb.BlbData, &hhb.BlbText)
	if err != nil {
		return nil, err
	}

	return &hhb, nil
}