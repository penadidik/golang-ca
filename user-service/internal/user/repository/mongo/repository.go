package mongo

import (
    "context"
    "https://github.com/penadidik/golang-ca/tree/c497d7ec117e47e1663371288d25dd38c13e8895/user-service/internal/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
    Create(context.Context, *model.User) error
    GetByID(context.Context, string) (*model.User, error)
    Update(context.Context, string, *model.User) error
    Delete(context.Context, string) error
    GetAll(context.Context) ([]model.User, error)
}

type repository struct {
    col *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
    return &repository{
        col: db.Collection("users"),
    }
}

func (r *repository) Create(ctx context.Context, user *model.User) error {
    _, err := r.col.InsertOne(ctx, user)
    return err
}

func (r *repository) GetByID(ctx context.Context, id string) (*model.User, error) {
    oid, _ := primitive.ObjectIDFromHex(id)
    var user model.User
    err := r.col.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
    return &user, err
}

func (r *repository) GetAll(ctx context.Context) ([]model.User, error) {
    cursor, err := r.col.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }

    var users []model.User
    err = cursor.All(ctx, &users)
    return users, err
}

func (r *repository) Update(ctx context.Context, id string, user *model.User) error {
    oid, _ := primitive.ObjectIDFromHex(id)
    _, err := r.col.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": user})
    return err
}

func (r *repository) Delete(ctx context.Context, id string) error {
    oid, _ := primitive.ObjectIDFromHex(id)
    _, err := r.col.DeleteOne(ctx, bson.M{"_id": oid})
    return err
}

