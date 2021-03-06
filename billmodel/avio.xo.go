// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// Avio represents a row from 'equinox.avios'.
type Avio struct {
	Avcustaccountno sql.NullString `json:"avcustaccountno"` // avcustaccountno
	Avaviosnumber   sql.NullString `json:"avaviosnumber"`   // avaviosnumber
	Avsurname       sql.NullString `json:"avsurname"`       // avsurname
	Avmtd           sql.NullInt64  `json:"avmtd"`           // avmtd
	Avytd           sql.NullInt64  `json:"avytd"`           // avytd
	Avservicelvl    sql.NullString `json:"avservicelvl"`    // avservicelvl
	Avsuspend       sql.NullInt64  `json:"avsuspend"`       // avsuspend
	Avsuspreason    sql.NullString `json:"avsuspreason"`    // avsuspreason
	EquinoxLrn      int64          `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`     // equinox_sec
}

func AllAvio(db XODB, callback func(x Avio) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`avcustaccountno, avaviosnumber, avsurname, avmtd, avytd, avservicelvl, avsuspend, avsuspreason, equinox_lrn, equinox_sec ` +
		`FROM equinox.avios `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		a := Avio{}

		// scan
		err = q.Scan(&a.Avcustaccountno, &a.Avaviosnumber, &a.Avsurname, &a.Avmtd, &a.Avytd, &a.Avservicelvl, &a.Avsuspend, &a.Avsuspreason, &a.EquinoxLrn, &a.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(a) {
			return nil
		}
	}

	return nil
}

// AvioByEquinoxLrn retrieves a row from 'equinox.avios' as a Avio.
//
// Generated from index 'avios_pkey'.
func AvioByEquinoxLrn(db XODB, equinoxLrn int64) (*Avio, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`avcustaccountno, avaviosnumber, avsurname, avmtd, avytd, avservicelvl, avsuspend, avsuspreason, equinox_lrn, equinox_sec ` +
		`FROM equinox.avios ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	a := Avio{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&a.Avcustaccountno, &a.Avaviosnumber, &a.Avsurname, &a.Avmtd, &a.Avytd, &a.Avservicelvl, &a.Avsuspend, &a.Avsuspreason, &a.EquinoxLrn, &a.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
