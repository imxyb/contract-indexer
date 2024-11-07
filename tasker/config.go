package tasker

var Conf Config

// Config tasker config
type Config struct {
	// Tasker tasker config
	Subscriber struct {
		// Type subscriber type
		Protocol string `yaml:"protocol"`
		// Conf subscriber config
		Conf map[string]any `yaml:"conf"`
	} `yaml:"subscriber"`

	// Tasker tasker config
	Storage struct {
		// Type storage type
		Type string `yaml:"type"`
		// Conf storage config
		Conf map[string]any `yaml:"conf"`
	}
}
