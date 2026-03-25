package chat_completions

import (
	. "github.com/webplode/CLIProxyAPI/v6/internal/constant"
	"github.com/webplode/CLIProxyAPI/v6/internal/interfaces"
	"github.com/webplode/CLIProxyAPI/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		OpenAI,
		Claude,
		ConvertOpenAIRequestToClaude,
		interfaces.TranslateResponse{
			Stream:    ConvertClaudeResponseToOpenAI,
			NonStream: ConvertClaudeResponseToOpenAINonStream,
		},
	)
}
