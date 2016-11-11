// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// PperrdbtBlb represents a row from 'equinox.pperrdbt_blb'.
type PperrdbtBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllPperrdbtBlb(db XODB, callback func(x PperrdbtBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.pperrdbt_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		pb := PperrdbtBlb{}

		// scan
		err = q.Scan(&pb.BlbLrn, &pb.BlbData, &pb.BlbText)
		if err != nil {
			return err
		}
		if !callback(pb) {
			return nil
		}
	}

	return nil
}

// PperrdbtBlbByBlbLrn retrieves a row from 'equinox.pperrdbt_blb' as a PperrdbtBlb.
//
// Generated from index 'pperrdbt_blb_pkey'.
func PperrdbtBlbByBlbLrn(db XODB, blbLrn int64) (*PperrdbtBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.pperrdbt_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	pb := PperrdbtBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&pb.BlbLrn, &pb.BlbData, &pb.BlbText)
	if err != nil {
		return nil, err
	}

	return &pb, nil
}
