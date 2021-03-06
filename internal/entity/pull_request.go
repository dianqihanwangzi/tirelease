package entity

import (
	"time"

	"github.com/google/go-github/v41/github"
)

// Struct of Pull Request
type PullRequest struct {
	// DataBase columns
	ID            int64  `json:"id,omitempty"`
	PullRequestID string `json:"pull_request_id,omitempty"`
	Number        int    `json:"number,omitempty"`
	State         string `json:"state,omitempty"`
	Title         string `json:"title,omitempty"`
	Repo          string `json:"repo,omitempty"`
	HTMLURL       string `json:"html_url,omitempty"`
	HeadBranch    string `json:"head_branch,omitempty"`

	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	ClosedAt  *time.Time `json:"closed_at,omitempty"`
	MergedAt  *time.Time `json:"merged_at,omitempty"`

	Merged         bool   `json:"merged,omitempty"`
	Mergeable      bool   `json:"mergeable,omitempty"`
	MergeableState string `json:"mergeable_state,omitempty"`

	SourcePullRequestID string `json:"source_pull_request_id,omitempty"`

	LabelsString             string `json:"labels_string,omitempty"`
	AssigneeString           string `json:"assignee_string,omitempty"`
	AssigneesString          string `json:"assignees_string,omitempty"`
	RequestedReviewersString string `json:"requested_reviewers_string,omitempty"`

	// OutPut-Serial
	Labels             *[]github.Label `json:"labels,omitempty" gorm:"-"`
	Assignee           *github.User    `json:"assignee,omitempty" gorm:"-"`
	Assignees          *[]github.User  `json:"assignees,omitempty" gorm:"-"`
	RequestedReviewers *[]github.User  `json:"requested_reviewers,omitempty" gorm:"-"`
}

// List Option
type PullRequestOption struct {
	ID                  int64  `json:"id"`
	PullRequestID       string `json:"pull_request_id,omitempty"`
	Number              int    `json:"number,omitempty"`
	State               string `json:"state,omitempty"`
	Repo                string `json:"repo,omitempty"`
	HeadBranch          string `json:"head_branch,omitempty"`
	SourcePullRequestID string `json:"source_pull_request_id,omitempty"`
}

// DB-Table
func (PullRequest) TableName() string {
	return "pull_request"
}

/**

CREATE TABLE IF NOT EXISTS pull_request (
	id INT(11) NOT NULL AUTO_INCREMENT COMMENT '??????',
	pull_request_id VARCHAR(255) COMMENT 'Pr??????ID',
	number INT(11) NOT NULL COMMENT '?????????ID',
	state VARCHAR(32) NOT NULL COMMENT '??????',
	title VARCHAR(1024) COMMENT '??????',
	repo VARCHAR(255) COMMENT '??????',
	html_url VARCHAR(1024) COMMENT '??????',
	head_branch VARCHAR(255) COMMENT '??????',

	closed_at TIMESTAMP COMMENT '????????????',
	created_at TIMESTAMP COMMENT '????????????',
	updated_at TIMESTAMP COMMENT '????????????',
	merged_at TIMESTAMP COMMENT '????????????',

	merged BOOLEAN COMMENT '???????????????',
	mergeable BOOLEAN COMMENT '???????????????',
	mergeable_state VARCHAR(32) COMMENT '???????????????',

	source_pull_request_id VARCHAR(255) COMMENT '??????ID',

	labels_string TEXT COMMENT '??????',
	assignee_string TEXT COMMENT '?????????',
	assignees_string TEXT COMMENT '???????????????',
	requested_reviewers_string TEXT COMMENT '???????????????',

	PRIMARY KEY (id),
	UNIQUE KEY uk_prid (pull_request_id),
	INDEX idx_state (state),
	INDEX idx_repo (repo),
	INDEX idx_createdat (created_at),
	INDEX idx_sourceprid (source_pull_request_id)
)
ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT 'pull_request?????????';

**/
