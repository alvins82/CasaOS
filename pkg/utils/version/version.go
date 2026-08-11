/*
 * @Author: LinkLeong link@icewhale.com
 * @Date: 2022-05-13 18:15:46
 * @LastEditors: LinkLeong
 * @LastEditTime: 2022-07-21 15:27:53
 * @FilePath: /CasaOS/pkg/utils/version/version.go
 * @Description:
 * @Website: https://www.casaos.io
 * Copyright (c) 2022 by icewhale, All Rights Reserved.
 */
package version

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/IceWhaleTech/CasaOS/common"
	"github.com/IceWhaleTech/CasaOS/model"
)

var numericVersionPart = regexp.MustCompile(`[0-9]+`)

func versionParts(value string) ([]uint64, bool) {
	matches := numericVersionPart.FindAllString(value, -1)
	if len(matches) == 0 {
		return nil, false
	}

	parts := make([]uint64, 0, len(matches))
	for _, match := range matches {
		part, err := strconv.ParseUint(match, 10, 64)
		if err != nil {
			return nil, false
		}
		parts = append(parts, part)
	}
	return parts, true
}

func IsVersionNewer(latest string, current string) bool {
	latestParts, latestOK := versionParts(latest)
	currentParts, currentOK := versionParts(current)
	if !latestOK || !currentOK {
		return false
	}

	length := len(latestParts)
	if len(currentParts) > length {
		length = len(currentParts)
	}
	for i := 0; i < length; i++ {
		var latestPart uint64
		var currentPart uint64
		if i < len(latestParts) {
			latestPart = latestParts[i]
		}
		if i < len(currentParts) {
			currentPart = currentParts[i]
		}
		if latestPart != currentPart {
			return latestPart > currentPart
		}
	}
	return false
}

func CurrentVersion() string {
	return currentVersionFromFile(common.FORK_RELEASE_FILE)
}

func currentVersionFromFile(path string) string {
	data, err := os.ReadFile(path)
	if err == nil {
		if installedVersion := strings.TrimSpace(string(data)); installedVersion != "" {
			return installedVersion
		}
	}
	return common.FORK_RELEASE_VERSION
}

func IsNeedUpdate(version model.Version) (bool, model.Version) {
	return IsVersionNewer(version.Version, CurrentVersion()), version
}
