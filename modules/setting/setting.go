package setting

type Settings struct {
	Server struct {
		Cache string `json:"cache"`
		Host  string `json:"host"`
		Port  string `json:"port"`
	} `json:"server"`
}

var (
	AppVer        = "0.0.2"
	CacheLocation = "/var/tilr/cache"
	CacheFiletype = "png"

	TileFileType = "png"
	Host         = ""
	Port         = "5555"
)

func init() {

}

func Load(configPath string) {

}
