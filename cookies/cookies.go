package cookies

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type ExportableCookieJar struct {
	jar        *cookiejar.Jar
	allCookies map[url.URL][]*http.Cookie
}

func NewExportableCookieJar() *ExportableCookieJar {
	realJar, _ := cookiejar.New(nil)

	e := &ExportableCookieJar{
		jar:        realJar,
		allCookies: make(map[url.URL][]*http.Cookie),
	}

	return e
}

func (jar *ExportableCookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {

	jar.allCookies[*u] = cookies
	jar.jar.SetCookies(u, cookies)
}

func (jar *ExportableCookieJar) Cookies(u *url.URL) []*http.Cookie {
	return jar.jar.Cookies(u)
}

func (jar *ExportableCookieJar) ExportAllCookies() map[url.URL][]*http.Cookie {

	copied := make(map[url.URL][]*http.Cookie)

	for u, c := range jar.allCookies {
		copied[u] = c
	}

	return copied
}

type Cookie struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func GetCookie(jar *ExportableCookieJar, name string) string {

	for _, cookies := range jar.ExportAllCookies() {
		for _, cookie := range cookies {
			if name == cookie.Name {
				return cookie.Value
			}
		}
	}

	return ""
}

func SetCookie(jar *ExportableCookieJar, name, value, domain, URL string) (*ExportableCookieJar, error) {

	new := NewExportableCookieJar()

	for host, cookies := range jar.ExportAllCookies() {

		cks := make([]*http.Cookie, 0)

		for _, cookie := range cookies {
			if name != cookie.Name {
				cks = append(cks, cookie)
			}
		}

		new.SetCookies(&host, cks)
	}

	cookie := http.Cookie{
		Name:  name,
		Value: value,
	}

	u, err := url.Parse(URL)

	if err != nil {
		return nil, err
	}

	new.SetCookies(u, []*http.Cookie{&cookie})

	return new, nil
}

func PopCookie(jar *ExportableCookieJar, name string) *ExportableCookieJar {

	new := NewExportableCookieJar()

	for host, cookies := range jar.ExportAllCookies() {

		cks := make([]*http.Cookie, 0)

		for _, cookie := range cookies {
			if name != cookie.Name {
				cks = append(cks, cookie)
			}
		}

		new.SetCookies(&host, cks)
	}

	return new
}

func ExtractChallenge(jar *ExportableCookieJar) string {

	for _, cookies := range jar.ExportAllCookies() {
		for _, cookie := range cookies {

			c := Cookie{
				Name:  cookie.Name,
				Value: cookie.Value,
			}

			if strings.Contains(cookie.Value, "||") {

				return fmt.Sprintf("%v=%v", c.Name, c.Value)
			} else {
				continue
			}
		}
	}
	return ""
}

func ExtractCookies(jar *ExportableCookieJar) []Cookie {

	r := make([]Cookie, 0)

	for _, cookies := range jar.ExportAllCookies() {
		for _, cookie := range cookies {

			c := Cookie{
				Name:  cookie.Name,
				Value: cookie.Value,
			}
			r = append(r, c)
		}
	}

	return r
}
