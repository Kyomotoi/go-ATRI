package device

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Kyomotoi/go-ATRI/lib"
	log "github.com/sirupsen/logrus"
)

const resourceDIR = "data/protocols/gocqhttp/"

func genDeviceInfo(account int64, proc int) DeviceInfo {
	gener := newGener(account)

	id, device := gener.androidDevice()
	mac := gener.macAddress()
	ssid := gener.SSID("TP-LINK_")
	bootid := gener.bootID()
	pv := gener.procVersion()
	ipa := gener.ipAddress()
	imei := gener.IMEI()
	imsiMD5 := lib.MD5(imei)
	inc := gener.incremental()

	return DeviceInfo{
		Display:      id,
		Product:      device.Name,
		Device:       device.Device,
		Board:        device.Device,
		Brand:        device.Branding,
		Model:        device.Model,
		WifiBSSID:    mac,
		WifiSSID:     ssid,
		AndroidID:    id,
		BootID:       bootid,
		ProcVersion:  pv,
		MacAddress:   mac,
		IPAddress:    ipa,
		IMEI:         imei,
		IMSIMD5:      imsiMD5,
		Incremental:  inc,
		Protocol:     proc,
		BootLodaer:   "unknown",
		FingerPrint:  fmt.Sprintf("%s/%s/%s:10/%s/%s:user/release-keys", device.Branding, device.Name, device.Device, id, inc),
		BaseBand:     "",
		SIM:          "T-Mobile",
		OSType:       "android",
		APN:          "wifi",
		VendorName:   device.Branding,
		VendorOSName: "android",
		Version: Version{
			Incremental: inc,
			Release:     "10",
			CodeName:    "REL",
			SDK:         29,
		},
	}
}

func InitDevice(version string, account int64, proc int) {
	device := genDeviceInfo(account, proc)

	filePath := resourceDIR + version + "/device.json"
	if !lib.IsExists(filePath) {
		log.Warn("Driver: 未找到 gocqhttp 设备文件, 将自动生成")

		data, _ := json.Marshal(device)
		err := os.WriteFile(filePath, data, 0777)
		if err != nil {
			log.Warn("Driver: 生成 gocqhhtp 设备文件失败")
			log.Warn("Driver: 将由 gocqhhtp 生成具备 'mirai' 标识的设备文件. 账号风控概率将增大")
		}
		log.Info("Driver: 生成成功")
	}

	log.Info("Driver: 成功加载适用于 gocqhttp 的设备文件")
}
