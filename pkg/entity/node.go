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
	Nodes     []struct {
		ID    string `json:"_id"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
		Allowed           string `json:"allowed"`
		AllowedStatusCode any    `json:"allowed_status_code"`
		Client            struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"client"`
		Extra struct {
			Note  any `json:"note"`
			Other struct {
				AccessToken string `json:"access_token"`
				MicroMeta   struct {
					MicroAttempts  int  `json:"micro_attempts"`
					MicroSentCount int  `json:"micro_sent_count"`
					SkipMicro      bool `json:"skip_micro"`
				} `json:"micro_meta"`
				UpdatedOn int64 `json:"updated_on"`
			} `json:"other"`
			SuppID string `json:"supp_id"`
		} `json:"extra"`
		Info struct {
			AccountNum string `json:"account_num"`
			Address    string `json:"address"`
			Balance    struct {
				Amount    string `json:"amount"`
				Currency  string `json:"currency"`
				UpdatedOn int64  `json:"updated_on"`
			} `json:"balance"`
			BankCode     string `json:"bank_code"`
			BankHlogo    string `json:"bank_hlogo"`
			BankLogo     string `json:"bank_logo"`
			BankLongName string `json:"bank_long_name"`
			BankName     string `json:"bank_name"`
			BankURL      string `json:"bank_url"`
			Class        string `json:"class"`
			MatchInfo    struct {
				EmailMatch       string `json:"email_match"`
				NameMatch        string `json:"name_match"`
				PhonenumberMatch string `json:"phonenumber_match"`
			} `json:"match_info"`
			NameOnAccount string `json:"name_on_account"`
			Nickname      string `json:"nickname"`
			RoutingNum    string `json:"routing_num"`
			Type          string `json:"type"`
		} `json:"info"`
		IsActive bool `json:"is_active"`
		Timeline []struct {
			Date int64  `json:"date"`
			Note string `json:"note"`
		} `json:"timeline"`
		Type   string `json:"type"`
		UserID string `json:"user_id"`
	} `json:"nodes"`
	Page      int  `json:"page"`
	PageCount int  `json:"page_count"`
	Success   bool `json:"success"`
}

type ViewNodeReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	NodeID        string `json:"-"`
	ForceRefresh  string `json:"-"`
	FullDehydrate string `json:"-"`
}

