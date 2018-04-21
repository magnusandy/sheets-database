package auth

type AuthComplete struct {
	Status string
	Body   string
}

func AuthFailure(e error) AuthComplete {
	return AuthComplete{
		"Failed",
		e.Error(),
	}
}

func AuthSuccess() AuthComplete {
	return AuthComplete{"Succeeded",
		"You are now authenticated to make reads and writes",
	}
}
