// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// ExpcallsBlb represents a row from 'equinox.expcalls_blb'.
type ExpcallsBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllExpcallsBlb(db XODB, callback func(x ExpcallsBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.expcalls_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		eb := ExpcallsBlb{}

		// scan
		err = q.Scan(&eb.BlbLrn, &eb.BlbData, &eb.BlbText)
		if err != nil {
			return err
		}
		if !callback(eb) {
			return nil
		}
	}

	return nil
}

// ExpcallsBlbByBlbLrn retrieves a row from 'equinox.expcalls_blb' as a ExpcallsBlb.
//
// Generated from index 'expcalls_blb_pkey'.
func ExpcallsBlbByBlbLrn(db XODB, blbLrn int64) (*ExpcallsBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.expcalls_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	eb := ExpcallsBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&eb.BlbLrn, &eb.BlbData, &eb.BlbText)
	if err != nil {
		return nil, err
	}

	return &eb, nil
}
