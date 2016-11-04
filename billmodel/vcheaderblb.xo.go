// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// VcheaderBlb represents a row from 'equinox.vcheader_blb'.
type VcheaderBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the VcheaderBlb exists in the database.
func (vb *VcheaderBlb) Exists() bool {
	return vb._exists
}

// Deleted provides information if the VcheaderBlb has been deleted from the database.
func (vb *VcheaderBlb) Deleted() bool {
	return vb._deleted
}

// Insert inserts the VcheaderBlb to the database.
func (vb *VcheaderBlb) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if vb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.vcheader_blb (` +
		`blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING blb_lrn`

	// run query
	XOLog(sqlstr, vb.BlbData, vb.BlbText)
	err = db.QueryRow(sqlstr, vb.BlbData, vb.BlbText).Scan(&vb.BlbLrn)
	if err != nil {
		return err
	}

	// set existence
	vb._exists = true

	return nil
}

// Update updates the VcheaderBlb in the database.
func (vb *VcheaderBlb) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !vb._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if vb._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.vcheader_blb SET (` +
		`blb_data, blb_text` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE blb_lrn = $3`

	// run query
	XOLog(sqlstr, vb.BlbData, vb.BlbText, vb.BlbLrn)
	_, err = db.Exec(sqlstr, vb.BlbData, vb.BlbText, vb.BlbLrn)
	return err
}

// Save saves the VcheaderBlb to the database.
func (vb *VcheaderBlb) Save(db XODB) error {
	if vb.Exists() {
		return vb.Update(db)
	}

	return vb.Insert(db)
}

// Upsert performs an upsert for VcheaderBlb.
//
// NOTE: PostgreSQL 9.5+ only
func (vb *VcheaderBlb) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if vb._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.vcheader_blb (` +
		`blb_lrn, blb_data, blb_text` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (blb_lrn) DO UPDATE SET (` +
		`blb_lrn, blb_data, blb_text` +
		`) = (` +
		`EXCLUDED.blb_lrn, EXCLUDED.blb_data, EXCLUDED.blb_text` +
		`)`

	// run query
	XOLog(sqlstr, vb.BlbLrn, vb.BlbData, vb.BlbText)
	_, err = db.Exec(sqlstr, vb.BlbLrn, vb.BlbData, vb.BlbText)
	if err != nil {
		return err
	}

	// set existence
	vb._exists = true

	return nil
}

// Delete deletes the VcheaderBlb from the database.
func (vb *VcheaderBlb) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !vb._exists {
		return nil
	}

	// if deleted, bail
	if vb._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.vcheader_blb WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, vb.BlbLrn)
	_, err = db.Exec(sqlstr, vb.BlbLrn)
	if err != nil {
		return err
	}

	// set deleted
	vb._deleted = true

	return nil
}

// VcheaderBlbByBlbLrn retrieves a row from 'equinox.vcheader_blb' as a VcheaderBlb.
//
// Generated from index 'vcheader_blb_pkey'.
func VcheaderBlbByBlbLrn(db XODB, blbLrn int64) (*VcheaderBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.vcheader_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	vb := VcheaderBlb{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&vb.BlbLrn, &vb.BlbData, &vb.BlbText)
	if err != nil {
		return nil, err
	}

	return &vb, nil
}
