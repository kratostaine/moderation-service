package repository

type Repository struct{}

var objectionableData map[string][]string = map[string][]string{
	"poop":   {},
	"murder": {},
	"shit":   []string{"dikshit"},
	"nazi":   {},
	"kill":   []string{"overkill"},
}

func (r Repository) GetObjectionableData() map[string][]string {
	return objectionableData
}

func (r Repository) GetMaxLengthOfObjectionableText() int {
	return 6 //TODO: To be modified to dynamically obtain the value
}
