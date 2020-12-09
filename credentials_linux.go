package jazz

import "go.uber.org/zap"

// Decrypt master password
func decryptMasterKey(data string) string {
	zap.L().Warn("decryptMasterKey not implemented for linux")
	return ""
}
