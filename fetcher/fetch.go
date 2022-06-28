package fetcher

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var agents = []string{
	"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.75.14 (KHTML, like Gecko) Version/7.0.3 Safari/537.75.14",
}

func Fetch(targetUrl string) ([]byte, error) {
	newUrl := strings.Replace(targetUrl, "http://", "https://", 1)
	rand.Seed(time.Now().Unix())
	rand2 := rand.Intn(4)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, newUrl, nil)

	req.Header.Set("User-Agent", agents[rand2])
	cookie1 := "FSSBBIl1UgzbN7NO=51dbnEdF8fNggVtF7wwmk5HyCtjApglevSbfHF3Q4elgFmj4JcVMIVVx4BTXsjepPVOq1QR4P52pHV.7CMngacA; sid=535455ec-5515-46c0-ba5d-a249a714dd62; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1656394630; ec=a3YDH4hg-1606829845757-598260863f5051800642939; _efmdata=%2BXLbWAJbBGTPGM31Uy6a7ZPT1p%2FejDjQ5%2BnS4OhqYsYBsCG2Zn9ZxWGyVPSS5r8uZZLDZz2vxUnZfoRcWGm2QblBbr7mCFnMRljfhQp%2Bv4o%3D; _exid=akZA4N5y91U1X1Kcs9D%2Fmn2fgJmc7ILyzugMAFclKLbJ5QIm6B0ieVqIvWoKw1r9j5rxyQ8dTxiOZrZxGyVXWQ%3D%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1656405697; FSSBBIl1UgzbN7NP=53.j_zKt_XjLqqqDrjU.PXqXvR2oVDYOvy3n52mLNmowOrtUmG7bMtgOrX73jtfSOstvjbJJQOXrki29tS3058FmvkZl7JIxtkcsRpnGdV0AZ4USSw51W0l6tgcwkkA6udSLLNgYR7GncFJc8MRAIHA..eXlAS3F5K3iGTwtUwG0Bz2VypmnXW1h7s1QY6szX.IUBEuo3oySM9pZtWB8TnlxmC69ooEnnWqpf29DS9uVFG91zPf96RHkCyv9TedUNLO2VvgPYzBbXAHN6uq3K9ybrgtlasBKlQvOeffJC3YO3EhLZySBb4Y5mb_vML1fD9"
	req.Header.Add("cookie", cookie1)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
