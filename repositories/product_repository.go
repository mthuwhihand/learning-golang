package repositories

import (
	"context"
	"firstproject/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(productID string) error
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Client) ProductRepository {
	return &productRepository{collection: db.Database("elec_equipment").Collection("products")}
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	// product.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(context.TODO(), product)
	return err
}

func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			log.Println("Error decoding product:", err)
			continue
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *productRepository) UpdateProduct(product *models.Product) error {
	id, err := primitive.ObjectIDFromHex(product.ID.Hex())
	if err != nil {
		log.Println("Invalid product ID:", err)
		return err
	}

	update := bson.M{"$set": product}
	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	return err
}

func (r *productRepository) DeleteProduct(productID string) error {
	id, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		log.Println("Invalid product ID:", err)
		return err
	}

	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
