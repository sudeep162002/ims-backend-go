package models

type User struct {
	UserID       string `json:"userId" bson:"userId,omitempty"`
	FullName     string `json:"fullName" bson:"fullName"`
	RitwickName  string `json:"ritwickName" bson:"ritwickName"`
	Swastyayani  string `json:"swastyayani" bson:"swastyayani"`
	Istavrity    string `json:"istavrity" bson:"istavrity"`
	Acharyavrity string `json:"acharyavrity" bson:"acharyavrity"`
	Dakshina     string `json:"dakshina" bson:"dakshina"`
	Sangathani   string `json:"sangathani" bson:"sangathani"`
	Ritwicki     string `json:"ritwicki" bson:"ritwicki"`
	Proname      string `json:"proname" bson:"proname"`
	Anandabazar  string `json:"anandabazar" bson:"anandabazar"`
	Srimandir    string `json:"srimandir" bson:"srimandir"`
	Parivrity    string `json:"parivrity" bson:"parivrity"`
	Misc         string `json:"misc" bson:"misc"`
	Address      string `json:"address" bson:"address"`
}
