// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// CallsBlb represents a row from 'equinox.calls_blb'.
type CallsBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllCallsBlb(db XODB, callback func(x CallsBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.calls_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		cb := CallsBlb{}

		// scan
		err = q.Scan(&cb.BlbLrn, &cb.BlbData, &cb.BlbText)
		if err != nil {
			return err
		}
		if !callback(cb) {
			return nil
		}
	}

	return nil
}

// CallsBlbByBlbLrn retrieves a row from 'equinox.calls_blb' as a CallsBlb.
//
// Generated from index 'calls_blb_pkey'.
func CallsBlbByBlbLrn(db XODB, blbLrn int64) (*CallsBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.calls_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	cb := CallsBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&cb.BlbLrn, &cb.BlbData, &cb.BlbText)
	if err != nil {
		return nil, err
	}

	return &cb, nil
}
