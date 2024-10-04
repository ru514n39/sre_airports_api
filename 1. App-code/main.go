package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Initialize logger to write to both standard output and a file
func init() {
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	log.SetOutput(file) // Log to file
	log.Println("Server started at", time.Now())
}

// Airport represents the structure of an airport
type Airport struct {
	Name     string `json:"name"`
	City     string `json:"city"`
	IATA     string `json:"iata"`
	ImageURL string `json:"image_url"`
}

// AirportV2 includes runway length along with basic airport details
type AirportV2 struct {
	Airport
	RunwayLength int `json:"runway_length"`
}

// Mock data for airports in Bangladesh
var airports = []Airport{
	{"Hazrat Shahjalal International Airport", "Dhaka", "DAC", "https://my-s3-bucket-name.s3.amazonaws.com/images/dac.jpg"},
	{"Shah Amanat International Airport", "Chittagong", "CGP", "https://my-s3-bucket-name.s3.amazonaws.com/images/cgp.jpg"},
	{"Osmani International Airport", "Sylhet", "ZYL", "https://my-s3-bucket-name.s3.amazonaws.com/images/zyl.jpg"},
}

var airportsV2 = []AirportV2{
	{Airport{"Hazrat Shahjalal International Airport", "Dhaka", "DAC", "https://my-s3-bucket-name.s3.amazonaws.com/images/dac.jpg"}, 3200},
	{Airport{"Shah Amanat International Airport", "Chittagong", "CGP", "https://my-s3-bucket-name.s3.amazonaws.com/images/cgp.jpg"}, 2900},
	{Airport{"Osmani International Airport", "Sylhet", "ZYL", "https://my-s3-bucket-name.s3.amazonaws.com/images/zyl.jpg"}, 2500},
}

// HealthCheck handler
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("HealthCheck endpoint hit")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Healthcheck: OK"))
}

// HomePage handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received: %s %s", r.Method, r.URL)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status: OK"))
}

// Airports handler
func Airports(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received: %s %s", r.Method, r.URL)
	if r.Method != http.MethodGet {
		log.Printf("Invalid method: %s. Only GET is allowed.", r.Method)
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(airports)
	if err != nil {
		log.Println("Failed to encode airports data:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// AirportsV2 handler
func AirportsV2(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received: %s %s", r.Method, r.URL)
	if r.Method != http.MethodGet {
		log.Printf("Invalid method: %s. Only GET is allowed.", r.Method)
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(airportsV2)
	if err != nil {
		log.Println("Failed to encode airportsV2 data:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Function to upload a file to S3
func uploadImageToS3(file multipart.File, fileName string) (string, error) {
	bucketName := "my-s3-bucket-name"
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Println("Failed to create AWS session:", err)
		return "", fmt.Errorf("failed to create AWS session: %v", err)
	}

	svc := s3.New(sess)

	buf := bytes.NewBuffer(nil)
	if _, err := buf.ReadFrom(file); err != nil {
		log.Println("Failed to read file:", err)
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	objectKey := fmt.Sprintf("images/%s_%d%s", fileName, time.Now().Unix(), filepath.Ext(fileName))

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectKey),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String("image/jpeg"),
		ACL:         aws.String("public-read"),
	})
	if err != nil {
		log.Println("Failed to upload image to S3:", err)
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}

	imageURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, objectKey)
	log.Println("Image successfully uploaded to S3:", imageURL)
	return imageURL, nil
}

// UpdateAirportImage handler
func UpdateAirportImage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received: %s %s", r.Method, r.URL)
	if r.Method != http.MethodPost {
		log.Printf("Invalid method: %s. Only POST is allowed.", r.Method)
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("Failed to parse form data:", err)
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	airportName := r.FormValue("name")
	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Println("Failed to retrieve image from form:", err)
		http.Error(w, "Failed to retrieve image from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	var airport *Airport
	for i := range airports {
		if airports[i].Name == airportName {
			airport = &airports[i]
			break
		}
	}

	if airport == nil {
		log.Println("Airport not found:", airportName)
		http.Error(w, "Airport not found", http.StatusNotFound)
		return
	}

	imageURL, err := uploadImageToS3(file, handler.Filename)
	if err != nil {
		log.Println("Failed to upload image:", err)
		http.Error(w, "Failed to upload image", http.StatusInternalServerError)
		return
	}

	airport.ImageURL = imageURL
	log.Printf("Successfully updated image for airport: %s", airportName)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":   "Image updated successfully",
		"image_url": imageURL,
	})
}

func main() {
	// Setup routes
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/airports", Airports)
	http.HandleFunc("/airports_v2", AirportsV2)
	http.HandleFunc("/healthcheck", HealthCheck)
	http.HandleFunc("/update_airport_image", UpdateAirportImage)

	// Start the server
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
