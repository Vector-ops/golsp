package lsp

type InitializeRequest struct {
	Request `json:"request,omitempty"`
	Params  InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewIntializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{},
			ServerInfo: ServerInfo{
				Name:    "golsp",
				Version: "0.0.0.0.0.0-beta1.final",
			},
		},
	}
}