package controller

import (
	"net/http"

	"github.com/go-squads/genrevan-scheduler/model"
)

var imageListModel model.ImageList

func GetAllAvailableImages(w http.ResponseWriter, r *http.Request) {
	imageList, err := imageListModel.GetAllImages()
	
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	RespondWithJSON(w, http.StatusOK, imageList)
}
