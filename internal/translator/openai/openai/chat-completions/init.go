package chat_completions

import (
	. "github.com/webplode/CLIProxyAPI/v6/internal/constant"
	"github.com/webplode/CLIProxyAPI/v6/internal/interfaces"
	"github.com/webplode/CLIProxyAPI/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		OpenAI,
		OpenAI,
		ConvertOpenAIRequestToOpenAI,
		interfaces.TranslateResponse{
			Stream:    ConvertOpenAIResponseToOpenAI,
			NonStream: ConvertOpenAIResponseToOpenAINonStream,
		},
	)
}
