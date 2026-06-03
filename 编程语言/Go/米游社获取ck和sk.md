```go
package main

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"log"
	mrand "math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// --------------------------- 配置与常量 ---------------------------

const publicKeyPEM = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDvekdPMHN3AYhm/vktJT+YJr7cI5DcsNKqdsx5DZX0gDuWFuIjzdwButrIYPNmRJ1G8ybDIF7oDW2eEpm5sMbL9zs
9ExXCdvqrn51qELbqj0XxtMTIpaCHFSI50PfPpTFV9Xt/hmyVwokoOXFlAEgCn+Q
CgGs52bFoYMtyi+xEQIDAQAB
-----END PUBLIC KEY-----`

var regex = `^#?(米哈?游社?登(录|陆|入)|登(录|陆|入)米哈?游社?)`

const errorTips = "登录失败，请检查日志\nhttps://Yunzai.TRSS.me"

var httpClient = &http.Client{
	Timeout: 20 * time.Second,
}

// --------------------------- 全局状态 ---------------------------

type Event struct {
	UserID    int64
	Msg       string
	IsPrivate bool
	SelfID    int64
	MessageID int64
}

var accounts = struct {
	sync.RWMutex
	m map[int64]*Event
}{m: make(map[int64]*Event)}

var Running = struct {
	sync.RWMutex
	m map[int64]interface{}
}{m: make(map[int64]interface{})}

// --------------------------- 帮助函数 ---------------------------

func randomString(n int) string {
	const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[mrand.Intn(len(letters))]
	}
	return string(b)
}

func encryptData(plain []byte) (string, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return "", errors.New("failed to parse PEM block containing the public key")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return "", errors.New("not RSA public key")
	}
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, pub, plain)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func md5Hex(data string) string {
	h := md5.Sum([]byte(data))
	return hex.EncodeToString(h[:])
}

func ds(data string) string {
	t := time.Now().Unix()
	r := randomString(6)
	salt := "JwYDpKvLj6MrMqqYU6jTKF17KNO2PXoS"
	h := md5Hex(fmt.Sprintf("salt=%s&t=%d&r=%s&b=%s&q=", salt, t, r, data))
	return fmt.Sprintf("%d,%s,%s", t, r, h)
}

func request(url string, data string, aigis string) (*http.Response, error) {
	reqBody := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-rpc-app_version", "2.41.0")
	req.Header.Set("DS", ds(data))
	if aigis != "" {
		req.Header.Set("x-rpc-aigis", aigis)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-rpc-game_biz", "bbs_cn")
	req.Header.Set("x-rpc-sys_version", "12")
	req.Header.Set("x-rpc-device_id", randomString(16))
	req.Header.Set("x-rpc-device_fp", randomString(13))
	req.Header.Set("x-rpc-device_name", randomString(16))
	req.Header.Set("x-rpc-device_model", randomString(16))
	req.Header.Set("x-rpc-app_id", "bll8iq97cem8")
	req.Header.Set("x-rpc-client_type", "2")
	req.Header.Set("User-Agent", "okhttp/4.8.0")

	return httpClient.Do(req)
}

// --------------------------- 稳健的 JSON 解析辅助 ---------------------------

func nestedLookup(m map[string]interface{}, path []string) interface{} {
	cur := interface{}(m)
	for _, k := range path {
		if curMap, ok := cur.(map[string]interface{}); ok {
			if v, exists := curMap[k]; exists {
				cur = v
				continue
			} else {
				return nil
			}
		} else {
			return nil
		}
	}
	return cur
}

func interfaceToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch t := v.(type) {
	case string:
		return t
	case float64:
		return fmt.Sprintf("%.0f", t)
	case json.Number:
		return t.String()
	case int:
		return fmt.Sprintf("%d", t)
	case int64:
		return fmt.Sprintf("%d", t)
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return ""
		}
		return string(b)
	}
}

