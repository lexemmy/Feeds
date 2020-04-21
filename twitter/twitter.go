package twitter

type Twitter struct{
	URL		string
	Name	string
	Followers int
}

func (t *Twitter) Feed() []string{
	return []string{
		"My twitter feeds",
		"hey!, i just posted a tweet",
		"twitter is savage",
	}
}

func (t *Twitter) Fame() int {
	return t.Followers	
}