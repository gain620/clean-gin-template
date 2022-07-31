package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

const (
	USA   = "USA"
	INDIA = "India"
	CHINA = "China"
)

type Employee interface {
	PaySalary()
}

type employee struct {
	Name, Position, Location string
	Salary                   int
	Currency                 string
}

type EmployeeFactory func(name string) Employee

func (e *employee) PaySalary() {
	fmt.Printf("Pay %s Salary of %s %v\n", e.Name, e.Currency, e.Salary)
}

func NewEmployeeFactory(location, position string, salary int) EmployeeFactory {
	return func(name string) Employee {
		return NewEmployee(name, position, location)
	}
}

func NewEmployee(name, position, location string) Employee {
	switch location {
	case USA:
		return &employee{
			Name:     name,
			Position: position,
			Location: USA,
			Salary:   200000,
			Currency: "USD",
		}
	case INDIA:
		return &employee{
			Name:     name,
			Position: position,
			Location: INDIA,
			Salary:   1500000,
			Currency: "INR",
		}
	case CHINA:
		return &employee{
			Name:     name,
			Position: position,
			Location: CHINA,
			Salary:   130000,
			Currency: "Chinese Yuan",
		}
	default:
		return &employee{
			Name:     name,
			Position: position,
			Salary:   20000,
			Currency: "USD",
		}
	}
}

func main() {
	chinaFactory := NewEmployeeFactory("China", "DevOps Engineer", 20000)
	xinEmployee := chinaFactory("Xin")
	log.Infof("Employee : %v", xinEmployee)
}
