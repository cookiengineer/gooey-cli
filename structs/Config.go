package structs

type Config struct {
	Name            string `json:"name"`
	Folder          string `json:"folder"`
	UseTheme        bool   `json:"use_theme"`

	Controllers     []ControllerConfig `json:"controllers"`
	Components      []ComponentConfig  `json:"components"`
	Views           []ViewConfig       `json:"views"`

}
