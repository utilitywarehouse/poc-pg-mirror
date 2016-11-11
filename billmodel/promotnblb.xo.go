// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// PromotnBlb represents a row from 'equinox.promotn_blb'.
type PromotnBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllPromotnBlb(db XODB, callback func(x PromotnBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.promotn_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		pb := PromotnBlb{}

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

// PromotnBlbByBlbLrn retrieves a row from 'equinox.promotn_blb' as a PromotnBlb.
//
// Generated from index 'promotn_blb_pkey'.
func PromotnBlbByBlbLrn(db XODB, blbLrn int64) (*PromotnBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.promotn_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	pb := PromotnBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&pb.BlbLrn, &pb.BlbData, &pb.BlbText)
	if err != nil {
		return nil, err
	}

	return &pb, nil
}
