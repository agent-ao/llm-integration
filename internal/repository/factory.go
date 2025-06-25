package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type Repos struct {
	HealthRepo *HealthRepo
	// Add other repositories here as needed
	// Example: UserRepo *UserRepo
}

func NewRepos(db *mongo.Database) *Repos {
	return &Repos{
		HealthRepo: NewHealthRepo(db),
	}
}
