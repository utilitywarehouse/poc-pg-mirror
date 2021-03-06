// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// PperrorBlb represents a row from 'equinox.pperror_blb'.
type PperrorBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllPperrorBlb(db XODB, callback func(x PperrorBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.pperror_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		pb := PperrorBlb{}

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

// PperrorBlbByBlbLrn retrieves a row from 'equinox.pperror_blb' as a PperrorBlb.
//
// Generated from index 'pperror_blb_pkey'.
func PperrorBlbByBlbLrn(db XODB, blbLrn int64) (*PperrorBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.pperror_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	pb := PperrorBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&pb.BlbLrn, &pb.BlbData, &pb.BlbText)
	if err != nil {
		return nil, err
	}

	return &pb, nil
}
