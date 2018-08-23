package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
)

const (
	IMAGE_LIST_URL = "http://cloud-images.ubuntu.com/daily/streams/v1/index.json"
	INDEX = "index"
	GCE_SERVER = "com.ubuntu.cloud:daily:gce"
	PRODUCTS = "products"
)

type ImageList struct {}

func (il *ImageList) GetAllImages() ([]string, error) {
	var imageList []string
	response, err := http.Get(IMAGE_LIST_URL)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	imageList, err = il.getImageListByServerName(body, GCE_SERVER)

	if err != nil {
		return nil, err
	}

	return imageList, nil
}

func (il *ImageList) getImageListByServerName(body []byte, serverName string) ([]string, error) {
	var unparsedImageList []string
	var jsonResponse interface{}

	err := json.Unmarshal(body, &jsonResponse)

	if err != nil {
		return nil, err
	}

	serverInfoMap := jsonResponse.(map[string]interface{})[INDEX]

	gceServerInfo := serverInfoMap.(map[string]interface{})[GCE_SERVER]

	gceServerProducts := gceServerInfo.(map[string]interface{})[PRODUCTS]
	
	for _, val := range gceServerProducts.([]interface{}) {
		unparsedImageList = append(unparsedImageList, val.(string))
	}

	parsedImageList := il.getParsedImageList(unparsedImageList)

	var parsedImageListSlice sort.StringSlice = parsedImageList

	sort.Sort(sort.Reverse(parsedImageListSlice[:]))

	return parsedImageList, nil
}

func (il *ImageList) getParsedImageList(unparsedImageList []string) []string {
	var imageList []string
	for _, unparsedImage := range unparsedImageList {
		imageList = append(imageList, ParseImageName(unparsedImage))
	}
	return imageList
}
