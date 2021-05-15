package devices

type Device struct {
	ID struct {
		Oid string `json:"$oid"`
	} `json:"_id"`
	Navigator struct {
		AppCodeName                   string   `json:"appCodeName"`
		AppName                       string   `json:"appName"`
		AppVersion                    string   `json:"appVersion"`
		Language                      string   `json:"language"`
		Languages                     []string `json:"languages"`
		UserAgent                     string   `json:"UserAgent"`
		Vendor                        string   `json:"vendor"`
		Product                       string   `json:"product"`
		ProductSub                    string   `json:"productSub"`
		Platform                      string   `json:"platform"`
		Vibrate                       bool     `json:"vibrate"`
		GetBattery                    bool     `json:"getBattery"`
		Credentials                   bool     `json:"credentials"`
		AppMinorVersion               bool     `json:"appMinorVersion"`
		Bluetooth                     bool     `json:"bluetooth"`
		Storage                       bool     `json:"storage"`
		GetGamepads                   bool     `json:"getGamepads"`
		GetStorageUpdates             bool     `json:"getStorageUpdates"`
		HardwareConcurrency           bool     `json:"hardwareConcurrency"`
		MediaDevices                  bool     `json:"mediaDevices"`
		MozAlarms                     bool     `json:"mozAlarms"`
		MozConnection                 bool     `json:"mozConnection"`
		MozIsLocallyAvailable         bool     `json:"mozIsLocallyAvailable"`
		MozPhoneNumberService         bool     `json:"mozPhoneNumberService"`
		MsManipulationViewsEnabled    bool     `json:"msManipulationViewsEnabled"`
		Permissions                   bool     `json:"permissions"`
		RegisterProtocolHandler       bool     `json:"registerProtocolHandler"`
		RequestMediaKeySystemAccess   bool     `json:"requestMediaKeySystemAccess"`
		RequestWakeLock               bool     `json:"requestWakeLock"`
		SendBeacon                    bool     `json:"sendBeacon"`
		ServiceWorker                 bool     `json:"serviceWorker"`
		StoreWebWideTrackingException bool     `json:"storeWebWideTrackingException"`
		WebkitGetGamepads             bool     `json:"webkitGetGamepads"`
		WebkitTemporaryStorage        bool     `json:"webkitTemporaryStorage"`
		CookieEnabled                 bool     `json:"cookieEnabled"`
		JavaEnabled                   bool     `json:"javaEnabled"`
		DoNotTrack                    int      `json:"doNotTrack"`
		Plugins                       []string `json:"plugins"`
	} `json:"navigator"`
	Window struct {
		ActiveXObject          bool   `json:"ActiveXObject"`
		InnerHeight            int    `json:"innerHeight"`
		InnerWidth             int    `json:"innerWidth"`
		OuterWidth             int    `json:"outerWidth"`
		OuterHeight            int    `json:"outerHeight"`
		Fmz                    int    `json:"devicePixelRatio"`
		AddEventListener       bool   `json:"addEventListener"`
		XMLHTTPRequest         bool   `json:"XMLHttpRequest"`
		XDomainRequest         bool   `json:"XDomainRequest"`
		Orientation            string `json:"orientation"`
		DeviceOrientationEvent bool   `json:"DeviceOrientationEvent"`
		DeviceMotionEvent      bool   `json:"DeviceMotionEvent"`
		TouchEvent             bool   `json:"TouchEvent"`
		Chrome                 struct {
			App struct {
				IsInstalled  bool `json:"isInstalled"`
				InstallState struct {
					DISABLED     string `json:"DISABLED"`
					INSTALLED    string `json:"INSTALLED"`
					NOTINSTALLED string `json:"NOT_INSTALLED"`
				} `json:"InstallState"`
				RunningState struct {
					CANNOTRUN  string `json:"CANNOT_RUN"`
					READYTORUN string `json:"READY_TO_RUN"`
					RUNNING    string `json:"RUNNING"`
				} `json:"RunningState"`
			} `json:"app"`
			Runtime struct {
				OnInstalledReason struct {
					CHROMEUPDATE       string `json:"CHROME_UPDATE"`
					INSTALL            string `json:"INSTALL"`
					SHAREDMODULEUPDATE string `json:"SHARED_MODULE_UPDATE"`
					UPDATE             string `json:"UPDATE"`
				} `json:"OnInstalledReason"`
				OnRestartRequiredReason struct {
					APPUPDATE string `json:"APP_UPDATE"`
					OSUPDATE  string `json:"OS_UPDATE"`
					PERIODIC  string `json:"PERIODIC"`
				} `json:"OnRestartRequiredReason"`
				PlatformArch struct {
					ARM    string `json:"ARM"`
					ARM64  string `json:"ARM64"`
					MIPS   string `json:"MIPS"`
					MIPS64 string `json:"MIPS64"`
					X8632  string `json:"X86_32"`
					X8664  string `json:"X86_64"`
				} `json:"PlatformArch"`
				PlatformNaclArch struct {
					ARM    string `json:"ARM"`
					MIPS   string `json:"MIPS"`
					MIPS64 string `json:"MIPS64"`
					X8632  string `json:"X86_32"`
					X8664  string `json:"X86_64"`
				} `json:"PlatformNaclArch"`
				PlatformOs struct {
					ANDROID string `json:"ANDROID"`
					CROS    string `json:"CROS"`
					LINUX   string `json:"LINUX"`
					MAC     string `json:"MAC"`
					OPENBSD string `json:"OPENBSD"`
					WIN     string `json:"WIN"`
				} `json:"PlatformOs"`
				RequestUpdateCheckStatus struct {
					NOUPDATE        string `json:"NO_UPDATE"`
					THROTTLED       string `json:"THROTTLED"`
					UPDATEAVAILABLE string `json:"UPDATE_AVAILABLE"`
				} `json:"RequestUpdateCheckStatus"`
			} `json:"runtime"`
		} `json:"chrome"`
		PrototypeBind   bool `json:"prototype_bind"`
		PointerEvent    bool `json:"PointerEvent"`
		SessionStorage  bool `json:"sessionStorage"`
		LocalStorage    bool `json:"localStorage"`
		IndexedDB       bool `json:"indexedDB"`
		FileReader      bool `json:"FileReader"`
		HTMLElement     bool `json:"HTMLElement"`
		WebRTC          bool `json:"webRTC"`
		MozInnerScreenY int  `json:"mozInnerScreenY"`
	} `json:"window"`
	Document struct {
		DocumentMode string `json:"documentMode"`
		Webdriver    bool   `json:"webdriver"`
		Driver       bool   `json:"driver"`
		Selenium     bool   `json:"selenium"`
		Hidden       bool   `json:"hidden"`
		WebkitHidden bool   `json:"webkitHidden"`
	} `json:"document"`
	Other struct {
		CCON             bool `json:"CC_ON"`
		InstallTrigger   bool `json:"InstallTrigger"`
		PrototypeForEach bool `json:"prototype_forEach"`
		Imul             bool `json:"imul"`
		ParseInt         bool `json:"parseInt"`
		Hypot            bool `json:"hypot"`
		Value1           bool `json:"value1"`
		XPathResult      bool `json:"XPathResult"`
	} `json:"other"`
	Performance struct {
		TimeOrigin float64 `json:"timeOrigin"`
		Timing     struct {
			ConnectStart               int64 `json:"connectStart"`
			NavigationStart            int64 `json:"navigationStart"`
			LoadEventEnd               int64 `json:"loadEventEnd"`
			DomLoading                 int64 `json:"domLoading"`
			SecureConnectionStart      int64 `json:"secureConnectionStart"`
			FetchStart                 int64 `json:"fetchStart"`
			DomContentLoadedEventStart int64 `json:"domContentLoadedEventStart"`
			ResponseStart              int64 `json:"responseStart"`
			ResponseEnd                int64 `json:"responseEnd"`
			DomInteractive             int64 `json:"domInteractive"`
			DomainLookupEnd            int64 `json:"domainLookupEnd"`
			RedirectStart              int   `json:"redirectStart"`
			RequestStart               int64 `json:"requestStart"`
			UnloadEventEnd             int   `json:"unloadEventEnd"`
			UnloadEventStart           int   `json:"unloadEventStart"`
			DomComplete                int64 `json:"domComplete"`
			DomainLookupStart          int64 `json:"domainLookupStart"`
			LoadEventStart             int64 `json:"loadEventStart"`
			DomContentLoadedEventEnd   int64 `json:"domContentLoadedEventEnd"`
			RedirectEnd                int   `json:"redirectEnd"`
			ConnectEnd                 int64 `json:"connectEnd"`
		} `json:"timing"`
		Navigation struct {
			Type          int `json:"type"`
			RedirectCount int `json:"redirectCount"`
		} `json:"navigation"`
	} `json:"performance"`
	Canvas struct {
		Value1 string `json:"value1"`
		Value2 string `json:"value2"`
	} `json:"canvas"`
	FontsOptm string `json:"fonts_optm"`
	Fonts     string `json:"fonts"`
	RCFP      string `json:"rCFP"`
	SSH       string `json:"ssh"`
	Mr        string `json:"mr"`
	Brave     string `json:"brave"`
	Wv        string `json:"wv"`
	Wr        string `json:"wr"`
	Weh       string `json:"weh"`
	Wl        int    `json:"wl"`
	Fmh       string `json:"fmh"`
	FpValstr  string `json:"fpValstr"`
	Np        string `json:"np"`
	Screen    struct {
		AvailHeight int `json:"availHeight"`
		AvailLeft   int `json:"availLeft"`
		AvailTop    int `json:"availTop"`
		AvailWidth  int `json:"availWidth"`
		ColorDepth  int `json:"colorDepth"`
		Height      int `json:"height"`
		PixelDepth  int `json:"pixelDepth"`
		Width       int `json:"width"`
	} `json:"screen"`
	V int `json:"__v"`
}
