package dao

type ConfigFile struct {
	FilePath string `json:"file_path"`
	FileName string `json:"file_name"`
}

func NewConfigFile(path string, filename string) *ConfigFile {
	return &ConfigFile{
		FilePath: path,
		FileName: filename,
	}
}
