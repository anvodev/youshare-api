package main

import (
	"fmt"
	"net/http"
	"time"

	"youshare-api.anvo.dev/internal/data"
)

func (app *application) createVideoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create video")
}

func (app *application) showVideoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	video := data.Video{
		ID: id,
		Url: "https://www.youtube.com/watch?v=KLuTLF3x9sA",
		Title: "Norway 4K • Scenic Relaxation Film with Peaceful Relaxing Music and Nature Video Ultra HD",
		Description: `12 hours of healing music and relaxation with beautiful views of Norway in 4K Ultra HD
		🎹 Richard Nomad`,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelop{"video": video}, nil)
	if err != nil {
		app.logger.Print(err)
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateVideoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	fmt.Fprintf(w, "update video %d \n", id)
}

func (app *application) deleteVideoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	fmt.Fprintf(w, "delete video %d \n", id)
}
