package database

import "github.com/DarioRoman01/photos/models"

// DatabaseRepository is an interface that defines the methods that a database must implement.
type DatabaseRepository interface {
	// InsertImage inserts an image into the database.
	InsertImage(image *models.Image) error
	// InsertFolder inserts a folder into the database.
	InsertFolder(folder *models.Folder) error
	// InsertUser inserts a user into the database.
	InsertUser(user *models.User) error
	// GetImage retrieves an image from the database.
	GetImage(id string) (*models.Image, error)
	// GetFolder retrieves a folder from the database.
	GetFolder(id string) (*models.Folder, error)
	// GetImages retrieves all images from the database  from the given user.
	GetImages(userID string) ([]*models.Image, error)
	// GetImagesByFolder retrieves all images from the database  from the given folder.
	GetImagesByFolder(folderID string) ([]*models.Image, error)
	// GetFolders retrieves all folders from the database from the given user.
	GetFolders(userID string) ([]*models.Folder, error)
	// GetUserByUsername retrieves a user from the database by username.
	GetUserByUsername(username string) (*models.User, error)
	// GetUserByEmail retrieves a user from the database by email.
	GetUserByEmail(email string) (*models.User, error)
	// GetUserByID retrieves a user from the database by ID.
	GetUserByID(id string) (*models.User, error)
	// DeleteImage deletes an image from the database.
	DeleteImage(id string) error
	// DeleteFolder deletes a folder from the database.
	DeleteFolder(id string) error
	// DeleteUser deletes a user from the database.
	DeleteUser(id string) error
}

var databaseRepository DatabaseRepository

func SetDatabaseRepository(repository DatabaseRepository) {
	databaseRepository = repository
}

func InsertImage(image *models.Image) error {
	return databaseRepository.InsertImage(image)
}

func InsertFolder(folder *models.Folder) error {
	return databaseRepository.InsertFolder(folder)
}

func InsertUser(user *models.User) error {
	return databaseRepository.InsertUser(user)
}

func GetImage(id string) (*models.Image, error) {
	return databaseRepository.GetImage(id)
}

func GetFolder(id string) (*models.Folder, error) {
	return databaseRepository.GetFolder(id)
}

func GetImages(userID string) ([]*models.Image, error) {
	return databaseRepository.GetImages(userID)
}

func GetFolders(userID string) ([]*models.Folder, error) {
	return databaseRepository.GetFolders(userID)
}

func GetImagesByFolder(folderID string) ([]*models.Image, error) {
	return databaseRepository.GetImagesByFolder(folderID)
}

func GetUserByUsername(username string) (*models.User, error) {
	return databaseRepository.GetUserByUsername(username)
}

func GetUserByEmail(email string) (*models.User, error) {
	return databaseRepository.GetUserByEmail(email)
}

func GetUserByID(id string) (*models.User, error) {
	return databaseRepository.GetUserByID(id)
}

func DeleteImage(id string) error {
	return databaseRepository.DeleteImage(id)
}

func DeleteFolder(id string) error {
	return databaseRepository.DeleteFolder(id)
}

func DeleteUser(id string) error {
	return databaseRepository.DeleteUser(id)
}