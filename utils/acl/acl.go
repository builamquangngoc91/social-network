package acl

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"social-network/utils/xerror"

	"github.com/dgrijalva/jwt-go"
)

type (
	AuthType     int
	ResponseType int
)

const (
	None = AuthType(0)
	User = AuthType(1)

	JSON = ResponseType(0)
	HTML = ResponseType(1)
	File = ResponseType(2)
)

type Decl struct {
	Auth         AuthType
	HandlerFunc  interface{}
	ResponseType ResponseType
}

func ServeHTTP(resp http.ResponseWriter, req *http.Request, acl map[string]map[string]Decl, jwtKey string) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	handler, ok := acl[req.URL.Path]
	if !ok {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	decl, ok := handler[req.Method]
	if !ok {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	handlerFunc := decl.HandlerFunc
	auth := decl.Auth

	// authorization
	switch auth {
	case User:
		var ok bool
		req, ok = validToken(req, jwtKey)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}
	case None:
		// no-op
	}

	switch decl.ResponseType {
	case JSON:
		// Call functions of services
		// Todo: check function (type of arguments, outputs)
		typ := reflect.TypeOf(handlerFunc)
		funcArgs := reflect.New(typ.In(1).Elem()).Interface()

		// parse body and request values
		switch req.Method {
		case http.MethodPost:
			err := json.NewDecoder(req.Body).Decode(funcArgs)
			defer req.Body.Close()
			if err != nil {
				log.Panicf("error %v", err)
				return
			}

		case http.MethodGet:
			funcArgsTyp := reflect.TypeOf(funcArgs).Elem()
			for i := 0; i < funcArgsTyp.NumField(); i++ {
				jsonTag := funcArgsTyp.Field(i).Tag.Get("json")
				reflect.ValueOf(funcArgs).Elem().Field(i).Set(reflect.ValueOf(req.FormValue(jsonTag)))
			}
		}

		ctx := req.Context()
		outs := reflect.ValueOf(handlerFunc).Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(funcArgs)})
		resultFunc := outs[0].Interface()
		errFunc := outs[1].Interface()

		resp.Header().Set("Content-Type", "application/json")
		if errFunc != nil {
			err := errFunc.(xerror.XError)
			resp.WriteHeader(err.HttpStatus())
			json.NewEncoder(resp).Encode(map[string]string{
				"error": err.Message,
			})
			return
		} else {
			// json.NewEncoder(resp).Encode(map[string]interface{}{
			// 	"data": resultFunc,
			// })

			json.NewEncoder(resp).Encode(resultFunc)
		}

	case HTML:
		reflect.ValueOf(handlerFunc).Call([]reflect.Value{reflect.ValueOf(resp), reflect.ValueOf(req)})
	}
}

func validToken(req *http.Request, jwtKey string) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		log.Println(err)
		return req, false
	}

	if !t.Valid {
		return req, false
	}

	id, ok := claims["id"].(string)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), accountAuthKey(0), id))
	return req, true
}

type accountAuthKey int8

func AccountIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(accountAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
