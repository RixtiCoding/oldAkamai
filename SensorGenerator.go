package main

import (
	"encoding/json"
	"fmt"
	"main/devices"
	"strconv"
)

func (t *Task) GenerateSensor(sensorNumber int, site string, device devices.Device) string {
	initiateValues()
	e := t.getabck(site)
	n := Gd(device)
	i := "do_en,dm_en,t_en"
	b := "https://www.bestbuy.com/?intl=nosplash"
	d := "-1,0"
	s := bm.KeVel + bm.MeVel + bm.DoeVel + bm.DmeVel + bm.TeVel + bm.PeVel
	l := "PiZtE"
	u := jrs(int(bm.StartTs))
	// underline := getCfDate() - bm.StartTs
	bm.D2 = bm.Z1 / 23
	f := bm.D2 / 6
	p, _ := fas(device)
	var y string
	var w string
	var v int
	// var h int
	S := ""
	C := ""
	x := ""
	nck, _ := strconv.Atoi(bm.NCk)

	if sensorNumber == 1 {
		y = "0,0,0,0,1,0,0"
		v = 0
		d = "0,0"
		w = fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,", bm.KeVel+1, bm.MeVel+32, bm.TeVel+32, bm.DoeVel, bm.DmeVel, bm.PeVel, s, updatet(), bm.InitTime, int(bm.StartTs), -999999, bm.D2, bm.KeCnt, bm.MeCnt, f, bm.PeCnt, bm.TeCnt, randomNumberInt(1, 4), bm.Ta, nck, e, Ab(e), "-1", "-1", p, l, u[0], u[1], v)
		FirstSensorData := fmt.Sprintf("7a74G7m23Vrp0o5c9254531.7-1,2,-94,-100,%v-1,2,-94,-101,%v-1,2,-94,-105,%v-1,2,-94,-102,%v-1,2,-94,-108,-1,2,-94,-110,-1,2,-94,-117,-1,2,-94,-111,-1,2,-94,-109,-1,2,-94,-114,-1,2,-94,-103,-1,2,-94,-112,%v-1,2,-94,-115,%v-1,2,-94,-106,%v-1,2,-94,-119,%v-1,2,-94,-122,%v-1,2,-94,-123,%v-1,2,-94,-124,%v-1,2,-94,-126,%v-1,2,-94,-127,8", n, i, getforminfo(site+"US"), getforminfo(site+"US"), b, w, d, "1", y, S, C, x)
		P := 24 ^ Ab(FirstSensorData)
		SecondSensorData := fmt.Sprintf(FirstSensorData+"-1,2,-94,-70,-1-1,2,-94,-80,94-1,2,-94,-116,%v-1,2,-94,-118,%v-1,2,-94,-129,-1,2,-94,-121,;%v;-1;0", o9(), P, getCfDate()-getCfDate())
		return SecondSensorData

	} else if sensorNumber == 2 {
		d = "8,1"
		w = fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,", bm.KeVel+1, bm.MeVel+32, bm.TeVel+32, bm.DoeVel, bm.DmeVel, bm.PeVel, s, updatet(), bm.InitTime, int(bm.StartTs), randomNumberInt(1, 40), bm.D2, bm.KeCnt, bm.MeCnt, f, bm.PeCnt, bm.TeCnt, randomNumberInt(1, 4), bm.Ta, nck, e, Ab(e), "-1", "-1", p, l, u[0], u[1], v)
		FirstSensorData := fmt.Sprintf("7a74G7m23Vrp0o5c9253601.7-1,2,-94,-100,%v-1,2,-94,-101,%v-1,2,-94,-105,%v-1,2,-94,-102,%v-1,2,-94,-108,-1,2,-94,-110,-1,2,-94,-117,-1,2,-94,-111,-1,2,-94,-109,-1,2,-94,-114,-1,2,-94,-103,-1,2,-94,-112,%v-1,2,-94,-115,%v-1,2,-94,-106,%v-1,2,-94,-119,%v-1,2,-94,-122,%v-1,2,-94,-123,%v-1,2,-94,-124,%v-1,2,-94,-126,%v-1,2,-94,-127,8", n, i, getforminfo(site+"US"), getforminfo(site+"US"), b, w, d, "1", y, S, C, x)
		P := 24 ^ Ab(FirstSensorData)
		SecondSensorData := fmt.Sprintf(FirstSensorData+"-1,2,-94,-70,-1-1,2,-94,-80,94-1,2,-94,-116,%v-1,2,-94,-118,%v-1,2,-94,-129,-1,2,-94,-121,;%v;-1;0", o9(), P, getCfDate()-getCfDate())
		fmt.Println(SecondSensorData)

	}
	return ""

}