func extractTokenFromResponse(res map[string]interface{}) string {
	tryPaths := [][]string{
		{"data", "token", "token"},
		{"data", "token"},
		{"data", "token", "access_token"},
		{"data", "token", "stoken"},
		{"token"},
		{"data", "stoken"},
	}
	for _, p := range tryPaths {
		if v := nestedLookup(res, p); v != nil {
			s := interfaceToString(v)
			if s != "" && s != "null" {
				return s
			}
		}
	}
	return ""
}

func extractAidMid(res map[string]interface{}) (aid, mid string) {
	tryAid := [][]string{
		{"data", "user_info", "aid"},
		{"data", "user_info", "aid_str"},
		{"data", "user_info", "aidNumber"},
	}
	tryMid := [][]string{
		{"data", "user_info", "mid"},
		{"data", "user_info", "mid_str"},
	}
	for _, p := range tryAid {
		if v := nestedLookup(res, p); v != nil {
			if s := interfaceToString(v); s != "" {
				aid = s
				break
			}
		}
	}
	for _, p := range tryMid {
		if v := nestedLookup(res, p); v != nil {
			if s := interfaceToString(v); s != "" {
				mid = s
				break
			}
		}
	}
	return
}

func extractCookieToken(res map[string]interface{}) string {
	tryPaths := [][]string{
		{"data", "cookie_token"},
		{"data", "data", "cookie_token"},
		{"cookie_token"},
	}
	for _, p := range tryPaths {
		if v := nestedLookup(res, p); v != nil {
			if s := interfaceToString(v); s != "" {
				return s
			}
		}
	}
	return ""
}

// --------------------------- 主要逻辑结构体 ---------------------------

type MiHoYoLogin struct {
	Reply func(e *Event, message interface{}, at bool)
}

func NewMiHoYoLogin(reply func(e *Event, message interface{}, at bool)) *MiHoYoLogin {
	return &MiHoYoLogin{Reply: reply}
}

func (m *MiHoYoLogin) MiHoYoLoginDetect(e *Event) {
	accounts.Lock()
	accounts.m[e.UserID] = e
	accounts.Unlock()
	if m.Reply != nil {
		m.Reply(e, "请发送密码", true)
	}
}

func (m *MiHoYoLogin) CrackGeetest(gt, challenge string, e *Event) (map[string]interface{}, error) {
	link := fmt.Sprintf("https://challenge.minigg.cn/manual/index.html?gt=%s&challenge=%s", gt, challenge)
	if m.Reply != nil {
		m.Reply(e, fmt.Sprintf("请完成验证：%s", link), true)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	tick := time.NewTicker(5 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("验证超时")
		case <-tick.C:
			// TODO: 查询你的验证服务
		}
	}
}

