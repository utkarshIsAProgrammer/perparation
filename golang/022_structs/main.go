package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Skills struct {
	SkillOne   string
	SkillTwo   string
	SkillThree string
}

// value receiver of struct "Skills"
func (s Skills) showSkills() {
	fmt.Println(s.SkillOne)
	fmt.Println(s.SkillTwo)
	fmt.Println(s.SkillThree)
}

// original pointer address of the struct
func (s *Skills) pointerSkills() {
	s.SkillOne = "Changed by pointer"
	fmt.Println("Inside pointer:", s.SkillOne)
}

func main() {
	u1 := User{
		ID:    1,
		Name:  "Indiedev",
		Email: "indiedev@gmail.com",
		Age:   20,
	}
	fmt.Println(u1.ID)

	mySkills := Skills{
		SkillOne:   "Golang",
		SkillTwo:   "Python",
		SkillThree: "Javascript",
	}
	fmt.Println(mySkills)

	mySkills.pointerSkills()
	fmt.Println("After pointer:", mySkills.SkillOne)
}
