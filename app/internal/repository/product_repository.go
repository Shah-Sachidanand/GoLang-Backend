package repository

import (
	"context"
	"learning-golang/app/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetProducts() ([]models.Product, error)
	GetProductByID(id string) (*models.Product, error)
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(client *mongo.Client, dbName string) ProductRepository {
	collection := client.Database(dbName).Collection("products")
	return &productRepository{collection: collection}
}

func (r *productRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	product.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(context.Background(), product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *productRepository) GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
