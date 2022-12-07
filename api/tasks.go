package api

import (
	"TodoList/pkg/utils"
	"TodoList/service"
	"fmt"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// 新增
func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	fmt.Println("Authorzation1=======", c.GetHeader("Authorization"))

	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	fmt.Println("Authorzation2=======", claim)
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id)
		c.JSON(200, res) //返回
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
		// c.JSON(400, err)
	}
}

// 查询备忘录详细信息
func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	// , _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

// 查询所有备忘录
func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}

}

// 查询所有备忘录
func ListTask1(c *gin.Context) {
	var listTask service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List1(claim.Id, c.Param("page_num"), c.Param("page_size"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}

}

// 修改备忘录
func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService
	fmt.Println("======", c)

	if err := c.ShouldBind(&updateTask); err == nil {
		res := updateTask.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

// 模糊查询
func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	fmt.Println("searchTask=====", searchTask)
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTask); err == nil {
		res := searchTask.Search(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

// 删除
func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	if err := c.ShouldBind(&deleteTask); err == nil {
		res := deleteTask.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}
