package main

import (
	"errors"
	"fmt"
	"main/devices"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// initialize the Bmak Struct
var bm = Bmak{}

// Forminfo contains all Forminfo data from akamai sites. site:value format
var Forminfo = map[string]string{
	"converseUS": "0,0,0,0,1248,113,0;0,-1,0,0,-1,1932,0;",
	"bestbuyUS":  "0,0,0,0,1487,231,0;0,0,0,1,1627,1627,0;",
	"fedex":      "0,-1,0,1,2588,1468,0;0,-1,0,1,1386,447,0;1,0,0,1,1649,331,0;0,-1,0,0,1498,-1,0;0,-1,0,1,-1,1500,0;0,-1,0,1,-1,1684,0;0,-1,0,1,-1,1684,0;0,-1,0,1,-1,1684,0;0,-1,0,1,2563,-1,0;",
}

// getforminfo gets the Forminfo value for the site parameter
func getforminfo(site string) string {
	if site == "converseUS" {
		return Forminfo["converseUS"]
	} else if site == "bestbuyUS" {
		return Forminfo["bestbuyUS"]
	} else if site == "fedex" {
		return Forminfo["fedex"]
	} else {
		return "not yet available"
	}
}

// initiateValues initiate all needed bmak values
func initiateValues() {
	bm.StartTs = time.Now().UTC().UnixNano() / 1e6
	bm.Kact = ""
	bm.KeCnt = 0
	bm.Mact = ""
	bm.MmeCnt = 0
	bm.MduceCnt = 0
	bm.MeVel = 0
	bm.Pact = ""
	bm.PmeCnt = 0
	bm.PduceCnt = 0
	bm.PeVel = 0
	bm.Tact = ""
	bm.TmeCnt = 0
	bm.TduceCnt = 0
	bm.TeVel = 0
	bm.Doact = ""
	bm.DoeCnt = 0
	bm.DoeVel = 0
	bm.Dmact = ""
	bm.DmeCnt = 0
	bm.DmeVel = 0
	bm.Vcact = ""
	bm.VcCnt = 0
	bm.AjIndx = 0
	bm.AjSs = 0
	bm.AjType = -1
	bm.AjIndxDoact = 0
	bm.AjIndxDmact = 0
	bm.AjIndxTact = 0
	bm.MeCnt = 0
	bm.PeCnt = 0
	bm.TeCnt = 0
	bm.NavPerm = ""
	bm.Brv = 0
	bm.Hbcalc = false
	bm.Fmh = ""
	bm.Fmz = ""
	bm.SSH = ""
	bm.Wv = ""
	bm.Wr = ""
	bm.Weh = ""
	bm.Wl = 0
	bm.Ver = 1.7
	bm.KeCntLmt = 150
	bm.MmeCntLmt = 100
	bm.Ckie = "_abck"
	bm.Firstload = true
	bm.APIPublicKey = "afSbep8yjnZUjq3aL010jO15Sawj2VZfdYK8uY90uxq"
	bm.InitTime = 0
	bm.Tst = -1
	bm.Z1 = z1()
}

func pi(t float64) int64 {
	return int64(t)
}

// GetCfDate returns the current time in 13-digits format
func getCfDate() int64 {
	return (time.Now().UTC().UnixNano() / 1e6)
}

func updatet() int64 {
	return (getCfDate()) - (bm.StartTs)
}

// uar returns the userAgent
func uar(userAgent string) string {
	return regexp.MustCompile(`/\\|"/g`).ReplaceAllString(userAgent, "")
}

// fas returns an int that corresponds with the browser you have. Always static values
func fas(device devices.Device) (int, error) {
	if strings.Contains(device.Navigator.UserAgent, "Chrome") {
		return 30261693, nil
	} else if strings.Contains(device.Navigator.UserAgent, "Firefox") {
		return 26067385, nil
	} else {
		return 0, errors.New("fas function: use Chrome or Firefox")
	}
}

// returns z1
func z1() int64 {
	return bm.StartTs / (2016 * 2016)

}

// returns d3
func d3() int64 {
	return (getCfDate() % 1e7)
}

// Ab somehow converts a string to an int
func Ab(t string) int {

	var a = 0
	for e := 0; e < len(t); e++ {
		n := []rune(t)[e]
		if n < 128 {
			a += int(n)
		}
	}
	return a
}

// calDis takes a []int and returns a float64
func calDis(t []int) float64 {
	a := t[0] - t[1]
	e := t[2] - t[3]
	n := t[4] - t[5]
	return math.Floor(math.Sqrt(float64(a*a + e*e + n*n)))
}

// jrs takes an int and returns a []float64
func jrs(t int) []float64 {
	a := int(math.Floor(100000*randomNumberFloat() + 10000))
	e := strconv.Itoa(t * a)
	m := len(e) >= 18
	n := 0
	var o []int
	for len(o) < 6 {
		num, err := strconv.Atoi(e[n:(n + 2)])
		if err != nil {
			return []float64{}
		}

		o = append(o, num)

		if m {
			n += 3
		} else {
			n += 2
		}

	}
	return []float64{float64(a), calDis(o)}
}

// cc is a weird function:)
func cc(t int64) func(int64, int64) int64 {
	a := t % 4
	if 2 == a {
		a = 3
	}
	e := 42 + a
	var n func(int64, int64) int64
	n = func(t int64, a int64) int64 {
		return 0
	}
	if 42 == e {
		n = func(t int64, a int64) int64 {
			return t * a
		}
	} else if 43 == e {
		n = func(t int64, a int64) int64 {
			return t + a
		}
	} else {
		n = func(t int64, a int64) int64 {
			return t - a
		}
	}
	return n
}

// o9 returns an int
func o9() int64 {
	t := d3()
	a := t
	for n := 0; n < 5; n++ {
		o := math.Mod(math.Floor(float64(t)/math.Pow(10, float64(n))), 10)
		m := o + 1
		op := cc(int64(o))
		a = op(int64(a), int64(m))
	}
	return a * 3
}

// rir takes 4 ints and returns 1
func rir(t int, a int, e int, n int) int {
	if t > a && t <= e {
		t += n % (e - a)
		if t > e {
			t = t - e + a
		}
	}
	return t
}

// od takes 2 strings and returns 1
func od(t string, a string) string {
	var e []string
	n := len(a)
	if n > 0 {
		for o := 0; o < len(t); o++ {
			m := []rune(t)[o]
			r := string(t[o])
			i := a[(o % n)]
			m = rune(rir(int(m), 47, 57, int(i)))
			if m != []rune(t)[o] {
				r = string(m)
			}
			e = append(e, r)
		}
		if len(e) > 0 {
			return strings.Join(e[:], "")
		}
	}
	return t
}

func Gd(device devices.Device) string {
	userAgent := uar(device.Navigator.UserAgent)
	d := randomNumberFloat()
	xagg := 0
	psub := 0
	plen := 0
	bd := ""

	if strings.Contains(device.Navigator.UserAgent, "Chrome") {
		xagg = 12147
		psub = 20030107
		plen = 3
		bd = ",cpen:0,i1:0,dm:0,cwen:0,non:1,opc:0,fc:0,sc:0,wrc:1,isc:0,vib:1,bat:1,x11:0,x12:1"

	} else if strings.Contains(device.Navigator.UserAgent, "Firefox") {
		xagg = 11059
		psub = 20100101
		plen = 0
		bd = fmt.Sprintf(",cpen:0,i1:0,dm:0,cwen:0,non:1,opc:0,fc:1,sc:0,wrc:1,isc:%v,vib:1,bat:0,x11:0,x12:1", randomNumberInt(60, 100))
	} else {
		return ""
	}

	// this is the final return string
	return fmt.Sprintf(
		"%v,uaend,%v,%v,%v,Gecko,%v,0,0,0,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v.5,0,loc:",
		userAgent,
		xagg,
		psub,
		device.Navigator.Language,
		plen,
		z1(),
		d3(),
		device.Screen.AvailWidth,
		device.Screen.AvailHeight,
		device.Screen.Width,
		device.Screen.Height,
		device.Window.InnerWidth,
		device.Window.InnerHeight,
		device.Window.OuterWidth,
		bd,
		Ab(userAgent),
		floatToString(d, 16)[0:11]+floatToString(math.Floor(1000*d/2), -1),
		bm.StartTs/2)

}
