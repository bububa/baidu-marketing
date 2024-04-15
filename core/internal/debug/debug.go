package debug

import (
	"bytes"
	"encoding/json"
	"io"
	"log"

	"github.com/bububa/baidu-marketing/util"
)

func PrintError(err error, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [ERROR]", err)
}

func PrintStringResponse(str string, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [RESPONSE]", str)
}

func PrintGetRequest(url string, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [API] GET", url)
}

func PrintPostJSONRequest(url string, body []byte, debug bool) {
	if !debug {
		return
	}
	const format = "[DEBUG] [API] JSON POST %s\n" +
		"http request body:\n%s\n"

	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	json.Indent(buf, body, "", "\t")
	log.Printf(format, url, body)
}

func PrintPostMultipartRequest(url string, mp map[string]string, debug bool) {
	if !debug {
		return
	}
	const format = "[DEBUG] [API] multipart/form-data POST %s\n" +
		"http request body:\n%s\n"

	bs, _ := json.MarshalIndent(mp, "", "\t")
	log.Printf(format, url, bs)
}

func DecodeJSONHttpResponse(r io.Reader, v interface{}, debug bool) error {
	if !debug {
		return json.NewDecoder(r).Decode(v)
	}
	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	body2 := body
	buf := bytes.NewBuffer(make([]byte, 0, len(body2)+1024))
	if err := json.Indent(buf, body2, "", "    "); err == nil {
		body2 = buf.Bytes()
	}
	log.Printf("[DEBUG] [API] http response body:\n%s\n", body2)

	return json.Unmarshal(body, v)
}
