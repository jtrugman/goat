package model

func init() {

}

type Kid struct {
	Job struct {
		Command struct {
			Port      string `yaml:"port"`
			Operation string `yaml:"operation"`
			Bitrate   struct {
				BitrateValue float64 `yaml:"bitrateValue"`
				BitrateUnit  string  `yaml:"bitrateUnit"`
			} `yaml:"bitrate"`
			Latency float64 `yaml:"latency"`
			PktLoss float64 `yaml:"pktLoss"`
			Jitter  float64 `yaml:"jitter"`
		} `yaml:"command"`
		Timer struct {
			TimeValue int    `yaml:"timeValue"`
			TimeUnit  string `yaml:"timeUnit"`
		} `yaml:"timer"`
		Link string `yaml:"link"`
	} `yaml:"job"`
}
