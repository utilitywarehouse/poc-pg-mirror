// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// DeptmntBlb represents a row from 'equinox.deptmnt_blb'.
type DeptmntBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// DeptmntBlbByBlbLrn retrieves a row from 'equinox.deptmnt_blb' as a DeptmntBlb.
//
// Generated from index 'deptmnt_blb_pkey'.
func DeptmntBlbByBlbLrn(db XODB, blbLrn int64) (*DeptmntBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.deptmnt_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	dbVal := DeptmntBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&dbVal.BlbLrn, &dbVal.BlbData, &dbVal.BlbText)
	if err != nil {
		return nil, err
	}

	return &dbVal, nil
}
