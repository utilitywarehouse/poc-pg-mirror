// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Ctcomp represents a row from 'equinox.ctcomp'.
type Ctcomp struct {
	Ctcompname       sql.NullString  `json:"ctcompname"`       // ctcompname
	Ctstandingelec   sql.NullFloat64 `json:"ctstandingelec"`   // ctstandingelec
	Ctstandinggas    sql.NullFloat64 `json:"ctstandinggas"`    // ctstandinggas
	Ctstandinge7     sql.NullFloat64 `json:"ctstandinge7"`     // ctstandinge7
	Ctstandingdfg    sql.NullFloat64 `json:"ctstandingdfg"`    // ctstandingdfg
	Ctunit1          sql.NullFloat64 `json:"ctunit1"`          // ctunit1
	Ctunit2          sql.NullFloat64 `json:"ctunit2"`          // ctunit2
	Ctunit1e7        sql.NullFloat64 `json:"ctunit1e7"`        // ctunit1e7
	Ctunit2e7        sql.NullFloat64 `json:"ctunit2e7"`        // ctunit2e7
	Ctunitraten      sql.NullFloat64 `json:"ctunitraten"`      // ctunitraten
	Ctunit1gas       sql.NullFloat64 `json:"ctunit1gas"`       // ctunit1gas
	Ctunit2gas       sql.NullFloat64 `json:"ctunit2gas"`       // ctunit2gas
	Ctunit1dfg       sql.NullFloat64 `json:"ctunit1dfg"`       // ctunit1dfg
	Ctunit2dfg       sql.NullFloat64 `json:"ctunit2dfg"`       // ctunit2dfg
	Ctunitthreshgas  sql.NullInt64   `json:"ctunitthreshgas"`  // ctunitthreshgas
	Ctunitthreshelec sql.NullInt64   `json:"ctunitthreshelec"` // ctunitthreshelec
	Ctunitthreshe7   sql.NullInt64   `json:"ctunitthreshe7"`   // ctunitthreshe7
	Ctunitthreshdfg  sql.NullInt64   `json:"ctunitthreshdfg"`  // ctunitthreshdfg
	Ctddfixedelec    sql.NullFloat64 `json:"ctddfixedelec"`    // ctddfixedelec
	Ctddfixedgas     sql.NullFloat64 `json:"ctddfixedgas"`     // ctddfixedgas
	Ctddfixede7      sql.NullFloat64 `json:"ctddfixede7"`      // ctddfixede7
	Ctddfixeddfg     sql.NullFloat64 `json:"ctddfixeddfg"`     // ctddfixeddfg
	Ctddpercelec     sql.NullFloat64 `json:"ctddpercelec"`     // ctddpercelec
	Ctddpercgas      sql.NullFloat64 `json:"ctddpercgas"`      // ctddpercgas
	Ctddperce7       sql.NullFloat64 `json:"ctddperce7"`       // ctddperce7
	Ctddpercdfg      sql.NullFloat64 `json:"ctddpercdfg"`      // ctddpercdfg
	Ctddthreshelec   sql.NullFloat64 `json:"ctddthreshelec"`   // ctddthreshelec
	Ctddthreshgas    sql.NullFloat64 `json:"ctddthreshgas"`    // ctddthreshgas
	Ctddthreshe7     sql.NullFloat64 `json:"ctddthreshe7"`     // ctddthreshe7
	Ctddthreshdfg    sql.NullFloat64 `json:"ctddthreshdfg"`    // ctddthreshdfg
	Ctdffixed        sql.NullFloat64 `json:"ctdffixed"`        // ctdffixed
	Ctdfperc         sql.NullFloat64 `json:"ctdfperc"`         // ctdfperc
	Ctgasperc        sql.NullFloat64 `json:"ctgasperc"`        // ctgasperc
	Ctdualperc       sql.NullFloat64 `json:"ctdualperc"`       // ctdualperc
	Ctgasbilll       sql.NullFloat64 `json:"ctgasbilll"`       // ctgasbilll
	Ctgasbills       sql.NullFloat64 `json:"ctgasbills"`       // ctgasbills
	Ctgasbillh       sql.NullFloat64 `json:"ctgasbillh"`       // ctgasbillh
	Ctelecbilll      sql.NullFloat64 `json:"ctelecbilll"`      // ctelecbilll
	Ctelecbills      sql.NullFloat64 `json:"ctelecbills"`      // ctelecbills
	Ctelecbillh      sql.NullFloat64 `json:"ctelecbillh"`      // ctelecbillh
	Cte7billl        sql.NullFloat64 `json:"cte7billl"`        // cte7billl
	Cte7bills        sql.NullFloat64 `json:"cte7bills"`        // cte7bills
	Cte7billh        sql.NullFloat64 `json:"cte7billh"`        // cte7billh
	Ctdualbilll      sql.NullFloat64 `json:"ctdualbilll"`      // ctdualbilll
	Ctdualbills      sql.NullFloat64 `json:"ctdualbills"`      // ctdualbills
	Ctdualbillh      sql.NullFloat64 `json:"ctdualbillh"`      // ctdualbillh
	Ctdualbille7l    sql.NullFloat64 `json:"ctdualbille7l"`    // ctdualbille7l
	Ctdualbille7s    sql.NullFloat64 `json:"ctdualbille7s"`    // ctdualbille7s
	Ctdualbille7h    sql.NullFloat64 `json:"ctdualbille7h"`    // ctdualbille7h
	Cteleperc        sql.NullFloat64 `json:"cteleperc"`        // cteleperc
	Ctn1             sql.NullFloat64 `json:"ctn1"`             // ctn1
	Ctn2             sql.NullFloat64 `json:"ctn2"`             // ctn2
	Ctn3             sql.NullFloat64 `json:"ctn3"`             // ctn3
	Ctd1             pq.NullTime     `json:"ctd1"`             // ctd1
	Ctd2             pq.NullTime     `json:"ctd2"`             // ctd2
	Ctc1             sql.NullString  `json:"ctc1"`             // ctc1
	Ctc2             sql.NullString  `json:"ctc2"`             // ctc2
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllCtcomp(db XODB, callback func(x Ctcomp) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`ctcompname, ctstandingelec, ctstandinggas, ctstandinge7, ctstandingdfg, ctunit1, ctunit2, ctunit1e7, ctunit2e7, ctunitraten, ctunit1gas, ctunit2gas, ctunit1dfg, ctunit2dfg, ctunitthreshgas, ctunitthreshelec, ctunitthreshe7, ctunitthreshdfg, ctddfixedelec, ctddfixedgas, ctddfixede7, ctddfixeddfg, ctddpercelec, ctddpercgas, ctddperce7, ctddpercdfg, ctddthreshelec, ctddthreshgas, ctddthreshe7, ctddthreshdfg, ctdffixed, ctdfperc, ctgasperc, ctdualperc, ctgasbilll, ctgasbills, ctgasbillh, ctelecbilll, ctelecbills, ctelecbillh, cte7billl, cte7bills, cte7billh, ctdualbilll, ctdualbills, ctdualbillh, ctdualbille7l, ctdualbille7s, ctdualbille7h, cteleperc, ctn1, ctn2, ctn3, ctd1, ctd2, ctc1, ctc2, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.ctcomp `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Ctcomp{}

		// scan
		err = q.Scan(&c.Ctcompname, &c.Ctstandingelec, &c.Ctstandinggas, &c.Ctstandinge7, &c.Ctstandingdfg, &c.Ctunit1, &c.Ctunit2, &c.Ctunit1e7, &c.Ctunit2e7, &c.Ctunitraten, &c.Ctunit1gas, &c.Ctunit2gas, &c.Ctunit1dfg, &c.Ctunit2dfg, &c.Ctunitthreshgas, &c.Ctunitthreshelec, &c.Ctunitthreshe7, &c.Ctunitthreshdfg, &c.Ctddfixedelec, &c.Ctddfixedgas, &c.Ctddfixede7, &c.Ctddfixeddfg, &c.Ctddpercelec, &c.Ctddpercgas, &c.Ctddperce7, &c.Ctddpercdfg, &c.Ctddthreshelec, &c.Ctddthreshgas, &c.Ctddthreshe7, &c.Ctddthreshdfg, &c.Ctdffixed, &c.Ctdfperc, &c.Ctgasperc, &c.Ctdualperc, &c.Ctgasbilll, &c.Ctgasbills, &c.Ctgasbillh, &c.Ctelecbilll, &c.Ctelecbills, &c.Ctelecbillh, &c.Cte7billl, &c.Cte7bills, &c.Cte7billh, &c.Ctdualbilll, &c.Ctdualbills, &c.Ctdualbillh, &c.Ctdualbille7l, &c.Ctdualbille7s, &c.Ctdualbille7h, &c.Cteleperc, &c.Ctn1, &c.Ctn2, &c.Ctn3, &c.Ctd1, &c.Ctd2, &c.Ctc1, &c.Ctc2, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CtcompByEquinoxLrn retrieves a row from 'equinox.ctcomp' as a Ctcomp.
//
// Generated from index 'ctcomp_pkey'.
func CtcompByEquinoxLrn(db XODB, equinoxLrn int64) (*Ctcomp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ctcompname, ctstandingelec, ctstandinggas, ctstandinge7, ctstandingdfg, ctunit1, ctunit2, ctunit1e7, ctunit2e7, ctunitraten, ctunit1gas, ctunit2gas, ctunit1dfg, ctunit2dfg, ctunitthreshgas, ctunitthreshelec, ctunitthreshe7, ctunitthreshdfg, ctddfixedelec, ctddfixedgas, ctddfixede7, ctddfixeddfg, ctddpercelec, ctddpercgas, ctddperce7, ctddpercdfg, ctddthreshelec, ctddthreshgas, ctddthreshe7, ctddthreshdfg, ctdffixed, ctdfperc, ctgasperc, ctdualperc, ctgasbilll, ctgasbills, ctgasbillh, ctelecbilll, ctelecbills, ctelecbillh, cte7billl, cte7bills, cte7billh, ctdualbilll, ctdualbills, ctdualbillh, ctdualbille7l, ctdualbille7s, ctdualbille7h, cteleperc, ctn1, ctn2, ctn3, ctd1, ctd2, ctc1, ctc2, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.ctcomp ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Ctcomp{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Ctcompname, &c.Ctstandingelec, &c.Ctstandinggas, &c.Ctstandinge7, &c.Ctstandingdfg, &c.Ctunit1, &c.Ctunit2, &c.Ctunit1e7, &c.Ctunit2e7, &c.Ctunitraten, &c.Ctunit1gas, &c.Ctunit2gas, &c.Ctunit1dfg, &c.Ctunit2dfg, &c.Ctunitthreshgas, &c.Ctunitthreshelec, &c.Ctunitthreshe7, &c.Ctunitthreshdfg, &c.Ctddfixedelec, &c.Ctddfixedgas, &c.Ctddfixede7, &c.Ctddfixeddfg, &c.Ctddpercelec, &c.Ctddpercgas, &c.Ctddperce7, &c.Ctddpercdfg, &c.Ctddthreshelec, &c.Ctddthreshgas, &c.Ctddthreshe7, &c.Ctddthreshdfg, &c.Ctdffixed, &c.Ctdfperc, &c.Ctgasperc, &c.Ctdualperc, &c.Ctgasbilll, &c.Ctgasbills, &c.Ctgasbillh, &c.Ctelecbilll, &c.Ctelecbills, &c.Ctelecbillh, &c.Cte7billl, &c.Cte7bills, &c.Cte7billh, &c.Ctdualbilll, &c.Ctdualbills, &c.Ctdualbillh, &c.Ctdualbille7l, &c.Ctdualbille7s, &c.Ctdualbille7h, &c.Cteleperc, &c.Ctn1, &c.Ctn2, &c.Ctn3, &c.Ctd1, &c.Ctd2, &c.Ctc1, &c.Ctc2, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
