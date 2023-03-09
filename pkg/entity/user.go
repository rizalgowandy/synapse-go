package entity

type GetDuplicatesReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
}

type GetDuplicatesResp struct {
	ClosedUsersID []string `json:"closed_users_id"`
	OpenUsersID   []string `json:"open_users_id"`
	Reason        string   `json:"reason"`
	Status        string   `json:"status"`
	UserID        string   `json:"user_id"`
}

type AllowedDocumentTypesResp struct {
	PhysicalDocs []struct {
		AllowedFileTypes []string `json:"allowed_file_types"`
		BaseDocTypes     []string `json:"base_doc_types"`
		CommonName       string   `json:"common_name"`
		Description      string   `json:"description"`
		MultipleAllowed  bool     `json:"multiple_allowed"`
		RequiredMeta     []any    `json:"required_meta"`
		Type             string   `json:"type"`
	} `json:"physical_docs"`
	SocialDocs []struct {
		BaseDocTypes    []string `json:"base_doc_types"`
		CommonName      string   `json:"common_name"`
		Description     string   `json:"description"`
		MultipleAllowed bool     `json:"multiple_allowed"`
		RequiredMeta    []any    `json:"required_meta"`
		Type            string   `json:"type"`
	} `json:"social_docs"`
	VirtualDocs []struct {
		BaseDocTypes    []string `json:"base_doc_types"`
		CommonName      string   `json:"common_name"`
		Description     string   `json:"description"`
		MultipleAllowed bool     `json:"multiple_allowed"`
		RequiredMeta    []any    `json:"required_meta"`
		Type            string   `json:"type"`
	} `json:"virtual_docs"`
}

type AllowedEntityTypesResp struct {
	Business []struct {
		CommonName string `json:"common_name"`
		Type       string `json:"type"`
	} `json:"BUSINESS"`
	Personal []struct {
		CommonName string `json:"common_name"`
		Type       string `json:"type"`
	} `json:"PERSONAL"`
}

type AllowedEntityScopesResp struct {
	Scopes []string `json:"scopes"`
}
