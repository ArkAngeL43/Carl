package Carl_DNS

import (
	"fmt"
	"net/http"
	"strings"
)

func (Config *Carl_Data) Parse_Dt() {
	switch Config.Record {
	case "mx":
		Config.MX()
		for _, l := range S.MX {
			fmt.Println("MX  Host  |  ", l.Host)
			fmt.Println("MX  Pref  |  ", l.Pref)
			fmt.Println(":::::::::::")
		}
	case "a":
		Config.A()
		for _, l := range S.A {
			fmt.Println("IP        | ", l)
			fmt.Println("===========")
			fmt.Println(" Is multicast                 ; ", l.IsMulticast())
			fmt.Println(" Is loopback                  ; ", l.IsLoopback())
			fmt.Println(" Is GlobalUnicast             ; ", l.IsGlobalUnicast())
			fmt.Println(" Is Link Local Multicast      ; ", l.IsLinkLocalMulticast())
			fmt.Println(" Is Link Local Unicast        ; ", l.IsLinkLocalUnicast())
			fmt.Println(" Is Interface Local Multicast ; ", l.IsInterfaceLocalMulticast())
			fmt.Println(" Is Private IP Address        ; ", l.IsPrivate())
			fmt.Println(" Is Unspecified               ; ", l.IsUnspecified())
			fmt.Println(" Defualt mask                 ; ", l.DefaultMask())
		}
	case "ns":
		Config.NS()
		for _, l := range S.NS {
			fmt.Println("NS    ; ", l)
		}
	case "txt":
		Config.TXT()
		for _, k := range S.TXT {
			fmt.Println("TXT   ; ", k)
		}
	case "cname":
		Config.CNAME()
		fmt.Println("CNAME ; ", string(fmt.Sprint(S.CNAME)))
	case "ptr":
		Config.PTR()
		for o, l := range S.PTR {
			fmt.Println("PTR [ ", o, " ]    ; ", l)
		}
	case "srv":
		Config.SRV()
		fmt.Print("\n\n")
		fmt.Println("SRV Target    SRV Port    SRV Priority   SRV Weight")
		fmt.Print("\n")
		for i := 0; i < len(S.SRV_Target); i++ {
			fmt.Printf("%v   **   %v   **   %v   **   %v \n", S.SRV_Target[i], S.SRV_Port[i], S.SRV_Priority[i], S.SRV_Weight[i])
		}
	case "server":
		url := "http://" + Config.Domain_Name
		d, x := http.Get(url)
		if x != nil {
			fmt.Println(x)
		}
		defer d.Body.Close()
		fmt.Println("\n\nServer -> ", strings.Trim(fmt.Sprint(d.Header.Values("server")), "[]"))
	case "head":
		url := "http://" + Config.Domain_Name
		d, x := http.Get(url)
		if x != nil {
			fmt.Println(x)
		}
		defer d.Body.Close()
		fmt.Println("\n\n\n\n\n")
		for _, v := range d.Header {
			fmt.Println(v)
		}
	}
}
