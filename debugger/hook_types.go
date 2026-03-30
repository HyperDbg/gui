package debugger

type HookFilter struct {
	ProcessNames  []string
	ProcessIds    []uint32
	ExcludeSystem bool
}

type HookScriptHandler func(ctx *HookContext)

type HookScript struct {
	ApiName  string
	HookType HookType
	Filter   *HookFilter
	OnMatch  HookScriptHandler
}

type HookContext struct {
	Args   any
	Return int32
}
