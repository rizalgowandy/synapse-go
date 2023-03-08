package enum

type Scope string

var (
	ScopeUserPatch         Scope = "USER|PATCH"
	ScopeUserGet           Scope = "USER|GET"
	ScopeNodesPost         Scope = "NODES|POST"
	ScopeNodesGet          Scope = "NODES|GET"
	ScopeNodeGet           Scope = "NODE|GET"
	ScopeNodePatch         Scope = "NODE|PATCH"
	ScopeNodeDelete        Scope = "NODE|DELETE"
	ScopeTransPost         Scope = "TRANS|POST"
	ScopeTransGet          Scope = "TRANS|GET"
	ScopeTranGet           Scope = "TRAN|GET"
	ScopeTranPatch         Scope = "TRAN|PATCH"
	ScopeTranDelete        Scope = "TRAN|DELETE"
	ScopeSubnetsPost       Scope = "SUBNETS|POST"
	ScopeSubnetsGet        Scope = "SUBNETS|GET"
	ScopeSubnetGet         Scope = "SUBNET|GET"
	ScopeSubnetPatch       Scope = "SUBNET|PATCH"
	ScopeStatementsGet     Scope = "STATEMENTS|GET"
	ScopeStatementGet      Scope = "STATEMENT|GET"
	ScopeStatementsPost    Scope = "STATEMENTS|POST"
	ScopeConversationsPost Scope = "CONVERSATIONS|POST"
	ScopeConversationsGet  Scope = "CONVERSATIONS|GET"
	ScopeConversationGet   Scope = "CONVERSATION|GET"
	ScopeConversationPatch Scope = "CONVERSATION|PATCH"
	ScopeMessagesPost      Scope = "MESSAGES|POST"
	ScopeMessagesGet       Scope = "MESSAGES|GET"
)
