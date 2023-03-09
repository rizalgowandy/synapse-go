package enum

type OAuthScope string

var (
	OAuthScopeUserPatch         OAuthScope = "USER|PATCH"
	OAuthScopeUserGet           OAuthScope = "USER|GET"
	OAuthScopeNodesPost         OAuthScope = "NODES|POST"
	OAuthScopeNodesGet          OAuthScope = "NODES|GET"
	OAuthScopeNodeGet           OAuthScope = "NODE|GET"
	OAuthScopeNodePatch         OAuthScope = "NODE|PATCH"
	OAuthScopeNodeDelete        OAuthScope = "NODE|DELETE"
	OAuthScopeTransPost         OAuthScope = "TRANS|POST"
	OAuthScopeTransGet          OAuthScope = "TRANS|GET"
	OAuthScopeTranGet           OAuthScope = "TRAN|GET"
	OAuthScopeTranPatch         OAuthScope = "TRAN|PATCH"
	OAuthScopeTranDelete        OAuthScope = "TRAN|DELETE"
	OAuthScopeSubnetsPost       OAuthScope = "SUBNETS|POST"
	OAuthScopeSubnetsGet        OAuthScope = "SUBNETS|GET"
	OAuthScopeSubnetGet         OAuthScope = "SUBNET|GET"
	OAuthScopeSubnetPatch       OAuthScope = "SUBNET|PATCH"
	OAuthScopeStatementsGet     OAuthScope = "STATEMENTS|GET"
	OAuthScopeStatementGet      OAuthScope = "STATEMENT|GET"
	OAuthScopeStatementsPost    OAuthScope = "STATEMENTS|POST"
	OAuthScopeConversationsPost OAuthScope = "CONVERSATIONS|POST"
	OAuthScopeConversationsGet  OAuthScope = "CONVERSATIONS|GET"
	OAuthScopeConversationGet   OAuthScope = "CONVERSATION|GET"
	OAuthScopeConversationPatch OAuthScope = "CONVERSATION|PATCH"
	OAuthScopeMessagesPost      OAuthScope = "MESSAGES|POST"
	OAuthScopeMessagesGet       OAuthScope = "MESSAGES|GET"
)
