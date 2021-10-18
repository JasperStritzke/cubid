package security

import (
	"encoding/base64"
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	"github.com/jasperstritzke/cubid/pkg/util"
	"github.com/jasperstritzke/cubid/pkg/util/fileutil"
	"github.com/jasperstritzke/cubid/pkg/util/random"
	"sync"
)

const (
	keyLength = 2048
	pathName  = "executor.cubidkey"
)

var syncOnce sync.Once
var (
	secretKey string
	hashedKey string
)

func InitControllerKey() {
	randomKey, _ := random.GenerateRandomString(keyLength)

	if fileutil.WriteIfNotExists(
		pathName,
		base64.StdEncoding.EncodeToString([]byte(randomKey)),
	) {
		logger.Info("Successfully created secret-key. Please copy the executor.cubidkey file to all executors.")
	}
}

func LoadControllerKey() {
	key := fileutil.ReadString(pathName)

	if len(key) == 0 {
		panic("Can't load secret-key. Please make sure to copy the executor.cubidkey file into the working directory of this service.")
	}

	secretKey = key
	logger.Info("Successfully loaded secret-key.")
}

func GetHashedKey() string {
	syncOnce.Do(func() {
		hashedKey = util.HashString(secretKey)
	})

	return hashedKey
}

func IsHashedKeyValid(hash string) bool {
	return GetHashedKey() == hash
}
