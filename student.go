package main

type Student struct {
	name    string
	email   string
	address map[string]string
	status  string
}

func newStudent(n string, e string, addr map[string]string) Student {
	stud := Student{
		name:    n,
		email:   e,
		address: addr,
		status:  "active",
	}
	return stud

}
func (s Student) updateStatus(stat string) {
	s.status = stat
}
