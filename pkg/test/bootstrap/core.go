package bootstrap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"testing"

	"log"

	"gitlab.com/mikrowezel/backend/granica/internal/config"
	"gitlab.com/mikrowezel/backend/granica/pkg/authentication"
)

// 	Config config
//  holds the configuration values from th.json file
var (
	cfg         *config.Config
	env         = "test"
	baseTest    = ""
	fixturesDir = "resources/fixtures"
	th          TestHandler
	apiPath     = "api"
	apiVersion  = "v1"
	logger      log.Logger
)

// BootParameters - Default boot parameters for tests
func BootParameters() map[string]string {
	params := make(map[string]string)
	// Envs: "dev", "test", "prod"
	// Migrations: "m", "r", "n" - migrate, rollback, none
	params["env"] = "test"
	params["app_home"] = os.Getenv("GRANICA_HOME")
	params["migration"] = "m"
	return params
}

// Reads th.yaml and decode into Config
func init() {
	baseTest = getBaseDir()
	configFile := fmt.Sprintf("configs/config.yaml", env)
	fullConfigPath := path.Join(baseTest, configFile)
	file, err := os.Open(fullConfigPath)
	defer file.Close()
	if err != nil {
		log.Fatalf("[ERROR]: %s\n", err)
	}

	th.Volume.BasePath = fullConfigPath

	decoder := json.NewDecoder(file)
	cfg = &config.Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("[ERROR]: %s\n", err)
	}
	log.Printf("Test config: %v", cfg)

	th := TestHandler{}
	th.MainConfig = cfg

	th.Volume.BasePath = baseTest
	th.Volume.FixturesPath = path.Join(baseTest, fixturesDir)

	// th.ServerInstance = httptest.NewServer(handler.AppHandler(TestHandler))
	svc := authentication.NewServiceForTests(cfg)
	th.ServerInstance = httptest.NewServer(authentication.SignUpHandler(svc))

	th.API.ServerURL = fmt.Sprintf("%s/%s/%s", th.ServerInstance.URL, th.API.Path, th.API.Version)
	th.API.Path = apiPath
	th.API.Version = apiVersion
}

func (th *TestHandler) Start(m *testing.M) {
	var err error
	// Open connection with the test database.
	// Do NOT import fixtures in a production database!
	// Existing data would be deleted
	// connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s", cfg.User, cfg.DB, cfg.SSL)
	mc, err := getConn(cfg)
	if err != nil {
		log.Fatal(err)
	}
	th.RepoClient = mc

	os.Exit(m.Run())
}

// AuthorizeRequest authorizes request.
func (th *TestHandler) AuthorizeRequest(req *http.Request, user, username, role string) {
	token, _ := GenerateJWT(user, username, role)
	var bearer = "Bearer " + token
	req.Header.Add("authorization", bearer)
}

func getBaseDir() string {
	exPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	parentPath := filepath.Dir(exPath)
	return parentPath
}
