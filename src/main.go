package main

import (
	"ApiServices/src/models"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/tiaguinho/gosoap"
)

func GetSaludo(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	fmt.Fprintf(w, "Bienvenido, capa de servicios de consultas, %s!", param["saludo"])
	fmt.Println("Endpoint: Capa de servicios--> GetSaludo: ", param["saludo"])
}

var criterioBusqueda string
var tipoBusqueda string

//instancia de estructura models/Models
var responseMusic []models.ReponseSearchMusic
var responseTv []models.ReponseSearchTv

func main() {
	router := mux.NewRouter()
	//Metodo de prueba, saludo from rest api
	router.HandleFunc("/api/GetSaludo/{saludo}", GetSaludo).Methods("GET")
	//router.HandleFunc("/api/GetBusqueda/{criterio}/{tipo}", GetCriterioBusqueda).Methods("GET")

	//metodo principal de busqueda
	router.HandleFunc("/api/GetBusqueda/{tipo}/{criterio}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint: Capa de servicios--> GetBusqueda ")

		param := mux.Vars(r)
		tipoBusqueda = param["tipo"]
		criterioBusqueda = param["criterio"]
		fmt.Printf("Tipo Busqueda: " + tipoBusqueda + "\n")

		switch tipoBusqueda {
		case "TELEVISION":
			fmt.Println("Iniciando busqueda ... ")
			var miUrl *url.URL
			miUrl, err := url.Parse("https://api.tvmaze.com/search/shows?")

			if err != nil {
				fmt.Println(err)
			}

			params := url.Values{}
			params.Add("q", criterioBusqueda)
			miUrl.RawQuery = params.Encode()

			resp, err := http.Get(miUrl.String())
			if err != nil {
				fmt.Println("Error response TELEVISION: ", err)
			}

			defer resp.Body.Close()
			responseBody, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Fatal(err)
			}

			var respJsTv []models.ResponseTV
			json.Unmarshal([]byte(responseBody), &respJsTv)

			for _, resp := range respJsTv {
				//fmt.Println(resp.Score)
				//fmt.Println(resp.Show.ID)
				//fmt.Println(resp.Show.Name)

				responseTv = append(responseTv, models.ReponseSearchTv{Name: resp.Show.Name, Type: resp.Show.Type,
					Language: resp.Show.Language, Status: resp.Show.Status, Premiered: resp.Show.Premiered,
					Ended: resp.Show.Ended, Summary: resp.Show.Summary, Time: resp.Show.Schedule.Time,
					Original: resp.Show.Image.Original, UrlOrigin: miUrl.String()})

			}

			//busqueda por persona
			miUrl, err = url.Parse("https://api.tvmaze.com/search/people?")

			if err != nil {
				fmt.Println(err)
			}

			params = url.Values{}
			params.Add("q", criterioBusqueda)
			miUrl.RawQuery = params.Encode()

			resp, err = http.Get(miUrl.String())
			if err != nil {
				fmt.Println("Error response TELEVISION: ", err)
			}

			defer resp.Body.Close()
			responseBody, err = ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Fatal(err)
			}

			var respJsTvP []models.ResponseTVPerson
			json.Unmarshal([]byte(responseBody), &respJsTvP)

			for _, resp := range respJsTvP {
				//fmt.Println("People: ", resp.Score)
				//fmt.Println(resp.ShowPerson.ID)
				//fmt.Println(resp.ShowPerson.Name)

				responseTv = append(responseTv, models.ReponseSearchTv{Name: resp.ShowPerson.Name, Type: resp.ShowPerson.Type,
					Language: resp.ShowPerson.Language, Status: resp.ShowPerson.Status, Premiered: resp.ShowPerson.Premiered,
					Ended: resp.ShowPerson.Ended, Summary: resp.ShowPerson.Summary, Time: resp.ShowPerson.Country.Name,
					Original: resp.ShowPerson.Url, UrlOrigin: miUrl.String()})

			}

			fmt.Println("Busqueda realizada con exito ... ")
			//se retorna respuesta consolidada
			json.NewEncoder(w).Encode(responseTv)

		case "MUSICA":
			fmt.Println("Iniciando busqueda ... ")

			var miUrl *url.URL
			var urlApi = "https://itunes.apple.com/search?"
			//Url de canciones
			miUrl, err := url.Parse(urlApi)

			if err != nil {
				fmt.Println(err)
			}

			params := url.Values{}
			params.Add("term", criterioBusqueda)
			params.Add("limit", "25")
			miUrl.RawQuery = params.Encode()

			resp, err := http.Get(miUrl.String())
			if err != nil {
				fmt.Println("Error response MUSICA: ", err)
			}

			defer resp.Body.Close()
			responseBody, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Fatal(err)
			}

			//se limpia la respuesta
			//jsonSSalto := strings.Replace(string(responseBody), `\n`, "", -1)
			//jsonResult := strings.Replace(string(jsonSSalto), `\`, "", -1)
			var respJs models.ResponseMusic
			json.Unmarshal([]byte(responseBody), &respJs)

			for _, resp := range respJs.Results {
				//fmt.Println("Musica: " + resp.WrapperType)
				//fmt.Println(resp.ArtistId)
				//fmt.Println(resp.CollectionPrice)

				responseMusic = append(responseMusic, models.ReponseSearchMusic{ArtistName: resp.ArtistName, CollectionName: resp.CollectionName,
					TrackName: resp.TrackName, CollectionPrice: resp.CollectionPrice, TrackPrice: resp.TrackPrice,
					ReleaseDate: resp.ReleaseDate, DiscCount: resp.DiscCount, Country: resp.Country,
					Currency: resp.Currency, UrlOrigin: miUrl.String()})

			}

			//Url de canciones con video
			miUrl, err = url.Parse(urlApi)
			if err != nil {
				fmt.Println(err)
			}

			params = url.Values{}
			params.Add("term", criterioBusqueda)
			params.Add("entity", "musicVideo")
			params.Add("limit", "25")
			miUrl.RawQuery = params.Encode()
			//fmt.Println("Url: " + miUrl.String())
			resp, err = http.Get(miUrl.String())
			if err != nil {
				fmt.Println("Error response TELEVISION: ", err)
			}

			defer resp.Body.Close()
			responseBody, err = ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Fatal(err)
			}

			var respMusicVideo models.ResponseMusic
			json.Unmarshal([]byte(responseBody), &respMusicVideo)

			for _, resp := range respMusicVideo.Results {
				//fmt.Println("Musica Video: " + resp.WrapperType)
				//fmt.Println(resp.ArtistId)
				//fmt.Println(resp.CollectionPrice)
				responseMusic = append(responseMusic, models.ReponseSearchMusic{ArtistName: resp.ArtistName, CollectionName: resp.CollectionName,
					TrackName: resp.TrackName, CollectionPrice: resp.CollectionPrice, TrackPrice: resp.TrackPrice,
					ReleaseDate: resp.ReleaseDate, DiscCount: resp.DiscCount, Country: resp.Country,
					Currency: resp.Currency, UrlOrigin: miUrl.String()})
			}

			//Url de canciones por pais
			miUrl, err = url.Parse(urlApi)
			if err != nil {
				fmt.Println(err)
			}

			params = url.Values{}
			params.Add("term", criterioBusqueda)
			params.Add("country", "ca")
			params.Add("limit", "25")
			miUrl.RawQuery = params.Encode()
			//fmt.Println("Url: " + miUrl.String())
			resp, err = http.Get(miUrl.String())
			if err != nil {
				fmt.Println("Error response TELEVISION: ", err)
			}

			defer resp.Body.Close()
			responseBody, err = ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Fatal(err)
			}

			var respMusicXCountry models.ResponseMusic
			json.Unmarshal([]byte(responseBody), &respMusicXCountry)

			for _, resp := range respMusicXCountry.Results {
				//fmt.Println("Musica por pais: " + resp.WrapperType)
				//fmt.Println(resp.ArtistId)
				//fmt.Println(resp.CollectionPrice)
				responseMusic = append(responseMusic, models.ReponseSearchMusic{ArtistName: resp.ArtistName, CollectionName: resp.CollectionName,
					TrackName: resp.TrackName, CollectionPrice: resp.CollectionPrice, TrackPrice: resp.TrackPrice,
					ReleaseDate: resp.ReleaseDate, DiscCount: resp.DiscCount, Country: resp.Country,
					Currency: resp.Currency, UrlOrigin: miUrl.String()})
			}

			fmt.Println("Busqueda realizada con exito ... ")
			//se retorna respuesta consolidada
			json.NewEncoder(w).Encode(responseMusic)

		case "PERSONA":
			fmt.Println("Iniciando busqueda ... ")
			httpClient := &http.Client{
				Timeout: time.Millisecond * 4500,
			}
			soap, err := gosoap.SoapClient("https://www.crcind.com/csp/samples/SOAP.Demo.cls?WSDL", httpClient)
			if err != nil {
				log.Fatalf("SoapClient error: %s", err)
			}

			//se envia el parametro a buscar
			params := gosoap.Params{
				"id": criterioBusqueda,
			}

			res, err := soap.Call("FindPerson", params)
			if err != nil {
				log.Fatalf("Error response PERSONA: %s", err)
			}

			model := models.FindPersonResponse{}
			//se setean los tags a los modelos
			err = xml.Unmarshal(res.Body, &model)

			if err != nil {
				fmt.Println("Error Unmarshal Soap ", err.Error())
			}

			/*fmt.Println(model.FindPersonResult.Name)
			fmt.Println(model.FindPersonResult.Home.Street)
			fmt.Println(model.FindPersonResult.Home.City)
			fmt.Println(model.FindPersonResult.Office.Street)*/

			fmt.Println("Busqueda realizada con exito ... ")
			//se retorna respuesta consolidada
			json.NewEncoder(w).Encode(model.FindPersonResult)

		default:
			fmt.Printf("Funcion no existe... ")
		}
	})

	fmt.Printf("Servidor iniciado... \n")

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8083",
		//Timeouts para el servidor
		WriteTimeout: time.Second * 20,
		ReadTimeout:  time.Second * 20,
		IdleTimeout:  time.Second * 60,
	}

	log.Fatal(srv.ListenAndServe())
}
