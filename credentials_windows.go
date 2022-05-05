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
	"encoding/base64"
	"syscall"
	"unsafe"
)

// Decrypt master password
func decryptMasterKey(data string) string {
	if data == "" {
		return ""
	}

	// get only the encrypted part (others should be empty)
	_, _, _, data = splitEncryptedData(data)

	// decode double base64
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	decodedData, err = base64.StdEncoding.DecodeString(string(decodedData))
	if err != nil {
		return ""
	}

	// decrypt/unprotect master password with win32 CryptUnprotectData
	pw, err := decrypt(decodedData)
	if err != nil {
		return ""
	}
	return string(pw)
}

// # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
// Win32 crypto
// # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #

var (
	dllcrypt32  = syscall.NewLazyDLL("Crypt32.dll")
	dllkernel32 = syscall.NewLazyDLL("Kernel32.dll")

	procDecryptData = dllcrypt32.NewProc("CryptUnprotectData")
	procLocalFree   = dllkernel32.NewProc("LocalFree")
)

type dataBlock struct {
	cbData uint32
	pbData *byte
}

func NewBlob(d []byte) *dataBlock {
	if len(d) == 0 {
		return &dataBlock{}
	}
	return &dataBlock{
		pbData: &d[0],
		cbData: uint32(len(d)),
	}
}

func (b *dataBlock) ToByteArray() []byte {
	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])
	return d
}

func decrypt(data []byte) ([]byte, error) {
	var outblob dataBlock
	r, _, err := procDecryptData.Call(uintptr(unsafe.Pointer(NewBlob(data))), 0, 0, 0, 0, 0, uintptr(unsafe.Pointer(&outblob)))
	if r == 0 {
		return nil, err
	}
	defer procLocalFree.Call(uintptr(unsafe.Pointer(outblob.pbData)))
	return outblob.ToByteArray(), nil
}
