package main

import (
  "flag"
  "reflect"
  
  "net/http"
  "github.com/gorilla/mux"
  "github.com/golang/glog"
  
  "github.com/Lunkov/lib-auth/base"

  "github.com/Lunkov/lib-model"
  "github.com/Lunkov/lib-cms"
)

var staticFS http.Handler

func notFound(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func main() {
  flag.Set("alsologtostderr", "true")
  flag.Set("log_dir", ".")
  // flag.Set("v", "9")
  flag.Parse()

  configPath := flag.String("config_path", "./etc/", "Config path")

  glog.Infof("LOG: Starting WebServer")
  
  cm := cms.New()
  cm.LoadConfig((*configPath) + "/config.yaml", 120)
 
  cm.InitAuth()
  defer cm.CloseAuth()
  
  cm.InitDB()
  
  cm.DB.BaseAdd("user_auth",            reflect.TypeOf(base.User{}))
  cm.DB.BaseAdd("news",                 reflect.TypeOf(NewsModel{}))
  cm.DB.BaseAdd("events",               reflect.TypeOf(EventsModel{}))
  cm.DB.BaseAdd("book",                 reflect.TypeOf(BookModel{}))
  cm.DB.BaseAdd("banners",              reflect.TypeOf(BannersModel{}))
  cm.DB.BaseAdd("user_group",           reflect.TypeOf(GroupModel{}))
  
  cm.DB.DBAutoMigrate(models.ConnectStr(cm.GetConfig().PostgresWrite))
  cm.DB.LoadData()
  defer cm.DB.Close()
  
  cm.InitUI()

  router := mux.NewRouter()
  router.NotFoundHandler = http.HandlerFunc(notFound)
  
  cm.HandleFuncs(router)

  glog.Infof("LOG: Starting HTTP server on %s", cm.GetConfig().Main.HTTPPort)
  
  err := http.ListenAndServe(":" + cm.GetConfig().Main.HTTPPort, router)
  if err != nil {
    glog.Errorf("ERR: HTTP server: %v", err)
  }
  glog.Infof("LOG: Finish AUTH server")
}
