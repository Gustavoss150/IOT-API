package services

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}
}

func StartMQTT() {
	LoadEnv()

	user := os.Getenv("MQTT_USER")
	password := os.Getenv("MQTT_PASSWORD")
	clientID := os.Getenv("MQTT_CLIENT_ID")
	broker := os.Getenv("MQTT_BROKER")

	if user == "" || password == "" || clientID == "" || broker == "" {
		log.Fatal("Variáveis de ambiente MQTT faltando. Verifique seu arquivo .env")
	}

	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetClientID(clientID).
		SetUsername(user).
		SetPassword(password).
		SetDefaultPublishHandler(onMessage).
		SetTLSConfig(&tls.Config{InsecureSkipVerify: true})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Erro ao conectar no broker: %v", token.Error())
	}

	log.Println("✅ Conectado ao HiveMQ Cloud!")

	if token := client.Subscribe("maquinas/status/#", 0, nil); token.Wait() && token.Error() != nil {
		log.Fatalf("Erro ao subscrever ao tópico: %v", token.Error())
	}
	log.Println("📡 Inscrito no tópico: maquinas/status/#")
}

func onMessage(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("📨 Recebi mensagem: %s do tópico: %s\n", msg.Payload(), msg.Topic())
}
