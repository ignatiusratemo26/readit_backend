package routes

import (
	"readit_backend/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	// r.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	// r.HandleFunc("/api/products/popular", controllers.GetPopularProducts).Methods("GET")
	// r.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/api/profile", controllers.ProfileHandler).Methods("GET")
	r.HandleFunc("/api/register", controllers.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/login", controllers.LoginHandler).Methods("POST")
	// r.HandleFunc("/api/cart", controllers.GetCart).Methods("GET")
	// r.HandleFunc("/api/cart", controllers.AddToCart).Methods("POST")
	// r.HandleFunc("/api/cart", controllers.RemoveFromCart).Methods("DELETE")
	// r.HandleFunc("/api/purchase", controllers.PurchaseCart).Methods("POST")
	// r.HandleFunc("/api/purchases", controllers.GetPurchases).Methods("GET")
	r.HandleFunc("/api/create-subscription", controllers.CreateSubscriptionPlanHandler).Methods("POST")
	r.HandleFunc("/api/subscriptions", controllers.GetSubscriptionPlansHandler).Methods("GET")
	r.HandleFunc("/api/products", controllers.GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/subscriptions/purchase", controllers.PurchaseSubscription).Methods("POST")
	r.HandleFunc("/api/subscriptions/view", controllers.ViewUserSubscription).Methods("GET")
	// r.HandleFunc("/api/products/popular", controllers.GetPopularProductsHandler).Methods("GET")

	// Contact routes
	r.HandleFunc("/api/contact", controllers.SubmitContactHandler).Methods("POST")

	return r
}
