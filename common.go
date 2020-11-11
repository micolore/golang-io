package tmst

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func GetNowTime() string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return t
}

func ReqTCP(url string) *net.TCPConn {
	server := url
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	return conn
}


func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
func HandlerData(reqType string, data []byte) []byte {

	length := len(reqType) + len(data) + 1
	var buf []byte = make([]byte, length+4)
	index := 0
	jsonLength := IntToBytes(length)
	for i := 0; i < 4; i++ {
		buf[index] = jsonLength[i]
		index++
	}
	typeByte := []byte(reqType)
	for i := 0; i < len(typeByte); i++ {
		buf[index] = typeByte[i]
		index++
	}
	buf[index] = ','
	index++
	for i := 0; i < len(data); i++ {
		buf[index] = data[i]
		index++
	}
	return buf
}

func CastStrToMap(jsonStr string) (dat map[string]interface{}) {

	var mapData map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &mapData)
	return mapData
}

func ReadTcpData(conn *net.TCPConn) (rb []byte) {
	rb = make([]byte, 1024*10)
	_, err := conn.Read(rb)
	if err != nil {
		return nil
	}
	return rb
}

// CheckError check error ...
func CheckError(err error) {
	if err != nil {
		log.Fatal("an error!", err.Error())
	}
}

func ReqTcp(url string) *net.TCPConn {
	server := url
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		//os.Exit(1)
		return nil
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	return conn
}

func CastByteToMap(data []byte) (dat map[string]interface{}) {
	str := string(data)
	firstIndex := strings.Index(str, ",{")
	lastIndex := strings.LastIndex(str, "}")
	jsonStr := str[firstIndex+1 : lastIndex+1]
	return CastStrToMap(jsonStr)
}

func WriteDataToFile(fileName string, content string) error {

	// 以只写的模式，打开文件
	//f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766)

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

// ReadLine 读取文件的每一行
func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

// StrToTime 字符串转time
func StrToTime(str string) time.Time {
	p, _ := time.Parse("2006-01-02 15:04:05", str)
	return p
}

// TimeToStr 日期转字符串
func TimeToStr(t time.Time) string {
	return t.Format("2006-01-02 03:04:05")
}
