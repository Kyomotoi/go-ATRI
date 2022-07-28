package device

import (
	_ "embed"

	"github.com/jszwec/csvutil"
)

//go:embed android_builds.csv
var androidBuildsCSVFile []byte

//go:embed android_devices.csv
var androidDevicesCSVFile []byte

func loadAndroidBuilds() []AndroidBuild {
	var ab []AndroidBuild

	_ = csvutil.Unmarshal(androidBuildsCSVFile, &ab)
	return ab
}

func loadAndroidDevices() []AndroidDevice {
	var ad []AndroidDevice

	_ = csvutil.Unmarshal(androidDevicesCSVFile, &ad)
	return ad
}
