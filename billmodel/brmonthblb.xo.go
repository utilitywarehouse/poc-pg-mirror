// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// BrmonthBlb represents a row from 'equinox.brmonth_blb'.
type BrmonthBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllBrmonthBlb(db XODB, callback func(x BrmonthBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.brmonth_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		bb := BrmonthBlb{}

		// scan
		err = q.Scan(&bb.BlbLrn, &bb.BlbData, &bb.BlbText)
		if err != nil {
			return err
		}
		if !callback(bb) {
			return nil
		}
	}

	return nil
}

// BrmonthBlbByBlbLrn retrieves a row from 'equinox.brmonth_blb' as a BrmonthBlb.
//
// Generated from index 'brmonth_blb_pkey'.
func BrmonthBlbByBlbLrn(db XODB, blbLrn int64) (*BrmonthBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.brmonth_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	bb := BrmonthBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&bb.BlbLrn, &bb.BlbData, &bb.BlbText)
	if err != nil {
		return nil, err
	}

	return &bb, nil
}