func createDevice1() devices.Device {
	device1json := `{"_id":{"$oid":"6086fb16254cd100152b36fa"},"navigator":{"appCodeName":"Mozilla","appName":"Netscape","appVersion":"5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.85 Safari/537.36","language":"en-US","languages":["en-US","en","ro"],"userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.85 Safari/537.36","vendor":"Google Inc.","product":"Gecko","productSub":"20030107","platform":"Win32","vibrate":true,"getBattery":true,"credentials":true,"appMinorVersion":false,"bluetooth":true,"storage":true,"getGamepads":true,"getStorageUpdates":false,"hardwareConcurrency":true,"mediaDevices":true,"mozAlarms":false,"mozConnection":false,"mozIsLocallyAvailable":false,"mozPhoneNumberService":false,"msManipulationViewsEnabled":false,"permissions":true,"registerProtocolHandler":true,"requestMediaKeySystemAccess":true,"requestWakeLock":false,"sendBeacon":true,"serviceWorker":true,"storeWebWideTrackingException":false,"webkitGetGamepads":false,"webkitTemporaryStorage":true,"cookieEnabled":true,"javaEnabled":false,"doNotTrack":-1,"plugins":["Chrome PDF Plugin","Chrome PDF Viewer","Native Client"]},"window":{"ActiveXObject":false,"innerHeight":937,"innerWidth":1920,"outerWidth":1920,"outerHeight":1040,"devicePixelRatio":1,"addEventListener":true,"XMLHttpRequest":true,"XDomainRequest":false,"orientation":"undefined","DeviceOrientationEvent":true,"DeviceMotionEvent":true,"TouchEvent":true,"chrome":{"app":{"isInstalled":false,"InstallState":{"DISABLED":"disabled","INSTALLED":"installed","NOT_INSTALLED":"not_installed"},"RunningState":{"CANNOT_RUN":"cannot_run","READY_TO_RUN":"ready_to_run","RUNNING":"running"}},"runtime":{"OnInstalledReason":{"CHROME_UPDATE":"chrome_update","INSTALL":"install","SHARED_MODULE_UPDATE":"shared_module_update","UPDATE":"update"},"OnRestartRequiredReason":{"APP_UPDATE":"app_update","OS_UPDATE":"os_update","PERIODIC":"periodic"},"PlatformArch":{"ARM":"arm","ARM64":"arm64","MIPS":"mips","MIPS64":"mips64","X86_32":"x86-32","X86_64":"x86-64"},"PlatformNaclArch":{"ARM":"arm","MIPS":"mips","MIPS64":"mips64","X86_32":"x86-32","X86_64":"x86-64"},"PlatformOs":{"ANDROID":"android","CROS":"cros","LINUX":"linux","MAC":"mac","OPENBSD":"openbsd","WIN":"win"},"RequestUpdateCheckStatus":{"NO_UPDATE":"no_update","THROTTLED":"throttled","UPDATE_AVAILABLE":"update_available"}}},"prototype_bind":true,"PointerEvent":true,"sessionStorage":true,"localStorage":true,"indexedDB":true,"FileReader":true,"HTMLElement":false,"webRTC":true,"mozInnerScreenY":0},"document":{"documentMode":"undefined","webdriver":false,"driver":false,"selenium":false,"hidden":false,"webkitHidden":false},"other":{"CC_ON":false,"InstallTrigger":false,"prototype_forEach":true,"imul":true,"parseInt":true,"hypot":true,"value1":false,"XPathResult":true},"performance":{"timeOrigin":1619458826528.4058,"timing":{"connectStart":1619458826589,"navigationStart":1619458826528,"loadEventEnd":1619458831024,"domLoading":1619458830846,"secureConnectionStart":1619458826644,"fetchStart":1619458826533,"domContentLoadedEventStart":1619458830999,"responseStart":1619458830837,"responseEnd":1619458830840,"domInteractive":1619458830999,"domainLookupEnd":1619458826589,"redirectStart":0,"requestStart":1619458826755,"unloadEventEnd":0,"unloadEventStart":0,"domComplete":1619458831024,"domainLookupStart":1619458826589,"loadEventStart":1619458831024,"domContentLoadedEventEnd":1619458830999,"redirectEnd":0,"connectEnd":1619458826755},"navigation":{"type":0,"redirectCount":0}},"canvas":{"value1":"-739578230","value2":"-1395479418"},"fonts_optm":"12,15,16,17,40,44","fonts":"12,15,16,17,40,44","rCFP":"-1444815886","ssh":"8ac13bd32cdcba031fb13564dbf5fe05473dd7da0579641e30ca66f64fa6d8d9","mr":"39,226,42,42,68,63,17,10,10,49,7,8,13,262,","brave":"0","wv":"Google Inc. (NVIDIA)","wr":"ANGLE (NVIDIA, NVIDIA GeForce GTX 970 Direct3D11 vs_5_0 ps_5_0, D3D11-27.21.14.6109)","weh":"95f5b71fe531f867faa814bdd4050dd8057206d53ecec1163523560525884870","wl":33,"fmh":"56c63bdbf3b4ec968f51613af09e44576bbe4a63fabcbf9785eaaac645c4d2a9","fpValstr":"-739578230;-1395479418;dis;,7,8;true;true;true;-180;true;24;24;true;false;-1","np":"11321144241322243122","screen":{"availHeight":1040,"availLeft":0,"availTop":0,"availWidth":1920,"colorDepth":24,"height":1080,"pixelDepth":24,"width":1920},"__v":0}`
	bs := []byte(device1json)
	var Device1 devices.Device
	err := json.Unmarshal(bs, &Device1)
	if err != nil {
		fmt.Println("Unmarshal Unsuccessful!")
	}
	return Device1
}
