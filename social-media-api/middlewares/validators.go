package middlewares

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"

// 	"github.com/asaskevich/govalidator"
// )

// func Validator(next http.HandlerFunc, s interface{},from string) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		err := parseRequest(r,s,from)

// 		if err != nil {
// 			http.Error(w,"Failed to Parse Request",http.StatusBadRequest);
// 			return
// 		}

// 		err := validateStruct(s);
// 		if err != nil {
// 			http.Error(w,"validation error",http.StatusBadRequest);
// 			return
// 		}

// 		next(w,r);
// 	}

// }

// func parseRequest(r *http.Request, s *interface{}, from string) (error) {
// 	switch from {
// 	case "body":
// 		body, err := ioutil.ReadAll(r.Body);
// 		if err != nil {
// 			return err;
// 		}

// 		err = json.Unmarshal(body, &s);
// 		if err != nil {
// 			return err;
// 		}
// 	return nil;
// }

// func validateStruct(s interface {}) (error) {
// 		err := govalidator.ValidateStruct(s);

// 		if err !=nil {
// 			return err;
// 		}
// 		return nil;
// }
