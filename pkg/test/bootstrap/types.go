/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package bootstrap

import (
	"fmt"
	"io"
	"net/http/httptest"

	"gitlab.com/mikrowezel/backend/granica/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	TestHandler struct {
		MainConfig     *config.Config
		Server         string
		Volume         VolumeConfig
		API            APIConfig
		RepoClient     *mongo.Client
		ServerInstance *httptest.Server
		Reader         io.Reader
	}

	VolumeConfig struct {
		BasePath      string
		FixturesPath  string
		PublicPath    string
		ResourcesPath string
	}

	APIConfig struct {
		ServerURL string
		Path      string
		Version   string
	}
)

func (th *TestHandler) GetServer() string {
	return th.Server
}

func (th *TestHandler) GetDBConnParamenters() map[string]interface{} {
	repoConf := make(map[string]interface{})
	repoConf["Host"] = th.MainConfig.Repo.MongoDB.Host
	repoConf["Port"] = fmt.Sprintf("%d", th.MainConfig.Repo.MongoDB.Port)
	repoConf["User"] = th.MainConfig.Repo.MongoDB.User
	repoConf["Password"] = th.MainConfig.Repo.MongoDB.Password
	// repoConf["SSL"] = th.MainConfig.MongoDB.DBSSL
	return repoConf
}

func (th *TestHandler) GetBaseDir() string {
	return th.Volume.BasePath
}

// GetResourcesDir returns the resouurces dir path.
func (th *TestHandler) GetResourcesDir() string {
	return th.Volume.ResourcesPath
}

// GetPublicDir returns the public dir path.
func (th *TestHandler) GetPublicDir() string {
	return th.Volume.PublicPath
}
