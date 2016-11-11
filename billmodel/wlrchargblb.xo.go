// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// WlrchargBlb represents a row from 'equinox.wlrcharg_blb'.
type WlrchargBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllWlrchargBlb(db XODB, callback func(x WlrchargBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.wlrcharg_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		wb := WlrchargBlb{}

		// scan
		err = q.Scan(&wb.BlbLrn, &wb.BlbData, &wb.BlbText)
		if err != nil {
			return err
		}
		if !callback(wb) {
			return nil
		}
	}

	return nil
}

// WlrchargBlbByBlbLrn retrieves a row from 'equinox.wlrcharg_blb' as a WlrchargBlb.
//
// Generated from index 'wlrcharg_blb_pkey'.
func WlrchargBlbByBlbLrn(db XODB, blbLrn int64) (*WlrchargBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.wlrcharg_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	wb := WlrchargBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&wb.BlbLrn, &wb.BlbData, &wb.BlbText)
	if err != nil {
		return nil, err
	}

	return &wb, nil
}
