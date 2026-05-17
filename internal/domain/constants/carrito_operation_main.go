package constants

// crea un nuevo tipo derivado de string
// mas seguro, mas explicito, mas entendible
type CarritoOperationResult string

const (
	Created CarritoOperationResult = "Created"
	Deleted CarritoOperationResult = "Delete"
	Updated CarritoOperationResult = "Updated"
	Error   CarritoOperationResult = "Error"
)

// convierte el tipo personalizado a string normal
func (operation CarritoOperationResult) String() string {
	return string(operation)
}
