package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// --- Models ---

// User represents a user in our system, with MongoDB BSON tags
type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"` // Hashed password
}

// Credentials represents the login request body
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims represents the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// --- Global Variables (Configuration) ---

// Define a secret key for signing JWTs. In a real application, this should be a strong,
// securely stored secret, perhaps loaded from environment variables or a secret management service.
var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// --- Repository Layer ---

// UserRepository defines the interface for user data operations
type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
}

// MongoUserRepository implements UserRepository for MongoDB
type MongoUserRepository struct {
	collection *mongo.Collection
}

// NewMongoUserRepository creates a new MongoUserRepository instance
func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{
		collection: db.Collection("users"), // "users" is the collection name
	}
}

// FindByUsername finds a user by their username in MongoDB
func (r *MongoUserRepository) FindByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	filter := bson.M{"username": username}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}
	return &user, nil
}

// CreateUser creates a new user in MongoDB
func (r *MongoUserRepository) CreateUser(ctx context.Context, user *User) error {
	// Hash the password before saving
	user.Password = hashPassword(user.Password)
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// --- Service Layer ---

// AuthService handles business logic related to authentication
type AuthService struct {
	userRepo UserRepository
}

// NewAuthService creates a new AuthService instance
func NewAuthService(repo UserRepository) *AuthService {
	return &AuthService{userRepo: repo}
}

// Login performs user authentication and generates a JWT
func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return "", fmt.Errorf("error finding user: %w", err)
	}
	if user == nil || !checkPasswordHash(password, user.Password) {
		return "", fmt.Errorf("invalid credentials")
	}

	// Set token expiration time (e.g., 5 minutes)
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "auth-service",
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("could not generate token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken validates a given JWT string
func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	return claims, nil
}

// --- Controller Layer ---

// AuthController handles HTTP requests for authentication
type AuthController struct {
	authService *AuthService
}

// NewAuthController creates a new AuthController instance
func NewAuthController(service *AuthService) *AuthController {
	return &AuthController{authService: service}
}

// LoginHandler handles user login requests.
func (c *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	tokenString, err := c.authService.Login(r.Context(), creds.Username, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// ProtectedHandler is an example of a protected endpoint.
func (c *AuthController) ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Missing authorization token", http.StatusUnauthorized)
		return
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims, err := c.authService.ValidateToken(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome, %s! You have accessed a protected resource.", claims.Username)
}

// --- Utility Functions ---

// hashPassword hashes a given password using bcrypt.
func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	return string(bytes)
}

// checkPasswordHash compares a hashed password with a plaintext password.
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// --- Main Function ---

func main() {
	// Ensure JWT_SECRET_KEY is set
	if os.Getenv("JWT_SECRET_KEY") == "" {
		log.Fatal("Error: JWT_SECRET_KEY environment variable not set.")
	}

	// Get MongoDB connection string from environment variable
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("Error: MONGO_URI environment variable not set. Example: mongodb://localhost:27017")
	}

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	// Ping the MongoDB server to verify connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	fmt.Println("Successfully connected to MongoDB!")

	// Get the database instance
	db := client.Database("authdb") // "authdb" is the database name

	// Initialize Repository, Service, and Controller
	userRepo := NewMongoUserRepository(db)
	authService := NewAuthService(userRepo)
	authController := NewAuthController(authService)

	// --- Seed initial user if not exists (for demonstration) ---
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = userRepo.FindByUsername(ctx, "admin")
	if err != nil && err.Error() == "failed to find user by username: mongo: no documents in result" {
		// Only create if user doesn't exist
		initialUser := &User{
			Username: "admin",
			Password: "password123", // This will be hashed by CreateUser
		}
		err = userRepo.CreateUser(ctx, initialUser)
		if err != nil {
			log.Printf("Warning: Failed to create initial 'admin' user: %v", err)
		} else {
			fmt.Println("Initial 'admin' user created.")
		}
	} else if err != nil {
		log.Printf("Error checking for initial user: %v", err)
	} else {
		fmt.Println("'admin' user already exists.")
	}
	// --- End Seed ---

	// Define HTTP routes
	http.HandleFunc("/login", authController.LoginHandler)
	http.HandleFunc("/protected", authController.ProtectedHandler)

	// Start the HTTP server
	port := ":8080"
	fmt.Printf("Auth microservice listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
