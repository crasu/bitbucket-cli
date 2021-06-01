package cli

import (
	"fmt"
	bitbucket "github.com/gfleury/go-bitbucket-v1"
	"strings"
)

type RepoPrListCmd struct {
	State string `arg:"-s,--state"`
}

func (b *BitbucketCLI) repoPrList(cmd *RepoCmd) {
	if cmd == nil || cmd.PrCmd == nil || cmd.PrCmd.List == nil {
		return
	}

	lCmd := cmd.PrCmd.List

	opts := map[string]interface{}{}
	if lCmd.State != "" {
		validStates := []string{
			"ALL",
			"OPEN",
			"DECLINED",
			"MERGED",
		}

		inputUpper := strings.ToUpper(lCmd.State)

		if !validValue(inputUpper, validStates) {
			b.logger.Fatalf("invalid value \"%s\" for state: accepted values are: \n%s",
				lCmd.State,
				strings.Join(validStates, "\n"),
			)
			return
		}
		opts["state"] = inputUpper
	}

	prs, err := b.client.DefaultApi.GetPullRequestsPage(cmd.ProjectKey, cmd.Slug, opts)
	if err != nil {
		b.logger.Fatalf("unable to get PRs: %v", err)
		return
	}

	prsResponse, err := bitbucket.GetPullRequestsResponse(prs)
	if err != nil {
		b.logger.Fatalf("unable to parse PRs response: %v", err)
	}

	for _, pr := range prsResponse {
		fmt.Printf("%s (ID: %d)", pr.Title, pr.ID)
	}

}
