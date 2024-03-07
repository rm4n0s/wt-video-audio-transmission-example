package main

import "github.com/rm4n0s/wt-video-audio-transmission-example/app/controllers"

func main() {
	const appID = "com.github.rm4n0s.wt-video-audio-transmission-example"
	app := controllers.NewApp(appID)
	app.Run()
}
