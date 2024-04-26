package notifier

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"testing"
)

func TestStringifyData(t *testing.T) {
	s := struct {
		Name  string
		Age   int
		Email string
	}{
		Name:  "John Doe",
		Age:   43,
		Email: "john.doe@example.com",
	}

	marshal, err := json.Marshal(s)
	if err != nil {
		return
	}

	fmt.Println(string(marshal))

	sha1Hash := sha1.New()
	sha1Hash.Write(marshal)
	hashBytes := sha1Hash.Sum(nil)
	sha1String := hex.EncodeToString(hashBytes)

	fmt.Println(sha1String)

	viper.SetConfigFile("../config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	bot, err := NewTelegramBot()
	if err != nil {
		return
	}

	bot.SendMessage(1247260748, string(marshal))
}
