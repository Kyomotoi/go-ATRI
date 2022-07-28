package device

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

const (
	digits       = "0123456789"
	asciiUpper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	asciiLower   = "abcdefghijklmnopqrstuvwxyz"
	asciiLetters = asciiUpper + asciiLower
	hexDigits    = "0123456789abcdefABCDEF"
)

type Gener struct {
	Seed int64
}

func newGener(account int64) *Gener {
	rand.Seed(int64(account))
	return &Gener{
		Seed: account,
	}
}

func (g *Gener) randomString(chr string, l int) string {
	b := make([]byte, l)
	for i := range b {
		b[i] = chr[rand.Intn(len(chr))]
	}
	return string(b)
}

func (g *Gener) IMEI() string {
	mem := g.randomString(digits, 14)
	return mem + checkSum(mem)
}

func checkSum(mem string) string {
	var sum int
	for _, v := range mem {
		sum += int(v - '0')
	}
	return fmt.Sprintf("%X", sum%10)
}

func (g *Gener) SSID(prefix string) string {
	return prefix + g.randomString(asciiUpper+digits, 6)
}

func (g *Gener) androidDevice() (string, AndroidDevice) {
	builds := loadAndroidBuilds()
	devices := loadAndroidDevices()

	build := builds[rand.Intn(len(builds))]
	device := devices[rand.Intn(len(devices))]

	return build.AndroidID, device
}

func (g *Gener) bootID() string {
	return uuid.New().String()
}

func (g *Gener) procVersion() string {
	majorVersion := rand.Intn(2) + 3
	minorVersion := rand.Intn(19)
	patchVersion := rand.Intn(99)

	buildID := g.randomString(asciiLetters, 8)
	ver := fmt.Sprintf("%d.%d.%d-%s", majorVersion, minorVersion, patchVersion, buildID)

	mailDomain := g.randomString(strings.Replace(hexDigits, "ABCDEF", "", -1), 12) + ".source.android.com"

	return fmt.Sprintf("Linux version %s (android-build@%s)", ver, mailDomain)
}

func (g *Gener) ipAddress() []int {
	return []int{
		rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256),
	}
}

func (g *Gener) macAddress() string {
	return fmt.Sprintf("%02X:%02X:%02X:%02X:%02X:%02X",
		rand.Intn(256), rand.Intn(256), rand.Intn(256),
		rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func (g *Gener) incremental() string {
	return fmt.Sprintf("%d", rand.Intn(2<<32))
}
