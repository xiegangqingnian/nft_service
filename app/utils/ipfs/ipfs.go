/**
@author: 铁柱
@date:2021/6/30
@note:
**/
package ipfs

import (
	shell "github.com/ipfs/go-ipfs-api"
	"net/http"
	"nft_service/app/global/variable"
)

func GetIpfsClient() *shell.Shell {
	//return shell.NewShellWithClient(variable.IpfsHost,NewClient(variable.InfuraIpfsProjectID,variable.InfuraProjectSecret))
	return shell.NewShell(variable.IpfsHost)
}

// NewClient creates an http.Client that automatically perform basic auth on each request.
func NewClient(projectId, projectSecret string) *http.Client {
	return &http.Client{
		Transport: authTransport{
			RoundTripper:  http.DefaultTransport,
			ProjectId:     projectId,
			ProjectSecret: projectSecret,
		},
	}
}

// authTransport decorates each request with a basic auth header.
type authTransport struct {
	http.RoundTripper
	ProjectId     string
	ProjectSecret string
}

func (t authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.ProjectId, t.ProjectSecret)
	return t.RoundTripper.RoundTrip(r)
}
