package entity

type ErrInvalidData struct {
	Message string
	Err     error
}

func (eid ErrInvalidData) Error() string {
	if eid.Err != nil {
		return eid.Err.Error()
	} else if eid.Message != "" {
		return eid.Message
	} else {
		return "Invalid request"
	}
}

type ErrUnauthorize struct {
	Message string
	Err     error
}

func (eu ErrUnauthorize) Error() string {
	if eu.Err != nil {
		return eu.Err.Error()
	} else if eu.Message != "" {
		return eu.Message
	} else {
		return "Unauthorize"
	}
}

type ErrForbidden struct {
	Message string
	Err     error
}

func (ef ErrForbidden) Error() string {
	if ef.Err != nil {
		return ef.Err.Error()
	} else if ef.Message != "" {
		return ef.Message
	} else {
		return "Forbidden"
	}
}

type ErrNotFound struct {
	Message string
	Err     error
}

func (ent ErrNotFound) Error() string {
	if ent.Err != nil {
		return ent.Err.Error()
	} else if ent.Message != "" {
		return ent.Message
	} else {
		return "Not Found"
	}
}

type ErrValidation struct {
	Message string
	Err     error
}

func (evl ErrValidation) Error() string {
	if evl.Err != nil {
		return evl.Err.Error()
	} else if evl.Message != "" {
		return evl.Message
	} else {
		return "Not Found"
	}
}
