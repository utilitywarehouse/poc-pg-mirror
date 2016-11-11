// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Extpay represents a row from 'equinox.extpay'.
type Extpay struct {
	Extpayperiod     sql.NullString  `json:"extpayperiod"`     // extpayperiod
	Extgoplus        sql.NullFloat64 `json:"extgoplus"`        // extgoplus
	Extpayperscvc    sql.NullFloat64 `json:"extpayperscvc"`    // extpayperscvc
	Extother         sql.NullFloat64 `json:"extother"`         // extother
	Extpayperscgb    sql.NullFloat64 `json:"extpayperscgb"`    // extpayperscgb
	Extpaycgbback    sql.NullFloat64 `json:"extpaycgbback"`    // extpaycgbback
	Extpaydealer     sql.NullString  `json:"extpaydealer"`     // extpaydealer
	Extpaytotbb      sql.NullFloat64 `json:"extpaytotbb"`      // extpaytotbb
	Extpayweb        sql.NullFloat64 `json:"extpayweb"`        // extpayweb
	Extpaytotalcvc   sql.NullFloat64 `json:"extpaytotalcvc"`   // extpaytotalcvc
	Extpaytotal      sql.NullFloat64 `json:"extpaytotal"`      // extpaytotal
	Extpaycarry      sql.NullFloat64 `json:"extpaycarry"`      // extpaycarry
	Extlevel1        sql.NullFloat64 `json:"extlevel1"`        // extlevel1
	Extlevel2        sql.NullFloat64 `json:"extlevel2"`        // extlevel2
	Extlevel3        sql.NullFloat64 `json:"extlevel3"`        // extlevel3
	Extlevel4        sql.NullFloat64 `json:"extlevel4"`        // extlevel4
	Extlevel5        sql.NullFloat64 `json:"extlevel5"`        // extlevel5
	Extlevel6        sql.NullFloat64 `json:"extlevel6"`        // extlevel6
	Extlevel7        sql.NullFloat64 `json:"extlevel7"`        // extlevel7
	Extlevel8        sql.NullFloat64 `json:"extlevel8"`        // extlevel8
	Extudb           sql.NullFloat64 `json:"extudb"`           // extudb
	Extsdb           sql.NullFloat64 `json:"extsdb"`           // extsdb
	Exterror         sql.NullFloat64 `json:"exterror"`         // exterror
	Extclawback      sql.NullFloat64 `json:"extclawback"`      // extclawback
	Extadjustment    sql.NullFloat64 `json:"extadjustment"`    // extadjustment
	Extnum1          sql.NullFloat64 `json:"extnum1"`          // extnum1
	Extnum2          sql.NullFloat64 `json:"extnum2"`          // extnum2
	Extdate1         pq.NullTime     `json:"extdate1"`         // extdate1
	Extloan          sql.NullFloat64 `json:"extloan"`          // extloan
	Exttraining      sql.NullFloat64 `json:"exttraining"`      // exttraining
	Extentered       sql.NullString  `json:"extentered"`       // extentered
	Extpot           sql.NullFloat64 `json:"extpot"`           // extpot
	Extrpc           sql.NullFloat64 `json:"extrpc"`           // extrpc
	Extspecialmobile sql.NullFloat64 `json:"extspecialmobile"` // extspecialmobile
	Extpayfast       sql.NullFloat64 `json:"extpayfast"`       // extpayfast
	Extpaypromobonus sql.NullFloat64 `json:"extpaypromobonus"` // extpaypromobonus
	Extpayuplift     sql.NullFloat64 `json:"extpayuplift"`     // extpayuplift
	Extpaysponsortbb sql.NullFloat64 `json:"extpaysponsortbb"` // extpaysponsortbb
	Extpayuplinetbb  sql.NullFloat64 `json:"extpayuplinetbb"`  // extpayuplinetbb
	Extlgcb          sql.NullFloat64 `json:"extlgcb"`          // extlgcb
	Extenergy        sql.NullFloat64 `json:"extenergy"`        // extenergy
	Exterrorsource   sql.NullString  `json:"exterrorsource"`   // exterrorsource
	Extmarginal      sql.NullFloat64 `json:"extmarginal"`      // extmarginal
	Extrenewal       sql.NullFloat64 `json:"extrenewal"`       // extrenewal
	Extpaycarrepay   sql.NullFloat64 `json:"extpaycarrepay"`   // extpaycarrepay
	Extpaycarbonus   sql.NullFloat64 `json:"extpaycarbonus"`   // extpaycarbonus
	Extpaypromo1     sql.NullFloat64 `json:"extpaypromo1"`     // extpaypromo1
	Extpaypromo2     sql.NullFloat64 `json:"extpaypromo2"`     // extpaypromo2
	Extpaypromo3     sql.NullFloat64 `json:"extpaypromo3"`     // extpaypromo3
	Extpaypromo4     sql.NullFloat64 `json:"extpaypromo4"`     // extpaypromo4
	Extratetel       sql.NullFloat64 `json:"extratetel"`       // extratetel
	Extrateenergy    sql.NullFloat64 `json:"extrateenergy"`    // extrateenergy
	Extstatus        sql.NullInt64   `json:"extstatus"`        // extstatus
	Extcustlastmonth sql.NullInt64   `json:"extcustlastmonth"` // extcustlastmonth
	Exttitle         sql.NullInt64   `json:"exttitle"`         // exttitle
	Extquickincome   sql.NullInt64   `json:"extquickincome"`   // extquickincome
	Extsendby        sql.NullString  `json:"extsendby"`        // extsendby
	Extsentdate      pq.NullTime     `json:"extsentdate"`      // extsentdate
	Extbatchno       sql.NullInt64   `json:"extbatchno"`       // extbatchno
	Extbatchindex    sql.NullInt64   `json:"extbatchindex"`    // extbatchindex
	Extgenie1start   sql.NullInt64   `json:"extgenie1start"`   // extgenie1start
	Extgenie1end     sql.NullInt64   `json:"extgenie1end"`     // extgenie1end
	Extgenie1level   sql.NullInt64   `json:"extgenie1level"`   // extgenie1level
	Extgenie2start   sql.NullInt64   `json:"extgenie2start"`   // extgenie2start
	Extgenie2end     sql.NullInt64   `json:"extgenie2end"`     // extgenie2end
	Extgenie2level   sql.NullInt64   `json:"extgenie2level"`   // extgenie2level
	Extpaysparec1    sql.NullString  `json:"extpaysparec1"`    // extpaysparec1
	Extpaysparen1    sql.NullFloat64 `json:"extpaysparen1"`    // extpaysparen1
	Extpayspared1    pq.NullTime     `json:"extpayspared1"`    // extpayspared1
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllExtpay(db XODB, callback func(x Extpay) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`extpayperiod, extgoplus, extpayperscvc, extother, extpayperscgb, extpaycgbback, extpaydealer, extpaytotbb, extpayweb, extpaytotalcvc, extpaytotal, extpaycarry, extlevel1, extlevel2, extlevel3, extlevel4, extlevel5, extlevel6, extlevel7, extlevel8, extudb, extsdb, exterror, extclawback, extadjustment, extnum1, extnum2, extdate1, extloan, exttraining, extentered, extpot, extrpc, extspecialmobile, extpayfast, extpaypromobonus, extpayuplift, extpaysponsortbb, extpayuplinetbb, extlgcb, extenergy, exterrorsource, extmarginal, extrenewal, extpaycarrepay, extpaycarbonus, extpaypromo1, extpaypromo2, extpaypromo3, extpaypromo4, extratetel, extrateenergy, extstatus, extcustlastmonth, exttitle, extquickincome, extsendby, extsentdate, extbatchno, extbatchindex, extgenie1start, extgenie1end, extgenie1level, extgenie2start, extgenie2end, extgenie2level, extpaysparec1, extpaysparen1, extpayspared1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.extpay `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		e := Extpay{}

		// scan
		err = q.Scan(&e.Extpayperiod, &e.Extgoplus, &e.Extpayperscvc, &e.Extother, &e.Extpayperscgb, &e.Extpaycgbback, &e.Extpaydealer, &e.Extpaytotbb, &e.Extpayweb, &e.Extpaytotalcvc, &e.Extpaytotal, &e.Extpaycarry, &e.Extlevel1, &e.Extlevel2, &e.Extlevel3, &e.Extlevel4, &e.Extlevel5, &e.Extlevel6, &e.Extlevel7, &e.Extlevel8, &e.Extudb, &e.Extsdb, &e.Exterror, &e.Extclawback, &e.Extadjustment, &e.Extnum1, &e.Extnum2, &e.Extdate1, &e.Extloan, &e.Exttraining, &e.Extentered, &e.Extpot, &e.Extrpc, &e.Extspecialmobile, &e.Extpayfast, &e.Extpaypromobonus, &e.Extpayuplift, &e.Extpaysponsortbb, &e.Extpayuplinetbb, &e.Extlgcb, &e.Extenergy, &e.Exterrorsource, &e.Extmarginal, &e.Extrenewal, &e.Extpaycarrepay, &e.Extpaycarbonus, &e.Extpaypromo1, &e.Extpaypromo2, &e.Extpaypromo3, &e.Extpaypromo4, &e.Extratetel, &e.Extrateenergy, &e.Extstatus, &e.Extcustlastmonth, &e.Exttitle, &e.Extquickincome, &e.Extsendby, &e.Extsentdate, &e.Extbatchno, &e.Extbatchindex, &e.Extgenie1start, &e.Extgenie1end, &e.Extgenie1level, &e.Extgenie2start, &e.Extgenie2end, &e.Extgenie2level, &e.Extpaysparec1, &e.Extpaysparen1, &e.Extpayspared1, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(e) {
			return nil
		}
	}

	return nil
}

// ExtpayByEquinoxLrn retrieves a row from 'equinox.extpay' as a Extpay.
//
// Generated from index 'extpay_pkey'.
func ExtpayByEquinoxLrn(db XODB, equinoxLrn int64) (*Extpay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`extpayperiod, extgoplus, extpayperscvc, extother, extpayperscgb, extpaycgbback, extpaydealer, extpaytotbb, extpayweb, extpaytotalcvc, extpaytotal, extpaycarry, extlevel1, extlevel2, extlevel3, extlevel4, extlevel5, extlevel6, extlevel7, extlevel8, extudb, extsdb, exterror, extclawback, extadjustment, extnum1, extnum2, extdate1, extloan, exttraining, extentered, extpot, extrpc, extspecialmobile, extpayfast, extpaypromobonus, extpayuplift, extpaysponsortbb, extpayuplinetbb, extlgcb, extenergy, exterrorsource, extmarginal, extrenewal, extpaycarrepay, extpaycarbonus, extpaypromo1, extpaypromo2, extpaypromo3, extpaypromo4, extratetel, extrateenergy, extstatus, extcustlastmonth, exttitle, extquickincome, extsendby, extsentdate, extbatchno, extbatchindex, extgenie1start, extgenie1end, extgenie1level, extgenie2start, extgenie2end, extgenie2level, extpaysparec1, extpaysparen1, extpayspared1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.extpay ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Extpay{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Extpayperiod, &e.Extgoplus, &e.Extpayperscvc, &e.Extother, &e.Extpayperscgb, &e.Extpaycgbback, &e.Extpaydealer, &e.Extpaytotbb, &e.Extpayweb, &e.Extpaytotalcvc, &e.Extpaytotal, &e.Extpaycarry, &e.Extlevel1, &e.Extlevel2, &e.Extlevel3, &e.Extlevel4, &e.Extlevel5, &e.Extlevel6, &e.Extlevel7, &e.Extlevel8, &e.Extudb, &e.Extsdb, &e.Exterror, &e.Extclawback, &e.Extadjustment, &e.Extnum1, &e.Extnum2, &e.Extdate1, &e.Extloan, &e.Exttraining, &e.Extentered, &e.Extpot, &e.Extrpc, &e.Extspecialmobile, &e.Extpayfast, &e.Extpaypromobonus, &e.Extpayuplift, &e.Extpaysponsortbb, &e.Extpayuplinetbb, &e.Extlgcb, &e.Extenergy, &e.Exterrorsource, &e.Extmarginal, &e.Extrenewal, &e.Extpaycarrepay, &e.Extpaycarbonus, &e.Extpaypromo1, &e.Extpaypromo2, &e.Extpaypromo3, &e.Extpaypromo4, &e.Extratetel, &e.Extrateenergy, &e.Extstatus, &e.Extcustlastmonth, &e.Exttitle, &e.Extquickincome, &e.Extsendby, &e.Extsentdate, &e.Extbatchno, &e.Extbatchindex, &e.Extgenie1start, &e.Extgenie1end, &e.Extgenie1level, &e.Extgenie2start, &e.Extgenie2end, &e.Extgenie2level, &e.Extpaysparec1, &e.Extpaysparen1, &e.Extpayspared1, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
