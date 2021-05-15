package main

import (
	"bytes"
	"fmt"
	tls "github.com/refraction-networking/utls"
	"github.com/x04/cclient"
	"io/ioutil"
	"main/cookies"
	"net/http"
	"strconv"
)

type Task struct {
	Client    http.Client
	CookieJar *cookies.ExportableCookieJar
}

func (t *Task) NewClient() http.Client {
	client, err := cclient.NewClient(tls.HelloChrome_Auto)
	if err != nil {
		fmt.Println(err)
	}
	return client
}

func (t *Task) SetupClient() {
	jar := cookies.NewExportableCookieJar()
	t.Client.Jar = jar
	t.CookieJar = jar
}

func (t *Task) getabck(site string) string {
	req, err := http.NewRequest("GET", "https://www.bestbuy.com/resource/e13635113ern224c6e851859da1219f2", nil)
	if err != nil {
		fmt.Println("Erroare", err)
	}
	resp, err := t.Client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	_ = resp.Body.Close()
	test1 := string(b)
	if len(test1) > 2 {
		fmt.Println("ok")
	}

	if err != nil {
		fmt.Println(err)
	}
	abckcookie := cookies.GetCookie(t.CookieJar, "_abck")
	return abckcookie
}

func (t *Task) SendFirstPost(sensor string) {
	fmt.Println(cookies.ExtractCookies(t.CookieJar))
	var textsensor = fmt.Sprintf(`{"sensor_data":"%v"}`, sensor)
	fmt.Println(textsensor)
	length := len(textsensor)
	fmt.Println(length)
	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Encoding": "gzip, deflate, br",
		"Accept-Language": "en-US,en;q=0.9,ro;q=0.8",
		"Cache-Control":   "no-cache",
		"Content-Length":  strconv.Itoa(length),
		"Content-Type":    "text/plain;charset=UTF-8",
		"Origin":          "https://www.bestbuy.com",
		"Pragma":          "no-cache",
		"Referer":         "https://www.bestbuy.com/?intl=nosplash",
		"sec-fetch-mode":  "cors",
		"sec-fetch-dest":  "empty",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.85 Safari/537.36",
	}
	var body = []byte(textsensor)
	req, err := http.NewRequest("POST", "https://www.bestbuy.com/resource/e13635113ern224c6e851859da1219f2", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := t.Client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	_ = resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	fmt.Println(resp.Status)
}

func main() {
	tsk := Task{}
	tsk.Client = tsk.NewClient()
	tsk.SetupClient()
	sensor := tsk.GenerateSensor(1, "bestbuy", createDevice1())
	tsk.SendFirstPost(sensor)
	fmt.Println(cookies.ExtractChallenge(tsk.CookieJar))

}
