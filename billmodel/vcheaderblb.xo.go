// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// VcheaderBlb represents a row from 'equinox.vcheader_blb'.
type VcheaderBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllVcheaderBlb(db XODB, callback func(x VcheaderBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.vcheader_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		vb := VcheaderBlb{}

		// scan
		err = q.Scan(&vb.BlbLrn, &vb.BlbData, &vb.BlbText)
		if err != nil {
			return err
		}
		if !callback(vb) {
			return nil
		}
	}

	return nil
}

// VcheaderBlbByBlbLrn retrieves a row from 'equinox.vcheader_blb' as a VcheaderBlb.
//
// Generated from index 'vcheader_blb_pkey'.
func VcheaderBlbByBlbLrn(db XODB, blbLrn int64) (*VcheaderBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.vcheader_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	vb := VcheaderBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&vb.BlbLrn, &vb.BlbData, &vb.BlbText)
	if err != nil {
		return nil, err
	}

	return &vb, nil
}
