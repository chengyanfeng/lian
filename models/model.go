package models

type Node struct {
	Data    string `json:"data"`
	TxIds   string `bson:"txIds"`
	AppKey  string `json:"appkey"`
	Created string `json:"created"`
	__v     string `json:"__v";`
	Number  int    `json:"number"`
}
