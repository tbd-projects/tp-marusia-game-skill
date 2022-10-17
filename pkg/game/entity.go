package game

import (
	"context"
	"encoding/json"
	"github.com/evrone/go-clean-template/pkg/game/scene"
)

type Result struct {
	Text          scene.Text
	Buttons       []scene.Button
	IsEndOfScript bool
}

type SceneDirectorConfig struct {
	StartScene     scene.Scene
	ErrorScene     scene.Scene
	GoodbyeScene   scene.Scene
	EndCommand     string
	GoodbyeMessage scene.Text
}

type UserInfo struct {
	UserId    string
	SessionId string
}

type SceneRequest struct {
	Command      string
	FullUserText string
	WasButton    bool
	Payload      json.RawMessage
	Info         UserInfo
}

func (sr *SceneRequest) toSceneContext(ctx context.Context) *scene.Context {
	return scene.NewContext(
		scene.Request{
			SearchedMessage: sr.Command,
			FullMessage:     sr.FullUserText,
			WasButton:       sr.WasButton,
			Payload:         sr.Payload,
		},
		scene.UserInfo{
			UserId:    sr.Info.UserId,
			SessionId: sr.Info.SessionId,
		},
		ctx,
	)
}
