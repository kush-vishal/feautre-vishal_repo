package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type employee3 struct {
	Id      int    `json:Id`
	Name    string `json:"Name`
	Address string `json:"Address"`
}

var database []employee3

func findEmployee1(id int) (employee3, int) {
	var a employee3
	var k int
	for i := range database {
		if database[i].Id == id {
			a = database[i]
			k++
			return a, k

		}
	}
	return a, k
}
func addEmployee(emp employee3) {
	database = append(database, emp)
}
func employeeData3() {
	data := []employee3{{1, "vishal", "adrress--howrah"}, {2, "bikash", "howrah-711107"}}
	database = append(database, data...)

}
func updateEmployee(emp employee3, id int) bool {
	for i := range database {
		if database[i].Id == id {
			database[i] = emp
			return true
		}
	}
	return false

}

func main() {
	employeeData3()

	e := echo.New()
	e.GET("/hellow", func(c echo.Context) error {
		return c.String(http.StatusOK, "hellow world")
	})

	e.GET("/employee/:id", findEmployeeHandler)
	e.PUT("/employee/replace/:id", updateEmployeeHandler)
	e.POST("/employee/add", addnewHandler)
	fmt.Println(database)
	e.Logger.Fatal(e.Start(":8080"))

}

func findEmployeeHandler(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	data, K := findEmployee1(id)
	if K == 1 {
		fmt.Println(K)

		return c.JSON(http.StatusOK, data)

	} else {
		return c.String(http.StatusOK, "not found")
	}

}

func updateEmployeeHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	var em employee3
	if err := c.Bind(&em); err != nil {
		return err
	}
	updateEmployee(em, id)
	return c.JSON(http.StatusOK, database)

}
func addnewHandler(c echo.Context) error {

	var em employee3
	if err := c.Bind(&em); err != nil {
		return err
	}
	addEmployee(em)
	return c.JSON(http.StatusOK, database)
}
