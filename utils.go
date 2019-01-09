package utils

import (
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"math/rand"
	"os"
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

func ConvXstrtoIntDecimal(xStr string) (value int) {
	valueint64, _ := strconv.ParseInt(xStr, 10, 0)
	value = int(valueint64)
	return value
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

func CreateStamp() (dataTimeStr string) {
	timeLayout := "2006-01-02 15:04:05"
	sr := time.Now().Unix()
	dataTimeStr = time.Unix(sr, 0).Format(timeLayout)
	dataTimeStr = strings.Replace(dataTimeStr, " ", "", -1)
	dataTimeStr = strings.Replace(dataTimeStr, ":", "", -1)
	dataTimeStr = strings.Replace(dataTimeStr, "-", "", -1)
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

//  生成最大范围内随机数
func GenerateRandnumWithin(max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max)

	return randNum
}

//  生成一个区间范围的随机数
func GenerateRangeNumBetween(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min

	return randNum
}

func appendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}

//
//GenRandomString generate a random string with specific length
func GenRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
