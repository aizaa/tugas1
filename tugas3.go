package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Job		  string
	Salary	  string
}

func (p *Person) SetFirst(firstname string) {
	p.FirstName = firstname
}
func (p *Person) SetLast(lastname string) {
	p.LastName = lastname
}
func (p *Person) SetJob(job string) {
	p.Job = job
}
func (p *Person) SetSalary(salary string) {
	p.Salary = salary
}

func (p Person) First() string {
	return p.FirstName 
}
func (p Person) Last() string {
	return p.LastName 
}
func (p Person) Jobs() string {
	return p.Job 
}
func (p Person) Salarys() string {
	return p.Salary 
}

func main() {
    p := Person{}
	p.SetFirst("Budi")
	p.SetLast("Pratama")
	p.SetJob("Marketing")
	p.SetSalary("2.000.000")

	firstname := p.First()
	lastname := p.Last()
	job := p.Jobs()
	salary := p.Salarys()

	fmt.Println("Firstname : ", firstname)
	fmt.Println("Lastname : ", lastname)
	fmt.Println("Job : ", job)
	fmt.Println("Salary : ", salary)
}