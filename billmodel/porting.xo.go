// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Porting represents a row from 'equinox.porting'.
type Porting struct {
	Porttitle        sql.NullString `json:"porttitle"`        // porttitle
	Portinits        sql.NullString `json:"portinits"`        // portinits
	Portsurname      sql.NullString `json:"portsurname"`      // portsurname
	Portname         sql.NullString `json:"portname"`         // portname
	Portadd1         sql.NullString `json:"portadd1"`         // portadd1
	Portadd2         sql.NullString `json:"portadd2"`         // portadd2
	Portadd3         sql.NullString `json:"portadd3"`         // portadd3
	Portadd4         sql.NullString `json:"portadd4"`         // portadd4
	Portadd5         sql.NullString `json:"portadd5"`         // portadd5
	Portpostcode     sql.NullString `json:"portpostcode"`     // portpostcode
	Portexecid       sql.NullString `json:"portexecid"`       // portexecid
	Portbusname      sql.NullString `json:"portbusname"`      // portbusname
	Portaccountno    sql.NullString `json:"portaccountno"`    // portaccountno
	Portreference    sql.NullString `json:"portreference"`    // portreference
	Portday1         pq.NullTime    `json:"portday1"`         // portday1
	Portexpect       pq.NullTime    `json:"portexpect"`       // portexpect
	Portday          sql.NullInt64  `json:"portday"`          // portday
	Portusrdate      pq.NullTime    `json:"portusrdate"`      // portusrdate
	Portusreject1    sql.NullInt64  `json:"portusreject1"`    // portusreject1
	Portusreject2    sql.NullInt64  `json:"portusreject2"`    // portusreject2
	Portusreject3    sql.NullInt64  `json:"portusreject3"`    // portusreject3
	Portusreject4    sql.NullInt64  `json:"portusreject4"`    // portusreject4
	Portusreject5    sql.NullInt64  `json:"portusreject5"`    // portusreject5
	Portusreject6    sql.NullInt64  `json:"portusreject6"`    // portusreject6
	Portusreject7    sql.NullInt64  `json:"portusreject7"`    // portusreject7
	Portusreject8    sql.NullInt64  `json:"portusreject8"`    // portusreject8
	Portstatus       sql.NullString `json:"portstatus"`       // portstatus
	Portdsp          sql.NullString `json:"portdsp"`          // portdsp
	Portdno          sql.NullString `json:"portdno"`          // portdno
	Portdsprequest   pq.NullTime    `json:"portdsprequest"`   // portdsprequest
	Portdspacknow    pq.NullTime    `json:"portdspacknow"`    // portdspacknow
	Portdspdecision  pq.NullTime    `json:"portdspdecision"`  // portdspdecision
	Portdsprejecttoc pq.NullTime    `json:"portdsprejecttoc"` // portdsprejecttoc
	Portwithdraw     pq.NullTime    `json:"portwithdraw"`     // portwithdraw
	Portwithdrawall  pq.NullTime    `json:"portwithdrawall"`  // portwithdrawall
	Portdspwithdraw  pq.NullTime    `json:"portdspwithdraw"`  // portdspwithdraw
	Portmanual       sql.NullInt64  `json:"portmanual"`       // portmanual
	Porttelephone    sql.NullInt64  `json:"porttelephone"`    // porttelephone
	Paperrec         sql.NullInt64  `json:"paperrec"`         // paperrec
	Portnoreturn     pq.NullTime    `json:"portnoreturn"`     // portnoreturn
	Portready        sql.NullInt64  `json:"portready"`        // portready
	Portcomplete     pq.NullTime    `json:"portcomplete"`     // portcomplete
	Portconnectacc   sql.NullString `json:"portconnectacc"`   // portconnectacc
	Portdouble       sql.NullInt64  `json:"portdouble"`       // portdouble
	Porttransferee   sql.NullString `json:"porttransferee"`   // porttransferee
	Portwelcome      pq.NullTime    `json:"portwelcome"`      // portwelcome
	Portcoverage     sql.NullInt64  `json:"portcoverage"`     // portcoverage
	Portstart        pq.NullTime    `json:"portstart"`        // portstart
	Portpac          sql.NullString `json:"portpac"`          // portpac
	Portpacexpire    pq.NullTime    `json:"portpacexpire"`    // portpacexpire
	Porttagged       pq.NullTime    `json:"porttagged"`       // porttagged
	Portlocked       pq.NullTime    `json:"portlocked"`       // portlocked
	Port00263        pq.NullTime    `json:"port00263"`        // port00263
	Porttype         sql.NullString `json:"porttype"`         // porttype
	Portedit         sql.NullInt64  `json:"portedit"`         // portedit
	Porttml          sql.NullInt64  `json:"porttml"`          // porttml
	EquinoxLrn       int64          `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64  `json:"equinox_sec"`      // equinox_sec
}

func AllPorting(db XODB, callback func(x Porting) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`porttitle, portinits, portsurname, portname, portadd1, portadd2, portadd3, portadd4, portadd5, portpostcode, portexecid, portbusname, portaccountno, portreference, portday1, portexpect, portday, portusrdate, portusreject1, portusreject2, portusreject3, portusreject4, portusreject5, portusreject6, portusreject7, portusreject8, portstatus, portdsp, portdno, portdsprequest, portdspacknow, portdspdecision, portdsprejecttoc, portwithdraw, portwithdrawall, portdspwithdraw, portmanual, porttelephone, paperrec, portnoreturn, portready, portcomplete, portconnectacc, portdouble, porttransferee, portwelcome, portcoverage, portstart, portpac, portpacexpire, porttagged, portlocked, port00263, porttype, portedit, porttml, equinox_lrn, equinox_sec ` +
		`FROM equinox.porting `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		p := Porting{}

		// scan
		err = q.Scan(&p.Porttitle, &p.Portinits, &p.Portsurname, &p.Portname, &p.Portadd1, &p.Portadd2, &p.Portadd3, &p.Portadd4, &p.Portadd5, &p.Portpostcode, &p.Portexecid, &p.Portbusname, &p.Portaccountno, &p.Portreference, &p.Portday1, &p.Portexpect, &p.Portday, &p.Portusrdate, &p.Portusreject1, &p.Portusreject2, &p.Portusreject3, &p.Portusreject4, &p.Portusreject5, &p.Portusreject6, &p.Portusreject7, &p.Portusreject8, &p.Portstatus, &p.Portdsp, &p.Portdno, &p.Portdsprequest, &p.Portdspacknow, &p.Portdspdecision, &p.Portdsprejecttoc, &p.Portwithdraw, &p.Portwithdrawall, &p.Portdspwithdraw, &p.Portmanual, &p.Porttelephone, &p.Paperrec, &p.Portnoreturn, &p.Portready, &p.Portcomplete, &p.Portconnectacc, &p.Portdouble, &p.Porttransferee, &p.Portwelcome, &p.Portcoverage, &p.Portstart, &p.Portpac, &p.Portpacexpire, &p.Porttagged, &p.Portlocked, &p.Port00263, &p.Porttype, &p.Portedit, &p.Porttml, &p.EquinoxLrn, &p.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(p) {
			return nil
		}
	}

	return nil
}

// PortingByEquinoxLrn retrieves a row from 'equinox.porting' as a Porting.
//
// Generated from index 'porting_pkey'.
func PortingByEquinoxLrn(db XODB, equinoxLrn int64) (*Porting, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`porttitle, portinits, portsurname, portname, portadd1, portadd2, portadd3, portadd4, portadd5, portpostcode, portexecid, portbusname, portaccountno, portreference, portday1, portexpect, portday, portusrdate, portusreject1, portusreject2, portusreject3, portusreject4, portusreject5, portusreject6, portusreject7, portusreject8, portstatus, portdsp, portdno, portdsprequest, portdspacknow, portdspdecision, portdsprejecttoc, portwithdraw, portwithdrawall, portdspwithdraw, portmanual, porttelephone, paperrec, portnoreturn, portready, portcomplete, portconnectacc, portdouble, porttransferee, portwelcome, portcoverage, portstart, portpac, portpacexpire, porttagged, portlocked, port00263, porttype, portedit, porttml, equinox_lrn, equinox_sec ` +
		`FROM equinox.porting ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	p := Porting{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&p.Porttitle, &p.Portinits, &p.Portsurname, &p.Portname, &p.Portadd1, &p.Portadd2, &p.Portadd3, &p.Portadd4, &p.Portadd5, &p.Portpostcode, &p.Portexecid, &p.Portbusname, &p.Portaccountno, &p.Portreference, &p.Portday1, &p.Portexpect, &p.Portday, &p.Portusrdate, &p.Portusreject1, &p.Portusreject2, &p.Portusreject3, &p.Portusreject4, &p.Portusreject5, &p.Portusreject6, &p.Portusreject7, &p.Portusreject8, &p.Portstatus, &p.Portdsp, &p.Portdno, &p.Portdsprequest, &p.Portdspacknow, &p.Portdspdecision, &p.Portdsprejecttoc, &p.Portwithdraw, &p.Portwithdrawall, &p.Portdspwithdraw, &p.Portmanual, &p.Porttelephone, &p.Paperrec, &p.Portnoreturn, &p.Portready, &p.Portcomplete, &p.Portconnectacc, &p.Portdouble, &p.Porttransferee, &p.Portwelcome, &p.Portcoverage, &p.Portstart, &p.Portpac, &p.Portpacexpire, &p.Porttagged, &p.Portlocked, &p.Port00263, &p.Porttype, &p.Portedit, &p.Porttml, &p.EquinoxLrn, &p.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
