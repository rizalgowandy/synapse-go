package entity

type GenerateUBODocReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	EntityInfo    struct {
		Cryptocurrency bool `json:"cryptocurrency"`
		Msb            struct {
			Federal bool     `json:"federal"`
			States  []string `json:"states"`
		} `json:"msb"`
		PublicCompany         bool   `json:"public_company"`
		MajorityOwnedByListed bool   `json:"majority_owned_by_listed"`
		RegisteredSEC         bool   `json:"registered_SEC"`
		RegulatedFinancial    bool   `json:"regulated_financial"`
		Gambling              bool   `json:"gambling"`
		DocumentID            string `json:"document_id"`
	} `json:"entity_info"`
	Signer struct {
		DocumentID           string `json:"document_id"`
		RelationshipToEntity string `json:"relationship_to_entity"`
	} `json:"signer"`
	ComplianceContact struct {
		DocumentID           string `json:"document_id"`
		RelationshipToEntity string `json:"relationship_to_entity"`
	} `json:"compliance_contact"`
	PrimaryControllingContact struct {
		DocumentID           string `json:"document_id"`
		RelationshipToEntity string `json:"relationship_to_entity"`
	} `json:"primary_controlling_contact"`
	Owners []struct {
		DocumentID string `json:"document_id"`
		Title      string `json:"title"`
		Ownership  int    `json:"ownership"`
	} `json:"owners"`
	Attested bool `json:"attested"`
}

type GenerateUBODocResp struct {
	ID    string `json:"_id"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	AccountClosureDate any `json:"account_closure_date"`
	Client             struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"client"`
	Documents []struct {
		EntityScope     string `json:"entity_scope"`
		EntityType      string `json:"entity_type"`
		ID              string `json:"id"`
		IDScore         any    `json:"id_score"`
		IsActive        bool   `json:"is_active"`
		Name            string `json:"name"`
		PermissionScope string `json:"permission_scope"`
		PhysicalDocs    []struct {
			DocumentType string `json:"document_type"`
			ID           string `json:"id"`
			LastUpdated  int64  `json:"last_updated"`
			Status       string `json:"status"`
		} `json:"physical_docs"`
		RequiredEddDocs []any `json:"required_edd_docs"`
		SocialDocs      []struct {
			DocumentType string `json:"document_type"`
			ID           string `json:"id"`
			LastUpdated  int64  `json:"last_updated"`
			Status       string `json:"status"`
			Info         struct {
				AddressCareOf      string `json:"address_care_of"`
				AddressCity        string `json:"address_city"`
				AddressCountryCode string `json:"address_country_code"`
				AddressPostalCode  string `json:"address_postal_code"`
				AddressStreet      string `json:"address_street"`
				AddressSubdivision string `json:"address_subdivision"`
			} `json:"info,omitempty"`
		} `json:"social_docs"`
		VirtualDocs []struct {
			DocumentType string `json:"document_type"`
			ID           string `json:"id"`
			LastUpdated  int64  `json:"last_updated"`
			Status       string `json:"status"`
		} `json:"virtual_docs"`
		Watchlists string `json:"watchlists"`
	} `json:"documents"`
	Emails []any `json:"emails"`
	Extra  struct {
		CipTag        int    `json:"cip_tag"`
		DateJoined    int64  `json:"date_joined"`
		ExtraSecurity bool   `json:"extra_security"`
		IsBusiness    bool   `json:"is_business"`
		IsTrusted     bool   `json:"is_trusted"`
		LastUpdated   int64  `json:"last_updated"`
		PublicNote    any    `json:"public_note"`
		SuppID        string `json:"supp_id"`
	} `json:"extra"`
	Flag       string   `json:"flag"`
	FlagCode   any      `json:"flag_code"`
	IsHidden   bool     `json:"is_hidden"`
	LegalNames []string `json:"legal_names"`
	Logins     []struct {
		Email string `json:"email"`
		Scope string `json:"scope"`
	} `json:"logins"`
	Permission     string   `json:"permission"`
	PermissionCode any      `json:"permission_code"`
	PhoneNumbers   []string `json:"phone_numbers"`
	Photos         []any    `json:"photos"`
	RefreshToken   string   `json:"refresh_token"`
	Watchlists     string   `json:"watchlists"`
}

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

type SwapDuplicateUsersReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	SwapToUserID  string `json:"swap_to_user_id"`
}

type SwapDuplicateUsersResp struct {
	UserSwapped string `json:"user_swapped"`
	UserID      string `json:"user_id"`
	SwappedTo   string `json:"swapped_to"`
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
