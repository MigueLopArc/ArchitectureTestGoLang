package responseCodes

var (
	UnknownError = ApiResponse{
		HttpStatusCode: 500,
		Detail: &CommonResponseDetail{
			ResponseDetail: ResponseDetail{Code: "unknown-error"},
			Message:        "Ocurrió un error desconocido",
		},
	}

	AuthorizarionMissing = ApiResponse{
		HttpStatusCode: 401,
		Detail: &CommonResponseDetail{
			ResponseDetail: ResponseDetail{Code: "auth-0001"},
			Message:        "Es necesario estar autenticado",
		},
	}

	AuthorizationFailed = ApiResponse{
		HttpStatusCode: 403,
		Detail: &CommonResponseDetail{
			ResponseDetail: ResponseDetail{Code: "auth-0002"},
			Message:        "Falló la autenticación",
		},
	}

	UserNotFound = ApiResponse{
		HttpStatusCode: 404,
		Detail: &CommonResponseDetail{
			ResponseDetail: ResponseDetail{Code: "auth-0003"},
			Message:        "No se encontro información del usuario",
		},
	}

	WrongPassword = ApiResponse{
		HttpStatusCode: 400,
		Detail: &CommonResponseDetail{
			ResponseDetail: ResponseDetail{Code: "auth-0004"},
			Message:        "La contraseña no es correcta",
		},
	}

	UserAlreadyExists = ApiResponse{
		HttpStatusCode: 400,
		Detail: &CommonResponseDetail{
			ResponseDetail: ResponseDetail{Code: "auth-0005"},
			Message:        "Este correo electrónico ya esta registrado",
		},
	}

	// Here start the codes for Bad Request

	InvalidEmail = &CommonResponseDetail{
		ResponseDetail: ResponseDetail{Code: "auth-0006"},
		Message:        "Correo electrónico inválido",
	}

	InvalidPassword = &CommonResponseDetail{
		ResponseDetail: ResponseDetail{Code: "auth-0007"},
		Message:        "La contraseña no es válida, debe tener al menos 6 caracteres",
	}

	UserFirstNameNotFound = &CommonResponseDetail{
		ResponseDetail: ResponseDetail{Code: "auth-0008"},
		Message:        "Se debe especificar el nombre de usuario",
	}

	UserLastNameNotFound = &CommonResponseDetail{
		ResponseDetail: ResponseDetail{Code: "auth-0009"},
		Message:        "Se debe especificar el apellido de usuario",
	}

	EntityNotFound = ApiResponse{
		HttpStatusCode: 404,
		Detail: &CommonResponseDetail{
			ResponseDetail: ResponseDetail{Code: "entity-not-found"},
			Message:        "No se encontro información del recurso solicitado",
		},
	}

	EntityDoesNotBelongToUser = ApiResponse{
		HttpStatusCode: 403,
		Detail: &CommonResponseDetail{
			ResponseDetail: ResponseDetail{Code: "entity-does-not-belong-to-you"},
			Message:        "No puedes editar ni consultar esta entidad, ya que no te pertenece",
		},
	}

	NoteTitleNotFound = &CommonResponseDetail{
		ResponseDetail: ResponseDetail{Code: "notes-0001"},
		Message:        "Es necesario especificar el titulo de la nota y debe tener al menos 3 caracteres",
	}

	NoteContentNotFound = &CommonResponseDetail{
		ResponseDetail: ResponseDetail{Code: "notes-0002"},
		Message:        "Se debe especificar el contenido de la nota",
	}
)

func BuildBadRequestMessage(errors []CommonResponseDetail) ApiResponse {
	return ApiResponse{
		HttpStatusCode: 400,
		Detail: &MultiDetailResponse{
			ResponseDetail: ResponseDetail{Code: "validation-errors"},
			Details:        errors,
		},
	}
}
