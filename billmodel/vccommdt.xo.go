// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Vccommdt represents a row from 'equinox.vccommdt'.
type Vccommdt struct {
	Detcperiodfrom   pq.NullTime     `json:"detcperiodfrom"`   // detcperiodfrom
	Detcperiodto     pq.NullTime     `json:"detcperiodto"`     // detcperiodto
	Detccrnumber     sql.NullInt64   `json:"detccrnumber"`     // detccrnumber
	Detccustaccno    sql.NullString  `json:"detccustaccno"`    // detccustaccno
	Detccliunique    sql.NullFloat64 `json:"detccliunique"`    // detccliunique
	Detcclino        sql.NullString  `json:"detcclino"`        // detcclino
	Detcclicpsdate   pq.NullTime     `json:"detcclicpsdate"`   // detcclicpsdate
	Detccliwlrdate   pq.NullTime     `json:"detccliwlrdate"`   // detccliwlrdate
	Detcclilivedate  pq.NullTime     `json:"detcclilivedate"`  // detcclilivedate
	Detccommlivedate pq.NullTime     `json:"detccommlivedate"` // detccommlivedate
	Detcservicetype  sql.NullString  `json:"detcservicetype"`  // detcservicetype
	Detccontractterm sql.NullInt64   `json:"detccontractterm"` // detccontractterm
	Detcservlvl      sql.NullString  `json:"detcservlvl"`      // detcservlvl
	Detcquantity     sql.NullInt64   `json:"detcquantity"`     // detcquantity
	Detcnet          sql.NullFloat64 `json:"detcnet"`          // detcnet
	Detcvat          sql.NullFloat64 `json:"detcvat"`          // detcvat
	Detctotal        sql.NullFloat64 `json:"detctotal"`        // detctotal
	Detcratetable    sql.NullFloat64 `json:"detcratetable"`    // detcratetable
	Detddpayments    sql.NullInt64   `json:"detddpayments"`    // detddpayments
	Detctobeclawed   pq.NullTime     `json:"detctobeclawed"`   // detctobeclawed
	Detccrystal      pq.NullTime     `json:"detccrystal"`      // detccrystal
	Detcconsec       sql.NullInt64   `json:"detcconsec"`       // detcconsec
	Detccomplete     sql.NullInt64   `json:"detccomplete"`     // detccomplete
	Detdateclawed    pq.NullTime     `json:"detdateclawed"`    // detdateclawed
	Detclawedbackbal sql.NullFloat64 `json:"detclawedbackbal"` // detclawedbackbal
	Detuniqueid      sql.NullInt64   `json:"detuniqueid"`      // detuniqueid
	Detcliservorig   sql.NullString  `json:"detcliservorig"`   // detcliservorig
	Detcrmoptid      sql.NullString  `json:"detcrmoptid"`      // detcrmoptid
	Detcstatus       sql.NullInt64   `json:"detcstatus"`       // detcstatus
	Detccbpending    sql.NullInt64   `json:"detccbpending"`    // detccbpending
	Detcpaid         pq.NullTime     `json:"detcpaid"`         // detcpaid
	Detcdead         pq.NullTime     `json:"detcdead"`         // detcdead
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllVccommdt(db XODB, callback func(x Vccommdt) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`detcperiodfrom, detcperiodto, detccrnumber, detccustaccno, detccliunique, detcclino, detcclicpsdate, detccliwlrdate, detcclilivedate, detccommlivedate, detcservicetype, detccontractterm, detcservlvl, detcquantity, detcnet, detcvat, detctotal, detcratetable, detddpayments, detctobeclawed, detccrystal, detcconsec, detccomplete, detdateclawed, detclawedbackbal, detuniqueid, detcliservorig, detcrmoptid, detcstatus, detccbpending, detcpaid, detcdead, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.vccommdt `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		v := Vccommdt{}

		// scan
		err = q.Scan(&v.Detcperiodfrom, &v.Detcperiodto, &v.Detccrnumber, &v.Detccustaccno, &v.Detccliunique, &v.Detcclino, &v.Detcclicpsdate, &v.Detccliwlrdate, &v.Detcclilivedate, &v.Detccommlivedate, &v.Detcservicetype, &v.Detccontractterm, &v.Detcservlvl, &v.Detcquantity, &v.Detcnet, &v.Detcvat, &v.Detctotal, &v.Detcratetable, &v.Detddpayments, &v.Detctobeclawed, &v.Detccrystal, &v.Detcconsec, &v.Detccomplete, &v.Detdateclawed, &v.Detclawedbackbal, &v.Detuniqueid, &v.Detcliservorig, &v.Detcrmoptid, &v.Detcstatus, &v.Detccbpending, &v.Detcpaid, &v.Detcdead, &v.EquinoxPrn, &v.EquinoxLrn, &v.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(v) {
			return nil
		}
	}

	return nil
}

// VccommdtByEquinoxLrn retrieves a row from 'equinox.vccommdt' as a Vccommdt.
//
// Generated from index 'vccommdt_pkey'.
func VccommdtByEquinoxLrn(db XODB, equinoxLrn int64) (*Vccommdt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`detcperiodfrom, detcperiodto, detccrnumber, detccustaccno, detccliunique, detcclino, detcclicpsdate, detccliwlrdate, detcclilivedate, detccommlivedate, detcservicetype, detccontractterm, detcservlvl, detcquantity, detcnet, detcvat, detctotal, detcratetable, detddpayments, detctobeclawed, detccrystal, detcconsec, detccomplete, detdateclawed, detclawedbackbal, detuniqueid, detcliservorig, detcrmoptid, detcstatus, detccbpending, detcpaid, detcdead, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.vccommdt ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	v := Vccommdt{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&v.Detcperiodfrom, &v.Detcperiodto, &v.Detccrnumber, &v.Detccustaccno, &v.Detccliunique, &v.Detcclino, &v.Detcclicpsdate, &v.Detccliwlrdate, &v.Detcclilivedate, &v.Detccommlivedate, &v.Detcservicetype, &v.Detccontractterm, &v.Detcservlvl, &v.Detcquantity, &v.Detcnet, &v.Detcvat, &v.Detctotal, &v.Detcratetable, &v.Detddpayments, &v.Detctobeclawed, &v.Detccrystal, &v.Detcconsec, &v.Detccomplete, &v.Detdateclawed, &v.Detclawedbackbal, &v.Detuniqueid, &v.Detcliservorig, &v.Detcrmoptid, &v.Detcstatus, &v.Detccbpending, &v.Detcpaid, &v.Detcdead, &v.EquinoxPrn, &v.EquinoxLrn, &v.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &v, nil
}
