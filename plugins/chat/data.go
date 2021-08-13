package chat

import (
	"encoding/json"
	"github.com/Kyomotoi/go-ATRI/utils"
	log "github.com/sirupsen/logrus"
	jieba "github.com/yanyiwu/gojieba"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func checkDIR() error {
	var fileDIR string
	fileDIR = "data/database/chat"
	_, err := os.Stat(fileDIR)
	if err != nil {
		if !os.IsExist(err) {
			err := os.MkdirAll(fileDIR, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func generateData() error {
	err := checkDIR()
	if err != nil {
		log.Error("无法创建文件夹（data/database/chat）")
		return err
	}

	url := "https://cdn.jsdelivr.net/gh/Kyomotoi/AnimeThesaurus/data.json"
	resp, err := http.Get(url)
	if err != nil {
		log.Error("请求失败: "+url)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("请求失败: "+url)
		return err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		return err
	}

	filePath := "data/database/chat/kimo.json"
	_ = ioutil.WriteFile(filePath, data, 0644)
	return nil
}

func UpdateData() error {
	err := checkDIR()
	if err != nil {
		log.Error("无法创建文件夹：data/database/chat")
		return err
	}

	filePath := "data/database/chat/kimo.json"
	if utils.IsExists(filePath) {
		err = os.Remove(filePath)
		if err != nil {
			log.Error("删除文件失败: "+filePath+"，请尝试手动删除")
			return err
		}
	}
	err = generateData()
	if err != nil {
		return err
	}
	return nil
}

func loadData() (map[string][]string, error) {
	var d map[string][]string
	err := checkDIR()
	if err != nil {
		log.Error("无法创建文件夹：data/database/chat")
		return map[string][]string{}, err
	}

	filePath := "data/database/chat/kimo.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		if !utils.IsExists(filePath) {
			err := generateData()
			if err != nil {
				return map[string][]string{}, err
			}
			data, err = ioutil.ReadFile(filePath)
			if err != nil {
				log.Error("读取文件失败: "+filePath)
				return map[string][]string{}, err
			}
		} else {
			log.Error("读取文件失败: "+filePath)
			return map[string][]string{}, err
		}
	}
	err = json.Unmarshal(data, &d)
	if err != nil {
		log.Error("")
		return map[string][]string{}, err
	}
	return d, nil
}

func StoreUserNickname(userID string, nickname string) error {
	var userNicknameData map[string]interface{}

	err := checkDIR()
	if err != nil {
		return err
	}
	filePath := "data/database/chat/users.json"
	isExist := utils.IsExists(filePath)
	if !isExist {
		_ = ioutil.WriteFile(filePath, []byte("{}"), 0644)
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error("读取文件失败: "+filePath)
		return err
	}

	err = json.Unmarshal(data, &userNicknameData)
	if err != nil {
		log.Error("解析JSON文件失败: "+filePath)
		return err
	}
	userNicknameData[userID] = nickname
	newData, err := json.Marshal(userNicknameData)
	_ = ioutil.WriteFile(filePath, newData, 0644)
	return nil
}

func loadUserNickname(userID string) string {
	var d map[string]string

	err := checkDIR()
	if err != nil {
		return "你"
	}
	filePath := "data/database/chat/users.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "你"
	}

	// If can't find object, will return null string.
	err = json.Unmarshal(data, &d)
	if err != nil {
		return "你"
	}
	result := d[userID]
	if result == "" {
		result = "你"
	}
	return result
}

func dealWord(word string) []string {
	var words []string
	x := jieba.NewJieba()
	defer x.Free()

	words = x.CutAll(word)
	return words
}

func Kimo(message string, userID string) (string, error) {
	kw := dealWord(message)
	rand.Shuffle(len(kw), func(i, j int) {
		kw[i], kw[j] = kw[j], kw[i]
	})

	data, err := loadData()
	if err != nil {
		log.Error("加载文件失败：data/database/chat/kimo.json")
		return "", err
	}

	repo := ""
	for i := range kw {
		w := kw[i]
		a := []string{kw[i]}
		if len(a) == 2 {
			if a[0] == a[1] {
				w = a[0]
			}
		}

		if _, in := data[w]; in {
			repo = data[w][rand.Intn(len(data[w]))]
		}
	}

	if repo == "" {
		var t []string
		for i := range data {
			t = append(t, i)
		}
		rand.Shuffle(len(t), func(i, j int) {
			t[i], t[j] = t[i], t[j]
		})

		for i := range t {
			w := t[i]
			if strings.Contains(message, w) {
				repo = data[w][rand.Intn(len(data[w]))]
			}
		}
	}

	userNickname := loadUserNickname(userID)
	result := strings.Replace(repo, "你", userNickname, -1)
	return result, nil
}
