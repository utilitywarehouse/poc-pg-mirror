// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// AclusersBlb represents a row from 'equinox.aclusers_blb'.
type AclusersBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AclusersBlb exists in the database.
func (ab *AclusersBlb) Exists() bool {
	return ab._exists
}

// Deleted provides information if the AclusersBlb has been deleted from the database.
func (ab *AclusersBlb) Deleted() bool {
	return ab._deleted
}

// Insert inserts the AclusersBlb to the database.
func (ab *AclusersBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ab._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.aclusers_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, ab.BlbData, ab.BlbText)
	err = db.QueryRow(sqlstr, ab.BlbData, ab.BlbText).Scan(&ab.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	ab._exists = true

	return nil
}

// Update updates the AclusersBlb in the database.
func (ab *AclusersBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ab._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ab._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.aclusers_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, ab.BlbData, ab.BlbText, ab.BlbLrn)
	_, err = db.Exec(sqlstr, ab.BlbData, ab.BlbText, ab.BlbLrn)
	return err
}

// Save saves the AclusersBlb to the database.
func (ab *AclusersBlb) Save(db XODB) error {
	if ab.Exists() {
		return ab.Update(db)
	}

	return ab.Insert(db)
}

// Upsert performs an upsert for AclusersBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (ab *AclusersBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if ab._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.aclusers_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, ab.BlbLrn, ab.BlbData, ab.BlbText)
	_, err = db.Exec(sqlstr, ab.BlbLrn, ab.BlbData, ab.BlbText)
	if err != nil {
		return err
	}

	// set existence
	ab._exists = true

	return nil
}

// Delete deletes the AclusersBlb from the database.
func (ab *AclusersBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ab._exists {
		return nil
	}

	// if deleted, bail
	if ab._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.aclusers_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, ab.BlbLrn)
	_, err = db.Exec(sqlstr, ab.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	ab._deleted = true

	return nil
}

// AclusersBlbByBlbLrn retrieves a row from 'equinox.aclusers_blb' as a AclusersBlb.
//
// Generated from index 'aclusers_blb_pkey'.
func AclusersBlbByBlbLrn(db XODB, blbLrn int64) (*AclusersBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.aclusers_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	ab := AclusersBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&ab.BlbLrn, &ab.BlbData, &ab.BlbText)
	if err != nil {
		return nil, err
	}

	return &ab, nil
}
