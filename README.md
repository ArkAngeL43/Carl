```
                    ____ ____ ____ _    /
                    |___ |--| |--< |___. 
                        Carl - V 1.0
```

Carl is a program that readers who read the Go cyber weapons development and advanced go article/book will make, this script is introducing advanced golang
and getting DNS records from a given domain using flags. Here is how you use CARL 


# Supported records 

```
 - mx       | Will get MX records
 - a        | Will get IPv4
 - ns       | Will get name servers
 - txt      | Will get the TXT records
 - cname    | Will get the Canonical Name
 - ptr      | Will reverse DNS
 - srv      | Will get the service record
 - server   | Will get the server
 - head     | Will get response headers
 - *        | Will run all of these options and is needed for JSON output
```

# Simple usage 

`go run main.go --domain some_domain.org --record mx/a/ns/txt/cname/ptr/srv/server/head/*`

# Advanced usage and JSON output 

`go run main.go --domain="www.scanme.org" --record="*" -o -f="Filename_22222.json"`

#### JSON OUTPUT ####

```json
{
 "MX": null,
 "NS": null,
 "A": [
  "2600:3c01::f03c:91ff:fe18:bb2f",
  "45.33.32.156",
  "2600:3c01::f03c:91ff:fe18:bb2f",
  "45.33.32.156"
 ],
 "TXT": null,
 "CNAME": "www.scanme.org.",
 "PTR": [
  "scanme.nmap.org.",
  "scanme.nmap.org.",
  "scanme.nmap.org.",
  "scanme.nmap.org."
 ],
 "SERVER": "Apache/2.4.7 (Ubuntu)",
 "STATUS": "200 OK",
 "METHOD": "GET",
 "Expires": "",
 "X_frame_opts": "",
 "Date": "Mon, 08 Aug 2022 04:32:55 GMT",
 "Content_Len": "-1",
 "Cache_Control": "",
 "Set_Cookie": "",
 "SRV_CNAME_BASE": "",
 "SRV_Target": null,
 "SRV_Port": null,
 "SRV_Priority": null,
 "SRV_Weight": null
}
```


