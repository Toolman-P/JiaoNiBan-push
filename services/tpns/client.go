package tpns

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	GuangzhouHost = "api.tpns.tencent.com"
	ShanghaiHost  = "api.tpns.sh.tencent.com"
	HongkongHost  = "api.tpns.hk.tencent.com"
	SingaporeHost = "api.tpns.sgp.tencent.com"
)

type Client struct {
	host string
	sign string
}

func NewClient(host string, accessId uint32, secretKey string) *Client {
	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		host = "https://" + host
	}

	cli := &Client{
		host: host,
		sign: base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d:%s", accessId, secretKey))),
	}
	return cli
}

func (c *Client) Do(req *Request) (Response, error) {
	return c.DoTimeout(req, 0)
}

func (c *Client) DoTimeout(req *Request, timeout time.Duration) (Response, error) {
	var resp Response
	if err := req.Validate(); err != nil {
		return resp, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, c.host+"/v3/push/app", bytes.NewReader(body))
	if err != nil {
		return resp, err
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.sign))

	cli := &http.Client{}
	if timeout > 0 {
		cli.Timeout = timeout
	}

	httpResp, err := cli.Do(httpReq)
	if err != nil {
		return resp, err
	}
	defer httpResp.Body.Close()
	body, err = ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return resp, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("response error, status: %s, body: %s", httpResp.Status, string(body))
	}

	if err = json.Unmarshal(body, &resp); err != nil {
		return resp, err
	}

	return resp, nil

}

func (c *Client) Upload(file string, duration time.Duration) (UploadResponse, error) {
	var resp UploadResponse
	fp, err := os.Open(file)
	if err != nil {
		return resp, err
	}

	defer fp.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()

	part, err := writer.CreateFormFile("file", filepath.Base(fp.Name()))
	if err != nil {
		return resp, err
	}
	io.Copy(part, fp)
	writer.Close()
	httpReq, err := http.NewRequest(http.MethodPost, c.host+"/v3/push/package/upload", body)
	if err != nil {
		return resp, err
	}
	httpReq.Header.Add("Content-Type", writer.FormDataContentType())
	httpReq.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.sign))

	cli := &http.Client{}
	if duration > 0 {
		cli.Timeout = duration
	}

	httpResp, err := cli.Do(httpReq)
	if err != nil {
		return resp, err
	}
	defer httpResp.Body.Close()
	data, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return resp, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("response error, status: %s, body: %s", httpResp.Status, string(data))
	}

	if err = json.Unmarshal(data, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
