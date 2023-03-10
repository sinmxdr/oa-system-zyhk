package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddOffice(c *gin.Context) {
	var office models.Office
	_ = c.ShouldBindJSON(&office)

	officeCre := models.Office{
		ID:                   office.ID,
		IsDelete:             false,
		Name:                 office.Name,
		Number:               office.Number,
		TaskLoad:             office.TaskLoad,
		RoleID:               office.RoleID,
		IsRanking:            office.IsRanking,
		RankingNo:            office.RankingNo,
		PushMoneyPercentages: 0,
		BusinessMoney:        0,
		Money:                0,
		MoneyCold:            0,
		TargetLoad:           0,
		YWTargetLoad:         0,
		ZYTargetLoad:         0,
		QDTargetLoad:         0,
		IsSubmit:             false,
	}

	code = models.GeneralInsert(&officeCre)
	msg.Message(c, code, nil)
}

func DelOffice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		code = models.GeneralDelete(&models.Office{}, id)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func EditOfficeBase(c *gin.Context) {
	var office models.Office
	_ = c.ShouldBindJSON(&office)
	var maps = make(map[string]interface{})
	maps["name"] = office.Name
	maps["number"] = office.Number
	maps["task_load"] = office.TaskLoad
	if office.RoleID == 0 {
		maps["role_id"] = nil
	} else {
		maps["role_id"] = office.RoleID
	}
	maps["ranking_no"] = office.RankingNo
	maps["push_money_percentages"] = office.PushMoneyPercentages
	code = models.GeneralUpdate(&models.Office{}, office.ID, maps)
	msg.Message(c, code, nil)
}

func EditOfficeMoney(c *gin.Context) {
	var office models.Office
	_ = c.ShouldBindJSON(&office)
	code = models.UpdateOfficeMoney(&office, c.MustGet("employeeID").(int))
	msg.Message(c, code, nil)
}

func QueryOffices(c *gin.Context) {
	var officeQuery models.Office
	_ = c.ShouldBindJSON(&officeQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectOffices(&officeQuery, &xForms)
	msg.Message(c, code, xForms)
}

func QueryAllOffice(c *gin.Context) {
	var offices []models.Office
	code = models.GeneralSelectAll(&offices, nil)
	msg.Message(c, code, offices)
}
