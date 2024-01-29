package dto

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Relation struct {
	Id           int    `json:"id"`
	RelationName string `json:"relation_name,omitempty"`
}

type Relationship struct {
	Id          int    `json:"id"`
	RelatedFrom string `json:"relatedFrom,omitempty"`
	RelatedAs   string `json:"relatedAs,omitempty"`
	RelatedTo   string `json:"relatedTo,omitempty"`
}

type Count struct {
	Number int `json:"number,omitempty"`
}
