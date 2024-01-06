package app

import (
	"net/http"
	"strconv"
	"strings"

	"event_booking_api/internal/app/handlers"
	"event_booking_api/pkg/errors"
)

/*     /login 						POST 							*/
/*     /register 					POST 							*/
/*     /events 						GET 							*/
/*     /events 						POST    => Auth required 		*/
/*     /events/<id> 				GET 							*/
/*     /events/<id> 				PUT   	=> Auth required 		*/
/*     /events/<id> 				DELETE  => Auth required 		*/
/*     /events/<id>/register 		POST 	=> Auth required 	 	*/
/*     /events/<id>/register 		DELETE 	=> Auth required 		*/

func Router(w http.ResponseWriter, r *http.Request) {
    path := strings.Trim(r.URL.Path, "/")
    pathSegments := strings.Split(path, "/")

    switch pathSegments[0] {
		case "login":
			handlers.LoginHandler(w, r)
			return
		case "register":
			handlers.RegisterHandler(w, r)
			return
		case "events":
			if len(pathSegments) > 1 && !isValidEventID(pathSegments[1]) {
				errors.HandleError(w, errors.NewAPIError(nil, "Invalid ID. Must be an integer.", http.StatusBadRequest))
                return
			}

			switch {
				case len(pathSegments) == 1:
					handlers.EventsHandler(w, r)
				case len(pathSegments) == 2:
					handlers.SingleEventHandler(w, r)
				case len(pathSegments) == 3 && pathSegments[2] == "register":
					handlers.BookingHandler(w, r)
				default:
					http.NotFound(w, r)
			}
			return 
    }

    // If no route matches, return 404 Not Found
    http.NotFound(w, r)
}

func isValidEventID(eventIDStr string) bool {
    _, err := strconv.Atoi(eventIDStr)
    return err == nil
}
