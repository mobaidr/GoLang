package  main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type fixedbilling struct {
	projectName string
	biddedAmount int
}

type timeAndMaterial struct {
	projectName string
	noOfHours int
	hourRate int
}

func (fb fixedbilling) source() string {
	return fb.projectName
}

func (fb fixedbilling) calculate() int {
	return fb.biddedAmount
}

func (tm timeAndMaterial) source() string {
	return tm.projectName
}

func (tm timeAndMaterial) calculate() int {
	return tm.hourRate*tm.noOfHours
}

func calculateNetIncome(ic []Income) {
	var netIncome int = 0

	for _, income := range ic {
		fmt.Printf("Income from %s = %d\n", income.source(), income.calculate())
		netIncome += income.calculate()
	}

	fmt.Printf("Net income of the organization = $%d",netIncome)
}

func main() {
	project1 := fixedbilling{
		projectName:  "Project 1",
		biddedAmount: 5000,
	}

	project2 := fixedbilling{
		projectName:  "Project 2 ",
		biddedAmount: 10000,
	}

	project3 := timeAndMaterial{
		projectName: "Project 3",
		noOfHours:   160,
		hourRate:    35,
	}

	projects := []Income{project1,project2,project3}

	calculateNetIncome(projects)
}