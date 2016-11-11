// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Cnxhm represents a row from 'equinox.cnxhm'.
type Cnxhm struct {
	Cnxhmc2dob2      pq.NullTime     `json:"cnxhmc2dob2"`      // cnxhmc2dob2
	Cnxhmenteredonc1 pq.NullTime     `json:"cnxhmenteredonc1"` // cnxhmenteredonc1
	Cnxhmstarted     pq.NullTime     `json:"cnxhmstarted"`     // cnxhmstarted
	Cnxhmcompleted   pq.NullTime     `json:"cnxhmcompleted"`   // cnxhmcompleted
	Cnxhmaccountno   sql.NullString  `json:"cnxhmaccountno"`   // cnxhmaccountno
	Cnxhmcustname    sql.NullString  `json:"cnxhmcustname"`    // cnxhmcustname
	Cnxhmbillgroup   sql.NullString  `json:"cnxhmbillgroup"`   // cnxhmbillgroup
	Cnxhmcusttitle   sql.NullString  `json:"cnxhmcusttitle"`   // cnxhmcusttitle
	Cnxhmcustiniti   sql.NullString  `json:"cnxhmcustiniti"`   // cnxhmcustiniti
	Cnxhmcustfirstn  sql.NullString  `json:"cnxhmcustfirstn"`  // cnxhmcustfirstn
	Cnxhmcustsurnam  sql.NullString  `json:"cnxhmcustsurnam"`  // cnxhmcustsurnam
	Cnxhmcustadd1    sql.NullString  `json:"cnxhmcustadd1"`    // cnxhmcustadd1
	Cnxhmcustadd2    sql.NullString  `json:"cnxhmcustadd2"`    // cnxhmcustadd2
	Cnxhmcustadd3    sql.NullString  `json:"cnxhmcustadd3"`    // cnxhmcustadd3
	Cnxhmcustadd4    sql.NullString  `json:"cnxhmcustadd4"`    // cnxhmcustadd4
	Cnxhmcustcounty  sql.NullString  `json:"cnxhmcustcounty"`  // cnxhmcustcounty
	Cnxhmcustpcode   sql.NullString  `json:"cnxhmcustpcode"`   // cnxhmcustpcode
	Cnxhmcustphone   sql.NullString  `json:"cnxhmcustphone"`   // cnxhmcustphone
	Cnxhmcustmobile  sql.NullString  `json:"cnxhmcustmobile"`  // cnxhmcustmobile
	Cnxhmdateofbirth pq.NullTime     `json:"cnxhmdateofbirth"` // cnxhmdateofbirth
	Cnxhmcustemail   sql.NullString  `json:"cnxhmcustemail"`   // cnxhmcustemail
	Cnxhmemailbills  sql.NullInt64   `json:"cnxhmemailbills"`  // cnxhmemailbills
	Cnxhmmoveindate  pq.NullTime     `json:"cnxhmmoveindate"`  // cnxhmmoveindate
	Cnxhmmoveoutdate pq.NullTime     `json:"cnxhmmoveoutdate"` // cnxhmmoveoutdate
	Cnxhmownertenant sql.NullString  `json:"cnxhmownertenant"` // cnxhmownertenant
	Cnxhmtenancyend  pq.NullTime     `json:"cnxhmtenancyend"`  // cnxhmtenancyend
	Cnxhmnewtitle    sql.NullString  `json:"cnxhmnewtitle"`    // cnxhmnewtitle
	Cnxhmnewfirstnam sql.NullString  `json:"cnxhmnewfirstnam"` // cnxhmnewfirstnam
	Cnxhmnewsurname  sql.NullString  `json:"cnxhmnewsurname"`  // cnxhmnewsurname
	Cnxhmnewadd1     sql.NullString  `json:"cnxhmnewadd1"`     // cnxhmnewadd1
	Cnxhmnewadd2     sql.NullString  `json:"cnxhmnewadd2"`     // cnxhmnewadd2
	Cnxhmnewadd3     sql.NullString  `json:"cnxhmnewadd3"`     // cnxhmnewadd3
	Cnxhmnewadd4     sql.NullString  `json:"cnxhmnewadd4"`     // cnxhmnewadd4
	Cnxhmnewcounty   sql.NullString  `json:"cnxhmnewcounty"`   // cnxhmnewcounty
	Cnxhmnewpostcode sql.NullString  `json:"cnxhmnewpostcode"` // cnxhmnewpostcode
	Cnxhmcurrtitle   sql.NullString  `json:"cnxhmcurrtitle"`   // cnxhmcurrtitle
	Cnxhmcurrfirstn  sql.NullString  `json:"cnxhmcurrfirstn"`  // cnxhmcurrfirstn
	Cnxhmcurrsurname sql.NullString  `json:"cnxhmcurrsurname"` // cnxhmcurrsurname
	Cnxhminvtitle    sql.NullString  `json:"cnxhminvtitle"`    // cnxhminvtitle
	Cnxhminvinitials sql.NullString  `json:"cnxhminvinitials"` // cnxhminvinitials
	Cnxhminvfirstn   sql.NullString  `json:"cnxhminvfirstn"`   // cnxhminvfirstn
	Cnxhminvsurname  sql.NullString  `json:"cnxhminvsurname"`  // cnxhminvsurname
	Cnxhminvadd1     sql.NullString  `json:"cnxhminvadd1"`     // cnxhminvadd1
	Cnxhminvadd2     sql.NullString  `json:"cnxhminvadd2"`     // cnxhminvadd2
	Cnxhminvadd3     sql.NullString  `json:"cnxhminvadd3"`     // cnxhminvadd3
	Cnxhminvadd4     sql.NullString  `json:"cnxhminvadd4"`     // cnxhminvadd4
	Cnxhminvcounty   sql.NullString  `json:"cnxhminvcounty"`   // cnxhminvcounty
	Cnxhminvpostcode sql.NullString  `json:"cnxhminvpostcode"` // cnxhminvpostcode
	Cnxhminvphone    sql.NullString  `json:"cnxhminvphone"`    // cnxhminvphone
	Cnxhminvmobile   sql.NullString  `json:"cnxhminvmobile"`   // cnxhminvmobile
	Cnxhminvemail    sql.NullString  `json:"cnxhminvemail"`    // cnxhminvemail
	Cnxhmc2title2    sql.NullString  `json:"cnxhmc2title2"`    // cnxhmc2title2
	Cnxhmc2firstnam2 sql.NullString  `json:"cnxhmc2firstnam2"` // cnxhmc2firstnam2
	Cnxhmc2surname2  sql.NullString  `json:"cnxhmc2surname2"`  // cnxhmc2surname2
	Cnxhmc2mobile2   sql.NullString  `json:"cnxhmc2mobile2"`   // cnxhmc2mobile2
	Cnxhmtempadd1    sql.NullString  `json:"cnxhmtempadd1"`    // cnxhmtempadd1
	Cnxhmtempadd2    sql.NullString  `json:"cnxhmtempadd2"`    // cnxhmtempadd2
	Cnxhmtempadd3    sql.NullString  `json:"cnxhmtempadd3"`    // cnxhmtempadd3
	Cnxhmtempadd4    sql.NullString  `json:"cnxhmtempadd4"`    // cnxhmtempadd4
	Cnxhmtempcounty  sql.NullString  `json:"cnxhmtempcounty"`  // cnxhmtempcounty
	Cnxhmtemppcode   sql.NullString  `json:"cnxhmtemppcode"`   // cnxhmtemppcode
	Cnxhmtempaddress sql.NullInt64   `json:"cnxhmtempaddress"` // cnxhmtempaddress
	Cnxhmnewphone    sql.NullString  `json:"cnxhmnewphone"`    // cnxhmnewphone
	Cnxhmbnksort     sql.NullString  `json:"cnxhmbnksort"`     // cnxhmbnksort
	Cnxhmbnkaccno    sql.NullString  `json:"cnxhmbnkaccno"`    // cnxhmbnkaccno
	Cnxhmbnkaccname  sql.NullString  `json:"cnxhmbnkaccname"`  // cnxhmbnkaccname
	Cnxhmbnkname     sql.NullString  `json:"cnxhmbnkname"`     // cnxhmbnkname
	Cnxhmnewbankdet  sql.NullInt64   `json:"cnxhmnewbankdet"`  // cnxhmnewbankdet
	Cnxhmbnksortnew  sql.NullString  `json:"cnxhmbnksortnew"`  // cnxhmbnksortnew
	Cnxhmbnkacnonew  sql.NullString  `json:"cnxhmbnkacnonew"`  // cnxhmbnkacnonew
	Cnxhmbnkacnamnew sql.NullString  `json:"cnxhmbnkacnamnew"` // cnxhmbnkacnamnew
	Cnxhmbnknamenew  sql.NullString  `json:"cnxhmbnknamenew"`  // cnxhmbnknamenew
	Cnxhmcctype      sql.NullInt64   `json:"cnxhmcctype"`      // cnxhmcctype
	Cnxhmccstartdate sql.NullString  `json:"cnxhmccstartdate"` // cnxhmccstartdate
	Cnxhmccenddate   sql.NullString  `json:"cnxhmccenddate"`   // cnxhmccenddate
	Cnxhmccname      sql.NullString  `json:"cnxhmccname"`      // cnxhmccname
	Cnxhmccnumber    sql.NullString  `json:"cnxhmccnumber"`    // cnxhmccnumber
	Cnxhmccseccode   sql.NullString  `json:"cnxhmccseccode"`   // cnxhmccseccode
	Cnxhmccissue     sql.NullString  `json:"cnxhmccissue"`     // cnxhmccissue
	Cnxhmcalltakerc1 sql.NullString  `json:"cnxhmcalltakerc1"` // cnxhmcalltakerc1
	Cnxhmcallsinadv  sql.NullInt64   `json:"cnxhmcallsinadv"`  // cnxhmcallsinadv
	Cnxhmlluenabled  sql.NullInt64   `json:"cnxhmlluenabled"`  // cnxhmlluenabled
	Cnxhmcablenumber sql.NullInt64   `json:"cnxhmcablenumber"` // cnxhmcablenumber
	Cnxhmsameexch    sql.NullInt64   `json:"cnxhmsameexch"`    // cnxhmsameexch
	Cnxhmnextaction  pq.NullTime     `json:"cnxhmnextaction"`  // cnxhmnextaction
	Cnxhmnextacttype sql.NullString  `json:"cnxhmnextacttype"` // cnxhmnextacttype
	Cnxhmgascomp     sql.NullInt64   `json:"cnxhmgascomp"`     // cnxhmgascomp
	Cnxhmc2vuln      sql.NullInt64   `json:"cnxhmc2vuln"`      // cnxhmc2vuln
	Cnxhmbcplus      sql.NullInt64   `json:"cnxhmbcplus"`      // cnxhmbcplus
	Cnxhmc2vulnage   sql.NullInt64   `json:"cnxhmc2vulnage"`   // cnxhmc2vulnage
	Cnxhmbcnumwith   sql.NullInt64   `json:"cnxhmbcnumwith"`   // cnxhmbcnumwith
	Cnxhmbcexdir     sql.NullInt64   `json:"cnxhmbcexdir"`     // cnxhmbcexdir
	Cnxhmbctps       sql.NullInt64   `json:"cnxhmbctps"`       // cnxhmbctps
	Cnxhmc2acname    sql.NullString  `json:"cnxhmc2acname"`    // cnxhmc2acname
	Cnxhmc2landl2    sql.NullString  `json:"cnxhmc2landl2"`    // cnxhmc2landl2
	Cnxhmc2vulnchage sql.NullInt64   `json:"cnxhmc2vulnchage"` // cnxhmc2vulnchage
	Cnxhmbcequipdel  sql.NullInt64   `json:"cnxhmbcequipdel"`  // cnxhmbcequipdel
	Cnxhmc2vulndisab sql.NullString  `json:"cnxhmc2vulndisab"` // cnxhmc2vulndisab
	Cnxhmbccomp      sql.NullInt64   `json:"cnxhmbccomp"`      // cnxhmbccomp
	Cnxhmfeatureline sql.NullInt64   `json:"cnxhmfeatureline"` // cnxhmfeatureline
	Cnxhmcable       sql.NullInt64   `json:"cnxhmcable"`       // cnxhmcable
	Cnxhmsmartbox    sql.NullInt64   `json:"cnxhmsmartbox"`    // cnxhmsmartbox
	Cnxhmspretainno  sql.NullInt64   `json:"cnxhmspretainno"`  // cnxhmspretainno
	Cnxhmmobsaver    sql.NullInt64   `json:"cnxhmmobsaver"`    // cnxhmmobsaver
	Cnxhmintsaver    sql.NullInt64   `json:"cnxhmintsaver"`    // cnxhmintsaver
	Cnxhmspnumwith   sql.NullInt64   `json:"cnxhmspnumwith"`   // cnxhmspnumwith
	Cnxhmspexdir     sql.NullInt64   `json:"cnxhmspexdir"`     // cnxhmspexdir
	Cnxhmsptps       sql.NullInt64   `json:"cnxhmsptps"`       // cnxhmsptps
	Cnxhmspcomp      sql.NullInt64   `json:"cnxhmspcomp"`      // cnxhmspcomp
	Cnxhmc2email2    sql.NullString  `json:"cnxhmc2email2"`    // cnxhmc2email2
	Cnxhmc2vulncarer sql.NullString  `json:"cnxhmc2vulncarer"` // cnxhmc2vulncarer
	Cnxhmc2vulncap   sql.NullString  `json:"cnxhmc2vulncap"`   // cnxhmc2vulncap
	Cnxhmc2vulnis    sql.NullString  `json:"cnxhmc2vulnis"`    // cnxhmc2vulnis
	Cnxhmc2vulnfinan sql.NullString  `json:"cnxhmc2vulnfinan"` // cnxhmc2vulnfinan
	Cnxhmgpequipdel  sql.NullInt64   `json:"cnxhmgpequipdel"`  // cnxhmgpequipdel
	Cnxhmc2vulnchild sql.NullString  `json:"cnxhmc2vulnchild"` // cnxhmc2vulnchild
	Cnxhmmy8comp     sql.NullInt64   `json:"cnxhmmy8comp"`     // cnxhmmy8comp
	Cnxhmmobtariff   sql.NullString  `json:"cnxhmmobtariff"`   // cnxhmmobtariff
	Cnxhmhandset     sql.NullString  `json:"cnxhmhandset"`     // cnxhmhandset
	Cnxhmrpc         sql.NullInt64   `json:"cnxhmrpc"`         // cnxhmrpc
	Cnxhminstbillval sql.NullFloat64 `json:"cnxhminstbillval"` // cnxhminstbillval
	Cnxhmvoicemail   sql.NullInt64   `json:"cnxhmvoicemail"`   // cnxhmvoicemail
	Cnxhmport        sql.NullInt64   `json:"cnxhmport"`        // cnxhmport
	Cnxhmpac         sql.NullString  `json:"cnxhmpac"`         // cnxhmpac
	Cnxhmmobequipdel sql.NullInt64   `json:"cnxhmmobequipdel"` // cnxhmmobequipdel
	Cnxhmmobcomp     sql.NullInt64   `json:"cnxhmmobcomp"`     // cnxhmmobcomp
	Cnxhmstatus      sql.NullString  `json:"cnxhmstatus"`      // cnxhmstatus
	Cnxhmnewaddexist sql.NullString  `json:"cnxhmnewaddexist"` // cnxhmnewaddexist
	Cnxhmuniquesys   sql.NullInt64   `json:"cnxhmuniquesys"`   // cnxhmuniquesys
	Cnxhmnewocacno   sql.NullString  `json:"cnxhmnewocacno"`   // cnxhmnewocacno
	Cnxhmc1p1servs   sql.NullString  `json:"cnxhmc1p1servs"`   // cnxhmc1p1servs
	Cnxhmc2p1servs   sql.NullString  `json:"cnxhmc2p1servs"`   // cnxhmc2p1servs
	Cnxhmnewownerten sql.NullString  `json:"cnxhmnewownerten"` // cnxhmnewownerten
	Cnxhmnewtenancye pq.NullTime     `json:"cnxhmnewtenancye"` // cnxhmnewtenancye
	Cnxhmnewinvpcode sql.NullString  `json:"cnxhmnewinvpcode"` // cnxhmnewinvpcode
	Cnxhmnewinvadd1  sql.NullString  `json:"cnxhmnewinvadd1"`  // cnxhmnewinvadd1
	Cnxhmnewinvadd2  sql.NullString  `json:"cnxhmnewinvadd2"`  // cnxhmnewinvadd2
	Cnxhmnewinvadd3  sql.NullString  `json:"cnxhmnewinvadd3"`  // cnxhmnewinvadd3
	Cnxhmnewinvadd4  sql.NullString  `json:"cnxhmnewinvadd4"`  // cnxhmnewinvadd4
	Cnxhmnewinvcount sql.NullString  `json:"cnxhmnewinvcount"` // cnxhmnewinvcount
	Cnxhmc1addchange pq.NullTime     `json:"cnxhmc1addchange"` // cnxhmc1addchange
	Cnxhmc1p2servs   sql.NullString  `json:"cnxhmc1p2servs"`   // cnxhmc1p2servs
	Cnxhmnewmpan     sql.NullString  `json:"cnxhmnewmpan"`     // cnxhmnewmpan
	Cnxhmnewmprn     sql.NullString  `json:"cnxhmnewmprn"`     // cnxhmnewmprn
	Cnxhmnewmsno     sql.NullString  `json:"cnxhmnewmsno"`     // cnxhmnewmsno
	Cnxhmnewlgas     sql.NullString  `json:"cnxhmnewlgas"`     // cnxhmnewlgas
	Cnxhmverbalcon   sql.NullInt64   `json:"cnxhmverbalcon"`   // cnxhmverbalcon
	Cnxhmclublevel   sql.NullInt64   `json:"cnxhmclublevel"`   // cnxhmclublevel
	Cnxhmputintoclub sql.NullInt64   `json:"cnxhmputintoclub"` // cnxhmputintoclub
	Cnxhmstatuscom   sql.NullInt64   `json:"cnxhmstatuscom"`   // cnxhmstatuscom
	Cnxhmfreetext    sql.NullInt64   `json:"cnxhmfreetext"`    // cnxhmfreetext
	Cnxhmp2complete  pq.NullTime     `json:"cnxhmp2complete"`  // cnxhmp2complete
	Cnxhmnewcctype   sql.NullString  `json:"cnxhmnewcctype"`   // cnxhmnewcctype
	Cnxhmnewccstart  sql.NullString  `json:"cnxhmnewccstart"`  // cnxhmnewccstart
	Cnxhmnewccend    sql.NullString  `json:"cnxhmnewccend"`    // cnxhmnewccend
	Cnxhmnewccnum    sql.NullString  `json:"cnxhmnewccnum"`    // cnxhmnewccnum
	Cnxhmnewccissue  sql.NullString  `json:"cnxhmnewccissue"`  // cnxhmnewccissue
	Cnxhmnewfreetext sql.NullInt64   `json:"cnxhmnewfreetext"` // cnxhmnewfreetext
	Cnxhmnewebills   sql.NullInt64   `json:"cnxhmnewebills"`   // cnxhmnewebills
	Cnxhmnewmarket   sql.NullInt64   `json:"cnxhmnewmarket"`   // cnxhmnewmarket
	Cnxhmnewclublvl  sql.NullString  `json:"cnxhmnewclublvl"`  // cnxhmnewclublvl
	Cnxhmnewdob      pq.NullTime     `json:"cnxhmnewdob"`      // cnxhmnewdob
	Cnxhmnewemail1   sql.NullString  `json:"cnxhmnewemail1"`   // cnxhmnewemail1
	Cnxhmcalltakerc2 sql.NullString  `json:"cnxhmcalltakerc2"` // cnxhmcalltakerc2
	Cnxhmenteredonc2 pq.NullTime     `json:"cnxhmenteredonc2"` // cnxhmenteredonc2
	Cnxhmfirstpoc    sql.NullInt64   `json:"cnxhmfirstpoc"`    // cnxhmfirstpoc
	Cnxhmbtsocket    sql.NullInt64   `json:"cnxhmbtsocket"`    // cnxhmbtsocket
	Cnxhmphase2      sql.NullInt64   `json:"cnxhmphase2"`      // cnxhmphase2
	Cnxhmc2status    sql.NullString  `json:"cnxhmc2status"`    // cnxhmc2status
	Cnxhmc2mobile    sql.NullString  `json:"cnxhmc2mobile"`    // cnxhmc2mobile
	Cnxhmc2email     sql.NullString  `json:"cnxhmc2email"`     // cnxhmc2email
	Cnxhmc2moveindte pq.NullTime     `json:"cnxhmc2moveindte"` // cnxhmc2moveindte
	Cnxhmc2landline  sql.NullString  `json:"cnxhmc2landline"`  // cnxhmc2landline
	Cnxhmcaller      sql.NullString  `json:"cnxhmcaller"`      // cnxhmcaller
	Cnxhmc2depagreed sql.NullInt64   `json:"cnxhmc2depagreed"` // cnxhmc2depagreed
	Cnxhminstantac   pq.NullTime     `json:"cnxhminstantac"`   // cnxhminstantac
	Cnxhminstantbill sql.NullInt64   `json:"cnxhminstantbill"` // cnxhminstantbill
	Cnxhmstatuscomc2 sql.NullInt64   `json:"cnxhmstatuscomc2"` // cnxhmstatuscomc2
	Cnxhminstacdone  pq.NullTime     `json:"cnxhminstacdone"`  // cnxhminstacdone
	Cnxhmc2vulndet   sql.NullInt64   `json:"cnxhmc2vulndet"`   // cnxhmc2vulndet
	Cnxhminstbilcalc sql.NullFloat64 `json:"cnxhminstbilcalc"` // cnxhminstbilcalc
	Cnxhminstbilpaid pq.NullTime     `json:"cnxhminstbilpaid"` // cnxhminstbilpaid
	Cnxhminstbilldet sql.NullInt64   `json:"cnxhminstbilldet"` // cnxhminstbilldet
	Cnxhmc1commpref  sql.NullString  `json:"cnxhmc1commpref"`  // cnxhmc1commpref
	Cnxhmdafcb       sql.NullString  `json:"cnxhmdafcb"`       // cnxhmdafcb
	Cnxhmdafavail    sql.NullString  `json:"cnxhmdafavail"`    // cnxhmdafavail
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllCnxhm(db XODB, callback func(x Cnxhm) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`cnxhmc2dob2, cnxhmenteredonc1, cnxhmstarted, cnxhmcompleted, cnxhmaccountno, cnxhmcustname, cnxhmbillgroup, cnxhmcusttitle, cnxhmcustiniti, cnxhmcustfirstn, cnxhmcustsurnam, cnxhmcustadd1, cnxhmcustadd2, cnxhmcustadd3, cnxhmcustadd4, cnxhmcustcounty, cnxhmcustpcode, cnxhmcustphone, cnxhmcustmobile, cnxhmdateofbirth, cnxhmcustemail, cnxhmemailbills, cnxhmmoveindate, cnxhmmoveoutdate, cnxhmownertenant, cnxhmtenancyend, cnxhmnewtitle, cnxhmnewfirstnam, cnxhmnewsurname, cnxhmnewadd1, cnxhmnewadd2, cnxhmnewadd3, cnxhmnewadd4, cnxhmnewcounty, cnxhmnewpostcode, cnxhmcurrtitle, cnxhmcurrfirstn, cnxhmcurrsurname, cnxhminvtitle, cnxhminvinitials, cnxhminvfirstn, cnxhminvsurname, cnxhminvadd1, cnxhminvadd2, cnxhminvadd3, cnxhminvadd4, cnxhminvcounty, cnxhminvpostcode, cnxhminvphone, cnxhminvmobile, cnxhminvemail, cnxhmc2title2, cnxhmc2firstnam2, cnxhmc2surname2, cnxhmc2mobile2, cnxhmtempadd1, cnxhmtempadd2, cnxhmtempadd3, cnxhmtempadd4, cnxhmtempcounty, cnxhmtemppcode, cnxhmtempaddress, cnxhmnewphone, cnxhmbnksort, cnxhmbnkaccno, cnxhmbnkaccname, cnxhmbnkname, cnxhmnewbankdet, cnxhmbnksortnew, cnxhmbnkacnonew, cnxhmbnkacnamnew, cnxhmbnknamenew, cnxhmcctype, cnxhmccstartdate, cnxhmccenddate, cnxhmccname, cnxhmccnumber, cnxhmccseccode, cnxhmccissue, cnxhmcalltakerc1, cnxhmcallsinadv, cnxhmlluenabled, cnxhmcablenumber, cnxhmsameexch, cnxhmnextaction, cnxhmnextacttype, cnxhmgascomp, cnxhmc2vuln, cnxhmbcplus, cnxhmc2vulnage, cnxhmbcnumwith, cnxhmbcexdir, cnxhmbctps, cnxhmc2acname, cnxhmc2landl2, cnxhmc2vulnchage, cnxhmbcequipdel, cnxhmc2vulndisab, cnxhmbccomp, cnxhmfeatureline, cnxhmcable, cnxhmsmartbox, cnxhmspretainno, cnxhmmobsaver, cnxhmintsaver, cnxhmspnumwith, cnxhmspexdir, cnxhmsptps, cnxhmspcomp, cnxhmc2email2, cnxhmc2vulncarer, cnxhmc2vulncap, cnxhmc2vulnis, cnxhmc2vulnfinan, cnxhmgpequipdel, cnxhmc2vulnchild, cnxhmmy8comp, cnxhmmobtariff, cnxhmhandset, cnxhmrpc, cnxhminstbillval, cnxhmvoicemail, cnxhmport, cnxhmpac, cnxhmmobequipdel, cnxhmmobcomp, cnxhmstatus, cnxhmnewaddexist, cnxhmuniquesys, cnxhmnewocacno, cnxhmc1p1servs, cnxhmc2p1servs, cnxhmnewownerten, cnxhmnewtenancye, cnxhmnewinvpcode, cnxhmnewinvadd1, cnxhmnewinvadd2, cnxhmnewinvadd3, cnxhmnewinvadd4, cnxhmnewinvcount, cnxhmc1addchange, cnxhmc1p2servs, cnxhmnewmpan, cnxhmnewmprn, cnxhmnewmsno, cnxhmnewlgas, cnxhmverbalcon, cnxhmclublevel, cnxhmputintoclub, cnxhmstatuscom, cnxhmfreetext, cnxhmp2complete, cnxhmnewcctype, cnxhmnewccstart, cnxhmnewccend, cnxhmnewccnum, cnxhmnewccissue, cnxhmnewfreetext, cnxhmnewebills, cnxhmnewmarket, cnxhmnewclublvl, cnxhmnewdob, cnxhmnewemail1, cnxhmcalltakerc2, cnxhmenteredonc2, cnxhmfirstpoc, cnxhmbtsocket, cnxhmphase2, cnxhmc2status, cnxhmc2mobile, cnxhmc2email, cnxhmc2moveindte, cnxhmc2landline, cnxhmcaller, cnxhmc2depagreed, cnxhminstantac, cnxhminstantbill, cnxhmstatuscomc2, cnxhminstacdone, cnxhmc2vulndet, cnxhminstbilcalc, cnxhminstbilpaid, cnxhminstbilldet, cnxhmc1commpref, cnxhmdafcb, cnxhmdafavail, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxhm `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Cnxhm{}

		// scan
		err = q.Scan(&c.Cnxhmc2dob2, &c.Cnxhmenteredonc1, &c.Cnxhmstarted, &c.Cnxhmcompleted, &c.Cnxhmaccountno, &c.Cnxhmcustname, &c.Cnxhmbillgroup, &c.Cnxhmcusttitle, &c.Cnxhmcustiniti, &c.Cnxhmcustfirstn, &c.Cnxhmcustsurnam, &c.Cnxhmcustadd1, &c.Cnxhmcustadd2, &c.Cnxhmcustadd3, &c.Cnxhmcustadd4, &c.Cnxhmcustcounty, &c.Cnxhmcustpcode, &c.Cnxhmcustphone, &c.Cnxhmcustmobile, &c.Cnxhmdateofbirth, &c.Cnxhmcustemail, &c.Cnxhmemailbills, &c.Cnxhmmoveindate, &c.Cnxhmmoveoutdate, &c.Cnxhmownertenant, &c.Cnxhmtenancyend, &c.Cnxhmnewtitle, &c.Cnxhmnewfirstnam, &c.Cnxhmnewsurname, &c.Cnxhmnewadd1, &c.Cnxhmnewadd2, &c.Cnxhmnewadd3, &c.Cnxhmnewadd4, &c.Cnxhmnewcounty, &c.Cnxhmnewpostcode, &c.Cnxhmcurrtitle, &c.Cnxhmcurrfirstn, &c.Cnxhmcurrsurname, &c.Cnxhminvtitle, &c.Cnxhminvinitials, &c.Cnxhminvfirstn, &c.Cnxhminvsurname, &c.Cnxhminvadd1, &c.Cnxhminvadd2, &c.Cnxhminvadd3, &c.Cnxhminvadd4, &c.Cnxhminvcounty, &c.Cnxhminvpostcode, &c.Cnxhminvphone, &c.Cnxhminvmobile, &c.Cnxhminvemail, &c.Cnxhmc2title2, &c.Cnxhmc2firstnam2, &c.Cnxhmc2surname2, &c.Cnxhmc2mobile2, &c.Cnxhmtempadd1, &c.Cnxhmtempadd2, &c.Cnxhmtempadd3, &c.Cnxhmtempadd4, &c.Cnxhmtempcounty, &c.Cnxhmtemppcode, &c.Cnxhmtempaddress, &c.Cnxhmnewphone, &c.Cnxhmbnksort, &c.Cnxhmbnkaccno, &c.Cnxhmbnkaccname, &c.Cnxhmbnkname, &c.Cnxhmnewbankdet, &c.Cnxhmbnksortnew, &c.Cnxhmbnkacnonew, &c.Cnxhmbnkacnamnew, &c.Cnxhmbnknamenew, &c.Cnxhmcctype, &c.Cnxhmccstartdate, &c.Cnxhmccenddate, &c.Cnxhmccname, &c.Cnxhmccnumber, &c.Cnxhmccseccode, &c.Cnxhmccissue, &c.Cnxhmcalltakerc1, &c.Cnxhmcallsinadv, &c.Cnxhmlluenabled, &c.Cnxhmcablenumber, &c.Cnxhmsameexch, &c.Cnxhmnextaction, &c.Cnxhmnextacttype, &c.Cnxhmgascomp, &c.Cnxhmc2vuln, &c.Cnxhmbcplus, &c.Cnxhmc2vulnage, &c.Cnxhmbcnumwith, &c.Cnxhmbcexdir, &c.Cnxhmbctps, &c.Cnxhmc2acname, &c.Cnxhmc2landl2, &c.Cnxhmc2vulnchage, &c.Cnxhmbcequipdel, &c.Cnxhmc2vulndisab, &c.Cnxhmbccomp, &c.Cnxhmfeatureline, &c.Cnxhmcable, &c.Cnxhmsmartbox, &c.Cnxhmspretainno, &c.Cnxhmmobsaver, &c.Cnxhmintsaver, &c.Cnxhmspnumwith, &c.Cnxhmspexdir, &c.Cnxhmsptps, &c.Cnxhmspcomp, &c.Cnxhmc2email2, &c.Cnxhmc2vulncarer, &c.Cnxhmc2vulncap, &c.Cnxhmc2vulnis, &c.Cnxhmc2vulnfinan, &c.Cnxhmgpequipdel, &c.Cnxhmc2vulnchild, &c.Cnxhmmy8comp, &c.Cnxhmmobtariff, &c.Cnxhmhandset, &c.Cnxhmrpc, &c.Cnxhminstbillval, &c.Cnxhmvoicemail, &c.Cnxhmport, &c.Cnxhmpac, &c.Cnxhmmobequipdel, &c.Cnxhmmobcomp, &c.Cnxhmstatus, &c.Cnxhmnewaddexist, &c.Cnxhmuniquesys, &c.Cnxhmnewocacno, &c.Cnxhmc1p1servs, &c.Cnxhmc2p1servs, &c.Cnxhmnewownerten, &c.Cnxhmnewtenancye, &c.Cnxhmnewinvpcode, &c.Cnxhmnewinvadd1, &c.Cnxhmnewinvadd2, &c.Cnxhmnewinvadd3, &c.Cnxhmnewinvadd4, &c.Cnxhmnewinvcount, &c.Cnxhmc1addchange, &c.Cnxhmc1p2servs, &c.Cnxhmnewmpan, &c.Cnxhmnewmprn, &c.Cnxhmnewmsno, &c.Cnxhmnewlgas, &c.Cnxhmverbalcon, &c.Cnxhmclublevel, &c.Cnxhmputintoclub, &c.Cnxhmstatuscom, &c.Cnxhmfreetext, &c.Cnxhmp2complete, &c.Cnxhmnewcctype, &c.Cnxhmnewccstart, &c.Cnxhmnewccend, &c.Cnxhmnewccnum, &c.Cnxhmnewccissue, &c.Cnxhmnewfreetext, &c.Cnxhmnewebills, &c.Cnxhmnewmarket, &c.Cnxhmnewclublvl, &c.Cnxhmnewdob, &c.Cnxhmnewemail1, &c.Cnxhmcalltakerc2, &c.Cnxhmenteredonc2, &c.Cnxhmfirstpoc, &c.Cnxhmbtsocket, &c.Cnxhmphase2, &c.Cnxhmc2status, &c.Cnxhmc2mobile, &c.Cnxhmc2email, &c.Cnxhmc2moveindte, &c.Cnxhmc2landline, &c.Cnxhmcaller, &c.Cnxhmc2depagreed, &c.Cnxhminstantac, &c.Cnxhminstantbill, &c.Cnxhmstatuscomc2, &c.Cnxhminstacdone, &c.Cnxhmc2vulndet, &c.Cnxhminstbilcalc, &c.Cnxhminstbilpaid, &c.Cnxhminstbilldet, &c.Cnxhmc1commpref, &c.Cnxhmdafcb, &c.Cnxhmdafavail, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CnxhmByEquinoxLrn retrieves a row from 'equinox.cnxhm' as a Cnxhm.
//
// Generated from index 'cnxhm_pkey'.
func CnxhmByEquinoxLrn(db XODB, equinoxLrn int64) (*Cnxhm, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cnxhmc2dob2, cnxhmenteredonc1, cnxhmstarted, cnxhmcompleted, cnxhmaccountno, cnxhmcustname, cnxhmbillgroup, cnxhmcusttitle, cnxhmcustiniti, cnxhmcustfirstn, cnxhmcustsurnam, cnxhmcustadd1, cnxhmcustadd2, cnxhmcustadd3, cnxhmcustadd4, cnxhmcustcounty, cnxhmcustpcode, cnxhmcustphone, cnxhmcustmobile, cnxhmdateofbirth, cnxhmcustemail, cnxhmemailbills, cnxhmmoveindate, cnxhmmoveoutdate, cnxhmownertenant, cnxhmtenancyend, cnxhmnewtitle, cnxhmnewfirstnam, cnxhmnewsurname, cnxhmnewadd1, cnxhmnewadd2, cnxhmnewadd3, cnxhmnewadd4, cnxhmnewcounty, cnxhmnewpostcode, cnxhmcurrtitle, cnxhmcurrfirstn, cnxhmcurrsurname, cnxhminvtitle, cnxhminvinitials, cnxhminvfirstn, cnxhminvsurname, cnxhminvadd1, cnxhminvadd2, cnxhminvadd3, cnxhminvadd4, cnxhminvcounty, cnxhminvpostcode, cnxhminvphone, cnxhminvmobile, cnxhminvemail, cnxhmc2title2, cnxhmc2firstnam2, cnxhmc2surname2, cnxhmc2mobile2, cnxhmtempadd1, cnxhmtempadd2, cnxhmtempadd3, cnxhmtempadd4, cnxhmtempcounty, cnxhmtemppcode, cnxhmtempaddress, cnxhmnewphone, cnxhmbnksort, cnxhmbnkaccno, cnxhmbnkaccname, cnxhmbnkname, cnxhmnewbankdet, cnxhmbnksortnew, cnxhmbnkacnonew, cnxhmbnkacnamnew, cnxhmbnknamenew, cnxhmcctype, cnxhmccstartdate, cnxhmccenddate, cnxhmccname, cnxhmccnumber, cnxhmccseccode, cnxhmccissue, cnxhmcalltakerc1, cnxhmcallsinadv, cnxhmlluenabled, cnxhmcablenumber, cnxhmsameexch, cnxhmnextaction, cnxhmnextacttype, cnxhmgascomp, cnxhmc2vuln, cnxhmbcplus, cnxhmc2vulnage, cnxhmbcnumwith, cnxhmbcexdir, cnxhmbctps, cnxhmc2acname, cnxhmc2landl2, cnxhmc2vulnchage, cnxhmbcequipdel, cnxhmc2vulndisab, cnxhmbccomp, cnxhmfeatureline, cnxhmcable, cnxhmsmartbox, cnxhmspretainno, cnxhmmobsaver, cnxhmintsaver, cnxhmspnumwith, cnxhmspexdir, cnxhmsptps, cnxhmspcomp, cnxhmc2email2, cnxhmc2vulncarer, cnxhmc2vulncap, cnxhmc2vulnis, cnxhmc2vulnfinan, cnxhmgpequipdel, cnxhmc2vulnchild, cnxhmmy8comp, cnxhmmobtariff, cnxhmhandset, cnxhmrpc, cnxhminstbillval, cnxhmvoicemail, cnxhmport, cnxhmpac, cnxhmmobequipdel, cnxhmmobcomp, cnxhmstatus, cnxhmnewaddexist, cnxhmuniquesys, cnxhmnewocacno, cnxhmc1p1servs, cnxhmc2p1servs, cnxhmnewownerten, cnxhmnewtenancye, cnxhmnewinvpcode, cnxhmnewinvadd1, cnxhmnewinvadd2, cnxhmnewinvadd3, cnxhmnewinvadd4, cnxhmnewinvcount, cnxhmc1addchange, cnxhmc1p2servs, cnxhmnewmpan, cnxhmnewmprn, cnxhmnewmsno, cnxhmnewlgas, cnxhmverbalcon, cnxhmclublevel, cnxhmputintoclub, cnxhmstatuscom, cnxhmfreetext, cnxhmp2complete, cnxhmnewcctype, cnxhmnewccstart, cnxhmnewccend, cnxhmnewccnum, cnxhmnewccissue, cnxhmnewfreetext, cnxhmnewebills, cnxhmnewmarket, cnxhmnewclublvl, cnxhmnewdob, cnxhmnewemail1, cnxhmcalltakerc2, cnxhmenteredonc2, cnxhmfirstpoc, cnxhmbtsocket, cnxhmphase2, cnxhmc2status, cnxhmc2mobile, cnxhmc2email, cnxhmc2moveindte, cnxhmc2landline, cnxhmcaller, cnxhmc2depagreed, cnxhminstantac, cnxhminstantbill, cnxhmstatuscomc2, cnxhminstacdone, cnxhmc2vulndet, cnxhminstbilcalc, cnxhminstbilpaid, cnxhminstbilldet, cnxhmc1commpref, cnxhmdafcb, cnxhmdafavail, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxhm ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cnxhm{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cnxhmc2dob2, &c.Cnxhmenteredonc1, &c.Cnxhmstarted, &c.Cnxhmcompleted, &c.Cnxhmaccountno, &c.Cnxhmcustname, &c.Cnxhmbillgroup, &c.Cnxhmcusttitle, &c.Cnxhmcustiniti, &c.Cnxhmcustfirstn, &c.Cnxhmcustsurnam, &c.Cnxhmcustadd1, &c.Cnxhmcustadd2, &c.Cnxhmcustadd3, &c.Cnxhmcustadd4, &c.Cnxhmcustcounty, &c.Cnxhmcustpcode, &c.Cnxhmcustphone, &c.Cnxhmcustmobile, &c.Cnxhmdateofbirth, &c.Cnxhmcustemail, &c.Cnxhmemailbills, &c.Cnxhmmoveindate, &c.Cnxhmmoveoutdate, &c.Cnxhmownertenant, &c.Cnxhmtenancyend, &c.Cnxhmnewtitle, &c.Cnxhmnewfirstnam, &c.Cnxhmnewsurname, &c.Cnxhmnewadd1, &c.Cnxhmnewadd2, &c.Cnxhmnewadd3, &c.Cnxhmnewadd4, &c.Cnxhmnewcounty, &c.Cnxhmnewpostcode, &c.Cnxhmcurrtitle, &c.Cnxhmcurrfirstn, &c.Cnxhmcurrsurname, &c.Cnxhminvtitle, &c.Cnxhminvinitials, &c.Cnxhminvfirstn, &c.Cnxhminvsurname, &c.Cnxhminvadd1, &c.Cnxhminvadd2, &c.Cnxhminvadd3, &c.Cnxhminvadd4, &c.Cnxhminvcounty, &c.Cnxhminvpostcode, &c.Cnxhminvphone, &c.Cnxhminvmobile, &c.Cnxhminvemail, &c.Cnxhmc2title2, &c.Cnxhmc2firstnam2, &c.Cnxhmc2surname2, &c.Cnxhmc2mobile2, &c.Cnxhmtempadd1, &c.Cnxhmtempadd2, &c.Cnxhmtempadd3, &c.Cnxhmtempadd4, &c.Cnxhmtempcounty, &c.Cnxhmtemppcode, &c.Cnxhmtempaddress, &c.Cnxhmnewphone, &c.Cnxhmbnksort, &c.Cnxhmbnkaccno, &c.Cnxhmbnkaccname, &c.Cnxhmbnkname, &c.Cnxhmnewbankdet, &c.Cnxhmbnksortnew, &c.Cnxhmbnkacnonew, &c.Cnxhmbnkacnamnew, &c.Cnxhmbnknamenew, &c.Cnxhmcctype, &c.Cnxhmccstartdate, &c.Cnxhmccenddate, &c.Cnxhmccname, &c.Cnxhmccnumber, &c.Cnxhmccseccode, &c.Cnxhmccissue, &c.Cnxhmcalltakerc1, &c.Cnxhmcallsinadv, &c.Cnxhmlluenabled, &c.Cnxhmcablenumber, &c.Cnxhmsameexch, &c.Cnxhmnextaction, &c.Cnxhmnextacttype, &c.Cnxhmgascomp, &c.Cnxhmc2vuln, &c.Cnxhmbcplus, &c.Cnxhmc2vulnage, &c.Cnxhmbcnumwith, &c.Cnxhmbcexdir, &c.Cnxhmbctps, &c.Cnxhmc2acname, &c.Cnxhmc2landl2, &c.Cnxhmc2vulnchage, &c.Cnxhmbcequipdel, &c.Cnxhmc2vulndisab, &c.Cnxhmbccomp, &c.Cnxhmfeatureline, &c.Cnxhmcable, &c.Cnxhmsmartbox, &c.Cnxhmspretainno, &c.Cnxhmmobsaver, &c.Cnxhmintsaver, &c.Cnxhmspnumwith, &c.Cnxhmspexdir, &c.Cnxhmsptps, &c.Cnxhmspcomp, &c.Cnxhmc2email2, &c.Cnxhmc2vulncarer, &c.Cnxhmc2vulncap, &c.Cnxhmc2vulnis, &c.Cnxhmc2vulnfinan, &c.Cnxhmgpequipdel, &c.Cnxhmc2vulnchild, &c.Cnxhmmy8comp, &c.Cnxhmmobtariff, &c.Cnxhmhandset, &c.Cnxhmrpc, &c.Cnxhminstbillval, &c.Cnxhmvoicemail, &c.Cnxhmport, &c.Cnxhmpac, &c.Cnxhmmobequipdel, &c.Cnxhmmobcomp, &c.Cnxhmstatus, &c.Cnxhmnewaddexist, &c.Cnxhmuniquesys, &c.Cnxhmnewocacno, &c.Cnxhmc1p1servs, &c.Cnxhmc2p1servs, &c.Cnxhmnewownerten, &c.Cnxhmnewtenancye, &c.Cnxhmnewinvpcode, &c.Cnxhmnewinvadd1, &c.Cnxhmnewinvadd2, &c.Cnxhmnewinvadd3, &c.Cnxhmnewinvadd4, &c.Cnxhmnewinvcount, &c.Cnxhmc1addchange, &c.Cnxhmc1p2servs, &c.Cnxhmnewmpan, &c.Cnxhmnewmprn, &c.Cnxhmnewmsno, &c.Cnxhmnewlgas, &c.Cnxhmverbalcon, &c.Cnxhmclublevel, &c.Cnxhmputintoclub, &c.Cnxhmstatuscom, &c.Cnxhmfreetext, &c.Cnxhmp2complete, &c.Cnxhmnewcctype, &c.Cnxhmnewccstart, &c.Cnxhmnewccend, &c.Cnxhmnewccnum, &c.Cnxhmnewccissue, &c.Cnxhmnewfreetext, &c.Cnxhmnewebills, &c.Cnxhmnewmarket, &c.Cnxhmnewclublvl, &c.Cnxhmnewdob, &c.Cnxhmnewemail1, &c.Cnxhmcalltakerc2, &c.Cnxhmenteredonc2, &c.Cnxhmfirstpoc, &c.Cnxhmbtsocket, &c.Cnxhmphase2, &c.Cnxhmc2status, &c.Cnxhmc2mobile, &c.Cnxhmc2email, &c.Cnxhmc2moveindte, &c.Cnxhmc2landline, &c.Cnxhmcaller, &c.Cnxhmc2depagreed, &c.Cnxhminstantac, &c.Cnxhminstantbill, &c.Cnxhmstatuscomc2, &c.Cnxhminstacdone, &c.Cnxhmc2vulndet, &c.Cnxhminstbilcalc, &c.Cnxhminstbilpaid, &c.Cnxhminstbilldet, &c.Cnxhmc1commpref, &c.Cnxhmdafcb, &c.Cnxhmdafavail, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
