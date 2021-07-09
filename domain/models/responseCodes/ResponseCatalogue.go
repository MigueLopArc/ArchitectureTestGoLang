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
			ResponseDetail: ResponseDetail{Code: "auth-0006"},
			Message:        "No se encontro información del usuario",
		},
	}

	WrongPassword = ApiResponse{
		HttpStatusCode: 400,
		Detail: &CommonResponseDetail{
			ResponseDetail: ResponseDetail{Code: "auth-0007"},
			Message:        "La contraseña no es correcta",
		},
	}

	UserAlreadyExists = ApiResponse{
		HttpStatusCode: 400,
		Detail: &CommonResponseDetail{
			ResponseDetail: ResponseDetail{Code: "auth-0010"},
			Message:        "Este correo electrónico ya esta registrado",
		},
	}

	EntityNotFound = ApiResponse{
		HttpStatusCode: 404,
		Detail: &CommonResponseDetail{
			ResponseDetail: ResponseDetail{Code: "entity-not-found"},
			Message:        "No se encontro información del recurso solicitado",
		},
	}

	// Here start the codes for Bad Request

	InvalidEmail = &CommonResponseDetail{
		ResponseDetail: ResponseDetail{Code: "auth-0004"},
		Message:        "Correo electrónico inválido",
	}

	InvalidPassword = &CommonResponseDetail{
		ResponseDetail: ResponseDetail{Code: "auth-0005"},
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

	NoteTitleNotFound = &CommonResponseDetail{
		ResponseDetail: ResponseDetail{Code: "notes-0001"},
		Message:        "Es necesario especificar el titulo de la nota y debe tener al menos 3 caracteres",
	}

	NoteContentNotFound = &CommonResponseDetail{
		ResponseDetail: ResponseDetail{Code: "notes-0002"},
		Message:        "Se debe especificar el contenido de la nota",
	}
	/*
		Detail1 = &CommonResponseDetail{
			ResponseDetail: ResponseDetail{Code: "notes-0001"},
			Message:        "Es necesario el id de nota",
		}

		Detail2 = &CommonResponseDetail{
			ResponseDetail: ResponseDetail{Code: "notes-0002"},
			Message:        "Es necesario el titulo de nota",
		}

		Test1BadRequest = ApiResponse{
			HttpStatusCode: 400,
			Detail: &BadRequestDetail{
				ResponseDetail: ResponseDetail{Code: "validation-errors"},
				Details:        []string{"Ocurrió un error desconocido", "Otro error", "mas errores"},
			},
		}

		Test2BadRequest = ApiResponse{
			HttpStatusCode: 400,
			Detail: &MultiDetailResponse{
				ResponseDetail: ResponseDetail{Code: "validation-errors"},
				Details:        []CommonResponseDetail{*Detail1, *Detail2},
			},
		}
	*/
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
