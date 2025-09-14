package main
import ("fmt"; "net/http"; "os"; "github.com/go-chi/chi/v5")
func main(){r:=chi.NewRouter();r.Get("/healthz",func(w http.ResponseWriter,r *http.Request){w.Write([]byte("ok"))});r.Get("/version",func(w http.ResponseWriter,r *http.Request){v:=os.Getenv("APP_VERSION");if v==""{v="dev"};w.Write([]byte(fmt.Sprintf("%s",v)))});http.ListenAndServe(":8080",r)}
