package pkg

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	s "strings"
)

type WorkspaceConfig struct {
	ConfigDirPath   string
	SrcDirPath      string
	BuildDirPath    string
	MessageDirPath  string
	ServicesDirPath string
	LaunchesDirPath string
}

func (cfg *WorkspaceConfig) GetAllMessages() []string {
	var messages []string

	_ = filepath.Walk(cfg.MessageDirPath, func(path string, info fs.FileInfo, err error) error {

		if s.HasSuffix(info.Name(), ".msg") {
			messages = append(messages, path)
		}

		return err
	})
	return messages
}

func workspaceRootPath() string {
	dir, _ := os.UserHomeDir()
	return path.Join(dir, "kos_workspace")
}

func GetWorkspaceConfig() WorkspaceConfig {
	return WorkspaceConfig{
		ConfigDirPath:   path.Join(workspaceRootPath(), ".kos"),
		SrcDirPath:      path.Join(workspaceRootPath(), "src"),
		BuildDirPath:    path.Join(workspaceRootPath(), "build"),
		MessageDirPath:  path.Join(workspaceRootPath(), "messages"),
		ServicesDirPath: path.Join(workspaceRootPath(), "services"),
		LaunchesDirPath: path.Join(workspaceRootPath(), "launches"),
	}
}
