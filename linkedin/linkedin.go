package linkedin

type Linkedin struct{
	URL		string
	Name	string
	connections int
}

func (l *Linkedin) Feed() []string{
	return []string{
		"My Linkedin feeds",
		"hey!, i just got new job",
		"what's your cv",
	}
}

func (l *Linkedin) Fame() int {
	return l.connections	
}