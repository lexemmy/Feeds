package feeds

type SocialMedia interface{
	Feed() []string
	Fame() int
}