package config

import "os"

var (
	API_BASE_URL      string
	CDN_BASE_URL      string
	APPLICATION_ID    string
	TESTING_GUILD_ID  string
	DISCORD_BOT_TOKEN string
)

func Load() {
	API_BASE_URL = os.Getenv("BASE_DISCORD_API_URL") + "/v" + os.Getenv("DISCORD_API_VERSION")
	APPLICATION_ID = os.Getenv("DISCORD_APPLICATION_ID")
	TESTING_GUILD_ID = os.Getenv("TESTING_GUILD_ID")
	DISCORD_BOT_TOKEN = os.Getenv("DISCORD_BOT_TOKEN")
	CDN_BASE_URL = os.Getenv("BASE_DISCORD_CDN_URL")
}
