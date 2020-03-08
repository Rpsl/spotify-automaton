package storage

type Track struct {
	ID         string
	Name       string
	Duration   int
	Explicit   bool
	Popularity int
	Artists    []string
}

type TrackStack struct {
	Tracks map[string]Track
}

func NewTrackStack() *TrackStack {
	ts := TrackStack{}
	ts.Tracks = make(map[string]Track)

	return &ts
}

func (t *TrackStack) Add(id string, name string, duration int, explicit bool, popularity int, artists []string) {
	t.Tracks[id] = Track{
		ID:         id,
		Name:       name,
		Duration:   duration,
		Explicit:   explicit,
		Popularity: popularity,
		Artists:    artists,
	}
}
