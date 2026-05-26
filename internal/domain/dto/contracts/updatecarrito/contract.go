package updatecarrito

type Body struct {
	Movements []Movement `json:"movements"`
}

type Movement struct {
	DulceID  uint64 `json:"dulce_id" validate:"required,number"`
	Unidades uint64 `json:"unidades" validate:"required,number"`
}
