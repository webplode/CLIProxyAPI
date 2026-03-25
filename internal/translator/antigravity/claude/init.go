package claude

import (
	. "github.com/webplode/CLIProxyAPI/v6/internal/constant"
	"github.com/webplode/CLIProxyAPI/v6/internal/interfaces"
	"github.com/webplode/CLIProxyAPI/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		Claude,
		Antigravity,
		ConvertClaudeRequestToAntigravity,
		interfaces.TranslateResponse{
			Stream:     ConvertAntigravityResponseToClaude,
			NonStream:  ConvertAntigravityResponseToClaudeNonStream,
			TokenCount: ClaudeTokenCount,
		},
	)
}
