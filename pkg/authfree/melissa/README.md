melissa email phone number look up

# types
- email types

```go
type Email struct {
	Recordid                      string
	Deliverabilityconfidencescore string
	Results                       string
	Emailaddress                  string
	Mailboxname                   string
	Domainname                    string
	Topleveldomain                string
	Topleveldomainname            string
	Datechecked                   string
	Emailageestimated             string
	Domainageestimated            string
	Domainexpirationdate          string
	Domaincreateddate             string
	Domainupdateddate             string
	Domainemail                   string
	Domainorganization            string
	Domainaddress1                string
	Domainlocality                string
	Domainadministrativearea      string
	Domainpostalcode              string
	Domaincountry                 string
	Domainavailability            string
	Domaincountrycode             string
	Domainprivateproxy            string
	Privacyflag                   string
	Mxserver                      string
}
```

- phone types

```go
type Phone struct {
	Recordid                     string
	Results                      string
	Phonenumber                  string
	Administrativearea           string
	Countryabbreviation          string
	Countryname                  string
	Carrier                      string
	Callerid                     string
	Dst                          string
	Internationalphonenumber     string
	Language                     string
	Latitude                     string
	Locality                     string
	Longitude                    string
	Phoneinternationalprefix     string
	Phonecountrydialingcode      string
	Phonenationprefix            string
	Phonenationaldestinationcode string
	Phonesubscribernumber        string
	Utc                          string
	Timezonecode                 string
	Timezonename                 string
	Postalcode                   string
}
```

# Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/krishpranav/infogather/pkg/authfree/melissa"
)

func main() {
	x, err := melissa.IPLookup("yourKeyHere", "1.1.1.1")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(x.Ispname)
}
```