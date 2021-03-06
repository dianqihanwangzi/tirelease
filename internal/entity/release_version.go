package entity

import (
	"time"
)

type ReleaseVersion struct {
	// Columns
	ID int64 `json:"id,omitempty"`

	CreateTime        time.Time  `json:"create_time,omitempty"`
	UpdateTime        time.Time  `json:"update_time,omitempty"`
	PlanReleaseTime   *time.Time `json:"plan_release_time,omitempty"`
	ActualReleaseTime *time.Time `json:"actual_release_time,omitempty"`

	Name        string               `json:"name,omitempty"`
	Description string               `json:"description,omitempty"`
	Owner       string               `json:"owner,omitempty"`
	Type        ReleaseVersionType   `json:"type,omitempty"`
	Status      ReleaseVersionStatus `json:"status,omitempty"`

	FatherReleaseVersionName string `json:"father_release_version_name,omitempty"`

	ReposString  string `json:"repos_string,omitempty"`
	LabelsString string `json:"labels_string,omitempty"`

	// OutPut-Serial
	Repos  *[]string `json:"repos,omitempty" gorm:"-"`
	Labels *[]string `json:"labels,omitempty" gorm:"-"`
}

// Enum status
type ReleaseVersionStatus string

const (
	ReleaseVersionStatusOpen     = ReleaseVersionStatus("Open")
	ReleaseVersionStatusClosed   = ReleaseVersionStatus("Closed")
	ReleaseVersionStatusReleased = ReleaseVersionStatus("Released")
)

// Enum type
type ReleaseVersionType string

const (
	ReleaseVersionTypeMajor = ReleaseVersionType("Major")
	ReleaseVersionTypeMinor = ReleaseVersionType("Minor")
	ReleaseVersionTypePatch = ReleaseVersionType("Patch")
)

// List Option
type ReleaseVersionOption struct {
	ID                       int64                `json:"id"`
	Name                     string               `json:"name,omitempty"`
	FatherReleaseVersionName string               `json:"father_release_version_name,omitempty"`
	Type                     ReleaseVersionType   `json:"type,omitempty"`
	Status                   ReleaseVersionStatus `json:"status,omitempty"`
}

// DB-Table
func (ReleaseVersion) TableName() string {
	return "release_version"
}

/**

CREATE TABLE IF NOT EXISTS release_version (
	id INT(11) NOT NULL AUTO_INCREMENT COMMENT '??????',
	create_time TIMESTAMP COMMENT '????????????',
	update_time TIMESTAMP COMMENT '????????????',
	plan_release_time TIMESTAMP COMMENT '??????????????????',
	actual_release_time TIMESTAMP COMMENT '??????????????????',

	name VARCHAR(255) NOT NULL COMMENT '?????????',
	description VARCHAR(1024) COMMENT '????????????',
	owner VARCHAR(255) COMMENT '???????????????',
	type VARCHAR(32) COMMENT '????????????',
	status VARCHAR(32) COMMENT '????????????',

	father_release_version_name VARCHAR(255) COMMENT '????????????',
	repos_string VARCHAR(1024) COMMENT '??????????????????',
	labels_string VARCHAR(1024) COMMENT '??????????????????',

	PRIMARY KEY (id),
	UNIQUE KEY uk_name (name),
	INDEX idx_createtime (create_time),
	INDEX idx_fathername (father_release_version_name)
)
ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT '???????????????';

**/
