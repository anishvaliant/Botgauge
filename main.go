package main

   import (
       "fmt"
       "net/http"
   )

   // helloHandler responds with "Hello from Botgauge"
   func helloHandler(w http.ResponseWriter, r *http.Request) {
       fmt.Fprintln(w, "Hello from Botgauge")
   }

   func main() {
       // Register the handler function for the /hello path
       http.HandleFunc("/hello", helloHandler)

       // Start the server on port 8080
       fmt.Println("Starting server at port 8080")
       if err := http.ListenAndServe(":8080", nil); err != nil {
           fmt.Printf("Error starting server: %s\n", err)
       }
   }