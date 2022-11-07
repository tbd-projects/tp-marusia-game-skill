// Code generated by scg 1, .
//
// LemonadeGame-SceneStructs
//
// Command:
// scg
//.

package scenes

import (
	"github.com/ThCompiler/go_game_constractor/director/scene"
	"github.com/evrone/go-clean-template/pkg/game/scg/lemonadegame/manager"
)

// FlyInfo scene
type FlyInfo struct {
	TextManager manager.TextManager
	NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *FlyInfo) React(_ *scene.Context) scene.Command {
	return scene.NoCommand
}

// Next function returning next scene
func (sc *FlyInfo) Next() scene.Scene {
	return &Shop{
		TextManager: sc.TextManager,
	}
}

// GetSceneInfo function returning info about scene
func (sc *FlyInfo) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {

	// TODO Write some actions for get data for texts

	text, _ := sc.TextManager.GetFlyInfoText()
	return scene.Info{
		Text:             text,
		ExpectedMessages: []scene.MessageMatcher{},
		Buttons:          []scene.Button{},
	}, false
}