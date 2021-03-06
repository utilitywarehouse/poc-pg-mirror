// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// PortcommBlb represents a row from 'equinox.portcomm_blb'.
type PortcommBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllPortcommBlb(db XODB, callback func(x PortcommBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.portcomm_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		pb := PortcommBlb{}

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

// PortcommBlbByBlbLrn retrieves a row from 'equinox.portcomm_blb' as a PortcommBlb.
//
// Generated from index 'portcomm_blb_pkey'.
func PortcommBlbByBlbLrn(db XODB, blbLrn int64) (*PortcommBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.portcomm_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	pb := PortcommBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&pb.BlbLrn, &pb.BlbData, &pb.BlbText)
	if err != nil {
		return nil, err
	}

	return &pb, nil
}
