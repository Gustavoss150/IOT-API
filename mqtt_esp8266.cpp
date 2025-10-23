// #include <ESP8266WiFi.h>
// #include <WiFiClientSecure.h>
// #include <PubSubClient.h>

// /*--- Configurações Wi-Fi ---*/
// const char* WIFI_SSID = "Euzinha";
// const char* WIFI_PW = "senha123";

// /*--- Configurações MQTT (HiveMQ Cloud) ---*/
// const char* MQTT_BROKER = "9d030438f89d4325bcff84c11cf9388c.s1.eu.hivemq.cloud";
// const uint16_t MQTT_PORT = 8883;
// const char* MQTT_USERNAME = "hivemq.webclient.1760040716694";
// const char* MQTT_PASSWORD = "i#ZOy!0f:Bg1<k57eJFP";

// /*--- Variáveis globais ---*/
// WiFiClientSecure secureClient;
// PubSubClient mqtt_client(secureClient);

// /*--- Função para conectar ao Wi-Fi ---*/
// void connectWiFi(const char* ssid, const char* password) {
//   Serial.println("Conectando ao WiFi...");
//   WiFi.begin(ssid, password);

//   while (WiFi.status() != WL_CONNECTED) {
//     delay(500);
//     Serial.print(".");
//   }

//   Serial.println("\n Conectado ao WiFi!");
//   Serial.print("Endereço IP: ");
//   Serial.println(WiFi.localIP());
// }

// String generateUUID() {
//   char uuid[37];
//   sprintf(uuid, "%04x%04x-%04x-%04x-%04x-%04x%04x%04x",
//           random(0xffff), random(0xffff),
//           random(0xffff),
//           (random(0x0fff)) | 0x4000,
//           (random(0x3fff)) | 0x8000,
//           random(0xffff), random(0xffff), random(0xffff));
//   return String(uuid);
// }

// String randomStatus() {
//   const char* statuses[] = {"ativo", "parado", "erro"};
//   int idx = random(0, 3);
//   return String(statuses[idx]);
// }

// /*--- Setup ---*/
// void setup() {
//   Serial.begin(9600);
//   connectWiFi(WIFI_SSID, WIFI_PW);

//   // TLS: ignorar verificação de certificado (ok para teste)
//   secureClient.setInsecure();

//   // Configuração MQTT
//   mqtt_client.setServer(MQTT_BROKER, MQTT_PORT);

//   Serial.println("Conectando ao HiveMQ Cloud...");
//   while (!mqtt_client.connected()) {
//     String client_id = "esp8266-client-" + String(WiFi.macAddress());
//     if (mqtt_client.connect(client_id.c_str(), MQTT_USERNAME, MQTT_PASSWORD)) {
//       Serial.println("Conectado ao HiveMQ Cloud!");
//     } else {
//       Serial.print("Falha na conexão MQTT. Código: ");
//       Serial.println(mqtt_client.state());
//       delay(2000);
//     }
//   }
// }

// /*--- Loop ---*/
// void loop() {
//   if (!mqtt_client.connected()) {
//     setup(); // reconecta se cair
//   }

//   mqtt_client.loop();

//   // Enviar mensagem a cada 60 segundos
//   String id = generateUUID();
//   String status = randomStatus();

//   String payload = "{\"id\": \"" + id + "\", \"status\": \"" + status + "\"}";

//   mqtt_client.publish("maquinas/status/1", payload.c_str());
//   Serial.println("📤 Enviado: " + payload);

//   delay(60000); // 1 minuto
// }
