package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type project struct {
	ID         string `json:"id"`
	Topic      string `json:"topic"`
	Student    string `json:"student"`
	Supervisor string `json:"supervisor"`
	Year       string `json:"year"`
}

var projects = []project{
	{
		ID:         "1",
		Topic:      "Design and fabrication of a self-driving car",
		Student:    "Ignatius Timothy",
		Supervisor: "Dr Bori Ige",
		Year:       "2023",
	},
	{
		ID:         "2",
		Topic:      "Design and fabrication of a robotic sports leg",
		Student:    "Ogbodo Patrick",
		Supervisor: "Dr. S. A. Ayo",
		Year:       "2023",
	},
	{
		ID:         "3",
		Topic:      "Finite Element Analysis of Carbon Fibre Engine Block",
		Student:    "Bamtefa Moses",
		Supervisor: "Prof O. O. Olugboji",
		Year:       "2023",
	},
}

func getProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projects)
}

func postProjects(c *gin.Context) {
	var newProject project

	// bind the received JSON object to newProject
	if err := c.BindJSON(&newProject); err != nil {
		return
	}

	// add the new project to the projects slice
	projects = append(projects, newProject)
	c.IndentedJSON(http.StatusCreated, newProject)
}

func getProjectById(c *gin.Context) {
	id := c.Param("id")
	for _, proj := range projects {
		if proj.ID == id {
			c.IndentedJSON(http.StatusOK, proj)
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "project not found"})
}

func main() {
	router := gin.Default()
	router.GET("/projects", getProjects)
	router.GET("/projects/:id", getProjectById)
	router.POST("/projects", postProjects)

	router.Run("localhost:8888")
}
