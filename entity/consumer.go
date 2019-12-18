package entity

type Recovery struct {
	VisibleTimeot int `json:"visible_timeout"`
	Count         int `json:"try_counter"`
}

type StateFullModels struct {
	Topic     string   `json:"topic"`
	Status    string   `json:"status"`
	Recovery  Recovery `json:"recovery"`
	Parameter string   `json:"parameter"`
	Body      string   `json:"body"`
}
