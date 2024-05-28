package firebase

import (
	"encoding/json"
	"log"
	"tanahore/configs"
	"tanahore/internal/model/domain"
)

func FirebaseConfigConverter(firebaseConfig *configs.FirebaseConfig) ([]byte, error) {
	var config domain.FirebaseAdminConfig
	err := json.Unmarshal([]byte(firebaseConfig.FirebaseKey), &config)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
		return nil, err
	}
	encodedJSON, err := json.Marshal(config)
	if err != nil {
		log.Fatalf("Failed to encode JSON: %v", err)
	}
	return encodedJSON, nil
}
