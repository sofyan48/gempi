package entity

type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type BodyValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Recovery struct {
	VisibleTimeot int `json:"visible_timeout"`
	Count         int `json:"try_counter"`
}

type StateFullModels struct {
	Topic     string      `json:"topic"`
	Status    string      `json:"status"`
	Recovery  Recovery    `json:"recovery"`
	Parameter []Parameter `json:"parameter"`
	Body      BodyValue   `json:"body"`
}
