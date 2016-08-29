package utils

import (
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/httplib"
)

func EncodeSha256(data string) (encodeStr string) {
	hash := sha256.New()
	hash.Write([]byte(data))
	encodeData := hash.Sum(nil)
	encodeStr = fmt.Sprintf("%x", encodeData)

	return
}

func ConvXstrtoInt(xStr string) (value int) {
	valueint64, _ := strconv.ParseInt(xStr, 16, 0)
	value = int(valueint64)
	return value
}

func ConvXstrtoFloat32(xStr string) (value float32) {
	valueFloat64, _ := strconv.ParseFloat(xStr, 16)
	value = float32(valueFloat64)
	return value
}

func GetDate() (dataStr string) {
	dateLayout := "2006-01-02"
	sr := time.Now().Unix()
	dataStr = time.Unix(sr, 0).Format(dateLayout)
	return
}

func GetTime() (timeStr string) {
	timeLayout := "2006-01-02 15:04:05"
	sr := time.Now().Unix()
	dataStr := time.Unix(sr, 0).Format(timeLayout)
	dateArr := strings.Split(dataStr, " ")
	timeStr = dateArr[1]
	return
}

func GetTimeMin() (timeStr string) {
	timeLayout := "2006-01-02 15:04:05"
	sr := time.Now().Unix()
	dataStr := time.Unix(sr, 0).Format(timeLayout)
	dateArr := strings.Split(dataStr, " ")
	timeStr = dateArr[1]
	rs := []rune(timeStr)
	timeStr = string(rs[0:5])
	return
}

func GetTimeStamp(dataTimeStr string) (timeStamp int64) {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	timePased, _ := time.ParseInLocation(timeLayout, dataTimeStr, loc)
	timeStamp = timePased.Unix()
	return
}

func HttpPost(url, req_str string) (recv string, err error) {

	req := httplib.Post(url).SetTimeout(5*time.Second, 5*time.Second)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Body(req_str)

	recv, err = req.String()

	//
	return recv, err

}

func SortMaps(sortMap map[string]string, appKey string) (sortedStr string) {

	sorted_keys := make([]string, 0)
	for k, _ := range sortMap {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	for v, k := range sorted_keys {
		fmt.Printf("k=%v, v=%v\n", k, sortMap[k])
		if v == 0 {
			sortedStr = k + "=" + sortMap[k]
		} else {
			sortedStr = sortedStr + "&" + k + "=" + sortMap[k]
		}
	}

	return
}
