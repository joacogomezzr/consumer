package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// HandleMessageAndSend procesa el mensaje y envía el ID del libro a la API de recomendaciones
func HandleMessageAndSend(message []byte) {
	// Parsear el mensaje recibido
	var data struct {
		ID           int64 `json:"id"`
		Recommendable bool  `json:"recommendable"`
	}
	err := json.Unmarshal(message, &data)
	if err != nil {
		log.Printf("Error parseando el mensaje: %v", err)
		return
	}

	log.Printf("ID recibido: %d, Recommendable: %t", data.ID, data.Recommendable)

	// Validar si el libro es Recommendable
	if data.Recommendable {
		log.Println("El libro es recommendable, enviando a la API de recomendaciones.")
		sendToRecommendationsAPI(data.ID)
	} else {
		log.Println("El libro no es recommendable, no se enviará a la API.")
	}
}

func sendToRecommendationsAPI(bookID int64) {
	apiURL := "http://44.197.18.81:4000/v1/recomendaciones/"
	payload := map[string]int64{"idBook": bookID}
	body, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error creando el payload para la API: %v", err)
		return
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error enviando solicitud a la API: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Respuesta de la API: %s", resp.Status)
}