package core

type Team struct {
	Name string
	Members []string
}

func (t *Team) AddMember(member string){
	t.Members = append(t.Members, member)
}