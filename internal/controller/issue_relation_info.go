package controller

import (
	"tirelease/internal/dto"
	"tirelease/internal/repository"

	"github.com/gin-gonic/gin"
)

func SelectIssueRelationInfos(c *gin.Context) {
	// Params
	option := dto.IssueRelationInfoQuery{}
	c.ShouldBind(&option)

	// Action
	issueRelationInfos, err := repository.SelectIssueRelationInfo(&option)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"data": issueRelationInfos})
}
