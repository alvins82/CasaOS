package service

import (
	json2 "encoding/json"
	"net/url"
	"strings"
	"time"

	"github.com/IceWhaleTech/CasaOS/common"
	"github.com/IceWhaleTech/CasaOS/model"
	"github.com/IceWhaleTech/CasaOS/pkg/config"
	"github.com/IceWhaleTech/CasaOS/pkg/utils/httper"
	"github.com/tidwall/gjson"
)

type CasaService interface {
	GetCasaosVersion() model.Version
}

type casaService struct{}

func validHTTPSURL(value string) bool {
	parsed, err := url.ParseRequestURI(value)
	return err == nil && parsed.Scheme == "https" && parsed.Host != ""
}

func resolveUpdateVersionURL() string {
	value := strings.TrimSpace(config.ServerInfo.UpdateVersionUrl)
	if validHTTPSURL(value) {
		return value
	}
	return common.FORK_VERSION_URL
}

func parseReleaseVersion(payload string) model.Version {
	var release model.Version
	if err := json2.Unmarshal([]byte(payload), &release); err == nil && release.Version != "" {
		return release
	}

	data := gjson.Get(payload, "data")
	if data.Exists() {
		_ = json2.Unmarshal([]byte(data.String()), &release)
		if release.Version != "" {
			return release
		}
	}

	return model.Version{
		Version:   gjson.Get(payload, "tag_name").String(),
		ChangeLog: gjson.Get(payload, "body").String(),
	}
}

/**
 * @description: get remote version
 * @return {model.Version}
 */
func (o *casaService) GetCasaosVersion() model.Version {
	versionURL := resolveUpdateVersionURL()
	keyName := "casa_version:" + versionURL
	var dataStr string
	var version model.Version
	if result, ok := Cache.Get(keyName); ok {
		dataStr, ok = result.(string)
		if ok {
			return parseReleaseVersion(dataStr)
		}
	}

	v := httper.Get(versionURL, map[string]string{
		"Accept":     "application/json",
		"User-Agent": "CasaOS-Fork-Updater",
	})
	version = parseReleaseVersion(v)

	if len(version.Version) > 0 {
		Cache.Set(keyName, v, time.Minute*20)
	}

	return version
}

func NewCasaService() CasaService {
	return &casaService{}
}
