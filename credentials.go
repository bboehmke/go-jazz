// Copyright 2022 Benjamin Böhmke <benjamin@boehmke.net>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jazz

import (
	"crypto"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/magiconair/properties"
	"go.uber.org/zap"
)

// Strip entries from split entry
func stripEntry(data string) string {
	return strings.TrimSpace(strings.Trim(data, "\t;,"))
}

var reSplitEncryptedData = regexp.MustCompile("(.*\t)?(.*;)?(.*,)?(.*)")

// Split encrypted data from eclipse.equinox.security
func splitEncryptedData(data string) (string, string, string, string) {
	res := reSplitEncryptedData.FindStringSubmatch(data)
	return stripEntry(res[1]), stripEntry(res[2]), stripEntry(res[3]), stripEntry(res[4])
}

// ReadEclipsePassword reads password from eclipse secure storage
func ReadEclipsePassword(user, url string) string {
	eclipseSecureStorage := path.Join(os.Getenv("USERPROFILE"),
		".eclipse/org.eclipse.equinox.security/secure_storage")

	// load secure storage
	props, err := properties.LoadFile(eclipseSecureStorage, properties.UTF8)
	if err != nil {
		// unable to load file -> it maybe does not exist
		return ""
	}

	// get master passwords
	masterKey32 := decryptMasterKey(props.GetString("/org.eclipse.equinox.secure.storage/windows/encryptedPassword", ""))
	masterKey64 := decryptMasterKey(props.GetString("/org.eclipse.equinox.secure.storage/windows64/encryptedPassword", ""))

	for _, key := range props.Keys() {
		// no RTC auth -> skip
		if !strings.Contains(key, "com.ibm.team.auth.info") {
			continue
		}
		value := props.MustGet(key)

		// key has format like:
		//  /com.ibm.team.auth.info//bboehmke@https://jazz.server.net/ccm/
		//  /com.ibm.team.auth.info//bboehmke@https://jazzccm.server:9443/ccm/

		// remove "com.ibm.team.auth.info" and "/"
		key = strings.TrimLeft(key[23:], "/")

		// remove specific parts from jazz
		key = strings.ReplaceAll(key, "jazzccm", "jazz")
		key = strings.ReplaceAll(key, ":9443", "")

		// check if key match with user and URL
		if !strings.HasPrefix(key, user+"@"+url) {
			continue
		}

		// split encrypted value
		module, iv, salt, encrypted := splitEncryptedData(value)

		// no IV support
		if iv != "" {
			zap.S().Errorf("eclipse secret: IV handling not implemented: %s", key)
			continue
		}

		// check which master key should be used
		var masterKey string
		if module == "org.eclipse.equinox.security.windowspasswordprovider64bit" {
			masterKey = masterKey64
		} else if module == "org.eclipse.equinox.security.windowspasswordprovider" {
			masterKey = masterKey32
		} else {
			zap.S().Errorf("eclipse secret: unsupported auth module: %s", module)
			continue
		}

		// no master key -> skip
		if masterKey == "" {
			zap.S().Debugf("eclipse secret: master key missing for %s", key)
			continue
		}

		zap.S().Debugf("eclipse secret: using %s", key)

		// prepare decryption key
		buf, _ := base64.StdEncoding.DecodeString(salt)
		keyIV := append([]byte(masterKey), buf...)
		for i := 0; i < 10; i++ {
			h := crypto.MD5.New()
			h.Write(keyIV)
			keyIV = h.Sum(nil)
		}
		keyIV = keyIV[:16]

		// created DES cipher
		block, err := des.NewCipher(keyIV[:8])
		if err != nil {
			continue
		}
		blockMode := cipher.NewCBCDecrypter(block, keyIV[8:])

		// decrypt password
		data, _ := base64.StdEncoding.DecodeString(encrypted)
		blockMode.CryptBlocks(data, data)

		// remove padded bytes

		c := data[len(data)-1]
		n := int(c)
		if n == 0 || n > len(data) {
			zap.L().Error("eclipse secret: invalid padding")
			continue
		}
		return string(data[:len(data)-n])
	}

	return ""
}
