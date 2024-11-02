package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}

func GenerateClientTokenUUID(userID int) string {
	newUUID := uuid.New()
	// convert uuid to string, remove -
	uuidString := strings.ReplaceAll((newUUID).String(), "-", "")
	return strconv.Itoa(userID) + "clitoken" + uuidString
}