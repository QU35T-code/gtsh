package manager

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type session struct {
	ID       int
	OS       string
	IP       string
	username string
	hostname string
}

func sessions() {
	fmt.Println("Sessions\n")
	tab_writer := tabwriter.NewWriter(os.Stdout, 20, 8, 5, '\t', tabwriter.AlignRight)

	// Example
	c1 := session{1, "Linux", "127.0.0.1", "root", "heist"}
	c2 := session{2, "Linux", "10.10.14.36", "yakei", "response"}
	c3 := session{3, "Windows", "192.168.42.132", "qu35t", "flight"}
	session_list := []session{c1, c2, c3}
	fmt.Fprintln(tab_writer, "ID\tOS\tIP\tUsername\tHostname\t")
	fmt.Fprintln(tab_writer, "-------------------------------------------------------------------------------------------------------------")
	for _, v := range session_list {
		// TODO FIX this hardcoded line
		fmt.Fprint(tab_writer, v.ID, "\t", v.OS, "\t", v.IP, "\t", v.username, "\t", v.hostname, "\t\n")
		fmt.Fprintln(tab_writer, "-------------------------------------------------------------------------------------------------------------")
	}
	tab_writer.Flush()
	fmt.Println()
}
