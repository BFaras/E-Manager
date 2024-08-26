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

func (r *storeRepository) FindById(id string) (*entity.Store ,error) {
    store:= &entity.Store{}
    query := `SELECT * FROM "public"."Store"stores WHERE id = $1;`
    err := r.db.QueryRow(query, id).Scan(&store.Id, &store.Name, &store.UserId, &store.CreatedAt, &store.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return store, nil
}

func (r *storeRepository) FindAllStores() ([]*entity.Store, error) {
    var stores []*entity.Store
    query := `SELECT * FROM "public"."Store"`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        store := &entity.Store{}
        err := rows.Scan(&store.Id, &store.Name, &store.UserId, &store.CreatedAt, &store.UpdatedAt)
        if err != nil {
            return nil, err
        }
        stores = append(stores, store)
    }
    return stores, nil
}

func (r *storeRepository) FindByUserId(userId string) (*entity.Store ,error) {
    store:= &entity.Store{}
    query := `SELECT * FROM "public"."Store" stores WHERE "userId" = $1 LIMIT 1;`
    err := r.db.QueryRow(query, userId).Scan(&store.Id, &store.Name, &store.UserId, &store.CreatedAt, &store.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return store, nil
}

func (r *storeRepository) FindAllByUserId(userId string) ([]*entity.Store, error) {
    var stores []*entity.Store
    query := `SELECT * FROM "public"."Store" WHERE "userId" = $1`
    rows, err := r.db.Query(query, userId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        store := &entity.Store{}
        err := rows.Scan(&store.Id, &store.Name, &store.UserId, &store.CreatedAt, &store.UpdatedAt)
        if err != nil {
            return nil, err
        }
        stores = append(stores, store)
    }
    return stores, nil
}

func (r *storeRepository) FindByIdAndUserId(id string ,userId string) (*entity.Store, error) {
    store:= &entity.Store{}
    query := `SELECT * FROM "public"."Store" stores WHERE id = $1 AND "userId" = $2;`
    err := r.db.QueryRow(query, id, userId).Scan(&store.Id, &store.Name, &store.UserId, &store.CreatedAt, &store.UpdatedAt)
    if err!= nil {
        return nil, err
    }
    return store, nil
}

func (r *storeRepository) Create(store *entity.Store) error {
    return nil
}

func (r *storeRepository) Update(store *entity.Store) (*entity.Store, error) {
    return nil, nil
}

func (r *storeRepository) Delete(id string) error {
    return nil
}

