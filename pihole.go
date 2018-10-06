package pihole

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Pihole struct {
	Url    string
	Pwhash string
}

type Version struct {
	Version int64 `json:"version,omitempty"`
}

type Status struct {
	Status string `json:"status,omitempty"`
}

type Type struct {
	Type string `json:"type,omitempty"`
}

type Period struct {
	Days    string `json:"days,omitempty"`
	Hours   string `json:"hours,omitempty"`
	Minutes string `json:"minutes,omitempty"`
}

type Gravity struct {
	FileExists bool   `json:"file_exists,omitempty"`
	Absolute   int64  `json:"absolute,omitempty"`
	Relative   Period `json:"relative,omitempty"`
}

type Summary struct {
	DomainsBeingBlocked  string  `json:"domains_being_blocked,omitempty"`
	DnsQueriesToday      string  `json:"dns_queries_today,omitempty"`
	AdsBlockedToday      string  `json:"ads_blocked_today,omitempty"`
	AdsPercentageToday   string  `json:"ads_percentage_today,omitempty"`
	UniqueDomains        string  `json:"unique_domains,omitempty"`
	QueriesForwarded     string  `json:"queries_forwarded,omitempty"`
	QueriesCached        string  `json:"queries_cached,omitempty"`
	ClientsEverSeen      string  `json:"clients_ever_seen,omitempty"`
	UniqueClients        string  `json:"unique_clients,omitempty"`
	DnsQueries_all_types string  `json:"dns_queries_all_types,omitempty"`
	ReplyNODATA          string  `json:"reply_NODATA,omitempty"`
	ReplyNXDOMAIN        string  `json:"reply_NXDOMAIN,omitempty"`
	ReplyCNAME           string  `json:"reply_CNAME,omitempty"`
	ReplyIP              string  `json:"reply_IP,omitempty"`
	Status               string  `json:"status,omitempty"`
	GravityLastUpdated   Gravity `json:"gravity_last_updated,omitempty"`
}

type SummaryRaw struct {
	DomainsBeingBlocked  int64   `json:"domains_being_blocked,omitempty"`
	DnsQueriesToday      int64   `json:"dns_queries_today,omitempty"`
	AdsBlockedToday      int64   `json:"ads_blocked_today,omitempty"`
	AdsPercentageToday   float64 `json:"ads_percentage_today,omitempty"`
	UniqueDomains        int64   `json:"unique_domains,omitempty"`
	QueriesForwarded     int64   `json:"queries_forwarded,omitempty"`
	QueriesCached        int64   `json:"queries_cached,omitempty"`
	ClientsEverSeen      int64   `json:"clients_ever_seen,omitempty"`
	UniqueClients        int64   `json:"unique_clients,omitempty"`
	DnsQueries_all_types int64   `json:"dns_queries_all_types,omitempty"`
	ReplyNODATA          int64   `json:"reply_NODATA,omitempty"`
	ReplyNXDOMAIN        int64   `json:"reply_NXDOMAIN,omitempty"`
	ReplyCNAME           int64   `json:"reply_CNAME,omitempty"`
	ReplyIP              int64   `json:"reply_IP,omitempty"`
	Status               string  `json:"status,omitempty"`
	GravityLastUpdated   Gravity `json:"gravity_last_updated,omitempty"`
}

type OverTimeData struct {
	DomainsOverTime map[string]int64 `json:"domains_over_time,omitempty"`
	AdsOverTime     map[string]int64 `json:"ads_over_time,omitempty"`
}

type TopItems struct {
	TopQueries map[string]int64 `json:"top_queries,omitempty"`
	TopAds     map[string]int64 `json:"top_ads,omitempty"`
}

type TopClients struct {
	TopSources map[string]int64 `json:"top_Sources,omitempty"`
}

type ForwardDestinations struct {
	ForwardDestinations map[string]float64 `json:"forward_destinations,omitempty"`
}

type QueryTypes struct {
	Querytypes map[string]float64 `json:"querytypes,omitempty"`
}

type Query struct {
	Timestamp int64
	Type      string
	Domain    string
	Client    string
	Answer    int
}

func New(url string, pwhash string) *Pihole {
	return &Pihole{Url: url, Pwhash: pwhash}
}

