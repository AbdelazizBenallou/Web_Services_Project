package usecase

import (
	"user_service/domain"
	"user_service/repository"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/argon2"
	"encoding/json"

)

type UserUseCase interface {
	Register(email, password, fullName string, role domain.Role, profile domain.Profile) (*domain.User, error)
	Login(email, password string) (*domain.User, error)
	GetUserByID(id int64) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
	ValidateUserCredentials(email, password string) (*domain.User, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
	publisher EventPublisher
}

func NewUserUseCase(
	userRepo repository.UserRepository,
	publisher EventPublisher,
) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
		publisher: publisher,
	}
}

type argon2Params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var argon2Config = &argon2Params{
	memory:      64 * 1024, // 64 MB
	iterations:  3,
	parallelism: 2,
	saltLength:  16,
	keyLength:   32,
}

func (uc *userUseCase) Register(email, password, fullName string, role domain.Role, profile domain.Profile) (*domain.User, error) {
	// 1. Validate inputs
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("email is required")
	}
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("password is required")
	}
	if strings.TrimSpace(fullName) == "" {
		return nil, errors.New("full name is required")
	}
	if !role.IsValid() {
		return nil, errors.New("invalid role")
	}
	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}

	// 2. Check if email already exists
	exists, err := uc.userRepo.EmailExists(email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email: %w", err)
	}
	if exists {
		return nil, errors.New("email already registered")
	}

	// 3. Hash password with Argon2
	hashedPassword, err := uc.hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// 4. Create user domain object
	user := &domain.User{
		FullName:  strings.TrimSpace(fullName),
		Email:     strings.ToLower(strings.TrimSpace(email)),
		Password:  hashedPassword,
		Role:      role,
		Profile:   profile,
		CreatedAt: time.Now(),
	}

	// 5. Save to repository
	err = uc.userRepo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}


	if uc.publisher != nil {
	event := map[string]interface{}{
		"user_id": user.ID,
	}

	data, _ := json.Marshal(event)
	_ = uc.publisher.Publish("user.registered", data)
}


	// 6. Clear sensitive data before returning
	user.Password = ""
	return user, nil
}

func (uc *userUseCase) Login(email, password string) (*domain.User, error) {
	// 1. Validate inputs
	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return nil, errors.New("email and password are required")
	}

	// 2. Get user from repository
	user, err := uc.userRepo.GetByEmail(strings.ToLower(strings.TrimSpace(email)))
	if err != nil {
		// Security: Return generic error to avoid user enumeration
		return nil, errors.New("invalid credentials")
	}

	// 3. Verify password with Argon2
	valid, err := uc.verifyPassword(password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("password verification failed: %w", err)
	}
	if !valid {
		return nil, errors.New("invalid credentials")
	}

	// 5. Clear sensitive data before returning
	user.Password = ""
	return user, nil
}

func (uc *userUseCase) ValidateUserCredentials(email, password string) (*domain.User, error) {
	return uc.Login(email, password)
}

func (uc *userUseCase) GetUserByID(id int64) (*domain.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}

	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Clear sensitive data
	user.Password = ""
	return user, nil
}

func (uc *userUseCase) GetAllUsers() ([]*domain.User, error) {
	users, err := uc.userRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	// Clear sensitive data from all users
	for _, user := range users {
		user.Password = ""
	}

	return users, nil
}

func (uc *userUseCase) hashPassword(password string) (string, error) {
	// Generate random salt
	salt := make([]byte, argon2Config.saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Hash password with Argon2
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		argon2Config.iterations,
		argon2Config.memory,
		argon2Config.parallelism,
		argon2Config.keyLength,
	)

	// Encode salt and hash to base64
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Format: $argon2id$v=19$m=65536,t=3,p=2$salt$hash
	encodedHash := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		argon2Config.memory,
		argon2Config.iterations,
		argon2Config.parallelism,
		b64Salt,
		b64Hash,
	)

	return encodedHash, nil
}

func (uc *userUseCase) verifyPassword(password, encodedHash string) (bool, error) {
	// Parse encoded hash
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	// Extract parameters
	var memory, iterations uint32
	var parallelism uint8
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
	if err != nil {
		return false, err
	}

	// Decode salt and hash
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	// Compute hash with provided password
	actualHash := argon2.IDKey(
		[]byte(password),
		salt,
		iterations,
		memory,
		parallelism,
		uint32(len(expectedHash)),
	)

	// Constant time comparison to prevent timing attacks
	return subtle.ConstantTimeCompare(actualHash, expectedHash) == 1, nil
}
