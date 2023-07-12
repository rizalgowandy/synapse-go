package entity

type ViewAllUserNodesReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	Type          string `json:"-"`
	PerPage       string `json:"-"`
	Page          string `json:"-"`
}

type ViewAllUserNodesResp struct {
	ErrorCode string `json:"error_code"`
	HTTPCode  string `json:"http_code"`
	Limit     int    `json:"limit"`
	NodeCount int    `json:"node_count"`
	Nodes     []Node `json:"nodes"`
	Page      int    `json:"page"`
	PageCount int    `json:"page_count"`
	Success   bool   `json:"success"`
}

type ViewNodeReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	NodeID        string `json:"-"`
	ForceRefresh  string `json:"-"`
	FullDehydrate string `json:"-"`
}

type ViewNodeResp Node

type CreateNodeReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	Type          string `json:"type"`
	Info          struct {
		Nickname   string `json:"nickname"`
		DocumentID string `json:"document_id"`
	} `json:"info"`
}

type CreateNodeResp struct {
	ErrorCode string `json:"error_code"`
	HTTPCode  string `json:"http_code"`
	Limit     int    `json:"limit"`
	NodeCount int    `json:"node_count"`
	Nodes     []Node `json:"nodes"`
	PageCount int    `json:"page_count"`
	Success   bool   `json:"success"`
}

type UpdateNodeReq struct {
	UserID            string `json:"-"`
	UserIPAddress     string `json:"-"`
	UserOAuthKey      string `json:"-"`
	NodeID            string `json:"-"`
	ReAuth            string `json:"-"`
	ResendMicro       string `json:"-"`
	Allowed           string `json:"allowed"`
	AllowedStatusCode string `json:"allowed_status_code"`
}
type UpdateNodeResp struct {
	ErrorCode string `json:"error_code"`
	HTTPCode  string `json:"http_code"`
	Limit     int    `json:"limit"`
	NodeCount int    `json:"node_count"`
	Nodes     []Node `json:"nodes"`
	PageCount int    `json:"page_count"`
	Success   bool   `json:"success"`
}

type GenerateECashBarcodeReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	NodeID        string `json:"-"`
	Amount        struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	} `json:"amount"`
	RetailerID int `json:"retailer_id"`
}

type GenerateECashBarcodeResp struct {
	ID            string `json:"_id"`
	Barcode       string `json:"barcode"`
	BarcodeBase64 string `json:"barcode_base64"`
}

type AllowedNodeTypesResp struct {
	Types []struct {
		Type string `json:"type"`
	} `json:"types"`
}

type ViewATMsResp struct {
	ATMs []struct {
		ATMLocation struct {
			DistanceMeters    string `json:"DistanceMeters"`
			DistanceUnit      string `json:"DistanceUnit"`
			LocationID        string `json:"LocationID"`
			LocationName      string `json:"LocationName"`
			LocationType      string `json:"LocationType"`
			LocationTypeLabel string `json:"LocationTypeLabel"`
			MapIcon           string `json:"MapIcon"`
			MapURL            string `json:"MapUrl"`
			SurchargeFree     string `json:"SurchargeFree"`
			Address           struct {
				City       string `json:"city"`
				Country    string `json:"country"`
				PostalCode string `json:"postalCode"`
				State      string `json:"state"`
				Street     string `json:"street"`
			} `json:"address"`
			Coordinates struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
			} `json:"coordinates"`
		} `json:"atmLocation"`
		ATMNetworkType string  `json:"atm_network_type"`
		Distance       float64 `json:"distance"`
	} `json:"atms"`
	ATMsCount int    `json:"atms_count"`
	ErrorCode string `json:"error_code"`
	HTTPCode  string `json:"http_code"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	PageCount int    `json:"page_count"`
	Success   bool   `json:"success"`
}
