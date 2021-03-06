// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// PortingBlb represents a row from 'equinox.porting_blb'.
type PortingBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllPortingBlb(db XODB, callback func(x PortingBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.porting_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		pb := PortingBlb{}

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

// PortingBlbByBlbLrn retrieves a row from 'equinox.porting_blb' as a PortingBlb.
//
// Generated from index 'porting_blb_pkey'.
func PortingBlbByBlbLrn(db XODB, blbLrn int64) (*PortingBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.porting_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	pb := PortingBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&pb.BlbLrn, &pb.BlbData, &pb.BlbText)
	if err != nil {
		return nil, err
	}

	return &pb, nil
}
