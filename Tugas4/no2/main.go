package main

import "fmt"

type Student struct {
	Name  []string
	Score []float64
}

func (s *Student) AddStudent(name string, score float64) {
	s.Name = append(s.Name, name)
	s.Score = append(s.Score, score)
}

func (s *Student) GetAverage() float64 {
	total := 0.0
	for _, val := range s.Score {
		total += val
	}
	average := total / float64(len(s.Score))

	return average
}

func (s *Student) GetMinScore() (string, float64) {
	min := s.Score[0]
	var index int
	for i, val := range s.Score {
		if val < min {
			min = val
			index = i
		}
	}
	return s.Name[index], min
}

func (s *Student) GetMaxScore() (string, float64) {
	max := 1.0
	var index int
	for i, val := range s.Score {
		if val > max {
			max = val
			index = i
		}
	}
	return s.Name[index], max
}

func main() {
	var students Student
	students.AddStudent("Pandu Pramudya", 99)
	students.AddStudent("MacAllister", 80)
	students.AddStudent("Zeni Haryadi", 79)
	students.AddStudent("Kai Havertz", 60)
	students.AddStudent("Martin Odegaard", 82)

	//fmt.Println("Data Siswa:")
	//for i := 0; i < len(students.Name); i++ {
	//	fmt.Printf("Siswa: %s\n", students.Name[i])
	//	fmt.Printf("Skor: %.2f\n", students.Score[i])
	//	fmt.Println()
	//}

	getAverage := students.GetAverage()
	getNameMin, getMinScore := students.GetMinScore()
	getNameMax, getMaxScore := students.GetMaxScore()

	fmt.Printf("Average Score : %v\n", getAverage)
	fmt.Printf("Max Score of Students : %s (%v)\n", getNameMax, getMaxScore)
	fmt.Printf("Min Score of Students : %s (%v)", getNameMin, getMinScore)
}
