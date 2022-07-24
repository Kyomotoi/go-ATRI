package antieffort

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Kyomotoi/go-ATRI/lib"
	"github.com/crackcell/gotabulate"
	log "github.com/sirupsen/logrus"
)

const (
	wakaURLpattern = `https:\/\/wakatime.com\/share\/@([a-zA-Z0-9].*?)\/([a-zA-Z0-9].*?).json`
	pluginDIR      = "data/plugins/anti_effort"
)

func init() {
	exi := lib.IsDir(pluginDIR)
	if !exi {
		err := os.MkdirAll(pluginDIR, 0777)
		if err != nil {
			panic("anti_effort: 创建文件夹: " + pluginDIR + " 失败, 请尝试手动创建")
		}
	}
}

func getData(groupID int64) AntiEffortModel {
	filePath := pluginDIR + "/" + strconv.FormatInt(groupID, 10) + ".json"
	if !lib.IsExists(filePath) {
		return AntiEffortModel{}
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Warning("anti_effort: 读取文件失败: " + filePath)
		return AntiEffortModel{}
	}
	var model AntiEffortModel
	_ = json.Unmarshal(data, &model)

	return model
}

func getEnabledGroups() []int64 {
	var groups []int64

	files, _ := os.ReadDir(pluginDIR)
	if files == nil {
		return groups
	}

	for _, file := range files {
		if !strings.Contains(file.Name(), "-ld") {
			groupID, _ := strconv.ParseInt(strings.Replace(file.Name(), ".json", "", -1), 10, 64)
			groups = append(groups, groupID)
		}
	}
	return groups
}

func addUser(groupID int64, userID int64, userNickname string, wakaURL string) (string, error) {
	match := regexp.MustCompile(wakaURLpattern).FindStringSubmatch(wakaURL)
	if match == nil {
		return "哥, 你提供的链接有问题啊", errors.New("wakaURL格式不正确")
	}

	wakaUserID := match[1]

	filePath := pluginDIR + "/" + strconv.FormatInt(groupID, 10) + ".json"
	if !lib.IsExists(filePath) {
		gen_data := &AntiEffortModel{
			UpdateTime: "",
			Data:       make([]AntiEffortUserModel, 0),
		}
		data, _ := json.Marshal(gen_data)

		err := os.WriteFile(filePath, data, 0777)
		if err != nil {
			log.Error("anti_effort: 写入文件失败: " + filePath)
			return "", err
		}
	}

	data := getData(groupID)
	for i := 0; i < len(data.Data); i++ {
		if data.Data[i].UserID == userID {
			return "你已经在卷王统计榜力！", nil
		}
	}

	resp, err := http.Get(wakaURL)
	if err != nil {
		log.Warning("anti_effort: 获取用户 " + strconv.FormatInt(groupID, 10) + " 的 wakatime 数据失败: " + wakaURL)
	}
	userWakaRawData, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	var wsed WakatimeShareEmbadData
	_ = json.Unmarshal(userWakaRawData, &wsed)

	recentCount, _ := strconv.ParseFloat(wsed.Data[len(wsed.Data)-1].GrandTotal.Decimal, 64)
	lastSevenDaysCount := float64(0)
	for i := 0; i < len(wsed.Data); i++ {
		_rd, _ := strconv.ParseFloat(wsed.Data[i].GrandTotal.Decimal, 64)
		lastSevenDaysCount += _rd
	}

	data.Data = append(data.Data, AntiEffortUserModel{
		UserID:             userID,
		UserNickname:       userNickname,
		WakaUserID:         wakaUserID,
		WakaURL:            wakaURL,
		LastSevenDaysCount: lastSevenDaysCount,
		RecentCount:        recentCount,
	})

	d, _ := json.Marshal(data)
	err = os.WriteFile(filePath, d, 0777)
	if err != nil {
		log.Error("anti_effort: 写入文件失败: " + filePath)
		return "", err
	}

	return "成功加入卷王统计榜！", nil
}

func delUser(groupID int64, userID int64) string {
	rawData := getData(groupID)
	if len(rawData.Data) > 0 {
		filePath := pluginDIR + "/" + strconv.FormatInt(groupID, 10) + ".json"
		data := rawData.Data
		for i := 0; i < len(data); i++ {
			if data[i].UserID == userID {
				data = append(data[:i], data[i+1:]...)
				break
			}
		}

		rawData.Data = data
		d, _ := json.Marshal(rawData)
		err := os.WriteFile(filePath, d, 0777)
		if err != nil {
			log.Error("anti_effort: 写入文件失败: " + filePath)
			log.Error("anti_effort: 删除用户 " + strconv.FormatInt(userID, 10) + " 失败")
			return ""
		}

		return "成功退出卷王统计榜！"
	}

	return "你未加入卷王统计榜捏"
}

