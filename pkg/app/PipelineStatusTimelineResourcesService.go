package app

import (
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/argoproj/gitops-engine/pkg/sync/common"
	"github.com/devtron-labs/devtron/internal/sql/repository/pipelineConfig"
	"github.com/devtron-labs/devtron/pkg/sql"
	"github.com/go-pg/pg"
	"go.uber.org/zap"
	"time"
)

type PipelineStatusTimelineResourcesService interface {
	SaveOrUpdateCdPipelineTimelineResources(cdWfrId int, application *v1alpha1.Application, timelineStage pipelineConfig.ResourceTimelineStage, tx *pg.Tx, userId int32) error
}

type PipelineStatusTimelineResourcesServiceImpl struct {
	dbConnection                              *pg.DB
	logger                                    *zap.SugaredLogger
	pipelineStatusTimelineResourcesRepository pipelineConfig.PipelineStatusTimelineResourcesRepository
}

func NewPipelineStatusTimelineResourcesServiceImpl(dbConnection *pg.DB, logger *zap.SugaredLogger,
	pipelineStatusTimelineResourcesRepository pipelineConfig.PipelineStatusTimelineResourcesRepository) *PipelineStatusTimelineResourcesServiceImpl {
	return &PipelineStatusTimelineResourcesServiceImpl{
		dbConnection: dbConnection,
		logger:       logger,
		pipelineStatusTimelineResourcesRepository: pipelineStatusTimelineResourcesRepository,
	}
}

func (impl *PipelineStatusTimelineResourcesServiceImpl) SaveOrUpdateCdPipelineTimelineResources(cdWfrId int, application *v1alpha1.Application, timelineStage pipelineConfig.ResourceTimelineStage, tx *pg.Tx, userId int32) error {
	//getting all timeline resources by cdWfrId
	timelineResources, err := impl.pipelineStatusTimelineResourcesRepository.GetByCdWfrIdAndTimelineStage(cdWfrId, timelineStage)
	if err != nil && err != pg.ErrNoRows {
		impl.logger.Errorw("error in getting timelineResources", "err", err)
		return err
	}
	//map of resourceName and its index
	oldTimelineResourceMap := make(map[string]int)

	for i, timelineResource := range timelineResources {
		oldTimelineResourceMap[timelineResource.ResourceName] = i
	}

	var timelineResourcesToBeSaved []*pipelineConfig.PipelineStatusTimelineResources
	var timelineResourcesToBeUpdated []*pipelineConfig.PipelineStatusTimelineResources

	if application != nil && application.Status.OperationState != nil && application.Status.OperationState.SyncResult != nil {
		for _, resource := range application.Status.OperationState.SyncResult.Resources {
			if index, ok := oldTimelineResourceMap[resource.Name]; ok {
				timelineResources[index].ResourceStatus = string(resource.HookPhase)
				timelineResources[index].StatusMessage = resource.Message
				timelineResources[index].UpdatedBy = userId
				timelineResources[index].UpdatedOn = time.Now()
				timelineResourcesToBeUpdated = append(timelineResourcesToBeUpdated, timelineResources[index])
			} else {
				newTimelineResource := &pipelineConfig.PipelineStatusTimelineResources{
					CdWorkflowRunnerId: cdWfrId,
					ResourceName:       resource.Name,
					ResourceKind:       resource.Kind,
					ResourceGroup:      resource.Group,
					ResourceStatus:     string(resource.HookPhase),
					StatusMessage:      resource.Message,
					TimelineStage:      timelineStage,
					AuditLog: sql.AuditLog{
						CreatedBy: userId,
						CreatedOn: time.Now(),
						UpdatedBy: userId,
						UpdatedOn: time.Now(),
					},
				}
				if resource.HookType != "" {
					newTimelineResource.ResourcePhase = string(resource.HookType)
				} else {
					//since hookType for non-hook resources is empty and always come under sync phase, hardcoding it
					newTimelineResource.ResourcePhase = string(common.HookTypeSync)
				}
				timelineResourcesToBeSaved = append(timelineResourcesToBeSaved, newTimelineResource)
			}
		}
	}
	if len(timelineResourcesToBeSaved) > 0 {
		if tx != nil {
			err = impl.pipelineStatusTimelineResourcesRepository.SaveTimelineResourcesWithTxn(timelineResourcesToBeSaved, tx)
			if err != nil {
				impl.logger.Errorw("error in saving timelineResources", "err", err, "timelineResources", timelineResourcesToBeSaved)
				return err
			}
		} else {
			err = impl.pipelineStatusTimelineResourcesRepository.SaveTimelineResources(timelineResourcesToBeSaved)
			if err != nil {
				impl.logger.Errorw("error in saving timelineResources", "err", err, "timelineResources", timelineResourcesToBeSaved)
				return err
			}
		}
	}
	if len(timelineResourcesToBeUpdated) > 0 {
		if tx != nil {
			err = impl.pipelineStatusTimelineResourcesRepository.UpdateTimelineResourcesWithTxn(timelineResourcesToBeUpdated, tx)
			if err != nil {
				impl.logger.Errorw("error in updating timelineResources", "err", err, "timelineResources", timelineResourcesToBeUpdated)
				return err
			}
		} else {
			err = impl.pipelineStatusTimelineResourcesRepository.UpdateTimelineResources(timelineResourcesToBeSaved)
			if err != nil {
				impl.logger.Errorw("error in updating timelineResources", "err", err, "timelineResources", timelineResourcesToBeUpdated)
				return err
			}
		}
	}
	return nil
}
