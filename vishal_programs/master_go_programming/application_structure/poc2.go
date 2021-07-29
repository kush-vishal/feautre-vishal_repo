package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type employee2 struct {
	employeeId    int    `json:employeeId`
	employeeName  string `json:"employeeName"`
	employeePhone int    `json:employeePhone`
}

var Emp employee2

func employeeData2() [2]employee2 {
	var employeeDetails [2]employee2
	employeeDetails[0].employeeId = 13
	employeeDetails[0].employeeName = "vishal kushwaha"
	employeeDetails[0].employeePhone = 1234
	employeeDetails[1].employeeId = 12
	employeeDetails[1].employeeName = "vikash kushwaha"
	employeeDetails[1].employeePhone = 4567
	return employeeDetails
}

func updateEmployeeData(emp employee2, index int, employeeDetails2 [2]employee2) [2]employee2 {
	employeeDetails2[index] = emp
	return employeeDetails2
}

func findEmployeeIndex(id int, employeeDetails [2]employee2) int {

	for i := range employeeDetails {
		if employeeDetails[i].employeeId == id {
			return i
		}
	}
	return -1
}

func main() {

	fmt.Println(employeeData2())
	server2()
	fmt.Println(employeeData2())

}
func Serveput(w http.ResponseWriter, r *http.Request) {
	a := findEmployeeIndex(12, employeeData2())
	var Emp employee2
	if a > -1 {
		err := json.NewDecoder(r.Body).Decode(&Emp)
		//b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		/*err = json.Unmarshal(b, &Emp)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}*/
		fmt.Fprintf(w, "Employee details: %v \n\n", Emp)
		fmt.Println(Emp.employeeId)
		fmt.Println(updateEmployeeData(Emp, a, employeeData2()))
		fmt.Fprintf(w, "Employee details are: %v\n\n", employeeData2())
		defer r.Body.Close()
	} else {
		fmt.Fprintf(w, "Employee details not found")
	}

}
func server2() {
	http.HandleFunc("/hola", Serveput)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
