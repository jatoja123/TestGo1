package TestZGo

import "strconv"

//zad 1
func Suma(a, b int) int {
	return a + b
}

//zad 2
func Concatenate(s1, s2 int) []int {
	return []int{s1, s2}
}

//zad 3
func Is18(age int) bool {
	return age > 17
}

//zad 4
func PrintDividedBy4(n int) []int {
	result := []int{}
	current := 4
	for current < n {
		result = append(result, current)
		current += 4
	}

	return result
}

//zad 5
func FilterEven(age []int) []int {
	result := []int{}
	for _, n := range age {
		if n%2 == 1 {
			result = append(result, n)
		}
	}
	return result
}

//zad 6
type User struct {
	name     string
	lastName string
	age      int
	isAdult  bool
}

func NewUser(name, lastName string, age int) *User {
	return &User{
		name:     name,
		lastName: lastName,
		age:      age,
		isAdult:  Is18(age),
	}
}

//zad 7
func (u *User) ToString() string {
	adultTxt := ""
	ageTxt := strconv.FormatInt(int64(u.age), 10)

	if u.isAdult {
		adultTxt = "Adult"
	} else {
		adultTxt = "Not an Adult"
	}

	return "| " + u.name + " " + u.lastName + " | " + ageTxt + " | " + adultTxt + " |"
}

func (u *User) ChangeAge(newAge int) {
	if newAge <= 120 && newAge >= 0 {
		u.age = newAge
	}
}
