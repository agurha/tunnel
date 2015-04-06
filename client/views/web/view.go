// interactive web user interface
package web

import (
	"net/http"
	"github.com/agurha/tunnel/client/mvc"
	"github.com/agurha/tunnel/pkg/log"
	"github.com/agurha/tunnel/pkg/proto"
	"github.com/agurha/tunnel/pkg/util"
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
//	"fmt"
)

type WebView struct {
	log.Logger

	ctl mvc.Controller

	// messages sent over this broadcast are sent to all websocket connections
	wsMessages *util.Broadcast
}


func NewWebView(ctl mvc.Controller, addr string) *WebView {
	wv := &WebView{
		Logger:     log.NewPrefixLogger("view", "web"),
		wsMessages: util.NewBroadcast(),
		ctl:        ctl,
	}


	boxHTTPFiles, err := rice.FindBox("public")

	log.Info("info about http files %s", boxHTTPFiles)

	if err != nil {
		log.Error("cannot find rice box. error: %s\n", err)
	}

	log.Info("quite good")

//	if flags.Verbose {
//		fmt.Printf("box http-files is appended: %t\n", boxHTTPFiles.IsAppended())
//	}

	fileServer := http.FileServer(boxHTTPFiles.HTTPBox())

	log.Info("info about fileserver %s", fileServer)

	// rootRouter is directly linked to the http server.
	rootRouter := mux.NewRouter()

	// serve static files on / and several subdirs
	// NOTE: only the exact path "/" will match. http.FileService will resolve this to index.html
	// NOTE: "index.html" itself wont match
	rootRouter.Methods("GET").Path("/").Handler(fileServer)
	rootRouter.Methods("GET").PathPrefix("/css/").Handler(fileServer)
	rootRouter.Methods("GET").PathPrefix("/fonts/").Handler(fileServer)
	rootRouter.Methods("GET").PathPrefix("/html/").Handler(fileServer)
	rootRouter.Methods("GET").PathPrefix("/js/").Handler(fileServer)
	rootRouter.Methods("GET").PathPrefix("/img/").Handler(fileServer)





	wv.Info("Serving web interface on %s", addr)
	wv.ctl.Go(func() { http.ListenAndServe(addr, rootRouter) })
	return wv
}

func (wv *WebView) NewHttpView(proto *proto.Http) *WebHttpView {
	return newWebHttpView(wv.ctl, wv, proto)
}

func (wv *WebView) Shutdown() {
}
