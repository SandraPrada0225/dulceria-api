package responses

type Response struct {
	NuevoCarritoID uint64 `json:"nuevo_carrito_id"`
}

func NewResponse(carritoID uint64) Response {
	return Response{
		NuevoCarritoID: carritoID,
	}
}
