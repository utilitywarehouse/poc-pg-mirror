// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Comhead represents a row from 'equinox.comhead'.
type Comhead struct {
	Chperiod         sql.NullString  `json:"chperiod"`         // chperiod
	Chdealer         sql.NullString  `json:"chdealer"`         // chdealer
	Chcrnumber       sql.NullInt64   `json:"chcrnumber"`       // chcrnumber
	Chinvoiceno      sql.NullString  `json:"chinvoiceno"`      // chinvoiceno
	Chpicture        sql.NullInt64   `json:"chpicture"`        // chpicture
	Chpaid           pq.NullTime     `json:"chpaid"`           // chpaid
	Chpaymethod      sql.NullString  `json:"chpaymethod"`      // chpaymethod
	Chpaidby         sql.NullString  `json:"chpaidby"`         // chpaidby
	Chnotes          sql.NullInt64   `json:"chnotes"`          // chnotes
	Chsent           pq.NullTime     `json:"chsent"`           // chsent
	Chsentby         sql.NullString  `json:"chsentby"`         // chsentby
	Chtotfixcall     sql.NullFloat64 `json:"chtotfixcall"`     // chtotfixcall
	Chtotmobcall     sql.NullFloat64 `json:"chtotmobcall"`     // chtotmobcall
	Chtotdatacall    sql.NullFloat64 `json:"chtotdatacall"`    // chtotdatacall
	Chtotfixconnect  sql.NullFloat64 `json:"chtotfixconnect"`  // chtotfixconnect
	Chtotmobconnect  sql.NullFloat64 `json:"chtotmobconnect"`  // chtotmobconnect
	Chtotdataconnect sql.NullFloat64 `json:"chtotdataconnect"` // chtotdataconnect
	Chtotcreditdebit sql.NullFloat64 `json:"chtotcreditdebit"` // chtotcreditdebit
	Chtotbonuses     sql.NullFloat64 `json:"chtotbonuses"`     // chtotbonuses
	Chtotresign      sql.NullFloat64 `json:"chtotresign"`      // chtotresign
	Chtotnett        sql.NullFloat64 `json:"chtotnett"`        // chtotnett
	Chtotvat         sql.NullFloat64 `json:"chtotvat"`         // chtotvat
	Chtotgross       sql.NullFloat64 `json:"chtotgross"`       // chtotgross
	Chtotaled        pq.NullTime     `json:"chtotaled"`        // chtotaled
	Chdowns          sql.NullInt64   `json:"chdowns"`          // chdowns
	Chdowne          sql.NullInt64   `json:"chdowne"`          // chdowne
	Chlevel          sql.NullInt64   `json:"chlevel"`          // chlevel
	Chstatus         sql.NullInt64   `json:"chstatus"`         // chstatus
	Chtitle          sql.NullInt64   `json:"chtitle"`          // chtitle
	Chpersonalcusts  sql.NullInt64   `json:"chpersonalcusts"`  // chpersonalcusts
	Chgroupcustomers sql.NullInt64   `json:"chgroupcustomers"` // chgroupcustomers
	Chcustlastmonth  sql.NullInt64   `json:"chcustlastmonth"`  // chcustlastmonth
	Chqcustlastmonth sql.NullFloat64 `json:"chqcustlastmonth"` // chqcustlastmonth
	Chcustinquarter  sql.NullInt64   `json:"chcustinquarter"`  // chcustinquarter
	Chqiregistered   sql.NullInt64   `json:"chqiregistered"`   // chqiregistered
	Chqiachieved     sql.NullInt64   `json:"chqiachieved"`     // chqiachieved
	Chcvcrateplan    sql.NullInt64   `json:"chcvcrateplan"`    // chcvcrateplan
	Chcvcrateplancod sql.NullString  `json:"chcvcrateplancod"` // chcvcrateplancod
	Chformationflags sql.NullString  `json:"chformationflags"` // chformationflags
	Chsparenum1      sql.NullFloat64 `json:"chsparenum1"`      // chsparenum1
	Chsparenum2      sql.NullFloat64 `json:"chsparenum2"`      // chsparenum2
	Chsparenum3      sql.NullFloat64 `json:"chsparenum3"`      // chsparenum3
	Chsparenum4      sql.NullFloat64 `json:"chsparenum4"`      // chsparenum4
	Chsponsorid      sql.NullString  `json:"chsponsorid"`      // chsponsorid
	Chexectype       sql.NullString  `json:"chexectype"`       // chexectype
	Chspared1        pq.NullTime     `json:"chspared1"`        // chspared1
	Chexecenddate    pq.NullTime     `json:"chexecenddate"`    // chexecenddate
	Chsparen1        sql.NullFloat64 `json:"chsparen1"`        // chsparen1
	Chsparen2        sql.NullFloat64 `json:"chsparen2"`        // chsparen2
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllComhead(db XODB, callback func(x Comhead) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`chperiod, chdealer, chcrnumber, chinvoiceno, chpicture, chpaid, chpaymethod, chpaidby, chnotes, chsent, chsentby, chtotfixcall, chtotmobcall, chtotdatacall, chtotfixconnect, chtotmobconnect, chtotdataconnect, chtotcreditdebit, chtotbonuses, chtotresign, chtotnett, chtotvat, chtotgross, chtotaled, chdowns, chdowne, chlevel, chstatus, chtitle, chpersonalcusts, chgroupcustomers, chcustlastmonth, chqcustlastmonth, chcustinquarter, chqiregistered, chqiachieved, chcvcrateplan, chcvcrateplancod, chformationflags, chsparenum1, chsparenum2, chsparenum3, chsparenum4, chsponsorid, chexectype, chspared1, chexecenddate, chsparen1, chsparen2, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.comhead `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Comhead{}

		// scan
		err = q.Scan(&c.Chperiod, &c.Chdealer, &c.Chcrnumber, &c.Chinvoiceno, &c.Chpicture, &c.Chpaid, &c.Chpaymethod, &c.Chpaidby, &c.Chnotes, &c.Chsent, &c.Chsentby, &c.Chtotfixcall, &c.Chtotmobcall, &c.Chtotdatacall, &c.Chtotfixconnect, &c.Chtotmobconnect, &c.Chtotdataconnect, &c.Chtotcreditdebit, &c.Chtotbonuses, &c.Chtotresign, &c.Chtotnett, &c.Chtotvat, &c.Chtotgross, &c.Chtotaled, &c.Chdowns, &c.Chdowne, &c.Chlevel, &c.Chstatus, &c.Chtitle, &c.Chpersonalcusts, &c.Chgroupcustomers, &c.Chcustlastmonth, &c.Chqcustlastmonth, &c.Chcustinquarter, &c.Chqiregistered, &c.Chqiachieved, &c.Chcvcrateplan, &c.Chcvcrateplancod, &c.Chformationflags, &c.Chsparenum1, &c.Chsparenum2, &c.Chsparenum3, &c.Chsparenum4, &c.Chsponsorid, &c.Chexectype, &c.Chspared1, &c.Chexecenddate, &c.Chsparen1, &c.Chsparen2, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// ComheadByEquinoxLrn retrieves a row from 'equinox.comhead' as a Comhead.
//
// Generated from index 'comhead_pkey'.
func ComheadByEquinoxLrn(db XODB, equinoxLrn int64) (*Comhead, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`chperiod, chdealer, chcrnumber, chinvoiceno, chpicture, chpaid, chpaymethod, chpaidby, chnotes, chsent, chsentby, chtotfixcall, chtotmobcall, chtotdatacall, chtotfixconnect, chtotmobconnect, chtotdataconnect, chtotcreditdebit, chtotbonuses, chtotresign, chtotnett, chtotvat, chtotgross, chtotaled, chdowns, chdowne, chlevel, chstatus, chtitle, chpersonalcusts, chgroupcustomers, chcustlastmonth, chqcustlastmonth, chcustinquarter, chqiregistered, chqiachieved, chcvcrateplan, chcvcrateplancod, chformationflags, chsparenum1, chsparenum2, chsparenum3, chsparenum4, chsponsorid, chexectype, chspared1, chexecenddate, chsparen1, chsparen2, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.comhead ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Comhead{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Chperiod, &c.Chdealer, &c.Chcrnumber, &c.Chinvoiceno, &c.Chpicture, &c.Chpaid, &c.Chpaymethod, &c.Chpaidby, &c.Chnotes, &c.Chsent, &c.Chsentby, &c.Chtotfixcall, &c.Chtotmobcall, &c.Chtotdatacall, &c.Chtotfixconnect, &c.Chtotmobconnect, &c.Chtotdataconnect, &c.Chtotcreditdebit, &c.Chtotbonuses, &c.Chtotresign, &c.Chtotnett, &c.Chtotvat, &c.Chtotgross, &c.Chtotaled, &c.Chdowns, &c.Chdowne, &c.Chlevel, &c.Chstatus, &c.Chtitle, &c.Chpersonalcusts, &c.Chgroupcustomers, &c.Chcustlastmonth, &c.Chqcustlastmonth, &c.Chcustinquarter, &c.Chqiregistered, &c.Chqiachieved, &c.Chcvcrateplan, &c.Chcvcrateplancod, &c.Chformationflags, &c.Chsparenum1, &c.Chsparenum2, &c.Chsparenum3, &c.Chsparenum4, &c.Chsponsorid, &c.Chexectype, &c.Chspared1, &c.Chexecenddate, &c.Chsparen1, &c.Chsparen2, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
