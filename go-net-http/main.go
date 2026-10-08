package main

import(
	"fmt"
	"net/http"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Go ^_^")
}

func hioGoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi Go ^_^")
}

func main (){
	http.HandleFunc("/hellogo", helloGoHandler)
	http.HandleFunc("/higo", hioGoHandler)

	fmt.Println("Server started at port 3030")

	//start server
	if err := http.ListenAndServe(":3030", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}