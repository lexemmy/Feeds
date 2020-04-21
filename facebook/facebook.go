package facebook

type Facebook struct{
	URL		string
	Name	string
	Friends int
}

func (f *Facebook) Feed() []string {
	return []string{
		"My facebook feeds",
		"hey!, here's my cool new selfie",
		"what's on your mind?",
	}
}

func (f *Facebook) Fame() int {
	return f.Friends	
}