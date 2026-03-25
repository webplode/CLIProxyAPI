package responses

import (
	. "github.com/webplode/CLIProxyAPI/v6/internal/constant"
	"github.com/webplode/CLIProxyAPI/v6/internal/interfaces"
	"github.com/webplode/CLIProxyAPI/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		OpenaiResponse,
		Codex,
		ConvertOpenAIResponsesRequestToCodex,
		interfaces.TranslateResponse{
			Stream:    ConvertCodexResponseToOpenAIResponses,
			NonStream: ConvertCodexResponseToOpenAIResponsesNonStream,
		},
	)
}
