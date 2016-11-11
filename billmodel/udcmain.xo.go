// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Udcmain represents a row from 'equinox.udcmain'.
type Udcmain struct {
	Udccustaccountno sql.NullString  `json:"udccustaccountno"` // udccustaccountno
	Udcaccountstatus sql.NullString  `json:"udcaccountstatus"` // udcaccountstatus
	Udcaccmanager    sql.NullString  `json:"udcaccmanager"`    // udcaccmanager
	Udcbatchnumber   sql.NullString  `json:"udcbatchnumber"`   // udcbatchnumber
	Udcimportbalance sql.NullFloat64 `json:"udcimportbalance"` // udcimportbalance
	Udcnewchase      sql.NullFloat64 `json:"udcnewchase"`      // udcnewchase
	Udcdatetoudc     pq.NullTime     `json:"udcdatetoudc"`     // udcdatetoudc
	Udcagency        sql.NullString  `json:"udcagency"`        // udcagency
	Udcpdvsent       pq.NullTime     `json:"udcpdvsent"`       // udcpdvsent
	Udcpdvreturned   pq.NullTime     `json:"udcpdvreturned"`   // udcpdvreturned
	Udccourtdategas  pq.NullTime     `json:"udccourtdategas"`  // udccourtdategas
	Udccourtdateelec pq.NullTime     `json:"udccourtdateelec"` // udccourtdateelec
	Udcisodategas    pq.NullTime     `json:"udcisodategas"`    // udcisodategas
	Udcisodateelec   pq.NullTime     `json:"udcisodateelec"`   // udcisodateelec
	Udcrebookisogas  pq.NullTime     `json:"udcrebookisogas"`  // udcrebookisogas
	Udcrebookisoelec pq.NullTime     `json:"udcrebookisoelec"` // udcrebookisoelec
	Udcisocancelgas  pq.NullTime     `json:"udcisocancelgas"`  // udcisocancelgas
	Udcisocancelelec pq.NullTime     `json:"udcisocancelelec"` // udcisocancelelec
	Udcisoreasongas  sql.NullString  `json:"udcisoreasongas"`  // udcisoreasongas
	Udcisoreasonelec sql.NullString  `json:"udcisoreasonelec"` // udcisoreasonelec
	Udcppmetermsngas sql.NullString  `json:"udcppmetermsngas"` // udcppmetermsngas
	Udcppmetermsnele sql.NullString  `json:"udcppmetermsnele"` // udcppmetermsnele
	Udcmetertypegas  sql.NullString  `json:"udcmetertypegas"`  // udcmetertypegas
	Udcmetertypeelec sql.NullString  `json:"udcmetertypeelec"` // udcmetertypeelec
	Udcprrgas        sql.NullString  `json:"udcprrgas"`        // udcprrgas
	Udcprrelec       sql.NullString  `json:"udcprrelec"`       // udcprrelec
	Udcfinaldebtgas  sql.NullFloat64 `json:"udcfinaldebtgas"`  // udcfinaldebtgas
	Udcfinaldebtelec sql.NullFloat64 `json:"udcfinaldebtelec"` // udcfinaldebtelec
	Udcppagreefeegas sql.NullFloat64 `json:"udcppagreefeegas"` // udcppagreefeegas
	Udppagreefeeelec sql.NullFloat64 `json:"udppagreefeeelec"` // udppagreefeeelec
	Udcisolationfgas sql.NullFloat64 `json:"udcisolationfgas"` // udcisolationfgas
	Udcisolationfele sql.NullFloat64 `json:"udcisolationfele"` // udcisolationfele
	Udcppstartdate   pq.NullTime     `json:"udcppstartdate"`   // udcppstartdate
	Udcppduration    sql.NullInt64   `json:"udcppduration"`    // udcppduration
	Udcppenddate     pq.NullTime     `json:"udcppenddate"`     // udcppenddate
	Udcagreedpayment sql.NullFloat64 `json:"udcagreedpayment"` // udcagreedpayment
	Udcpaymentfreq   sql.NullString  `json:"udcpaymentfreq"`   // udcpaymentfreq
	Udcpaymentmethod sql.NullString  `json:"udcpaymentmethod"` // udcpaymentmethod
	Udcpaymonth      sql.NullInt64   `json:"udcpaymonth"`      // udcpaymonth
	Udcexpectedpdate pq.NullTime     `json:"udcexpectedpdate"` // udcexpectedpdate
	Udcamountpaid    sql.NullFloat64 `json:"udcamountpaid"`    // udcamountpaid
	Udcexecfailgas   sql.NullString  `json:"udcexecfailgas"`   // udcexecfailgas
	Udcexecfailelec  sql.NullString  `json:"udcexecfailelec"`  // udcexecfailelec
	Udcactionreviewd pq.NullTime     `json:"udcactionreviewd"` // udcactionreviewd
	Udcimportseq     sql.NullInt64   `json:"udcimportseq"`     // udcimportseq
	Udcmodifieddate  pq.NullTime     `json:"udcmodifieddate"`  // udcmodifieddate
	Udcmodifiedtime  pq.NullTime     `json:"udcmodifiedtime"`  // udcmodifiedtime
	Udcmodifiedby    sql.NullString  `json:"udcmodifiedby"`    // udcmodifiedby
	Udcprogstatus    sql.NullString  `json:"udcprogstatus"`    // udcprogstatus
	Udccustomersalu  sql.NullString  `json:"udccustomersalu"`  // udccustomersalu
	Udcaccounttype   sql.NullString  `json:"udcaccounttype"`   // udcaccounttype
	Udcpdvdate       pq.NullTime     `json:"udcpdvdate"`       // udcpdvdate
	Udcpropertytype  sql.NullString  `json:"udcpropertytype"`  // udcpropertytype
	Udccourtfee      sql.NullFloat64 `json:"udccourtfee"`      // udccourtfee
	Udcinterest      sql.NullFloat64 `json:"udcinterest"`      // udcinterest
	Udcmartsonfee    sql.NullFloat64 `json:"udcmartsonfee"`    // udcmartsonfee
	Udclast4digits   sql.NullString  `json:"udclast4digits"`   // udclast4digits
	Udcdatesent      pq.NullTime     `json:"udcdatesent"`      // udcdatesent
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllUdcmain(db XODB, callback func(x Udcmain) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`udccustaccountno, udcaccountstatus, udcaccmanager, udcbatchnumber, udcimportbalance, udcnewchase, udcdatetoudc, udcagency, udcpdvsent, udcpdvreturned, udccourtdategas, udccourtdateelec, udcisodategas, udcisodateelec, udcrebookisogas, udcrebookisoelec, udcisocancelgas, udcisocancelelec, udcisoreasongas, udcisoreasonelec, udcppmetermsngas, udcppmetermsnele, udcmetertypegas, udcmetertypeelec, udcprrgas, udcprrelec, udcfinaldebtgas, udcfinaldebtelec, udcppagreefeegas, udppagreefeeelec, udcisolationfgas, udcisolationfele, udcppstartdate, udcppduration, udcppenddate, udcagreedpayment, udcpaymentfreq, udcpaymentmethod, udcpaymonth, udcexpectedpdate, udcamountpaid, udcexecfailgas, udcexecfailelec, udcactionreviewd, udcimportseq, udcmodifieddate, udcmodifiedtime, udcmodifiedby, udcprogstatus, udccustomersalu, udcaccounttype, udcpdvdate, udcpropertytype, udccourtfee, udcinterest, udcmartsonfee, udclast4digits, udcdatesent, equinox_lrn, equinox_sec ` +
		`FROM equinox.udcmain `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		u := Udcmain{}

		// scan
		err = q.Scan(&u.Udccustaccountno, &u.Udcaccountstatus, &u.Udcaccmanager, &u.Udcbatchnumber, &u.Udcimportbalance, &u.Udcnewchase, &u.Udcdatetoudc, &u.Udcagency, &u.Udcpdvsent, &u.Udcpdvreturned, &u.Udccourtdategas, &u.Udccourtdateelec, &u.Udcisodategas, &u.Udcisodateelec, &u.Udcrebookisogas, &u.Udcrebookisoelec, &u.Udcisocancelgas, &u.Udcisocancelelec, &u.Udcisoreasongas, &u.Udcisoreasonelec, &u.Udcppmetermsngas, &u.Udcppmetermsnele, &u.Udcmetertypegas, &u.Udcmetertypeelec, &u.Udcprrgas, &u.Udcprrelec, &u.Udcfinaldebtgas, &u.Udcfinaldebtelec, &u.Udcppagreefeegas, &u.Udppagreefeeelec, &u.Udcisolationfgas, &u.Udcisolationfele, &u.Udcppstartdate, &u.Udcppduration, &u.Udcppenddate, &u.Udcagreedpayment, &u.Udcpaymentfreq, &u.Udcpaymentmethod, &u.Udcpaymonth, &u.Udcexpectedpdate, &u.Udcamountpaid, &u.Udcexecfailgas, &u.Udcexecfailelec, &u.Udcactionreviewd, &u.Udcimportseq, &u.Udcmodifieddate, &u.Udcmodifiedtime, &u.Udcmodifiedby, &u.Udcprogstatus, &u.Udccustomersalu, &u.Udcaccounttype, &u.Udcpdvdate, &u.Udcpropertytype, &u.Udccourtfee, &u.Udcinterest, &u.Udcmartsonfee, &u.Udclast4digits, &u.Udcdatesent, &u.EquinoxLrn, &u.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(u) {
			return nil
		}
	}

	return nil
}

// UdcmainByEquinoxLrn retrieves a row from 'equinox.udcmain' as a Udcmain.
//
// Generated from index 'udcmain_pkey'.
func UdcmainByEquinoxLrn(db XODB, equinoxLrn int64) (*Udcmain, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`udccustaccountno, udcaccountstatus, udcaccmanager, udcbatchnumber, udcimportbalance, udcnewchase, udcdatetoudc, udcagency, udcpdvsent, udcpdvreturned, udccourtdategas, udccourtdateelec, udcisodategas, udcisodateelec, udcrebookisogas, udcrebookisoelec, udcisocancelgas, udcisocancelelec, udcisoreasongas, udcisoreasonelec, udcppmetermsngas, udcppmetermsnele, udcmetertypegas, udcmetertypeelec, udcprrgas, udcprrelec, udcfinaldebtgas, udcfinaldebtelec, udcppagreefeegas, udppagreefeeelec, udcisolationfgas, udcisolationfele, udcppstartdate, udcppduration, udcppenddate, udcagreedpayment, udcpaymentfreq, udcpaymentmethod, udcpaymonth, udcexpectedpdate, udcamountpaid, udcexecfailgas, udcexecfailelec, udcactionreviewd, udcimportseq, udcmodifieddate, udcmodifiedtime, udcmodifiedby, udcprogstatus, udccustomersalu, udcaccounttype, udcpdvdate, udcpropertytype, udccourtfee, udcinterest, udcmartsonfee, udclast4digits, udcdatesent, equinox_lrn, equinox_sec ` +
		`FROM equinox.udcmain ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	u := Udcmain{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&u.Udccustaccountno, &u.Udcaccountstatus, &u.Udcaccmanager, &u.Udcbatchnumber, &u.Udcimportbalance, &u.Udcnewchase, &u.Udcdatetoudc, &u.Udcagency, &u.Udcpdvsent, &u.Udcpdvreturned, &u.Udccourtdategas, &u.Udccourtdateelec, &u.Udcisodategas, &u.Udcisodateelec, &u.Udcrebookisogas, &u.Udcrebookisoelec, &u.Udcisocancelgas, &u.Udcisocancelelec, &u.Udcisoreasongas, &u.Udcisoreasonelec, &u.Udcppmetermsngas, &u.Udcppmetermsnele, &u.Udcmetertypegas, &u.Udcmetertypeelec, &u.Udcprrgas, &u.Udcprrelec, &u.Udcfinaldebtgas, &u.Udcfinaldebtelec, &u.Udcppagreefeegas, &u.Udppagreefeeelec, &u.Udcisolationfgas, &u.Udcisolationfele, &u.Udcppstartdate, &u.Udcppduration, &u.Udcppenddate, &u.Udcagreedpayment, &u.Udcpaymentfreq, &u.Udcpaymentmethod, &u.Udcpaymonth, &u.Udcexpectedpdate, &u.Udcamountpaid, &u.Udcexecfailgas, &u.Udcexecfailelec, &u.Udcactionreviewd, &u.Udcimportseq, &u.Udcmodifieddate, &u.Udcmodifiedtime, &u.Udcmodifiedby, &u.Udcprogstatus, &u.Udccustomersalu, &u.Udcaccounttype, &u.Udcpdvdate, &u.Udcpropertytype, &u.Udccourtfee, &u.Udcinterest, &u.Udcmartsonfee, &u.Udclast4digits, &u.Udcdatesent, &u.EquinoxLrn, &u.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
