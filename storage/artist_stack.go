package storage

type Artist struct {
	ID   string
	Name string
}

type ArtistStack struct {
	Artists map[string]Artist
}

func NewArtistStack() *ArtistStack {
	as := ArtistStack{}
	as.Artists = make(map[string]Artist)

	return &as
}

func (as *ArtistStack) Add(id string, name string) {
	as.Artists[id] = Artist{
		ID:   id,
		Name: name,
	}
}
