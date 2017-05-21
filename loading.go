package oak

import (
	"bitbucket.org/oakmoundstudio/oak/audio"
	"bitbucket.org/oakmoundstudio/oak/dlog"
	"bitbucket.org/oakmoundstudio/oak/render"
)

var (
	startupLoadComplete = make(chan bool)
	// LoadingR is a renderable that is displayed during loading screens.
	LoadingR render.Renderable
)

func loadAssets() {
	dlog.Info("Loading Images")
	err := render.BatchLoad(imageDir)
	if err != nil {
		dlog.Error(err)
		return
	}
	dlog.Info("Done Loading Images")
	dlog.Info("Loading Audio")
	err = audio.BatchLoad(audioDir)
	if err != nil {
		dlog.Error(err)
	}
	dlog.Info("Done Loading Audio")

	startupLoadComplete <- true
}
