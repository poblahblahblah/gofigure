package operatingsystem

import (
  "io/ioutil"
  "os"
  "strings"
)

func Load() string {

    // redhat/centos/scientific
    // if /etc/redhat-release exists then we're dealing with redhat, centos, or scientific
	if _, err := os.Stat("/etc/redhat-release"); err == nil {

	    results, err := ioutil.ReadFile("/etc/redhat-release")
	        if err != nil {
	            panic(err)
	        }
	    return(string(results))
		    
	}
	
    // ubuntu
    // if /etc/lsb-release exists then we're dealing with an ubuntu host
    if _, err := os.Stat("/etc/lsb-release"); err == nil {

        // Read in /etc/lsb-release
        results, err := ioutil.ReadFile("/etc/lsb-release")
        if err != nil {
            panic(err)
        }

        // scan through results w/ strings.Split
        for _, line :=  range strings.Split(string(results), "\n") {
            s := strings.Split(line, "=")
            if s[0] == "DISTRIB_DESCRIPTION" {
              return(strings.Replace(s[1], `"`, "", -1))
            }
        }
    }

	return(string("error: cannot detect operating system"))
}

