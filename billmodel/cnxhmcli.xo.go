// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Cnxhmcli represents a row from 'equinox.cnxhmcli'.
type Cnxhmcli struct {
	Cnxhmcliunisys   sql.NullFloat64 `json:"cnxhmcliunisys"`   // cnxhmcliunisys
	Cnxhmclinumber   sql.NullString  `json:"cnxhmclinumber"`   // cnxhmclinumber
	Cnxhmcliservtype sql.NullString  `json:"cnxhmcliservtype"` // cnxhmcliservtype
	Cnxhmcliaction   sql.NullInt64   `json:"cnxhmcliaction"`   // cnxhmcliaction
	Cnxhmclitransfer sql.NullInt64   `json:"cnxhmclitransfer"` // cnxhmclitransfer
	Cnxhmcligodate   pq.NullTime     `json:"cnxhmcligodate"`   // cnxhmcligodate
	Cnxhmclisnum1    sql.NullString  `json:"cnxhmclisnum1"`    // cnxhmclisnum1
	Cnxhmclisnum2    sql.NullString  `json:"cnxhmclisnum2"`    // cnxhmclisnum2
	Cnxhmcliomr      sql.NullInt64   `json:"cnxhmcliomr"`      // cnxhmcliomr
	Cnxhmclipaymeth  sql.NullInt64   `json:"cnxhmclipaymeth"`  // cnxhmclipaymeth
	Cnxhmclipromocd  sql.NullString  `json:"cnxhmclipromocd"`  // cnxhmclipromocd
	Cnxhmclimnumber  sql.NullString  `json:"cnxhmclimnumber"`  // cnxhmclimnumber
	Cnxhmclibcname   sql.NullInt64   `json:"cnxhmclibcname"`   // cnxhmclibcname
	Cnxhmclibcplus   sql.NullInt64   `json:"cnxhmclibcplus"`   // cnxhmclibcplus
	Cnxhmcliwadaptor sql.NullInt64   `json:"cnxhmcliwadaptor"` // cnxhmcliwadaptor
	Cnxhmclinumwithh sql.NullInt64   `json:"cnxhmclinumwithh"` // cnxhmclinumwithh
	Cnxhmcliexdir    sql.NullInt64   `json:"cnxhmcliexdir"`    // cnxhmcliexdir
	Cnxhmclitps      sql.NullInt64   `json:"cnxhmclitps"`      // cnxhmclitps
	Cnxhmcliequip    sql.NullString  `json:"cnxhmcliequip"`    // cnxhmcliequip
	Cnxhmclimaccode  sql.NullString  `json:"cnxhmclimaccode"`  // cnxhmclimaccode
	Cnxhmclifilters  sql.NullInt64   `json:"cnxhmclifilters"`  // cnxhmclifilters
	Cnxhmcliss07     sql.NullInt64   `json:"cnxhmcliss07"`     // cnxhmcliss07
	Cnxhmcliss01     sql.NullInt64   `json:"cnxhmcliss01"`     // cnxhmcliss01
	Cnxhmcliss1a     sql.NullInt64   `json:"cnxhmcliss1a"`     // cnxhmcliss1a
	Cnxhmcliss08     sql.NullInt64   `json:"cnxhmcliss08"`     // cnxhmcliss08
	Cnxhmcliss06     sql.NullInt64   `json:"cnxhmcliss06"`     // cnxhmcliss06
	Cnxhmcliss09     sql.NullInt64   `json:"cnxhmcliss09"`     // cnxhmcliss09
	Cnxhmcliss02     sql.NullInt64   `json:"cnxhmcliss02"`     // cnxhmcliss02
	Cnxhmcliss05     sql.NullInt64   `json:"cnxhmcliss05"`     // cnxhmcliss05
	Cnxhmcliss14     sql.NullInt64   `json:"cnxhmcliss14"`     // cnxhmcliss14
	Cnxhmclideladd   sql.NullInt64   `json:"cnxhmclideladd"`   // cnxhmclideladd
	Cnxhmclipcode    sql.NullString  `json:"cnxhmclipcode"`    // cnxhmclipcode
	Cnxhmcliadd1     sql.NullString  `json:"cnxhmcliadd1"`     // cnxhmcliadd1
	Cnxhmcliadd2     sql.NullString  `json:"cnxhmcliadd2"`     // cnxhmcliadd2
	Cnxhmcliadd3     sql.NullString  `json:"cnxhmcliadd3"`     // cnxhmcliadd3
	Cnxhmcliadd4     sql.NullString  `json:"cnxhmcliadd4"`     // cnxhmcliadd4
	Cnxhmclicounty   sql.NullString  `json:"cnxhmclicounty"`   // cnxhmclicounty
	Cnxhmclifeaturel sql.NullInt64   `json:"cnxhmclifeaturel"` // cnxhmclifeaturel
	Cnxhmclicable    sql.NullInt64   `json:"cnxhmclicable"`    // cnxhmclicable
	Cnxhmclismartbox sql.NullInt64   `json:"cnxhmclismartbox"` // cnxhmclismartbox
	Cnxhmcliretainno sql.NullInt64   `json:"cnxhmcliretainno"` // cnxhmcliretainno
	Cnxhmclimobsaver sql.NullInt64   `json:"cnxhmclimobsaver"` // cnxhmclimobsaver
	Cnxhmcliintsaver sql.NullInt64   `json:"cnxhmcliintsaver"` // cnxhmcliintsaver
	Cnxhmcliservlvl  sql.NullString  `json:"cnxhmcliservlvl"`  // cnxhmcliservlvl
	Cnxhmclitermno   sql.NullString  `json:"cnxhmclitermno"`   // cnxhmclitermno
	Cnxhmclihandset  sql.NullString  `json:"cnxhmclihandset"`  // cnxhmclihandset
	Cnxhmclirpc      sql.NullInt64   `json:"cnxhmclirpc"`      // cnxhmclirpc
	Cnxhmclicontterm sql.NullInt64   `json:"cnxhmclicontterm"` // cnxhmclicontterm
	Cnxhmclivmail    sql.NullInt64   `json:"cnxhmclivmail"`    // cnxhmclivmail
	Cnxhmcliport     sql.NullInt64   `json:"cnxhmcliport"`     // cnxhmcliport
	Cnxhmclipaccode  sql.NullString  `json:"cnxhmclipaccode"`  // cnxhmclipaccode
	Cnxhmcliservatp1 sql.NullString  `json:"cnxhmcliservatp1"` // cnxhmcliservatp1
	Cnxhmclitariff   sql.NullString  `json:"cnxhmclitariff"`   // cnxhmclitariff
	Cnxhmclicontdate pq.NullTime     `json:"cnxhmclicontdate"` // cnxhmclicontdate
	Cnxhmclimonthly  sql.NullString  `json:"cnxhmclimonthly"`  // cnxhmclimonthly
	Cnxhmcliorgread1 sql.NullInt64   `json:"cnxhmcliorgread1"` // cnxhmcliorgread1
	Cnxhmcliorgread2 sql.NullInt64   `json:"cnxhmcliorgread2"` // cnxhmcliorgread2
	Cnxhmcliaddread  sql.NullInt64   `json:"cnxhmcliaddread"`  // cnxhmcliaddread
	Cnxhmclicomplete pq.NullTime     `json:"cnxhmclicomplete"` // cnxhmclicomplete
	Cnxhmclip2appdun pq.NullTime     `json:"cnxhmclip2appdun"` // cnxhmclip2appdun
	Cnxhmclispared3  pq.NullTime     `json:"cnxhmclispared3"`  // cnxhmclispared3
	Cnxhmclihp2hplr  sql.NullInt64   `json:"cnxhmclihp2hplr"`  // cnxhmclihp2hplr
	Cnxhmclihponly   sql.NullInt64   `json:"cnxhmclihponly"`   // cnxhmclihponly
	Cnxhmclibalarm   sql.NullInt64   `json:"cnxhmclibalarm"`   // cnxhmclibalarm
	Cnxhmcliforcedr  sql.NullInt64   `json:"cnxhmcliforcedr"`  // cnxhmcliforcedr
	Cnxhmclisparet1  pq.NullTime     `json:"cnxhmclisparet1"`  // cnxhmclisparet1
	Cnxhmcliomr2     sql.NullInt64   `json:"cnxhmcliomr2"`     // cnxhmcliomr2
	Cnxhmclimatchp1  sql.NullInt64   `json:"cnxhmclimatchp1"`  // cnxhmclimatchp1
	Cnxhmclicontadvi sql.NullInt64   `json:"cnxhmclicontadvi"` // cnxhmclicontadvi
	Cnxhmclicanretno sql.NullInt64   `json:"cnxhmclicanretno"` // cnxhmclicanretno
	Cnxhmcliintaddeq sql.NullString  `json:"cnxhmcliintaddeq"` // cnxhmcliintaddeq
	Cnxhmclinumatp2  sql.NullString  `json:"cnxhmclinumatp2"`  // cnxhmclinumatp2
	Cnxhmclic2reques pq.NullTime     `json:"cnxhmclic2reques"` // cnxhmclic2reques
	Cnxhmcliannual   sql.NullString  `json:"cnxhmcliannual"`   // cnxhmcliannual
	Cnxhmcliactionc2 sql.NullInt64   `json:"cnxhmcliactionc2"` // cnxhmcliactionc2
	Cnxhmselectservs sql.NullString  `json:"cnxhmselectservs"` // cnxhmselectservs
	Cnxhmclicni      sql.NullInt64   `json:"cnxhmclicni"`      // cnxhmclicni
	Cnxhmclicniphone sql.NullString  `json:"cnxhmclicniphone"` // cnxhmclicniphone
	Cnxhmclietf      sql.NullInt64   `json:"cnxhmclietf"`      // cnxhmclietf
	Cnxhmclietfterms sql.NullInt64   `json:"cnxhmclietfterms"` // cnxhmclietfterms
	Cnxhmclinewphone sql.NullString  `json:"cnxhmclinewphone"` // cnxhmclinewphone
	Cnxhmclibtcable  sql.NullString  `json:"cnxhmclibtcable"`  // cnxhmclibtcable
	Cnxhmclicps      sql.NullString  `json:"cnxhmclicps"`      // cnxhmclicps
	Cnxhmclipendnrg  sql.NullInt64   `json:"cnxhmclipendnrg"`  // cnxhmclipendnrg
	Cnxhmclisrvlvlc2 sql.NullString  `json:"cnxhmclisrvlvlc2"` // cnxhmclisrvlvlc2
	Cnxhmclikwhrs    sql.NullString  `json:"cnxhmclikwhrs"`    // cnxhmclikwhrs
	Cnxhmcliecosave  sql.NullInt64   `json:"cnxhmcliecosave"`  // cnxhmcliecosave
	Cnxhmclilinetype sql.NullString  `json:"cnxhmclilinetype"` // cnxhmclilinetype
	Cnxhmclicontract sql.NullString  `json:"cnxhmclicontract"` // cnxhmclicontract
	Cnxhmcliultraava sql.NullInt64   `json:"cnxhmcliultraava"` // cnxhmcliultraava
	Cnxhmclimobilebb sql.NullInt64   `json:"cnxhmclimobilebb"` // cnxhmclimobilebb
	Cnxhmclinewcli   sql.NullString  `json:"cnxhmclinewcli"`   // cnxhmclinewcli
	Cnxhmclinewuniq  sql.NullFloat64 `json:"cnxhmclinewuniq"`  // cnxhmclinewuniq
	Cnxhmcliomrtype  sql.NullString  `json:"cnxhmcliomrtype"`  // cnxhmcliomrtype
	Cnxhmcliprefer   sql.NullInt64   `json:"cnxhmcliprefer"`   // cnxhmcliprefer
	Cnxhmclipeaks    sql.NullInt64   `json:"cnxhmclipeaks"`    // cnxhmclipeaks
	Cnxhmcliintsave  sql.NullInt64   `json:"cnxhmcliintsave"`  // cnxhmcliintsave
	Cnxhmcliinstnote sql.NullInt64   `json:"cnxhmcliinstnote"` // cnxhmcliinstnote
	Cnxhmcliss20     sql.NullInt64   `json:"cnxhmcliss20"`     // cnxhmcliss20
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllCnxhmcli(db XODB, callback func(x Cnxhmcli) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`cnxhmcliunisys, cnxhmclinumber, cnxhmcliservtype, cnxhmcliaction, cnxhmclitransfer, cnxhmcligodate, cnxhmclisnum1, cnxhmclisnum2, cnxhmcliomr, cnxhmclipaymeth, cnxhmclipromocd, cnxhmclimnumber, cnxhmclibcname, cnxhmclibcplus, cnxhmcliwadaptor, cnxhmclinumwithh, cnxhmcliexdir, cnxhmclitps, cnxhmcliequip, cnxhmclimaccode, cnxhmclifilters, cnxhmcliss07, cnxhmcliss01, cnxhmcliss1a, cnxhmcliss08, cnxhmcliss06, cnxhmcliss09, cnxhmcliss02, cnxhmcliss05, cnxhmcliss14, cnxhmclideladd, cnxhmclipcode, cnxhmcliadd1, cnxhmcliadd2, cnxhmcliadd3, cnxhmcliadd4, cnxhmclicounty, cnxhmclifeaturel, cnxhmclicable, cnxhmclismartbox, cnxhmcliretainno, cnxhmclimobsaver, cnxhmcliintsaver, cnxhmcliservlvl, cnxhmclitermno, cnxhmclihandset, cnxhmclirpc, cnxhmclicontterm, cnxhmclivmail, cnxhmcliport, cnxhmclipaccode, cnxhmcliservatp1, cnxhmclitariff, cnxhmclicontdate, cnxhmclimonthly, cnxhmcliorgread1, cnxhmcliorgread2, cnxhmcliaddread, cnxhmclicomplete, cnxhmclip2appdun, cnxhmclispared3, cnxhmclihp2hplr, cnxhmclihponly, cnxhmclibalarm, cnxhmcliforcedr, cnxhmclisparet1, cnxhmcliomr2, cnxhmclimatchp1, cnxhmclicontadvi, cnxhmclicanretno, cnxhmcliintaddeq, cnxhmclinumatp2, cnxhmclic2reques, cnxhmcliannual, cnxhmcliactionc2, cnxhmselectservs, cnxhmclicni, cnxhmclicniphone, cnxhmclietf, cnxhmclietfterms, cnxhmclinewphone, cnxhmclibtcable, cnxhmclicps, cnxhmclipendnrg, cnxhmclisrvlvlc2, cnxhmclikwhrs, cnxhmcliecosave, cnxhmclilinetype, cnxhmclicontract, cnxhmcliultraava, cnxhmclimobilebb, cnxhmclinewcli, cnxhmclinewuniq, cnxhmcliomrtype, cnxhmcliprefer, cnxhmclipeaks, cnxhmcliintsave, cnxhmcliinstnote, cnxhmcliss20, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxhmcli `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Cnxhmcli{}

		// scan
		err = q.Scan(&c.Cnxhmcliunisys, &c.Cnxhmclinumber, &c.Cnxhmcliservtype, &c.Cnxhmcliaction, &c.Cnxhmclitransfer, &c.Cnxhmcligodate, &c.Cnxhmclisnum1, &c.Cnxhmclisnum2, &c.Cnxhmcliomr, &c.Cnxhmclipaymeth, &c.Cnxhmclipromocd, &c.Cnxhmclimnumber, &c.Cnxhmclibcname, &c.Cnxhmclibcplus, &c.Cnxhmcliwadaptor, &c.Cnxhmclinumwithh, &c.Cnxhmcliexdir, &c.Cnxhmclitps, &c.Cnxhmcliequip, &c.Cnxhmclimaccode, &c.Cnxhmclifilters, &c.Cnxhmcliss07, &c.Cnxhmcliss01, &c.Cnxhmcliss1a, &c.Cnxhmcliss08, &c.Cnxhmcliss06, &c.Cnxhmcliss09, &c.Cnxhmcliss02, &c.Cnxhmcliss05, &c.Cnxhmcliss14, &c.Cnxhmclideladd, &c.Cnxhmclipcode, &c.Cnxhmcliadd1, &c.Cnxhmcliadd2, &c.Cnxhmcliadd3, &c.Cnxhmcliadd4, &c.Cnxhmclicounty, &c.Cnxhmclifeaturel, &c.Cnxhmclicable, &c.Cnxhmclismartbox, &c.Cnxhmcliretainno, &c.Cnxhmclimobsaver, &c.Cnxhmcliintsaver, &c.Cnxhmcliservlvl, &c.Cnxhmclitermno, &c.Cnxhmclihandset, &c.Cnxhmclirpc, &c.Cnxhmclicontterm, &c.Cnxhmclivmail, &c.Cnxhmcliport, &c.Cnxhmclipaccode, &c.Cnxhmcliservatp1, &c.Cnxhmclitariff, &c.Cnxhmclicontdate, &c.Cnxhmclimonthly, &c.Cnxhmcliorgread1, &c.Cnxhmcliorgread2, &c.Cnxhmcliaddread, &c.Cnxhmclicomplete, &c.Cnxhmclip2appdun, &c.Cnxhmclispared3, &c.Cnxhmclihp2hplr, &c.Cnxhmclihponly, &c.Cnxhmclibalarm, &c.Cnxhmcliforcedr, &c.Cnxhmclisparet1, &c.Cnxhmcliomr2, &c.Cnxhmclimatchp1, &c.Cnxhmclicontadvi, &c.Cnxhmclicanretno, &c.Cnxhmcliintaddeq, &c.Cnxhmclinumatp2, &c.Cnxhmclic2reques, &c.Cnxhmcliannual, &c.Cnxhmcliactionc2, &c.Cnxhmselectservs, &c.Cnxhmclicni, &c.Cnxhmclicniphone, &c.Cnxhmclietf, &c.Cnxhmclietfterms, &c.Cnxhmclinewphone, &c.Cnxhmclibtcable, &c.Cnxhmclicps, &c.Cnxhmclipendnrg, &c.Cnxhmclisrvlvlc2, &c.Cnxhmclikwhrs, &c.Cnxhmcliecosave, &c.Cnxhmclilinetype, &c.Cnxhmclicontract, &c.Cnxhmcliultraava, &c.Cnxhmclimobilebb, &c.Cnxhmclinewcli, &c.Cnxhmclinewuniq, &c.Cnxhmcliomrtype, &c.Cnxhmcliprefer, &c.Cnxhmclipeaks, &c.Cnxhmcliintsave, &c.Cnxhmcliinstnote, &c.Cnxhmcliss20, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CnxhmcliByEquinoxLrn retrieves a row from 'equinox.cnxhmcli' as a Cnxhmcli.
//
// Generated from index 'cnxhmcli_pkey'.
func CnxhmcliByEquinoxLrn(db XODB, equinoxLrn int64) (*Cnxhmcli, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cnxhmcliunisys, cnxhmclinumber, cnxhmcliservtype, cnxhmcliaction, cnxhmclitransfer, cnxhmcligodate, cnxhmclisnum1, cnxhmclisnum2, cnxhmcliomr, cnxhmclipaymeth, cnxhmclipromocd, cnxhmclimnumber, cnxhmclibcname, cnxhmclibcplus, cnxhmcliwadaptor, cnxhmclinumwithh, cnxhmcliexdir, cnxhmclitps, cnxhmcliequip, cnxhmclimaccode, cnxhmclifilters, cnxhmcliss07, cnxhmcliss01, cnxhmcliss1a, cnxhmcliss08, cnxhmcliss06, cnxhmcliss09, cnxhmcliss02, cnxhmcliss05, cnxhmcliss14, cnxhmclideladd, cnxhmclipcode, cnxhmcliadd1, cnxhmcliadd2, cnxhmcliadd3, cnxhmcliadd4, cnxhmclicounty, cnxhmclifeaturel, cnxhmclicable, cnxhmclismartbox, cnxhmcliretainno, cnxhmclimobsaver, cnxhmcliintsaver, cnxhmcliservlvl, cnxhmclitermno, cnxhmclihandset, cnxhmclirpc, cnxhmclicontterm, cnxhmclivmail, cnxhmcliport, cnxhmclipaccode, cnxhmcliservatp1, cnxhmclitariff, cnxhmclicontdate, cnxhmclimonthly, cnxhmcliorgread1, cnxhmcliorgread2, cnxhmcliaddread, cnxhmclicomplete, cnxhmclip2appdun, cnxhmclispared3, cnxhmclihp2hplr, cnxhmclihponly, cnxhmclibalarm, cnxhmcliforcedr, cnxhmclisparet1, cnxhmcliomr2, cnxhmclimatchp1, cnxhmclicontadvi, cnxhmclicanretno, cnxhmcliintaddeq, cnxhmclinumatp2, cnxhmclic2reques, cnxhmcliannual, cnxhmcliactionc2, cnxhmselectservs, cnxhmclicni, cnxhmclicniphone, cnxhmclietf, cnxhmclietfterms, cnxhmclinewphone, cnxhmclibtcable, cnxhmclicps, cnxhmclipendnrg, cnxhmclisrvlvlc2, cnxhmclikwhrs, cnxhmcliecosave, cnxhmclilinetype, cnxhmclicontract, cnxhmcliultraava, cnxhmclimobilebb, cnxhmclinewcli, cnxhmclinewuniq, cnxhmcliomrtype, cnxhmcliprefer, cnxhmclipeaks, cnxhmcliintsave, cnxhmcliinstnote, cnxhmcliss20, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxhmcli ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cnxhmcli{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cnxhmcliunisys, &c.Cnxhmclinumber, &c.Cnxhmcliservtype, &c.Cnxhmcliaction, &c.Cnxhmclitransfer, &c.Cnxhmcligodate, &c.Cnxhmclisnum1, &c.Cnxhmclisnum2, &c.Cnxhmcliomr, &c.Cnxhmclipaymeth, &c.Cnxhmclipromocd, &c.Cnxhmclimnumber, &c.Cnxhmclibcname, &c.Cnxhmclibcplus, &c.Cnxhmcliwadaptor, &c.Cnxhmclinumwithh, &c.Cnxhmcliexdir, &c.Cnxhmclitps, &c.Cnxhmcliequip, &c.Cnxhmclimaccode, &c.Cnxhmclifilters, &c.Cnxhmcliss07, &c.Cnxhmcliss01, &c.Cnxhmcliss1a, &c.Cnxhmcliss08, &c.Cnxhmcliss06, &c.Cnxhmcliss09, &c.Cnxhmcliss02, &c.Cnxhmcliss05, &c.Cnxhmcliss14, &c.Cnxhmclideladd, &c.Cnxhmclipcode, &c.Cnxhmcliadd1, &c.Cnxhmcliadd2, &c.Cnxhmcliadd3, &c.Cnxhmcliadd4, &c.Cnxhmclicounty, &c.Cnxhmclifeaturel, &c.Cnxhmclicable, &c.Cnxhmclismartbox, &c.Cnxhmcliretainno, &c.Cnxhmclimobsaver, &c.Cnxhmcliintsaver, &c.Cnxhmcliservlvl, &c.Cnxhmclitermno, &c.Cnxhmclihandset, &c.Cnxhmclirpc, &c.Cnxhmclicontterm, &c.Cnxhmclivmail, &c.Cnxhmcliport, &c.Cnxhmclipaccode, &c.Cnxhmcliservatp1, &c.Cnxhmclitariff, &c.Cnxhmclicontdate, &c.Cnxhmclimonthly, &c.Cnxhmcliorgread1, &c.Cnxhmcliorgread2, &c.Cnxhmcliaddread, &c.Cnxhmclicomplete, &c.Cnxhmclip2appdun, &c.Cnxhmclispared3, &c.Cnxhmclihp2hplr, &c.Cnxhmclihponly, &c.Cnxhmclibalarm, &c.Cnxhmcliforcedr, &c.Cnxhmclisparet1, &c.Cnxhmcliomr2, &c.Cnxhmclimatchp1, &c.Cnxhmclicontadvi, &c.Cnxhmclicanretno, &c.Cnxhmcliintaddeq, &c.Cnxhmclinumatp2, &c.Cnxhmclic2reques, &c.Cnxhmcliannual, &c.Cnxhmcliactionc2, &c.Cnxhmselectservs, &c.Cnxhmclicni, &c.Cnxhmclicniphone, &c.Cnxhmclietf, &c.Cnxhmclietfterms, &c.Cnxhmclinewphone, &c.Cnxhmclibtcable, &c.Cnxhmclicps, &c.Cnxhmclipendnrg, &c.Cnxhmclisrvlvlc2, &c.Cnxhmclikwhrs, &c.Cnxhmcliecosave, &c.Cnxhmclilinetype, &c.Cnxhmclicontract, &c.Cnxhmcliultraava, &c.Cnxhmclimobilebb, &c.Cnxhmclinewcli, &c.Cnxhmclinewuniq, &c.Cnxhmcliomrtype, &c.Cnxhmcliprefer, &c.Cnxhmclipeaks, &c.Cnxhmcliintsave, &c.Cnxhmcliinstnote, &c.Cnxhmcliss20, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
