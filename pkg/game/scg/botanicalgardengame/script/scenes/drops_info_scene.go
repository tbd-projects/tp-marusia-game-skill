// Code generated by scg 1, .
//
// BotanicalGardenGame-SceneStructs
//
// Command:
// scg
//.

package scenes

import (
	"github.com/ThCompiler/go_game_constractor/director/scene"
	"github.com/evrone/go-clean-template/pkg/game/scg/botanicalgardengame/manager"
)

// DropsInfo scene
type DropsInfo struct {
	TextManager manager.TextManager
	NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *DropsInfo) React(_ *scene.Context) scene.Command {
	return scene.NoCommand
}

// Next function returning next scene
func (sc *DropsInfo) Next() scene.Scene {
	return &Shop{
		TextManager: sc.TextManager,
	}
}

// GetSceneInfo function returning info about scene
func (sc *DropsInfo) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {
	var (
		dropsCount uint64
	)

	text, _ := sc.TextManager.GetDropsInfoText(
		dropsCount,
	)
	return scene.Info{
		Text:             text,
		ExpectedMessages: []scene.MessageMatcher{},
		Buttons:          []scene.Button{},
	}, false
}