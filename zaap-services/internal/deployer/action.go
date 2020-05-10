package deployer

import (
	"context"
	"errors"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/satori/go.uuid"
)

var (
	ErrApplicationNotFound = errors.New("application not found")
	ErrDeploymentNotFound  = errors.New("deployment not found")
	ErrRunnerNotFound      = errors.New("runner not found")
)

func (s *Server) deployApplication(ctx context.Context, applicationId uuid.UUID, deploymentId uuid.UUID) error {
	application, err := s.applicationStore.FindWithRunner(ctx, applicationId)
	if err != nil {
		return err
	} else if application == nil {
		return ErrApplicationNotFound
	}

	deployment, err := s.deploymentStore.Find(ctx, deploymentId)
	if err != nil {
		return err
	} else if deployment == nil {
		return ErrDeploymentNotFound
	}

	client, conn, err := s.runnerService.NewConnection(application.Runner)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = client.DeployApplication(ctx, &runnerpb.DeployApplicationRequest{
		Application: &runnerpb.Application{
			Id:               application.ID.String(),
			DeploymentId:     deployment.ID.String(),
			Name:             application.Name,
			Image:            deployment.Image,
			Replicas:         uint32(deployment.Replicas),
			Domains:          append(application.Domains, application.DefaultDomain),
			Environment:      deployment.Environment,
			Roles:            deployment.Roles,
			ImagePullSecrets: deployment.ImagePullSecrets,
		},
	})
	return err
}

func (s *Server) deleteApplication(ctx context.Context, applicationId uuid.UUID, applicationName string, runnerId uuid.UUID) error {
	runner, err := s.runnerStore.Find(ctx, runnerId)
	if err != nil {
		return err
	} else if runner == nil {
		return ErrRunnerNotFound
	}

	client, conn, err := s.runnerService.NewConnection(runner)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = client.DeleteApplication(ctx, &runnerpb.DeleteApplicationRequest{
		Id:   applicationId.String(),
		Name: applicationName,
	})
	return err
}
