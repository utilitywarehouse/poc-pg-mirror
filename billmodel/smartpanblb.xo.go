// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// SmartpanBlb represents a row from 'equinox.smartpan_blb'.
type SmartpanBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// SmartpanBlbByBlbLrn retrieves a row from 'equinox.smartpan_blb' as a SmartpanBlb.
//
// Generated from index 'smartpan_blb_pkey'.
func SmartpanBlbByBlbLrn(db XODB, blbLrn int64) (*SmartpanBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.smartpan_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	sb := SmartpanBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&sb.BlbLrn, &sb.BlbData, &sb.BlbText)
	if err != nil {
		return nil, err
	}

	return &sb, nil
}
