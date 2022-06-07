package accesstoken

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Lornzo/facebookgo"
	"github.com/Lornzo/facebookgo/internal/types"
)

func GetAppToken(appID string, appSecret string) (AppToken, error) {

	var (
		err             error
		apiRoute        string     = "/oauth/access_token"
		apiQuerys       url.Values = make(url.Values, 3)
		apiUrl          string
		apiResponse     *http.Response
		apiResponseBody []byte
		apiResponseData struct {
			AppToken
			types.APIError `json:"error"`
		}
	)

	apiQuerys.Add("client_id", appID)
	apiQuerys.Add("client_secret", appSecret)
	apiQuerys.Add("grant_tyep", "client_credentials")

	apiUrl = fmt.Sprintf("%s/%s%s?%s", facebookgo.API_HOST_GRAPH, facebookgo.API_VERSION, apiRoute, apiQuerys.Encode())

	if apiResponse, err = http.Get(apiUrl); err != nil {
		return AppToken{}, fmt.Errorf("request error : %w", err)
	}

	defer apiResponse.Body.Close()

	if apiResponseBody, err = ioutil.ReadAll(apiResponse.Body); err != nil {
		return AppToken{}, fmt.Errorf("requesting api error : %w", err)
	}

	if err = json.Unmarshal(apiResponseBody, &apiResponseData); err != nil {
		return AppToken{}, fmt.Errorf("api response error : %w", err)
	}

	if err = apiResponseData.Error(); err != nil {
		return AppToken{}, fmt.Errorf("api response error : %w", err)
	}

	return apiResponseData.AppToken, nil
}

func GetPageToken(userID string, userAccessToken string, tokens *PagesTokenList) error {
	return nil
}
