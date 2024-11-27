package dao

type Video struct {
	Key         string   `json:"_key" validate:"required"`
	Publishable bool     `json:"publishable" validate:"required"`
	Categories  []string `json:"categories" validate:"required,dive,required"`
	Description string   `json:"description,omitempty"`
	Name        string   `json:"name" validate:"required"`
	Views       int      `json:"views,omitempty" validate:"gte=0"`
	Type        string   `json:"type,omitempty" validate:"oneof=movie series tvshow"`
}

type VideoByName struct {
	Publishable bool     `json:"publishable" validate:"required"`
	Categories  []string `json:"categories" validate:"required,dive,required"`
	Description string   `json:"description,omitempty"`
	Name        string   `json:"name" validate:"required"`
	Views       int      `json:"views,omitempty" validate:"gte=0"`
	Type        string   `json:"type,omitempty" validate:"oneof=movie series tvshow"`
}
