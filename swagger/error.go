package swagger

type ErrRespType string

const (
	TypeClientReq   ErrRespType = "Client error"
	TypeInternalReq ErrRespType = "Internal"
	Err400          ErrRespType = "Bad Request"
	Err404          ErrRespType = "Not Found"
	Err500          ErrRespType = "Internal Server Error"
)

type Error struct {
	Description string      `json:"description,omitempty"`
	Error       ErrRespType `json:"error,omitempty"`
	Type        ErrRespType `json:"type,omitempty"`
}
