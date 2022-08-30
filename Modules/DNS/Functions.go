package Carl_DNS

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

type DNS_Record struct {
	MX             []*net.MX
	NS             []string
	A              []net.IP
	TXT            []string
	CNAME          string
	PTR            []string
	SERVER         string
	STATUS         string
	METHOD         string
	Expires        string
	X_frame_opts   string
	Date           string
	Content_Len    string
	Cache_Control  string
	Set_Cookie     string
	SRV_CNAME_BASE string
	SRV_Target     []string
	SRV_Port       []uint16
	SRV_Priority   []uint16
	SRV_Weight     []uint16
}

var S DNS_Record
var IPs []net.IP

// grab A-AAAA records (IPv4 / IPv6)
func (Config *Carl_Data) A() {
	r, x := net.LookupIP(string(Config.Domain_Name))
	if x != nil {
		fmt.Println("[>>] CARL    ::: Log - Could not lookup the A and AAAA records of the domain -> ", x)
	} else {
		for _, A := range r {
			IPs = append(IPs, A)
		}
		S.A = IPs
	}
}

// CNAME
func (Config *Carl_Data) CNAME() {
	r, x := net.LookupCNAME(string(Config.Domain_Name))
	if x != nil {
		fmt.Println("[>>] CARL    ::: Log - Could not lookup CNAME records of the domain -> ", x)
	}
	S.CNAME = r
}

// PTR

func (Config *Carl_Data) PTR() {
	Config.A()
	for _, Q := range S.A {
		v := fmt.Sprint(Q)
		p, x := net.LookupAddr(v)
		if x != nil {
			fmt.Println("[>>] CARL    ::: Log - Could not lookup the PTR records for the domain -> ", x)
		} else {
			for _, l := range p {
				S.PTR = append(S.PTR, l)
			}
		}
	}
}

// NS
func (Config *Carl_Data) NS() {
	ns, x := net.LookupNS(string(Config.Domain_Name))
	if x != nil {
		fmt.Println("[>>] CARL    ::: Log - Could not lookup name server for the domain -> ", ns)
	} else {
		for _, n := range ns {
			a := fmt.Sprint(n)
			S.NS = append(S.NS, strings.Trim(a, "&{}."))
		}
	}
}

// MX
func (Config *Carl_Data) MX() {
	mx, x := net.LookupMX(string(Config.Domain_Name))
	if x != nil {
		fmt.Println("[>>] CARL    ::: Log - Could not lookup MX records for the domain -> ", x)
	}
	for _, m := range mx {
		S.MX = append(S.MX, m)
	}
}

// TXT
func (Config *Carl_Data) TXT() {
	txt, x := net.LookupTXT(string(Config.Domain_Name))
	if x != nil {
		fmt.Println("[>>] CARL    ::: Log - Could not lookup TXT records for the domain -> ", x)
	}
	for _, t := range txt {
		S.TXT = append(S.TXT, t)
	}
}

// Header
func (Config *Carl_Data) Head() {
	dom := "http://" + Config.Domain_Name
	f, x := http.Get(dom)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Body.Close()
		S.METHOD = f.Request.Method
		S.Date = strings.Trim(fmt.Sprint(f.Header.Values("date")), "[]")
		S.SERVER = strings.Trim(fmt.Sprint(f.Header.Values("server")), "[]")
		S.Expires = strings.Trim(fmt.Sprint(f.Header.Values("expires")), "[]")
		S.Cache_Control = strings.Trim(fmt.Sprint(f.Header.Values("cache-control")), "[]")
		S.STATUS = f.Status
		S.Set_Cookie = strings.Trim(fmt.Sprint(f.Header.Values("set-cookie")), "[]")
		S.X_frame_opts = strings.Trim(fmt.Sprint(f.Header.Values("x-frame-options")), "[]")
		S.Content_Len = fmt.Sprint(f.ContentLength)

	}
}

// srv
func (Config *Carl_Data) SRV() {
	c, r, x := net.LookupSRV("xmpp-server", "tcp", Config.Domain_Name)
	if x != nil {
		fmt.Println("[>>] CARL   :: Could not get the SRv records from the server, got error -> ", x)
	} else {
		S.SRV_CNAME_BASE = c
		for _, l := range r {
			S.SRV_Port = append(S.SRV_Port, l.Port)
			S.SRV_Priority = append(S.SRV_Priority, l.Priority)
			S.SRV_Target = append(S.SRV_Target, l.Target)
			S.SRV_Weight = append(S.SRV_Weight, l.Weight)
		}
	}
}

func (Config *Carl_Data) Generate() {
	f, x := json.MarshalIndent(S, "", " ")
	if x != nil {
		log.Fatal(x)
	}
	x = ioutil.WriteFile(Config.Filepath, f, 0644)
	if x != nil {
		log.Fatal(x)
	}
}
