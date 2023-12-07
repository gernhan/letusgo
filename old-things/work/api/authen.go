package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"strings"
)

const (
	XApiKeyHeader       string = "x-api-key"
	AuthorizationHeader string = "Authorization"
	SecretKey           string = "secret-key"
	Bearer              string = "Bearer"
	ByPassAuthorHeader  string = "X-ByPass-Author"
	RequestSubHeader    string = "X-Sub-Request"
	RequestSubHeaderFE  string = "X-Sub-Request-FE"
)

var parser = new(jwt.Parser)

func validateHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("validateHeader %s", r.URL)
		transformJWT(r)
		next.ServeHTTP(w, r)
	})
}

func transformJWT(r *http.Request) {
	// omit header firstly for prevent attack
	r.Header.Del(ByPassAuthorHeader)
	r.Header.Del(RequestSubHeader)
	r.Header.Del(RequestSubHeaderFE)

	jwtXApiKey := strings.TrimSpace(r.Header.Get(XApiKeyHeader))
	if authorToken, err := parseJwt(jwtXApiKey); err == nil {
		r.Header.Set(ByPassAuthorHeader, "true")
		if sub, getSubErr := getSubClaim(authorToken); getSubErr == nil {
			log.Infof("author sub: %s", sub)
			r.Header.Set(RequestSubHeader, sub)
		}
	}

	//handle for service call from Frontend
	jwtAuthor := strings.TrimSpace(strings.TrimPrefix(r.Header.Get(AuthorizationHeader), Bearer))
	if authorToken, jwtAuthorErr := parseJwt(jwtAuthor); jwtAuthorErr == nil {
		if sub, getSubErr := getSubClaim(authorToken); getSubErr == nil {
			log.Infof("author sub: %s", sub)
			r.Header.Set(RequestSubHeaderFE, sub)
		}
	}
}

func parseJwt(jwtData string) (*jwt.Token, error) {
	token, _, err := parser.ParseUnverified(jwtData, jwt.MapClaims{})
	if err != nil {
		return nil, errors.Wrapf(err, "fail to parse jwt %s", jwtData)
	}

	return token, nil
}

func getSubClaim(token *jwt.Token) (string, error) {
	if token == nil {
		return "", errors.New("token must not be null")
	}

	if token.Claims == nil {
		return "", errors.New("token is invalid, claims must not be null")
	}

	mapClaims := token.Claims.(jwt.MapClaims)
	if mapClaims == nil {
		return "", errors.New("token is invalid, claims is invalid, claims must be mapClaims type")
	}

	subValue, isExisting := mapClaims["sub"]
	if !isExisting {
		return "", errors.New("token is invalid, sub claims must be existing")
	}

	subStr, ok := subValue.(string)
	if !ok {
		return "", errors.New("token is invalid, sub claims must be string type")
	}

	return subStr, nil
}

var tcinvestURL = []string{"/v1/([0-9]+)/balances", "/v1/([0-9]+)/transactions",}
var tcbackendURL = "/v1/memberships/update"
var healthCheckURL = "/"

//func validateAuthor(r *http.Request) bool {
//	if isInternalCall(r) == false {
//		return false
//	}
//
//	return true
//}

func validateAuthorForInternalCall(r *http.Request) (bool, string) {
	return isInternalCall(r)
}

func validateAuthorForExternalCall(r *http.Request) (bool, string) {

	//if an request is an external call => it have not to be an internal call
	// because it is called from frontend => it will have not an X-api-key
	isInternal, msg := isInternalCall(r)
	if isInternal == false {
		//not internal call => verify it is external call
		requestSubHeaderFE := strings.TrimSpace(r.Header.Get(RequestSubHeaderFE))
		log.Infof("%v : %v", RequestSubHeaderFE, requestSubHeaderFE)
		if requestSubHeaderFE != "" {
			urlPath := r.URL.Path
			log.Info("Path ", urlPath)
			isValidateUrlFE, message := validateTcbsIdFromUrlAndInHeader(urlPath, r)
			if isValidateUrlFE == false {
				log.Error(message)
				return false, message
			}
			log.Info("Validate header for external call success")

			return true, "validate fe request success"
		}
		log.Info("Validate header for external call failed")
		return false, "requestSubHeaderFE is empty"
	}
	return isInternal, msg
}

func validateTcbsIdFromUrlAndInHeader(currentUri string, r *http.Request) (bool, string) {

	for _, uri := range tcinvestURL {
		matched, _ := regexp.MatchString(uri, currentUri)
		if matched == true {
			arr := strings.Split(currentUri, "/")
			tcbsid := arr[len(arr)-2]
			if tcbsid == r.Header.Get(RequestSubHeaderFE) {
				return true, ""
			}
			return false, fmt.Sprintf("tcbsid in uri and in token are different uri: %s , jwt: %s", tcbsid, r.Header.Get(RequestSubHeaderFE))
		}
	}

	return false, fmt.Sprintf("Not matched any request allow tcinvest3 %v", tcinvestURL)
}

var securityAuthToken = ""

func isInternalCall(r *http.Request) (bool, string) {
	isPassXApiKey, msg := isPassXApiKey(r)
	if isPassXApiKey == false {
		return isPassXApiKey, msg
	}
	log.Info("Validate header x-api-key for internal call success")
	//verify Secret-Key in header to confirm request is from an authorized service
	secretKey := strings.TrimSpace(r.Header.Get(SecretKey))
	if secretKey != securityAuthToken {
		log.Infof("Validate securityAuthToken failed, wrong secretKey : %v", secretKey)
		return false, "Wrong secret key"
	}
	log.Info("Validate header secretKey for internal call success")
	return true, "success"
}

func isPassXApiKey(r *http.Request) (bool, string) {
	value := strings.TrimSpace(r.Header.Get(ByPassAuthorHeader))
	if value != "true" {
		log.Info("Validate header for internal call failed")
		return false, "Not pass author header"
	}
	return true, "Pass x-api-key"
}

func responseForbidden(w http.ResponseWriter, message string) {
	w.WriteHeader(403)
	w.Write([]byte(message))
	return
}
