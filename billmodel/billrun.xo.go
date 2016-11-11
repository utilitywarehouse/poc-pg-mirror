// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Billrun represents a row from 'equinox.billruns'.
type Billrun struct {
	Brnumber         sql.NullFloat64 `json:"brnumber"`         // brnumber
	Brperiod         sql.NullString  `json:"brperiod"`         // brperiod
	Brtype           sql.NullInt64   `json:"brtype"`           // brtype
	Brtypetxt        sql.NullString  `json:"brtypetxt"`        // brtypetxt
	Brrundate        pq.NullTime     `json:"brrundate"`        // brrundate
	Brrunstarttime   pq.NullTime     `json:"brrunstarttime"`   // brrunstarttime
	Brrunenddate     pq.NullTime     `json:"brrunenddate"`     // brrunenddate
	Brrunendtime     pq.NullTime     `json:"brrunendtime"`     // brrunendtime
	Brrunby          sql.NullString  `json:"brrunby"`          // brrunby
	Brcomplete       sql.NullInt64   `json:"brcomplete"`       // brcomplete
	Brcloseddate     pq.NullTime     `json:"brcloseddate"`     // brcloseddate
	Brclosedby       sql.NullString  `json:"brclosedby"`       // brclosedby
	Brlastbilldate   pq.NullTime     `json:"brlastbilldate"`   // brlastbilldate
	Brmonthstobill   sql.NullInt64   `json:"brmonthstobill"`   // brmonthstobill
	Brlowestbillno   sql.NullFloat64 `json:"brlowestbillno"`   // brlowestbillno
	Brhighestbillno  sql.NullFloat64 `json:"brhighestbillno"`  // brhighestbillno
	Brnoofbills      sql.NullFloat64 `json:"brnoofbills"`      // brnoofbills
	Brprevos         sql.NullFloat64 `json:"brprevos"`         // brprevos
	Brwholesale      sql.NullFloat64 `json:"brwholesale"`      // brwholesale
	Brretail         sql.NullFloat64 `json:"brretail"`         // brretail
	Brretail2        sql.NullFloat64 `json:"brretail2"`        // brretail2
	Broneoff         sql.NullFloat64 `json:"broneoff"`         // broneoff
	Brgreendeal      sql.NullFloat64 `json:"brgreendeal"`      // brgreendeal
	Brmonthly        sql.NullFloat64 `json:"brmonthly"`        // brmonthly
	Brmonthlyw       sql.NullFloat64 `json:"brmonthlyw"`       // brmonthlyw
	Brcreditall      sql.NullFloat64 `json:"brcreditall"`      // brcreditall
	Brcreditunits    sql.NullFloat64 `json:"brcreditunits"`    // brcreditunits
	Brdebitall       sql.NullFloat64 `json:"brdebitall"`       // brdebitall
	Brvattp          sql.NullFloat64 `json:"brvattp"`          // brvattp
	Brvatgp          sql.NullFloat64 `json:"brvatgp"`          // brvatgp
	Brvatep          sql.NullFloat64 `json:"brvatep"`          // brvatep
	Brnettp          sql.NullFloat64 `json:"brnettp"`          // brnettp
	Brnetgp          sql.NullFloat64 `json:"brnetgp"`          // brnetgp
	Brnetep          sql.NullFloat64 `json:"brnetep"`          // brnetep
	Brdiscount       sql.NullFloat64 `json:"brdiscount"`       // brdiscount
	Brsurcharge      sql.NullFloat64 `json:"brsurcharge"`      // brsurcharge
	Brclubfee        sql.NullFloat64 `json:"brclubfee"`        // brclubfee
	Brclubdisc       sql.NullFloat64 `json:"brclubdisc"`       // brclubdisc
	Brtotal          sql.NullFloat64 `json:"brtotal"`          // brtotal
	Brcgadiscount    sql.NullFloat64 `json:"brcgadiscount"`    // brcgadiscount
	Brchgecallcount  sql.NullFloat64 `json:"brchgecallcount"`  // brchgecallcount
	Brchgecallsecs   sql.NullFloat64 `json:"brchgecallsecs"`   // brchgecallsecs
	Brnoofcreditnote sql.NullFloat64 `json:"brnoofcreditnote"` // brnoofcreditnote
	Brcardcashbak    sql.NullFloat64 `json:"brcardcashbak"`    // brcardcashbak
	Braffcashbak     sql.NullFloat64 `json:"braffcashbak"`     // braffcashbak
	Brbydd           sql.NullFloat64 `json:"brbydd"`           // brbydd
	Brbydd1          sql.NullFloat64 `json:"brbydd1"`          // brbydd1
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllBillrun(db XODB, callback func(x Billrun) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`brnumber, brperiod, brtype, brtypetxt, brrundate, brrunstarttime, brrunenddate, brrunendtime, brrunby, brcomplete, brcloseddate, brclosedby, brlastbilldate, brmonthstobill, brlowestbillno, brhighestbillno, brnoofbills, brprevos, brwholesale, brretail, brretail2, broneoff, brgreendeal, brmonthly, brmonthlyw, brcreditall, brcreditunits, brdebitall, brvattp, brvatgp, brvatep, brnettp, brnetgp, brnetep, brdiscount, brsurcharge, brclubfee, brclubdisc, brtotal, brcgadiscount, brchgecallcount, brchgecallsecs, brnoofcreditnote, brcardcashbak, braffcashbak, brbydd, brbydd1, equinox_lrn, equinox_sec ` +
		`FROM equinox.billruns `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		b := Billrun{}

		// scan
		err = q.Scan(&b.Brnumber, &b.Brperiod, &b.Brtype, &b.Brtypetxt, &b.Brrundate, &b.Brrunstarttime, &b.Brrunenddate, &b.Brrunendtime, &b.Brrunby, &b.Brcomplete, &b.Brcloseddate, &b.Brclosedby, &b.Brlastbilldate, &b.Brmonthstobill, &b.Brlowestbillno, &b.Brhighestbillno, &b.Brnoofbills, &b.Brprevos, &b.Brwholesale, &b.Brretail, &b.Brretail2, &b.Broneoff, &b.Brgreendeal, &b.Brmonthly, &b.Brmonthlyw, &b.Brcreditall, &b.Brcreditunits, &b.Brdebitall, &b.Brvattp, &b.Brvatgp, &b.Brvatep, &b.Brnettp, &b.Brnetgp, &b.Brnetep, &b.Brdiscount, &b.Brsurcharge, &b.Brclubfee, &b.Brclubdisc, &b.Brtotal, &b.Brcgadiscount, &b.Brchgecallcount, &b.Brchgecallsecs, &b.Brnoofcreditnote, &b.Brcardcashbak, &b.Braffcashbak, &b.Brbydd, &b.Brbydd1, &b.EquinoxLrn, &b.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(b) {
			return nil
		}
	}

	return nil
}

// BillrunByEquinoxLrn retrieves a row from 'equinox.billruns' as a Billrun.
//
// Generated from index 'billruns_pkey'.
func BillrunByEquinoxLrn(db XODB, equinoxLrn int64) (*Billrun, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`brnumber, brperiod, brtype, brtypetxt, brrundate, brrunstarttime, brrunenddate, brrunendtime, brrunby, brcomplete, brcloseddate, brclosedby, brlastbilldate, brmonthstobill, brlowestbillno, brhighestbillno, brnoofbills, brprevos, brwholesale, brretail, brretail2, broneoff, brgreendeal, brmonthly, brmonthlyw, brcreditall, brcreditunits, brdebitall, brvattp, brvatgp, brvatep, brnettp, brnetgp, brnetep, brdiscount, brsurcharge, brclubfee, brclubdisc, brtotal, brcgadiscount, brchgecallcount, brchgecallsecs, brnoofcreditnote, brcardcashbak, braffcashbak, brbydd, brbydd1, equinox_lrn, equinox_sec ` +
		`FROM equinox.billruns ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	b := Billrun{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&b.Brnumber, &b.Brperiod, &b.Brtype, &b.Brtypetxt, &b.Brrundate, &b.Brrunstarttime, &b.Brrunenddate, &b.Brrunendtime, &b.Brrunby, &b.Brcomplete, &b.Brcloseddate, &b.Brclosedby, &b.Brlastbilldate, &b.Brmonthstobill, &b.Brlowestbillno, &b.Brhighestbillno, &b.Brnoofbills, &b.Brprevos, &b.Brwholesale, &b.Brretail, &b.Brretail2, &b.Broneoff, &b.Brgreendeal, &b.Brmonthly, &b.Brmonthlyw, &b.Brcreditall, &b.Brcreditunits, &b.Brdebitall, &b.Brvattp, &b.Brvatgp, &b.Brvatep, &b.Brnettp, &b.Brnetgp, &b.Brnetep, &b.Brdiscount, &b.Brsurcharge, &b.Brclubfee, &b.Brclubdisc, &b.Brtotal, &b.Brcgadiscount, &b.Brchgecallcount, &b.Brchgecallsecs, &b.Brnoofcreditnote, &b.Brcardcashbak, &b.Braffcashbak, &b.Brbydd, &b.Brbydd1, &b.EquinoxLrn, &b.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
