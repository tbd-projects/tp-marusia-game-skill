// Code generated by scg 1, .
//
// BotanicalGardenGame-SceneStructs
//
// Command:
// scg
//.

package scenes

import (
	base_matchers "github.com/ThCompiler/go_game_constractor/director/matchers"
	"github.com/ThCompiler/go_game_constractor/director/scene"
	"github.com/evrone/go-clean-template/pkg/game/scg/botanicalgardengame/manager"
	"github.com/evrone/go-clean-template/pkg/game/scg/botanicalgardengame/script/errors"
)

const (
	// AgreeEndGameButtonText - text for button Agree
	AgreeEndGameButtonText = "Понятно"
)

// EndGame scene
type EndGame struct {
	TextManager manager.TextManager
	NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *EndGame) React(ctx *scene.Context) scene.Command {
	switch {
	// Buttons select
	case ctx.Request.NameMatched == AgreeEndGameButtonText && ctx.Request.WasButton:
		sc.NextScene = ShopScene
		// Matcher select
	case ctx.Request.NameMatched == base_matchers.AgreeMatchedString:
		sc.NextScene = ShopScene
	}

	return scene.NoCommand
}

// Next function returning next scene
func (sc *EndGame) Next() scene.Scene {
	switch sc.NextScene {
	case ShopScene:
		return &Shop{
			TextManager: sc.TextManager,
		}
	}

	return &EndGame{
		TextManager: sc.TextManager,
	}
}

// GetSceneInfo function returning info about scene
func (sc *EndGame) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {
	var (
		score uint64
	)

	text, _ := sc.TextManager.GetEndGameText(
		score,
	)
	return scene.Info{
		Text: text,
		ExpectedMessages: []scene.MessageMatcher{
			base_matchers.Agree,
		},
		Buttons: []scene.Button{
			{
				Title: AgreeEndGameButtonText,
			},
		},
		Err: errors.NotUnderstandEndGameError,
	}, true
}