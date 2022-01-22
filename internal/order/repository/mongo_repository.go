package repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/romaxa83/hra/config"
	"github.com/romaxa83/hra/internal/order/models"
	"github.com/romaxa83/hra/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	logger logger.Logger
	cfg    *config.Config
	db     *mongo.Client
}

func NewMongoRepository(logger logger.Logger, cfg *config.Config, db *mongo.Client) *mongoRepository {
	return &mongoRepository{
		logger: logger,
		cfg:    cfg,
		db:     db,
	}
}

func (r *mongoRepository) CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {

	collection := r.db.Database(r.cfg.Mongo.Db).Collection(r.cfg.MongoCollections.Orders)

	_, err := collection.InsertOne(ctx, order, &options.InsertOneOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "InsertOne")
	}

	return order, nil
}
