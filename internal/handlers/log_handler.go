package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// HandleMessageAndSend procesa el mensaje y envía el ID a la API
func HandleMessageAndSend(message []byte) {
	// Parsear el mensaje recibido
	var data struct {
		ID int64 `json:"id"`
	}
	err := json.Unmarshal(message, &data)
	if err != nil {
		log.Printf("Error parseando el mensaje: %v", err)
		return
	}

	log.Printf("ID recibido: %d", data.ID)

	// Construir la solicitud para la API
	apiURL := "http://127.0.0.1:4000/v1/recomendaciones/"
	payload := map[string]int64{"id": data.ID}
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