func (m *MiHoYoLogin) MiHoYoLoginPassword(e *Event) error {
	if e == nil || strings.TrimSpace(e.Msg) == "" {
		return errors.New("没有收到密码")
	}

	Running.Lock()
	if _, ok := Running.m[e.UserID]; ok {
		Running.Unlock()
		if m.Reply != nil {
			m.Reply(e, "有正在进行的登录操作，请完成后再试……", true)
		}
		return errors.New("already running")
	}
	Running.m[e.UserID] = true
	Running.Unlock()

	defer func() {
		Running.Lock()
		delete(Running.m, e.UserID)
		Running.Unlock()
	}()

	password := strings.TrimSpace(e.Msg)

	accounts.RLock()
	prev, ok := accounts.m[e.UserID]
	accounts.RUnlock()
	if !ok {
		return errors.New("找不到对应的账号上下文")
	}
	account := strings.TrimSpace(prev.Msg)

	encAcc, err := encryptData([]byte(account))
	if err != nil {
		return err
	}
	encPwd, err := encryptData([]byte(password))
	if err != nil {
		return err
	}

	payloadObj := map[string]interface{}{
		"account":  encAcc,
		"password": encPwd,
	}
	payloadBytes, _ := json.Marshal(payloadObj)
	url := "https://passport-api.mihoyo.com/account/ma-cn-passport/app/loginByPassword"

	resp, err := request(url, string(payloadBytes), "")
	if err != nil {
		if m.Reply != nil {
			m.Reply(e, fmt.Sprintf("请求失败：%v", err), true)
		}
		return err
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Printf("getTokenByPassword raw: %s", string(bodyBytes))
	var resObj map[string]interface{}
	_ = json.Unmarshal(bodyBytes, &resObj)

	aigisHeader := resp.Header.Get("x-rpc-aigis")

	if codeFloat, ok := resObj["retcode"].(float64); ok && int(codeFloat) == -3101 {
		var aigisData map[string]interface{}
		if aigisHeader != "" {
			_ = json.Unmarshal([]byte(aigisHeader), &aigisData)
		}
		var aigisCaptchaData map[string]interface{}
		if dataStr, ok := aigisData["data"].(string); ok {
			_ = json.Unmarshal([]byte(dataStr), &aigisCaptchaData)
		}
		challenge := ""
		if v, ok := aigisCaptchaData["challenge"].(string); ok {
			challenge = v
		}
		gt := ""
		if v, ok := aigisCaptchaData["gt"].(string); ok {
			gt = v
		}

		validate, err := m.CrackGeetest(gt, challenge, e)
		if err != nil {
			if m.Reply != nil {
				m.Reply(e, "验证失败或超时", true)
			}
			return err
		}
		geetestValidate := ""
		if v, ok := validate["geetest_validate"].(string); ok {
			geetestValidate = v
		}
		if geetestValidate == "" {
			if m.Reply != nil {
				m.Reply(e, "验证结果无效", true)
			}
			return errors.New("invalid geetest validate")
		}

		sessionID := ""
		if v, ok := aigisData["session_id"].(string); ok {
			sessionID = v
		}
		aigisPayload := map[string]string{
			"geetest_challenge": challenge,
			"geetest_seccode":   geetestValidate + "|jordan",
			"geetest_validate":  geetestValidate,
		}
		apBytes, _ := json.Marshal(aigisPayload)
		aigis := fmt.Sprintf("%s;%s", sessionID, base64.StdEncoding.EncodeToString(apBytes))

		resp2, err := request(url, string(payloadBytes), aigis)
		if err != nil {
			if m.Reply != nil {
				m.Reply(e, fmt.Sprintf("二次请求失败：%v", err), true)
			}
			return err
		}
		body2, _ := ioutil.ReadAll(resp2.Body)
		resp2.Body.Close()
		log.Printf("getTokenByPassword retry raw: %s", string(body2))
		_ = json.Unmarshal(body2, &resObj)
	}

	// 使用稳健提取
	tokenToken := extractTokenFromResponse(resObj)
	aid, mid := extractAidMid(resObj)
	if tokenToken == "" {
		log.Printf("token empty after extract, full resp: %+v", resObj)
		if m.Reply != nil {
			m.Reply(e, errorTips, true)
		}
		return errors.New("empty token")
	}

	cookieURL := fmt.Sprintf("https://api-takumi.mihoyo.com/auth/api/getCookieAccountInfoBySToken?stoken=%s&mid=%s", tokenToken, mid)
	respCookie, err := httpClient.Get(cookieURL)
	if err != nil {
		return err
	}
	cookieBody, _ := ioutil.ReadAll(respCookie.Body)
	respCookie.Body.Close()
	log.Printf("getCookieAccountInfoBySToken raw: %s", string(cookieBody))
	var cookieObj map[string]interface{}
	_ = json.Unmarshal(cookieBody, &cookieObj)

	cookieToken := extractCookieToken(cookieObj)

	cookies := []string{
		fmt.Sprintf("ltoken=%s;ltuid=%s;cookie_token=%s;login_ticket=%s", tokenToken, aid, cookieToken, interfaceToString(nestedLookup(resObj, []string{"data", "login_ticket"}))),
		fmt.Sprintf("stoken=%s;stuid=%s;mid=%s", tokenToken, aid, mid),
	}

	for _, c := range cookies {
		if m.Reply != nil {
			m.Reply(e, c, e.IsPrivate)
		}
	}

	return nil
}

func (m *MiHoYoLogin) MiHoYoLoginQRCode(e *Event, appID int) error {
	Running.Lock()
	if _, ok := Running.m[e.UserID]; ok {
		Running.Unlock()
		if m.Reply != nil {
			m.Reply(e, "有正在进行的登录操作，请完成后再试……", true)
		}
		return errors.New("already running")
	}
	Running.m[e.UserID] = true
	Running.Unlock()
	defer func() {
		Running.Lock()
		delete(Running.m, e.UserID)
		Running.Unlock()
	}()

	device := randomString(64)
	fetchURL := "https://hk4e-sdk.mihoyo.com/hk4e_cn/combo/panda/qrcode/fetch"
	payload := map[string]interface{}{"app_id": appID, "device": device}
	payloadBytes, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", fetchURL, bytes.NewReader(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		if m.Reply != nil {
			m.Reply(e, fmt.Sprintf("请求二维码失败：%v", err), true)
		}
		return err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Printf("qrcode fetch raw: %s", string(body))
	var fetchRes map[string]interface{}
	_ = json.Unmarshal(body, &fetchRes)

	dataObj, _ := fetchRes["data"].(map[string]interface{})
	urlStr := ""
	if v, ok := dataObj["url"].(string); ok {
		urlStr = v
	}
	if urlStr == "" {
		if m.Reply != nil {
			m.Reply(e, errorTips, true)
		}
		return errors.New("no url in response")
	}

	parts := strings.Split(urlStr, "ticket=")
	if len(parts) < 2 {
		if m.Reply != nil {
			m.Reply(e, "二维码 URL 格式不对", true)
		}
		return errors.New("invalid url")
	}
	ticket := parts[1]

	tmpPNG := fmt.Sprintf("qrc_%d.png", time.Now().UnixNano())
	err = qrcode.WriteFile(urlStr, qrcode.Medium, 256, tmpPNG)
	if err != nil {
		if m.Reply != nil {
			m.Reply(e, fmt.Sprintf("生成二维码失败：%v", err), true)
		}
		return err
	}
	if m.Reply != nil {
		m.Reply(e, fmt.Sprintf("请使用米游社扫码登录（二维码图片已保存：%s）", tmpPNG), true)
	}

	queryURL := "https://hk4e-sdk.mihoyo.com/hk4e_cn/combo/panda/qrcode/query"
	var data map[string]interface{}
	scanned := false
	for n := 0; n < 60; n++ {
		time.Sleep(5 * time.Second)

		payload2 := map[string]interface{}{"app_id": appID, "device": device, "ticket": ticket}
		b2, _ := json.Marshal(payload2)
		req2, _ := http.NewRequest("POST", queryURL, bytes.NewReader(b2))
		req2.Header.Set("Content-Type", "application/json")
		resp2, err := httpClient.Do(req2)
		if err != nil {
			log.Printf("query err: %v", err)
			continue
		}
		body2, _ := ioutil.ReadAll(resp2.Body)
		resp2.Body.Close()
		log.Printf("qrcode query raw: %s", string(body2))
		var qrRes map[string]interface{}
		_ = json.Unmarshal(body2, &qrRes)

		if rc, ok := qrRes["retcode"].(float64); ok && int(rc) != 0 {
			Running.Lock()
			delete(Running.m, e.UserID)
			Running.Unlock()
			if m.Reply != nil {
				m.Reply(e, "二维码已过期，请重新登录", true)
			}
			return errors.New("qrcode expired")
		}

		if dataMap, ok := qrRes["data"].(map[string]interface{}); ok {
			if stat, ok := dataMap["stat"].(string); ok {
				if stat == "Scanned" && !scanned {
					scanned = true
					if m.Reply != nil {
						m.Reply(e, "二维码已扫描，请确认登录", true)
					}
				}
				if stat == "Confirmed" {
					if payloadRaw, ok := dataMap["payload"].(map[string]interface{}); ok {
						if raw, ok := payloadRaw["raw"].(string); ok {
							_ = json.Unmarshal([]byte(raw), &data)
							break
						}
					}
				}
			}
		}
	}

	if data == nil {
		if m.Reply != nil {
			m.Reply(e, errorTips, true)
		}
		return errors.New("no data after polling")
	}

	uid := ""
	token := ""
	if v, ok := data["uid"].(string); ok {
		uid = v
	} else if v, ok := data["uid"].(float64); ok {
		uid = fmt.Sprintf("%.0f", v)
	}
	if v, ok := data["token"].(string); ok {
		token = v
	}
	if uid == "" || token == "" {
		if m.Reply != nil {
			m.Reply(e, errorTips, true)
		}
		return errors.New("invalid payload data")
	}

	urlToken := "https://passport-api.mihoyo.com/account/ma-cn-session/app/getTokenByGameToken"
	// account_id must be a number, not a string — parse uid to int
	accountIDNum := 0
	if uid != "" {
		if v, err := strconv.ParseInt(uid, 10, 64); err == nil {
			accountIDNum = int(v)
		} else {
			log.Printf("warn: failed to parse uid to int: %v", err)
		}
	}

	reqPayload := map[string]interface{}{
		"account_id": accountIDNum,
		"game_token": token,
	}
	bToken, _ := json.Marshal(reqPayload)
	respToken, err := request(urlToken, string(bToken), "")
	if err != nil {
		return err
	}
	bodyToken, _ := ioutil.ReadAll(respToken.Body)
	respToken.Body.Close()
	log.Printf("getTokenByGameToken raw: %s", string(bodyToken))
	var tokenRes map[string]interface{}
	_ = json.Unmarshal(bodyToken, &tokenRes)

	cookieURL := fmt.Sprintf("https://api-takumi.mihoyo.com/auth/api/getCookieAccountInfoByGameToken?account_id=%s&game_token=%s", uid, token)
	respCookie2, err := httpClient.Get(cookieURL)
	if err != nil {
		return err
	}
	bodyCookie2, _ := ioutil.ReadAll(respCookie2.Body)
	respCookie2.Body.Close()
	log.Printf("getCookieAccountInfoByGameToken raw: %s", string(bodyCookie2))
	var cookieRes map[string]interface{}
	_ = json.Unmarshal(bodyCookie2, &cookieRes)

	dataMap2, _ := tokenRes["data"].(map[string]interface{})
	tokenToken := extractTokenFromResponse(tokenRes)
	userInfo2, _ := dataMap2["user_info"].(map[string]interface{})
	aid := interfaceToString(nestedLookup(userInfo2, []string{"aid"}))
	mid := interfaceToString(nestedLookup(userInfo2, []string{"mid"}))
	cookieToken2 := extractCookieToken(cookieRes)

	cookies := []string{
		fmt.Sprintf("ltoken=%s;ltuid=%s;cookie_token=%s", tokenToken, aid, cookieToken2),
		fmt.Sprintf("stoken=%s;stuid=%s;mid=%s", tokenToken, aid, mid),
	}
	for _, c := range cookies {
		if m.Reply != nil {
			m.Reply(e, fmt.Sprintf("登录完成：%s", c), e.IsPrivate)
		}
	}

	return nil
}

func main() {

	//var wg sync.WaitGroup
	
	mrand.Seed(time.Now().UnixNano())
	
	reply := func(e *Event, message interface{}, at bool) {
		log.Printf("[Reply to %d] %v (at=%v)", e.UserID, message, at)
	}
	
	login := NewMiHoYoLogin(reply)
	
	ev := &Event{UserID: 123456, Msg: "example_account", IsPrivate: true}
	login.MiHoYoLoginDetect(ev)
	ev2 := &Event{UserID: 123456, Msg: "mypassword123", IsPrivate: true}
	if err := login.MiHoYoLoginPassword(ev2); err != nil {
		log.Printf("登录出错: %v", err)
	}
	ev3 := &Event{UserID: 654321, Msg: "#扫码登录", IsPrivate: true}
	go func() {
		//defer wg.Done()
		if err := login.MiHoYoLoginQRCode(ev3, 2); err != nil {
			log.Printf("扫码登录出错: %v", err)
		}
	}()
	
	//wg.Wait() // 等待所有 goroutine 完成
	select {
	case <-time.After(320 * time.Second):
	
	}


}

```