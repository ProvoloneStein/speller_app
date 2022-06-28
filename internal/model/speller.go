package model

//input model
type Speller struct {
	Text string `json:"text" form:"text" binding:"required"`
}

//input model
type Spellers struct {
	Text []string `json:"texts" form:"texts" binding:"required"`
}

//speller output model
type SpellerResponse struct {
	//Code int `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}
