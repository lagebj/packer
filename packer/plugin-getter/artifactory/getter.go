package artifactory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/jfrog/jfrog-client-go/artifactory"
	"github.com/jfrog/jfrog-client-go/auth"
	plugingetter "github.com/hashicorp/packer/packer/plugin-getter"
)

const (
	artifactoryTokenAccessor = "PACKER_ARTIFACTORY_API_TOKEN"
	defaultUserAgent         = "packer-artifactory-plugin-getter"
)

type Getter struct {
	Client    *artifactory.Artifactory
	UserAgent string
}

var _ plugingetter.Getter = &Getter{}

func (g *Getter) Get(plugin *plugingetter.Plugin, target string) (io.ReadCloser, error) {
	switch target {
		case "releases":
			// search for all releases of the plugin
			searchResults, err := g.Client.Search.ArtifactsByName(plugin.Name)
			if err != nil {
				return nil, err
			}
			jsonResults, err := json.Marshal(searchResults)
			if err != nil {
				return nil, err
			}
			reader := bytes.NewReader(jsonResults)
			return ioutil.NopCloser(reader), nil
		case "sha256":
			// download the sha256 file
			sha256File, _, err := g.Client.Download.DownloadFile(plugin.Path + "/" + plugin.Name + "_" + plugin.Version + "_sha256")
			if err != nil {
				return nil, err
			}
			return sha256File, nil
		case "zip":
			// download the zip file
			zipFile, _, err := g.Client.Download.DownloadFile(plugin.Path + "/" + plugin.Name + "_" + plugin.Version + ".zip")
			if err != nil {
				return nil, err
			}
			return zipFile, nil
		default:
			return nil, fmt.Errorf("Invalid target: %s", target)
	}
}

func New(serverURL, token string) (*Getter, error) {
	if serverURL == "" {
		return nil, fmt.Errorf("Missing JFrog Artifactory server URL")
	}
	if token == "" {
		return nil, fmt.Errorf("Missing JFrog Artifactory API token")
	}
	transport := auth.NewTokenTransport(token, defaultUserAgent)
	client, err := artifactory.NewClient(serverURL, transport)
	if err != nil {
		return nil, err
	}
	return &Getter{
		Client:    client,
		UserAgent: defaultUserAgent,
	}, nil
}
