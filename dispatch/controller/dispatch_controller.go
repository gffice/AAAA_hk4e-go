package controller

import (
	"bytes"
	"encoding/base64"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"hk4e/common/region"
	httpapi "hk4e/dispatch/api"
	"hk4e/node/api"
	"hk4e/pkg/endec"
	"hk4e/pkg/logger"

	"github.com/gin-gonic/gin"
)

func (c *Controller) query_security_file(context *gin.Context) {
	return
	file, err := os.ReadFile("static/security_file")
	if err != nil {
		logger.Error("open security_file error")
		return
	}
	context.Header("Content-type", "text/html; charset=UTF-8")
	_, _ = context.Writer.WriteString(string(file))
}

func (c *Controller) query_region_list(context *gin.Context) {
	context.Header("Content-type", "text/html; charset=UTF-8")
	regionListBase64 := region.GetRegionListBase64(c.ec2b)
	_, _ = context.Writer.WriteString(regionListBase64)
}

func (c *Controller) query_cur_region(context *gin.Context) {
	versionName := context.Query("version")
	response := "CAESGE5vdCBGb3VuZCB2ZXJzaW9uIGNvbmZpZw=="
	if len(context.Request.URL.RawQuery) > 0 {
		addr, err := c.discovery.GetGateServerAddr(context.Request.Context(), &api.NullMsg{})
		if err != nil {
			logger.Error("get gate server addr error: %v", err)
			return
		}
		regionCurrBase64 := region.GetRegionCurrBase64(addr.IpAddr, int32(addr.Port), c.ec2b)
		response = regionCurrBase64
	}
	reg, err := regexp.Compile("[0-9]+")
	if err != nil {
		logger.Error("compile regexp error: %v", err)
		return
	}
	versionSlice := reg.FindAllString(versionName, -1)
	version := 0
	for index := 0; index < len(versionSlice); index++ {
		v, err := strconv.Atoi(versionSlice[index])
		if err != nil {
			logger.Error("parse client version error: %v", err)
			return
		}
		for i := 0; i < len(versionSlice)-1-index; i++ {
			v *= 10
		}
		version += v
	}
	if version >= 1000 {
		// 测试版本
		version /= 10
	}
	if version >= 275 {
		logger.Debug("do hk4e 2.8 rsa logic")
		if context.Query("dispatchSeed") == "" {
			rsp := &httpapi.QueryCurRegionRspJson{
				Content: response,
				Sign:    "TW9yZSBsb3ZlIGZvciBVQSBQYXRjaCBwbGF5ZXJz",
			}
			context.JSON(http.StatusOK, rsp)
			return
		}
		keyId := context.Query("key_id")
		encPubPrivKey, exist := c.encRsaKeyMap[keyId]
		if !exist {
			logger.Error("can not found key id: %v", keyId)
			return
		}
		regionInfo, err := base64.StdEncoding.DecodeString(response)
		if err != nil {
			logger.Error("decode region info error: %v", err)
			return
		}
		chunkSize := 256 - 11
		regionInfoLength := len(regionInfo)
		numChunks := int(math.Ceil(float64(regionInfoLength) / float64(chunkSize)))
		encryptedRegionInfo := make([]byte, 0)
		for i := 0; i < numChunks; i++ {
			from := i * chunkSize
			to := int(math.Min(float64((i+1)*chunkSize), float64(regionInfoLength)))
			chunk := regionInfo[from:to]
			pubKey, err := endec.RsaParsePubKeyByPrivKey(encPubPrivKey)
			if err != nil {
				logger.Error("parse rsa pub key error: %v", err)
				return
			}
			privKey, err := endec.RsaParsePrivKey(encPubPrivKey)
			if err != nil {
				logger.Error("parse rsa priv key error: %v", err)
				return
			}
			encrypt, err := endec.RsaEncrypt(chunk, pubKey)
			if err != nil {
				logger.Error("rsa enc error: %v", err)
				return
			}
			decrypt, err := endec.RsaDecrypt(encrypt, privKey)
			if err != nil {
				logger.Error("rsa dec error: %v", err)
				return
			}
			if bytes.Compare(decrypt, chunk) != 0 {
				logger.Error("rsa dec test fail")
				return
			}
			encryptedRegionInfo = append(encryptedRegionInfo, encrypt...)
		}
		signPrivkey, err := endec.RsaParsePrivKey(c.signRsaKey)
		if err != nil {
			logger.Error("parse rsa priv key error: %v", err)
			return
		}
		signData, err := endec.RsaSign(regionInfo, signPrivkey)
		if err != nil {
			logger.Error("rsa sign error: %v", err)
			return
		}
		ok, err := endec.RsaVerify(regionInfo, signData, &signPrivkey.PublicKey)
		if err != nil {
			logger.Error("rsa verify error: %v", err)
			return
		}
		if !ok {
			logger.Error("rsa verify test fail")
			return
		}
		rsp := &httpapi.QueryCurRegionRspJson{
			Content: base64.StdEncoding.EncodeToString(encryptedRegionInfo),
			Sign:    base64.StdEncoding.EncodeToString(signData),
		}
		context.JSON(http.StatusOK, rsp)
		return
	} else {
		context.Header("Content-type", "text/html; charset=UTF-8")
		_, _ = context.Writer.WriteString(response)
	}
}
