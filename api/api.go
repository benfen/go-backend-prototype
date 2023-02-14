package api

import (
	"fmt"
	"net/http"

	"github.com/benfen/go-prototype-backend/db"
	"github.com/gin-gonic/gin"
)

type Person struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

type API struct {
	db     *db.DB
	engine *gin.Engine
}

func (api *API) handleGet(c *gin.Context) {
	people, err := api.db.GetPeople()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	outgoingPeople := make([]Person, len(people))
	for i, person := range people {
		outgoingPeople[i] = Person{
			ID:    person.ID,
			Name:  person.Name,
			Score: person.Score,
		}
	}

	c.JSON(http.StatusOK, outgoingPeople)
}

func (api *API) handleAdd(c *gin.Context) {
	var person Person
	err := c.BindJSON(&person)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err = api.db.AddPerson(db.Person{
		Name:  person.Name,
		Score: person.Score,
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func (api *API) handleUpdate(c *gin.Context) {
	var person Person
	err := c.BindJSON(&person)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err = api.db.UpdatePerson(db.Person{
		Name:  person.Name,
		Score: person.Score,
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func (api *API) handleDelete(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("attempting delete")
	fmt.Println(id)

	err := api.db.DeletePerson(id)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func CreateAPI(engine *gin.Engine, dbHandle *db.DB) {
	api := API{
		db:     dbHandle,
		engine: engine,
	}

	rg := engine.Group("/api")

	rg.GET("/", api.handleGet)
	rg.POST("/", api.handleAdd)
	rg.PUT("/", api.handleUpdate)
	rg.DELETE("/:id", api.handleDelete)
}
