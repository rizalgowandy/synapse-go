package enum

type OAuthScope string

var (
	ScopeUserPatch         OAuthScope = "USER|PATCH"
	ScopeUserGet           OAuthScope = "USER|GET"
	ScopeNodesPost         OAuthScope = "NODES|POST"
	ScopeNodesGet          OAuthScope = "NODES|GET"
	ScopeNodeGet           OAuthScope = "NODE|GET"
	ScopeNodePatch         OAuthScope = "NODE|PATCH"
	ScopeNodeDelete        OAuthScope = "NODE|DELETE"
	ScopeTransPost         OAuthScope = "TRANS|POST"
	ScopeTransGet          OAuthScope = "TRANS|GET"
	ScopeTranGet           OAuthScope = "TRAN|GET"
	ScopeTranPatch         OAuthScope = "TRAN|PATCH"
	ScopeTranDelete        OAuthScope = "TRAN|DELETE"
	ScopeSubnetsPost       OAuthScope = "SUBNETS|POST"
	ScopeSubnetsGet        OAuthScope = "SUBNETS|GET"
	ScopeSubnetGet         OAuthScope = "SUBNET|GET"
	ScopeSubnetPatch       OAuthScope = "SUBNET|PATCH"
	ScopeStatementsGet     OAuthScope = "STATEMENTS|GET"
	ScopeStatementGet      OAuthScope = "STATEMENT|GET"
	ScopeStatementsPost    OAuthScope = "STATEMENTS|POST"
	ScopeConversationsPost OAuthScope = "CONVERSATIONS|POST"
	ScopeConversationsGet  OAuthScope = "CONVERSATIONS|GET"
	ScopeConversationGet   OAuthScope = "CONVERSATION|GET"
	ScopeConversationPatch OAuthScope = "CONVERSATION|PATCH"
	ScopeMessagesPost      OAuthScope = "MESSAGES|POST"
	ScopeMessagesGet       OAuthScope = "MESSAGES|GET"
)
