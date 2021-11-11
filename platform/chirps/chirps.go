package chirps

type Chirp struct {
	Content  string `json:"content"`
	Location string `json:"location"`
}

type Repo struct {
	Chirps []Chirp
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Add(chirp Chirp) {
	r.Chirps = append(r.Chirps, chirp)
}

func (r *Repo) GetAll() []Chirp {
	return r.Chirps
}