func updateUserData(groupID int64, userID int64, updateMap AntiEffortUpdateDataMap) error {
	rawData := getData(groupID)
	if rawData.Data == nil {
		return errors.New("anti_effort: 暂无群启用该插件")
	}

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	filePath := pluginDIR + "/" + strconv.FormatInt(groupID, 10) + ".json"
	data := rawData.Data
	for i := 0; i < len(data); i++ {
		if data[i].UserID == userID {
			data[i].LastSevenDaysCount = updateMap.LastSevenDaysCount
			data[i].RecentCount = updateMap.RecentCount
			rawData.UpdateTime = nowTime
			break
		}
	}

	rawData.Data = data
	d, _ := json.Marshal(rawData)
	err := os.WriteFile(filePath, d, 0777)
	if err != nil {
		log.Error("anti_effort: 写入文件失败: " + filePath)
		log.Error("anti_effort: 更新用户 " + strconv.FormatInt(userID, 10) + " 的数据失败")
		return err
	}

	return nil
}

func updateData() {
	groupData := getEnabledGroups()
	if groupData == nil {
		return
	}

	for i := 0; i < len(groupData); i++ {
		rawData := getData(groupData[i])
		if rawData.Data == nil {
			continue
		}

		for j := 0; j < len(rawData.Data); j++ {
			resp, err := http.Get(rawData.Data[j].WakaURL)
			if err != nil {
				log.Warning("anti_effort: 获取用户 " + strconv.FormatInt(rawData.Data[j].UserID, 10) + " 的 wakatime 数据失败: " + rawData.Data[j].WakaURL)
				log.Warning("anti_effort: 该用户将从卷王榜中移除")
				delUser(groupData[i], rawData.Data[j].UserID)
			}
			userWakaRawData, _ := io.ReadAll(resp.Body)
			resp.Body.Close()

			var wsed WakatimeShareEmbadData
			_ = json.Unmarshal(userWakaRawData, &wsed)

			recentCount, _ := strconv.ParseFloat(wsed.Data[len(wsed.Data)-1].GrandTotal.Decimal, 64)
			log.Info(recentCount)
			lastSevenDaysCount := float64(0)
			for i := 0; i < len(wsed.Data); i++ {
				_rd, _ := strconv.ParseFloat(wsed.Data[i].GrandTotal.Decimal, 64)
				lastSevenDaysCount += _rd
			}

			updateMap := AntiEffortUpdateDataMap{
				LastSevenDaysCount: lastSevenDaysCount,
				RecentCount:        recentCount,
			}
			updateUserData(groupData[i], rawData.Data[j].UserID, updateMap)
		}
	}
}

func genRank(data AntiEffortModel, userID int64, typ string) string {
	var rankType string
	var tableType = "Today"
	var userRank = int16(0)
	var userCount = float64(0)
	var rankMsg string

	d := data.Data
	switch typ {
	case "recent_week":
		rankType = "近一周"
		tableType = "Last 7 Days"
		sort.Slice(
			d,
			func(i, j int) bool { return d[i].LastSevenDaysCount > d[j].LastSevenDaysCount },
		)
	case "global_today":
		rankType = "今日公共"
		sort.Slice(
			d,
			func(i, j int) bool { return d[i].RecentCount > d[j].RecentCount },
		)
	case "global_week":
		rankType = "近一周公共"
		tableType = "Last 7 Days"
		sort.Slice(
			d,
			func(i, j int) bool { return d[i].LastSevenDaysCount > d[j].LastSevenDaysCount },
		)
	default:
		rankType = "今日"
		sort.Slice(
			d,
			func(i, j int) bool { return d[i].RecentCount > d[j].RecentCount },
		)
	}

	for i := 0; i < len(d); i++ {
		if d[i].UserID == userID {
			userRank = int16(i + 1)
			userCount = d[i].RecentCount
			break
		}
	}

	table := [][]string{
		{"Rank", tableType, "Member"},
	}
	for i := 0; i < len(d); i++ {
		table = append(table, []string{
			strconv.Itoa(i + 1),
			strconv.FormatFloat(d[i].RecentCount, 'f', -1, 64) + " hrs",
			d[i].UserNickname,
		})
	}

	tabu := gotabulate.NewTabulator()
	tabu.SetFirstRowHeader(true)
	tabu.SetFormat("plain")

	tabled := tabu.Tabulate(lib.GetSliceByRange(table, 0, 11))

	if userRank != 0 && userCount != 0 {
		rankMsg = "你位于第 " + strconv.Itoa(int(userRank)) + " 名"
	} else {
		rankMsg = "暂无你的记录"
	}
	repo := "《" + rankType + "十佳卷王》\nUpdate Time: " + data.UpdateTime + "\n" + tabled + rankMsg
	return repo
}
