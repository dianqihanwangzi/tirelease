package controller

import (
	"tirelease/commons/configs"
	"tirelease/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v41/github"
)

func GithubWebhookHandler(c *gin.Context) {
	// parse webhook payload
	payload, err := github.ValidatePayload(c.Request, []byte(configs.Config.Github.WebhookSecret))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	event, err := github.ParseWebHook(github.WebHookType(c.Request), payload)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	// handle event
	switch event := event.(type) {

	case *github.IssuesEvent:
		err := service.WebhookRefreshIssueV4(event.Issue)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

	case *github.IssueCommentEvent:
		err := service.WebhookRefreshIssueV4(event.Issue)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

	case *github.PullRequestEvent:
		err := service.WebhookRefreshPullRequestV3(event.PullRequest)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

	default:
		c.JSON(200, gin.H{"status": "ok", "data": "no listener, return directly"})

	}
}
