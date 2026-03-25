package gemini

import (
	. "github.com/webplode/CLIProxyAPI/v6/internal/constant"
	"github.com/webplode/CLIProxyAPI/v6/internal/interfaces"
	"github.com/webplode/CLIProxyAPI/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		Gemini,
		OpenAI,
		ConvertGeminiRequestToOpenAI,
		interfaces.TranslateResponse{
			Stream:     ConvertOpenAIResponseToGemini,
			NonStream:  ConvertOpenAIResponseToGeminiNonStream,
			TokenCount: GeminiTokenCount,
		},
	)
}
