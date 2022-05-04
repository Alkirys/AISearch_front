package model

type Find struct {
	Url       string     `json:"url"`
	Boundings [][]uint64 `json:"boundings"`
}

type Detect struct {
	Query     string   `json:"query"`
	Filename  string   `json:"filename"`
	Boundings []uint64 `json:"boundings"`
}

type Img struct {
	Name  string
	Bound []uint64
}

type Faces struct {
	Faces [][]interface{} `json:"faces"`
}
