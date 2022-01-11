package entity

import (
	"time"
)

type VersionTriage struct {
	ID           int64               `json:"id,omitempty"`
	VersionName  string              `json:"version_name,omitempty"`
	IssueID      string              `json:"issue_id,omitempty"`
	TriageResult VersionTriageResult `json:"triage_result,omitempty"`

	CreateTime time.Time  `json:"create_time"`
	UpdateTime time.Time  `json:"update_time"`
	DueTime    *time.Time `json:"due_time,omitempty"`
	Comment    string     `json:"comment,omitempty"`
}

// Enum type
type VersionTriageResult string

const (
	VersionTriageResultUnKnown  = VersionTriageResult("unknown")
	VersionTriageResultAccept   = VersionTriageResult("accept")
	VersionTriageResultLater    = VersionTriageResult("later")
	VersionTriageResultWontFix  = VersionTriageResult("won't-fix")
	VersionTriageResultReleased = VersionTriageResult("released")
)

// List Option
type VersionTriageOption struct {
	ID           int64               `json:"id"`
	VersionName  string              `json:"version_name,omitempty"`
	IssueID      string              `json:"issue_id,omitempty"`
	TriageResult VersionTriageResult `json:"triage_result,omitempty"`
}

// DB-Table
func (VersionTriage) TableName() string {
	return "version_triage"
}

/**

CREATE TABLE IF NOT EXISTS version_triage (
	id INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
	version_name VARCHAR(255) NOT NULL COMMENT '版本号',
	issue_id VARCHAR(255) COMMENT 'Issue全局ID',
	triage_result VARCHAR(32) COMMENT 'Triage状态',

	create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	due_time TIMESTAMP COMMENT '延期时间',
	comment VARCHAR(1024) COMMENT '评论',

	PRIMARY KEY (id),
	UNIQUE KEY uk_versionname_issueid (version_name, issue_id),
	INDEX idx_issueid (issue_id)
)
ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT '版本Triage信息表';

**/