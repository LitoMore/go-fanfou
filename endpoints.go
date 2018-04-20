package fanfou

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiBase                      = "http://api.fanfou.com"
	apiUserShow                  = apiBase + "/users/show.json"
	apiStatusesUpdate            = apiBase + "/statuses/update.json"
	apiAccountVerifyCredentials  = apiBase + "/account/verify_credentials.json"
	apiAccountUpdateProfile      = apiBase + "/account/update_profile.json"
	apiAccountUpdateNotifyNum    = apiBase + "/account/update_notify_num.json"
	apiSavedSearchesCreate       = apiBase + "/saved_searches/create.json"
	apiPhotosUpload              = apiBase + "/photos/upload.json"
	apiAccountUpdateProfileImage = apiBase + "/account/update_profile_image.json"
)

func (client *baseClient) UserShow(params *ReqParams) (*responseUser, []byte, error) {
	requestURL := fmt.Sprintf("%s?id=%s", apiUserShow, params.ID)
	data, err := client.makeRequest(http.MethodGet, requestURL, nil)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request UserShow: %+v", err)
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) StatusesUpdate(params *ReqParams) (*responseStatus, []byte, error) {
	data, err := client.makeRequest(http.MethodPost, apiStatusesUpdate, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to request StatusesUpdate: %+v", err)
	}

	ret := responseStatus{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}
