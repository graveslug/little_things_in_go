package main

//Parsing log file that reports back the amount of visits
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		sum map[string]int
		//stores unquie names in a slice
		domains []string
		total   int
		lines   int
	)

	sum = make(map[string]int)
	//scans log file line by line
	in := bufio.NewScanner(os.Stdin)
	//scans the log file
	for in.Scan() {
		lines++
		//since there are two fields we split them.
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("Wrong input: %v (line #%d)\n", fields, lines)
			return
		}
		// fmt.Printf("domain:%s\n - visits: %s\n", fields[0], fields[1])

		domain := fields[0]
		//converts visits to an int
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", fields[1], lines)
			return
		}
		//checks if domain exists if not it will create a new slice.
		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}
		total += visits
		sum[domain] += visits
	}
	//draws a line
	fmt.Println(strings.Repeat("-", 45))
	//sorts the domains
	sort.Strings(domains)
	//loops of the domain variable so it only grabs the unquie values
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
	fmt.Printf("%-30s %10d\n", "TOTAL", total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}

}
