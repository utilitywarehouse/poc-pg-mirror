// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// SmartcfgBlb represents a row from 'equinox.smartcfg_blb'.
type SmartcfgBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// SmartcfgBlbByBlbLrn retrieves a row from 'equinox.smartcfg_blb' as a SmartcfgBlb.
//
// Generated from index 'smartcfg_blb_pkey'.
func SmartcfgBlbByBlbLrn(db XODB, blbLrn int64) (*SmartcfgBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.smartcfg_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	sb := SmartcfgBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&sb.BlbLrn, &sb.BlbData, &sb.BlbText)
	if err != nil {
		return nil, err
	}

	return &sb, nil
}
