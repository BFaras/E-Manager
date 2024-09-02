package db

import (
	"back-end/internal/domain/entity"
	"back-end/internal/domain/repository"
	"back-end/internal/infrastructure/logger"
	"database/sql"
    "go.uber.org/zap"
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

func (r *storeRepository) IsOwnerOfStore(id string, userId string) (bool) {
    var count int
    query := `
        SELECT COUNT(*)
        FROM "public"."Store"
        WHERE "id" = $1 AND "userId" = $2
    `
    err := r.db.QueryRow(query, id, userId).Scan(&count)
    if err != nil {
        logger.Error("Error while checking ownership of store: ", zap.Error(err))
        return false
    }
    return count > 0
}

func (r *storeRepository) Create(store *entity.Store) error {
    query := `
    INSERT INTO "public"."Store" ("id","name","userId","createdAt", "updatedAt")
    VALUES ($1, $2, $3, $4, $5)
    `
    _, err := r.db.Exec(query, store.Id, store.Name, store.UserId, store.CreatedAt, store.UpdatedAt)
    if err != nil {
        logger.Error("Error : ",zap.Error(err))
        return err
}

return nil
}

func (r *storeRepository) Update(store *entity.Store) (error) {
    query := `
    UPDATE "public"."Store"
    SET "name" = $1, "userId" = $2, "createdAt" = $3, "updatedAt" = $4
    WHERE "id" = $5
    `
    _, err := r.db.Exec(query, store.Name, store.UserId, store.CreatedAt, store.UpdatedAt, store.Id)
    if err != nil {
        logger.Error("Error while updating store : ", zap.Error(err))
        return err
}

return nil
}

func (r *storeRepository) Delete(id string) error {
    query := `DELETE FROM "public"."Store" WHERE id = $1;`
    result, err := r.db.Exec(query, id)
    if err != nil {
        logger.Error("Error while deleting store : ",zap.Error(err))
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        logger.Error("Error : ",zap.Error(err))
        return err
    }

    if rowsAffected == 0 {
        logger.Error("No row has been affected : ",zap.String("storeId",id))
    }

    return nil
}

