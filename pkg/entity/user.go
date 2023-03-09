package entity

type ViewUserReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	FullDehydrate string `json:"-"`
}

type ViewUserResp struct {
	ID    string `json:"_id"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	V                  int `json:"_v"`
	AccountClosureDate any `json:"account_closure_date"`
	Client             struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"client"`
	Documents []struct {
		AddressCity        string  `json:"address_city"`
		AddressCountryCode string  `json:"address_country_code"`
		AddressPostalCode  string  `json:"address_postal_code"`
		AddressStreet      string  `json:"address_street"`
		AddressSubdivision string  `json:"address_subdivision"`
		Alias              string  `json:"alias"`
		Day                int     `json:"day"`
		DesiredScope       any     `json:"desired_scope"`
		DocOptionKey       any     `json:"doc_option_key"`
		DocsKey            any     `json:"docs_key"`
		DocsTitle          any     `json:"docs_title"`
		Email              string  `json:"email"`
		EntityScope        string  `json:"entity_scope"`
		EntityType         string  `json:"entity_type"`
		ID                 string  `json:"id"`
		IDScore            float64 `json:"id_score"`
		IP                 string  `json:"ip"`
		IsActive           bool    `json:"is_active"`
		Month              int     `json:"month"`
		Name               string  `json:"name"`
		PermissionScope    string  `json:"permission_scope"`
		PhoneNumber        string  `json:"phone_number"`
		PhysicalDocs       []struct {
			DocumentType  string `json:"document_type"`
			DocumentValue string `json:"document_value"`
			ID            string `json:"id"`
			LastUpdated   int64  `json:"last_updated"`
			Status        string `json:"status"`
		} `json:"physical_docs"`
		RequiredEddDocs  []any `json:"required_edd_docs"`
		ScreeningResults struct {
			Num561                              string `json:"561"`
			Aucl                                string `json:"aucl"`
			ConcernLocation                     string `json:"concern_location"`
			Dpl                                 string `json:"dpl"`
			Dtc                                 string `json:"dtc"`
			El                                  string `json:"el"`
			Eucl                                string `json:"eucl"`
			FatfNonCooperativeJurisdiction      string `json:"fatf_non_cooperative_jurisdiction"`
			FbiBankRobbers                      string `json:"fbi_bank_robbers"`
			FbiCounterIntelligence              string `json:"fbi_counter_intelligence"`
			FbiCrimesAgainstChildren            string `json:"fbi_crimes_against_children"`
			FbiCriminalEnterpriseInvestigations string `json:"fbi_criminal_enterprise_investigations"`
			FbiCyber                            string `json:"fbi_cyber"`
			FbiDomesticTerrorism                string `json:"fbi_domestic_terrorism"`
			FbiHumanTrafficking                 string `json:"fbi_human_trafficking"`
			FbiMurders                          string `json:"fbi_murders"`
			FbiViolentCrimes                    string `json:"fbi_violent_crimes"`
			FbiWantedTerrorists                 string `json:"fbi_wanted_terrorists"`
			FbiWhiteCollar                      string `json:"fbi_white_collar"`
			FincenRedList                       string `json:"fincen_red_list"`
			Fse                                 string `json:"fse"`
			FtoSanctions                        string `json:"fto_sanctions"`
			FuturesSanctions                    string `json:"futures_sanctions"`
			HkmaSanctions                       string `json:"hkma_sanctions"`
			HmTreasurySanctions                 string `json:"hm_treasury_sanctions"`
			Isn                                 string `json:"isn"`
			MasSanctions                        string `json:"mas_sanctions"`
			MonitoredLocation                   string `json:"monitored_location"`
			NsIsa                               string `json:"ns-isa"`
			Ofac561List                         string `json:"ofac_561_list"`
			OfacEo13645                         string `json:"ofac_eo13645"`
			OfacFse                             string `json:"ofac_fse"`
			OfacFseIr                           string `json:"ofac_fse_ir"`
			OfacFseSy                           string `json:"ofac_fse_sy"`
			OfacIsa                             string `json:"ofac_isa"`
			OfacNsIsa                           string `json:"ofac_ns_isa"`
			OfacPlc                             string `json:"ofac_plc"`
			OfacSdn                             string `json:"ofac_sdn"`
			OfacSsi                             string `json:"ofac_ssi"`
			OfacSyria                           string `json:"ofac_syria"`
			OfacUkraineEo13662                  string `json:"ofac_ukraine_eo13662"`
			Osfi                                string `json:"osfi"`
			Pep                                 string `json:"pep"`
			Plc                                 string `json:"plc"`
			PrimaryConcern                      string `json:"primary_concern"`
			Sdn                                 string `json:"sdn"`
			Ssi                                 string `json:"ssi"`
			TelSanctions                        string `json:"tel_sanctions"`
			Ukcl                                string `json:"ukcl"`
			Uvl                                 string `json:"uvl"`
		} `json:"screening_results"`
		SocialDocs []struct {
			DocumentType  string `json:"document_type"`
			DocumentValue string `json:"document_value"`
			ID            string `json:"id"`
			LastUpdated   int64  `json:"last_updated"`
			Status        string `json:"status"`
		} `json:"social_docs"`
		VirtualDocs []struct {
			DocumentType  string `json:"document_type"`
			DocumentValue string `json:"document_value"`
			ID            string `json:"id"`
			LastUpdated   int64  `json:"last_updated"`
			Status        string `json:"status"`
		} `json:"virtual_docs"`
		Watchlists string `json:"watchlists"`
		Year       int    `json:"year"`
	} `json:"documents"`
	Emails []any `json:"emails"`
	Extra  struct {
		CipTag        int    `json:"cip_tag"`
		DateJoined    int64  `json:"date_joined"`
		ExtraSecurity bool   `json:"extra_security"`
		IsBusiness    bool   `json:"is_business"`
		IsTrusted     bool   `json:"is_trusted"`
		LastUpdated   int64  `json:"last_updated"`
		Note          any    `json:"note"`
		PublicNote    any    `json:"public_note"`
		SuppID        string `json:"supp_id"`
	} `json:"extra"`
	Flag       string   `json:"flag"`
	FlagCode   any      `json:"flag_code"`
	Ips        []string `json:"ips"`
	IsHidden   bool     `json:"is_hidden"`
	LegalNames []string `json:"legal_names"`
	Logins     []struct {
		Email        string `json:"email"`
		OldPasswords []any  `json:"old_passwords"`
		Scope        string `json:"scope"`
		UpdatedOn    struct {
			Date int64 `json:"$date"`
		} `json:"updated_on"`
	} `json:"logins"`
	Permission     string   `json:"permission"`
	PermissionCode any      `json:"permission_code"`
	PhoneNumbers   []string `json:"phone_numbers"`
	Photos         []any    `json:"photos"`
	RefreshToken   string   `json:"refresh_token"`
	Watchlists     string   `json:"watchlists"`
}

type CreateUserReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	Logins        []struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"logins"`
	PhoneNumbers []string `json:"phone_numbers"`
	LegalNames   []string `json:"legal_names"`
	Documents    []struct {
		Email              string `json:"email"`
		PhoneNumber        string `json:"phone_number"`
		IP                 string `json:"ip"`
		Name               string `json:"name"`
		Alias              string `json:"alias"`
		EntityType         string `json:"entity_type"`
		EntityScope        string `json:"entity_scope"`
		Day                int    `json:"day"`
		Month              int    `json:"month"`
		Year               int    `json:"year"`
		AddressStreet      string `json:"address_street"`
		AddressCity        string `json:"address_city"`
		AddressSubdivision string `json:"address_subdivision"`
		AddressPostalCode  string `json:"address_postal_code"`
		AddressCountryCode string `json:"address_country_code"`
		VirtualDocs        []struct {
			DocumentValue string `json:"document_value"`
			DocumentType  string `json:"document_type"`
			Meta          struct {
				CountryCode string `json:"country_code"`
			} `json:"meta"`
		} `json:"virtual_docs"`
		PhysicalDocs []struct {
			DocumentValue string `json:"document_value"`
			DocumentType  string `json:"document_type"`
			Meta          struct {
				CountryCode string `json:"country_code"`
				StateCode   string `json:"state_code"`
			} `json:"meta"`
		} `json:"physical_docs"`
		SocialDocs []struct {
			DocumentValue string `json:"document_value"`
			DocumentType  string `json:"document_type"`
			Meta          struct {
				AddressStreet      string `json:"address_street"`
				AddressCity        string `json:"address_city"`
				AddressSubdivision string `json:"address_subdivision"`
				AddressPostalCode  string `json:"address_postal_code"`
				AddressCountryCode string `json:"address_country_code"`
				AddressCareOf      string `json:"address_care_of"`
			} `json:"meta"`
		} `json:"social_docs"`
	} `json:"documents"`
	Extra struct {
		SuppID     string `json:"supp_id"`
		CipTag     int    `json:"cip_tag"`
		IsBusiness bool   `json:"is_business"`
	} `json:"extra"`
}

type CreateUserResp struct {
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

type UpdateUserReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	UserOAuthKey  string `json:"-"`
	Documents     []struct {
		Email              string `json:"email"`
		PhoneNumber        string `json:"phone_number"`
		IP                 string `json:"ip"`
		Name               string `json:"name"`
		Alias              string `json:"alias"`
		EntityType         string `json:"entity_type"`
		EntityScope        string `json:"entity_scope"`
		Day                int    `json:"day"`
		Month              int    `json:"month"`
		Year               int    `json:"year"`
		AddressStreet      string `json:"address_street"`
		AddressCity        string `json:"address_city"`
		AddressSubdivision string `json:"address_subdivision"`
		AddressPostalCode  string `json:"address_postal_code"`
		AddressCountryCode string `json:"address_country_code"`
		VirtualDocs        []struct {
			DocumentValue string `json:"document_value"`
			DocumentType  string `json:"document_type"`
			Meta          struct {
				CountryCode string `json:"country_code"`
			} `json:"meta"`
		} `json:"virtual_docs"`
		PhysicalDocs []struct {
			DocumentValue string `json:"document_value"`
			DocumentType  string `json:"document_type"`
			Meta          struct {
				CountryCode string `json:"country_code"`
				StateCode   string `json:"state_code"`
			} `json:"meta"`
		} `json:"physical_docs"`
		SocialDocs []struct {
			DocumentValue string `json:"document_value"`
			DocumentType  string `json:"document_type"`
			Meta          struct {
				AddressStreet      string `json:"address_street"`
				AddressCity        string `json:"address_city"`
				AddressSubdivision string `json:"address_subdivision"`
				AddressPostalCode  string `json:"address_postal_code"`
				AddressCountryCode string `json:"address_country_code"`
				AddressCareOf      string `json:"address_care_of"`
			} `json:"meta"`
		} `json:"social_docs"`
	} `json:"documents"`
}

type UpdateUserResp struct {
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
