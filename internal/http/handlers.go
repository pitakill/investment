package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pitakill/investment/internal/mongo"
	"github.com/pitakill/investment/pkg/investment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Handlers() *chi.Mux {
	r := chi.NewRouter()

	// Only allow json
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		welcomeMessage := fmt.Sprintf(
			"Welcome to the %s server.\nMaybe you want to test your chances on /credit-assignment with a POST request with a body like {%q: 3000}",
			appName,
			"investment",
		)

		w.Write([]byte(welcomeMessage))
	})

	r.Post("/credit-assignment", func(w http.ResponseWriter, r *http.Request) {
		investment := new(investment.Request)

		if err := json.NewDecoder(r.Body).Decode(&investment); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			msg := BadRequest{Error: errInvalidJSON}
			d, _ := json.Marshal(msg)
			w.Write(d)

			return
		}

		// Handle db connection
		// TODO: get the databaes and collection from environment
		// TODO: pass a client configured to only handle the CRUD operations to the
		// DB
		collection := mng.Database("data").Collection("statistics")

		response, err := investment.Process()
		if err != nil {
			// Assignation rejected
			document := mongo.Investment{
				Amount:     investment.Investment,
				Successful: false,
			}
			collection.InsertOne(context.TODO(), document)

			w.WriteHeader(http.StatusBadRequest)
			msg := BadRequest{Error: err.Error()}
			d, _ := json.Marshal(msg)
			w.Write(d)

			return
		}

		data, _ := json.Marshal(response)

		// Assignation accepted
		document := mongo.Investment{
			Amount:     investment.Investment,
			Successful: true,
		}
		collection.InsertOne(context.TODO(), document)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	})

	r.Post("/statistics", func(w http.ResponseWriter, r *http.Request) {
		// Handle db connection
		// TODO: get the databaes and collection from environment
		// TODO: pass a client configured to only handle the CRUD operations to the
		// DB
		collection := mng.Database("data").Collection("statistics")

		var results []*mongo.Investment

		cur, _ := collection.Find(context.TODO(), bson.D{{}}, options.Find())
		for cur.Next(context.TODO()) {
			var element mongo.Investment
			err := cur.Decode(&element)
			if err != nil {
				log.Fatal(err)
			}

			results = append(results, &element)
		}
		cur.Close(context.TODO())

		response := struct {
			Total                          int     `json:"total"`
			Successful                     int     `json:"successful"`
			Failed                         int     `json:"failed"`
			PercentageInvestmentSuccessful float64 `json:"percentage_investment_successful"`
			PercentageInvestmentFailed     float64 `json:"percentage_investment_failed"`
			TotalSuccessful                int     `json:"-"`
			TotalFailed                    int     `json:"-"`
		}{
			Total: len(results),
		}

		for _, i := range results {
			if i.Successful {
				response.Successful++
				response.TotalSuccessful += int(i.Amount)
			} else {
				response.Failed++
				response.TotalFailed += int(i.Amount)
			}
		}

		response.PercentageInvestmentSuccessful = float64(response.TotalSuccessful) / float64(response.Successful)
		response.PercentageInvestmentFailed = float64(response.TotalFailed) / float64(response.Failed)

		data, _ := json.Marshal(response)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")

		w.Write([]byte(data))
	})

	return r
}
