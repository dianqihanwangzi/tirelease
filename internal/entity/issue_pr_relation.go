package entity

import (
	"time"
)

// Struct of Issue_Pr_Relation
type IssuePrRelation struct {
	ID         int64     `json:"id,omitempty"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`

	IssueID       int `json:"issue_id,omitempty"`
	PullRequestID int `json:"pull_request_id,omitempty"`
}

// List Option
type IssuePrRelationOption struct {
	ID            int64 `json:"id"`
	IssueID       int   `json:"issue_id,omitempty"`
	PullRequestID int   `json:"pull_request_id,omitempty"`
}

// DB-Table
func (IssuePrRelation) TableName() string {
	return "issue_pr_relation"
}

/**

CREATE TABLE IF NOT EXISTS issue_pr_releation (
	id INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
	create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

	issue_id INT(11) NOT NULL COMMENT 'issue的ID',
	pull_request_id INT(11) NOT NULL COMMENT 'pr的ID',

	PRIMARY KEY (id),
	UNIQUE KEY uk_issueid_prid (issue_id, pull_request_id)
)
ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT 'issue与pull_request关联表';

**/
