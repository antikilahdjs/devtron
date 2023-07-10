package repository

import (
	"github.com/go-pg/pg"
	"go.uber.org/zap"
	"time"
)

type EphemeralContainer struct {
	tableName           struct{} `sql:"ephemeral_container" pg:",discard_unknown_columns"`
	Id                  int      `sql:"id,pk"`
	Name                string   `sql:"name"`
	ClusterID           int      `sql:"cluster_id"`
	Namespace           string   `sql:"namespace"`
	PodName             string   `sql:"pod_name"`
	TargetContainer     string   `sql:"target_container"`
	Config              string   `sql:"config"`
	IsExternallyCreated bool     `sql:"is_externally_created"`
}

type EphemeralContainerAction struct {
	tableName            struct{}  `sql:"ephemeral_container_actions" pg:",discard_unknown_columns"`
	Id                   int       `sql:"id,pk"`
	EphemeralContainerID int       `sql:"ephemeral_container_id"`
	ActionType           int       `sql:"action_type"`
	PerformedBy          int       `sql:"performed_by"`
	PerformedAt          time.Time `sql:"performed_at"`
}

type EphemeralContainersRepository interface {
	SaveData(model *EphemeralContainer) error
	SaveAction(model *EphemeralContainerAction) error
	FindAllActive(clusterId int, namespace, podName string) ([]string, error)
}

func NewEphemeralContainersRepositoryImpl(dbConnection *pg.DB, logger *zap.SugaredLogger) *EphemeralContainersImpl {
	return &EphemeralContainersImpl{
		dbConnection: dbConnection,
		logger:       logger,
	}
}

type EphemeralContainersImpl struct {
	dbConnection *pg.DB
	logger       *zap.SugaredLogger
}

func (impl EphemeralContainersImpl) SaveData(model *EphemeralContainer) error {
	return impl.dbConnection.Insert(model)
}

func (impl EphemeralContainersImpl) SaveAction(model *EphemeralContainerAction) error {
	return impl.dbConnection.Insert(model)
}

func (impl EphemeralContainersImpl) FindAllActive(clusterId int, namespace, podName string) ([]string, error) {
	var activeContainers []EphemeralContainer
	err := impl.dbConnection.Model(&EphemeralContainer{}).
		Column("name").
		Where("id NOT IN (SELECT ephemeral_container_id FROM ephemeral_container_actions WHERE action_type = 2)").
		Where("cluster_id = ?", clusterId).
		Where("namespace = ?", namespace).
		Where("pod_name = ?", podName).
		Select(&activeContainers)

	if err != nil {
		return nil, err
	}

	names := make([]string, len(activeContainers))
	for i, container := range activeContainers {
		names[i] = container.Name
	}

	return names, nil
}
