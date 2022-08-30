package main

import (
	"fmt"
	Carl_Bann "main/Modules/Banner"
	Carl_DNS "main/Modules/DNS"
	Carl_Opts "main/Modules/Options"
	"os"

	"github.com/spf13/pflag"
)

var (
	flags = pflag.FlagSet{SortFlags: false}
	CLIS  Carl_Opts.Options
	Carl  Carl_DNS.Carl_Data
)

func main() {}

//
//
// flag base
//
//
//

func init() {
	flags.StringVarP(&CLIS.Domain, "domain", "d", "", "Specify a domain name to scan  | Required")
	flags.StringVarP(&CLIS.Record, "record", "r", "", "Specify a DNS record to search | Required")
	flags.BoolVarP(&CLIS.Output, "output", "o", false, "Enable JSON output File       | Optional")
	flags.StringVarP(&CLIS.Output_Directory, "filepath", "f", "", "The output file directory | Needed if you use -o/--output")
	flags.Parse(os.Args[1:])
	Carl_Bann.Out()

	if CLIS.Domain == "" {
		fmt.Println("[!] ERORR: You will need t o specify a domain name here using the `--domain/-d` flag")
		os.Exit(0)
	} else {
		Carl.Domain_Name = CLIS.Domain
	}
	if CLIS.Record == "" {
		fmt.Println("[!] ERROR: You need to specify a DNS record you want to lookup an example being MX")
		fmt.Println("[*] All possible lookup options are")
		fmt.Println(" - mx       | Will get MX records")
		fmt.Println(" - a        | Will get IPv4")
		fmt.Println(" - ns       | Will get name servers")
		fmt.Println(" - txt      | Will get the TXT records")
		fmt.Println(" - cname    | Will get the Canonical Name")
		fmt.Println(" - ptr      | Will reverse DNS")
		fmt.Println(" - srv      | Will get the service record")
		fmt.Println(" - server   | Will get the server")
		fmt.Println(" - head     | Will get response headers")
		fmt.Println(" - *        | Will run all of these options and is needed for JSON output")
		os.Exit(0)
	} else {
		Carl.Record = CLIS.Record
		Carl.Parse_Dt()
	}
	if CLIS.Output {
		if CLIS.Record != "*" {
			fmt.Println("[!] Error: This option is only reserved for the option (*) as a record")
			fmt.Println("[!] Error: You will need to specify the `--record` flag with * if you want")
			fmt.Println("[!] Error: to continue, this is due to output and formatting reasons")
			os.Exit(0)
		} else {
			if CLIS.Output_Directory == "" {
				fmt.Println("[!] Error: Can not run this script, the output directory you specified is NIL / Empty. ")
				fmt.Println("[!] Error: Please use (--filepath/-f) to specify the directory of the filepath you want")
				fmt.Println("[*] Example: ./main --domain example.com --record * --output --filepath /home/Desktop/File.json")
				fmt.Println("[*] NOTE   ---- You MUST specify the FULL path with the filename at the end of the json file")
				os.Exit(0)
			} else {
				Carl.Filepath = CLIS.Output_Directory
				Carl.A()
				Carl.CNAME()
				Carl.MX()
				Carl.NS()
				Carl.TXT()
				Carl.PTR()
				Carl.Head()
				Carl.SRV()
				Carl.Generate()
				fmt.Printf("File [ %s ] Has been generated\n\n", Carl.Filepath)
			}
		}
	}
}
