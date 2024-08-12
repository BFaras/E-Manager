package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"database/sql"
)

type storeRepository struct {
    db *sql.DB
}

func NewStoreRepository(db *sql.DB) repository.StoreRepository {
    return &storeRepository{db: db}
}

func (r *storeRepository) FindByID(id string) (*entity.Store ,error) {
    store:= &entity.Store{}
    query := `SELECT * FROM "public"."Store"stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&store.ID, &store.Name, &store.UserID, &store.CreatedAt, &store.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return store, nil
}

func (r *storeRepository) Create(store *entity.Store) error {

}


func (r *storeRepository) Update(store *entity.Store) (*entity.Store, error) {

}

func (r *storeRepository) Delete(id string) error {

}