func (p *Pihole) GetType() (Type, error) {
	t := Type{}
	r, err := http.Get(p.Url + "?type")
	if err != nil {
		return t, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&t)
	return t, err
}

func (p *Pihole) GetVersion() (Version, error) {
	v := Version{}
	r, err := http.Get(p.Url + "?version")
	if err != nil {
		return v, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&v)
	return v, err
}

func (p *Pihole) GetStatus() (Status, error) {
	s := Status{}
	r, err := http.Get(p.Url + "?status")
	if err != nil {
		return s, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&s)
	return s, err
}

func (p *Pihole) GetSummary() (Summary, error) {
	s := Summary{}
	r, err := http.Get(p.Url + "?summary")
	if err != nil {
		return s, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&s)
	return s, err
}

func (p *Pihole) GetSummaryRaw() (SummaryRaw, error) {
	s := SummaryRaw{}
	r, err := http.Get(p.Url + "?summaryRaw")
	if err != nil {
		return s, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&s)
	return s, err
}

func (p *Pihole) GetOverTimeData10mins() (OverTimeData, error) {
	o := OverTimeData{}
	r, err := http.Get(p.Url + "?overTimeData10mins")
	if err != nil {
		return o, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&o)
	return o, err
}

func (p *Pihole) GetTopItems(items int64) (TopItems, error) {
	var url string
	if items <= 0 {
		url = fmt.Sprintf("%s?topItems&auth=%p", p.Url, p.Pwhash)
	} else {
		url = fmt.Sprintf("%s?topItems=%d&auth=%p", p.Url, items, p.Pwhash)
	}
	t := TopItems{}
	r, err := http.Get(url)
	if err != nil {
		return t, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&t)
	return t, err
}

func (p *Pihole) GetTopClients(items int64) (TopClients, error) {
	var url string
	if items <= 0 {
		url = fmt.Sprintf("%s?topClients&auth=%p", p.Url, p.Pwhash)
	} else {
		url = fmt.Sprintf("%s?topClients=%d&auth=%p", p.Url, items, p.Pwhash)
	}
	t := TopClients{}
	r, err := http.Get(url)
	if err != nil {
		return t, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&t)
	return t, err
}

func (p *Pihole) GetForwardDestinations() (ForwardDestinations, error) {
	f := ForwardDestinations{}
	r, err := http.Get(p.Url + "?getForwardDestinations&auth=" + p.Pwhash)
	if err != nil {
		return f, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&f)
	return f, err
}

func (p *Pihole) GetQueryTypes() (QueryTypes, error) {
	q := QueryTypes{}
	r, err := http.Get(p.Url + "?getQueryTypes&auth=" + p.Pwhash)
	if err != nil {
		return q, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&q)
	return q, err
}

type TwoD [][]string

type Qry struct {
	Data TwoD `json:"data,omitempty"`
}

func (p *Pihole) GetQueries() ([]Query, error) {
	q := make([]Query, 0)
	r, err := http.Get(p.Url + "?getAllQueries&auth=" + p.Pwhash)
	if err != nil {
		return q, err
	}
	defer r.Body.Close()
	qi := Qry{}
	err = json.NewDecoder(r.Body).Decode(&qi)

	for i := 0; i < len(qi.Data); i++ {
		query := Query{}
		query.Timestamp, _ = strconv.ParseInt(qi.Data[i][0], 10, 64)
		query.Type = qi.Data[i][1]
		query.Domain = qi.Data[i][2]
		query.Client = qi.Data[i][3]
		query.Answer, err = strconv.Atoi(qi.Data[i][4])
		q = append(q, query)
	}
	return q, err
}

func (p *Pihole) Enable() (Status, error) {
	s := Status{}
	r, err := http.Get(p.Url + "?enable&auth=" + p.Pwhash)
	if err != nil {
		return s, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&s)
	return s, err
}

func (p *Pihole) Disable(duration int64) (Status, error) {
	s := Status{}
	var url string
	if duration > 0 {
		url = fmt.Sprintf("%s?disable=%d&auth=%s", p.Url, duration, p.Pwhash)
	} else {
		url = fmt.Sprintf("%s?disable&auth=%s", p.Url, p.Pwhash)
	}
	r, err := http.Get(url)
	if err != nil {
		return s, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&s)
	return s, err
}
