package validation

import (
	"net/url"
	"regexp"
	"strings"
)

type UrlInfo struct {
	Domain string
	Http   string
	Https  string
}

func IsValid(_url string) (bool, UrlInfo, error) {

	IsDomain := CheckDomain(_url)
	IsUrl := CheckUrl(_url)

	_UrlInfo := UrlInfo{}
	isError := false

	if IsDomain {

		us := url.URL{
			Scheme: "https",
			Host:   _url,
			Path:   "/",
		}

		up := url.URL{
			Scheme: "http",
			Host:   _url,
			Path:   "/",
		}

		_UrlInfo = UrlInfo{
			Https:  us.String(),
			Http:   up.String(),
			Domain: _url,
		}

	} else if IsUrl {

		us, _ := url.ParseRequestURI(_url)
		us.Scheme = "https"
		us.Path = "/"

		up, _ := url.ParseRequestURI(_url)
		up.Scheme = "http"
		up.Path = "/"

		uh, _ := url.ParseRequestURI(_url)

		_UrlInfo = UrlInfo{
			Https:  us.String(),
			Http:   up.String(),
			Domain: uh.Hostname(),
		}

	} else {
		isError = true
	}

	return (!isError && (IsUrl || IsDomain)), _UrlInfo, nil
}

func CheckUrl(_url string) bool {
	urlInfo, err := url.Parse(_url)
	return (err == nil && urlInfo.Host != "" && urlInfo.Scheme != "")
}

func CheckDomain(_url string) bool {
	RegExp := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z
		]{2,3})$`)
	return RegExp.MatchString(_url)
}

func CompareDomainForUrl(url1 string, url2 string) bool {
	urlInfo1, err1 := url.Parse(url1)
	urlInfo2, err2 := url.Parse(url2)
	return (err1 == nil && err2 == nil && urlInfo1.Host != "" && urlInfo1.Scheme != "" && urlInfo2.Host != "" && urlInfo2.Scheme != "" && (strings.Compare(urlInfo1.Hostname(), urlInfo2.Hostname()) > 0 && strings.Compare(urlInfo1.Hostname(), urlInfo2.Hostname()) == 0))
}
