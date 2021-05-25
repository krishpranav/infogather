package userlookup

import "sync"

type Website struct {
	Title  string
	Domain string
	Delete string
	Valid  bool
	Extra  string
}

var results []Website

func UserLookup(userres string) []Website {
	var wg sync.WaitGroup
	var arrNo []Website = noRedirSites(userres)
	for _, v := range arrNo {
		go checkUser(v, userres, false, &wg)
	}
	var arrYes []Website = redirSites(userres)
	for _, v := range arrYes {
		go checkUser(v, userres, true, &wg)
	}

	wg.Wait()
	return results
}
