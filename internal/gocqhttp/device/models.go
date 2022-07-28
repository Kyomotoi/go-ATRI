package device

type AndroidBuild struct {
	AndroidID string `csv:"Android_id"`
	Version   string `csv:"Version"`
}

type AndroidDevice struct {
	Branding string `csv:"Branding,omitempty"`
	Name     string `csv:"Name,omitempty"`
	Device   string `csv:"Device"`
	Model    string `csv:"Model"`
}

type Version struct {
	Incremental string
	Release     string
	CodeName    string
	SDK         int
}

type DeviceInfo struct {
	Display      string  `json:"display"`
	Product      string  `json:"product"`
	Device       string  `json:"device"`
	Board        string  `json:"board"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	WifiBSSID    string  `json:"wifi_bssid"`
	WifiSSID     string  `json:"wifi_ssid"`
	AndroidID    string  `json:"android_id"`
	BootID       string  `json:"boot_id"`
	ProcVersion  string  `json:"proc_version"`
	MacAddress   string  `json:"mac_address"`
	IPAddress    []int   `json:"ip_address"`
	IMEI         string  `json:"imei"`
	IMSIMD5      string  `json:"imsi_md5"`
	Incremental  string  `json:"incremental"`
	Protocol     int     `json:"protocol"`
	BootLodaer   string  `json:"bootloader"`
	FingerPrint  string  `json:"finger_print"`
	BaseBand     string  `json:"baseband"`
	SIM          string  `json:"sim"`
	OSType       string  `json:"os_type"`
	APN          string  `json:"apn"`
	VendorName   string  `json:"vendor_name"`
	VendorOSName string  `json:"vendor_os_name"`
	Version      Version `json:"version"`
}
