// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// TaskdestBlb represents a row from 'equinox.taskdest_blb'.
type TaskdestBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

// TaskdestBlbByBlbLrn retrieves a row from 'equinox.taskdest_blb' as a TaskdestBlb.
//
// Generated from index 'taskdest_blb_pkey'.
func TaskdestBlbByBlbLrn(db XODB, blbLrn int64) (*TaskdestBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.taskdest_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	tb := TaskdestBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&tb.BlbLrn, &tb.BlbData, &tb.BlbText)
	if err != nil {
		return nil, err
	}

	return &tb, nil
}
