package responseCodes

type ApiResponse struct {
	HttpStatusCode int
	Detail         IResponseDetail
}

type ResponseDetail struct {
	Code string `json:"code"`
}

type CommonResponseDetail struct {
	ResponseDetail
	Message string `json:"message"`
}

/*
type BadRequestDetail struct {
	ResponseDetail
	Details []string `json:"details"`
}
*/
type MultiDetailResponse struct {
	ResponseDetail
	Details []CommonResponseDetail `json:"details"`
}

type IResponseDetail interface {
	GetCode() string
}

func (cr *ResponseDetail) GetCode() string {
	return cr.Code
}

/*
Status: 500
{
	"code": "unknown_error",
	"message": "Ocurrió un error desconocido"
}
Status: 401
{
	"code": "auth-0001",
	"message": "Es necesario estar autenticado"
}
Status: 400
{
	"code": "validation-errors",
	"details": ["Es necesario especificar un título de nota", "Es necesario especificar el dueño de la entidad"]
}
OR
Status: 400
{
	"code": "validation-errors",
	"details": [
		{
			"code": "notes-0001",
			message: "Es necesario especificar un título de nota"
		},
		{
			"code": "notes-0002",
			message: "Es necesario especificar el dueño de la entidad"
		}
	]
}
*/
