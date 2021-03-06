// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// AccffileBlb represents a row from 'equinox.accffile_blb'.
type AccffileBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllAccffileBlb(db XODB, callback func(x AccffileBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.accffile_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		ab := AccffileBlb{}

		// scan
		err = q.Scan(&ab.BlbLrn, &ab.BlbData, &ab.BlbText)
		if err != nil {
			return err
		}
		if !callback(ab) {
			return nil
		}
	}

	return nil
}

// AccffileBlbByBlbLrn retrieves a row from 'equinox.accffile_blb' as a AccffileBlb.
//
// Generated from index 'accffile_blb_pkey'.
func AccffileBlbByBlbLrn(db XODB, blbLrn int64) (*AccffileBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.accffile_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	ab := AccffileBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&ab.BlbLrn, &ab.BlbData, &ab.BlbText)
	if err != nil {
		return nil, err
	}

	return &ab, nil
}