type ViewNodeResp struct {
	ID    string `json:"_id"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Allowed           string `json:"allowed"`
	AllowedStatusCode string `json:"allowed_status_code"`
	Client            struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"client"`
	Extra struct {
		Note  any `json:"note"`
		Other struct {
			AccessToken any `json:"access_token"`
			Info        struct {
			} `json:"info"`
			MicroMeta struct {
				MicroAttempts  int  `json:"micro_attempts"`
				MicroSentCount int  `json:"micro_sent_count"`
				SkipMicro      bool `json:"skip_micro"`
			} `json:"micro_meta"`
			Transactions []any `json:"transactions"`
			UpdatedOn    any   `json:"updated_on"`
		} `json:"other"`
		SuppID string `json:"supp_id"`
	} `json:"extra"`
	Info struct {
		AccountNum string `json:"account_num"`
		Address    string `json:"address"`
		Balance    struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"balance"`
		BankCode     string `json:"bank_code"`
		BankHLogo    string `json:"bank_hlogo"`
		BankLogo     string `json:"bank_logo"`
		BankLongName string `json:"bank_long_name"`
		BankName     string `json:"bank_name"`
		BankURL      string `json:"bank_url"`
		Class        string `json:"class"`
		MatchInfo    struct {
			EmailMatch       string `json:"email_match"`
			NameMatch        string `json:"name_match"`
			PhoneNumberMatch string `json:"phonenumber_match"`
		} `json:"match_info"`
		NameOnAccount string `json:"name_on_account"`
		Nickname      string `json:"nickname"`
		RoutingNum    string `json:"routing_num"`
		Type          string `json:"type"`
	} `json:"info"`
	IsActive bool `json:"is_active"`
	Timeline []struct {
		Date int64  `json:"date"`
		Note string `json:"note"`
	} `json:"timeline"`
	Type   string `json:"type"`
	UserID string `json:"user_id"`
}

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
	Nodes     []struct {
		ID    string `json:"_id"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
		Allowed           string `json:"allowed"`
		AllowedStatusCode any    `json:"allowed_status_code"`
		Client            struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"client"`
		Extra struct {
			Note  any `json:"note"`
			Other struct {
			} `json:"other"`
			SuppID string `json:"supp_id"`
		} `json:"extra"`
		Info struct {
			Agreements []struct {
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"agreements"`
			Balance struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"balance"`
			BankCode      string `json:"bank_code"`
			DocumentID    string `json:"document_id"`
			NameOnAccount string `json:"name_on_account"`
			Nickname      string `json:"nickname"`
		} `json:"info"`
		IsActive bool `json:"is_active"`
		Timeline []struct {
			Date int64  `json:"date"`
			Note string `json:"note"`
		} `json:"timeline"`
		Type   string `json:"type"`
		UserID string `json:"user_id"`
	} `json:"nodes"`
	PageCount int  `json:"page_count"`
	Success   bool `json:"success"`
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
	Nodes     []struct {
		ID    string `json:"_id"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
		Allowed           string `json:"allowed"`
		AllowedStatusCode any    `json:"allowed_status_code"`
		Client            struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"client"`
		Extra struct {
			Note  any `json:"note"`
			Other struct {
			} `json:"other"`
			SuppID string `json:"supp_id"`
		} `json:"extra"`
		Info struct {
			Agreements []struct {
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"agreements"`
			Balance struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"balance"`
			BankCode      string `json:"bank_code"`
			DocumentID    string `json:"document_id"`
			NameOnAccount string `json:"name_on_account"`
			Nickname      string `json:"nickname"`
		} `json:"info"`
		IsActive bool `json:"is_active"`
		Timeline []struct {
			Date int64  `json:"date"`
			Note string `json:"note"`
		} `json:"timeline"`
		Type   string `json:"type"`
		UserID string `json:"user_id"`
	} `json:"nodes"`
	PageCount int  `json:"page_count"`
	Success   bool `json:"success"`
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

type ViewCryptoQuotesResp struct {
	Aaveusd  float64 `json:"AAVEUSD"`
	Algousd  float64 `json:"ALGOUSD"`
	Ausdcusd float64 `json:"AUSDCUSD"`
	Ausdtusd float64 `json:"AUSDTUSD"`
	Avaxusd  float64 `json:"AVAXUSD"`
	Batusd   float64 `json:"BATUSD"`
	Btcusd   float64 `json:"BTCUSD"`
	Busdusd  float64 `json:"BUSDUSD"`
	Compusd  float64 `json:"COMPUSD"`
	Crvusd   float64 `json:"CRVUSD"`
	Daiusd   float64 `json:"DAIUSD"`
	Ethusd   float64 `json:"ETHUSD"`
	Fusdusd  float64 `json:"FUSDUSD"`
	Gusdusd  float64 `json:"GUSDUSD"`
	Lbtcusd  bool    `json:"LBTCUSD"`
	Lethusd  float64 `json:"LETHUSD"`
	Linkusd  float64 `json:"LINKUSD"`
	Lusdcusd float64 `json:"LUSDCUSD"`
	Mkrusd   float64 `json:"MKRUSD"`
	Paxusd   float64 `json:"PAXUSD"`
	Raiusd   float64 `json:"RAIUSD"`
	Snxusd   float64 `json:"SNXUSD"`
	Umausd   float64 `json:"UMAUSD"`
	Uniusd   float64 `json:"UNIUSD"`
	Usdaave  float64 `json:"USDAAVE"`
	Usdalgo  float64 `json:"USDALGO"`
	Usdausdc float64 `json:"USDAUSDC"`
	Usdausdt float64 `json:"USDAUSDT"`
	Usdavax  float64 `json:"USDAVAX"`
	Usdbat   float64 `json:"USDBAT"`
	Usdbtc   float64 `json:"USDBTC"`
	Usdbusd  int     `json:"USDBUSD"`
	Usdcomp  float64 `json:"USDCOMP"`
	Usdcrv   float64 `json:"USDCRV"`
	Usdcusd  float64 `json:"USDCUSD"`
	Usddai   float64 `json:"USDDAI"`
	Usdeth   float64 `json:"USDETH"`
	Usdfusd  int     `json:"USDFUSD"`
	Usdgusd  int     `json:"USDGUSD"`
	Usdlbtc  bool    `json:"USDLBTC"`
	Usdleth  float64 `json:"USDLETH"`
	Usdlink  float64 `json:"USDLINK"`
	Usdlusdc int     `json:"USDLUSDC"`
	Usdmkr   float64 `json:"USDMKR"`
	Usdpax   int     `json:"USDPAX"`
	Usdrai   float64 `json:"USDRAI"`
	Usdsnx   float64 `json:"USDSNX"`
	Usdsusd  float64 `json:"USDSUSD"`
	Usdtusd  float64 `json:"USDTUSD"`
	Usduma   float64 `json:"USDUMA"`
	Usduni   float64 `json:"USDUNI"`
	Usdusdc  int     `json:"USDUSDC"`
	Usdusds  int     `json:"USDUSDS"`
	Usdusdt  float64 `json:"USDUSDT"`
	Usdwbtc  float64 `json:"USDWBTC"`
	Usdxlm   float64 `json:"USDXLM"`
	Usdyfi   float64 `json:"USDYFI"`
	Usdzusd  float64 `json:"USDZUSD"`
	USDmUSDC bool    `json:"USDmUSDC"`
	USDpDAI  bool    `json:"USDpDAI"`
	USDsUSDC bool    `json:"USDsUSDC"`
	Wbtcusd  float64 `json:"WBTCUSD"`
	Xlmusd   float64 `json:"XLMUSD"`
	Yfiusd   float64 `json:"YFIUSD"`
	Zusdusd  float64 `json:"ZUSDUSD"`
	MUSDCUSD bool    `json:"mUSDCUSD"`
	PDAIUSD  bool    `json:"pDAIUSD"`
	SUSDCUSD bool    `json:"sUSDCUSD"`
}
