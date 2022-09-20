package v7action

import (
	"io"

	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	"code.cloudfoundry.org/cli/resources"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . CloudControllerClient

// CloudControllerClient is the interface to the cloud controller V3 API.
type CloudControllerClient interface {
	ApplyOrganizationQuota(quotaGUID string, orgGUID string) (resources.RelationshipList, ccv3.Warnings, error)
	ApplySpaceQuota(quotaGUID string, spaceGUID string) (resources.RelationshipList, ccv3.Warnings, error)
	AppSSHEndpoint() string
	AppSSHHostKeyFingerprint() string
	CheckRoute(domainGUID string, hostname string, path string, port int) (bool, ccv3.Warnings, error)
	CloudControllerAPIVersion() string
	CancelDeployment(deploymentGUID string) (ccv3.Warnings, error)
	CopyPackage(sourcePackageGUID string, targetAppGUID string) (resources.Package, ccv3.Warnings, error)
	CreateApplication(app resources.Application) (resources.Application, ccv3.Warnings, error)
	CreateApplicationDeployment(appGUID string, dropletGUID string) (string, ccv3.Warnings, error)
	CreateApplicationDeploymentByRevision(appGUID string, revisionGUID string) (string, ccv3.Warnings, error)
	CreateApplicationProcessScale(appGUID string, process resources.Process) (resources.Process, ccv3.Warnings, error)
	CreateApplicationTask(appGUID string, task resources.Task) (resources.Task, ccv3.Warnings, error)
	CreateBuild(build resources.Build) (resources.Build, ccv3.Warnings, error)
	CreateBuildpack(bp resources.Buildpack) (resources.Buildpack, ccv3.Warnings, error)
	CreateDomain(domain resources.Domain) (resources.Domain, ccv3.Warnings, error)
	CreateDroplet(appGUID string) (resources.Droplet, ccv3.Warnings, error)
	CreateIsolationSegment(isolationSegment resources.IsolationSegment) (resources.IsolationSegment, ccv3.Warnings, error)
	CreateOrganization(orgName string) (resources.Organization, ccv3.Warnings, error)
	CreateOrganizationQuota(orgQuota resources.OrganizationQuota) (resources.OrganizationQuota, ccv3.Warnings, error)
	CreatePackage(pkg resources.Package) (resources.Package, ccv3.Warnings, error)
	CreateRole(role resources.Role) (resources.Role, ccv3.Warnings, error)
	CreateRoute(route resources.Route) (resources.Route, ccv3.Warnings, error)
	CreateServiceBroker(serviceBroker resources.ServiceBroker) (ccv3.JobURL, ccv3.Warnings, error)
	CreateSecurityGroup(securityGroup resources.SecurityGroup) (resources.SecurityGroup, ccv3.Warnings, error)
	CreateSpace(space resources.Space) (resources.Space, ccv3.Warnings, error)
	CreateSpaceQuota(spaceQuota resources.SpaceQuota) (resources.SpaceQuota, ccv3.Warnings, error)
	CreateUser(userGUID string) (resources.User, ccv3.Warnings, error)
	DeleteApplication(guid string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteApplicationProcessInstance(appGUID string, processType string, instanceIndex int) (ccv3.Warnings, error)
	DeleteBuildpack(buildpackGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteDomain(domainGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteIsolationSegment(guid string) (ccv3.Warnings, error)
	DeleteIsolationSegmentOrganization(isolationSegmentGUID string, organizationGUID string) (ccv3.Warnings, error)
	DeleteOrganization(orgGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteOrganizationQuota(quotaGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteOrphanedRoutes(spaceGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteRole(roleGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteRoute(routeGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteSecurityGroup(securityGroupGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteServiceBroker(serviceBrokerGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteServiceInstanceRelationshipsSharedSpace(serviceInstanceGUID string, sharedToSpaceGUID string) (ccv3.Warnings, error)
	DeleteSpaceQuota(spaceQuotaGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteSpace(guid string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteUser(userGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DownloadDroplet(dropletGUID string) ([]byte, ccv3.Warnings, error)
	EntitleIsolationSegmentToOrganizations(isoGUID string, orgGUIDs []string) (resources.RelationshipList, ccv3.Warnings, error)
	GetApplicationByNameAndSpace(appName string, spaceGUID string) (resources.Application, ccv3.Warnings, error)
	GetApplicationDropletCurrent(appGUID string) (resources.Droplet, ccv3.Warnings, error)
	GetApplicationEnvironment(appGUID string) (ccv3.Environment, ccv3.Warnings, error)
	GetApplicationManifest(appGUID string) ([]byte, ccv3.Warnings, error)
	GetApplicationProcessByType(appGUID string, processType string) (resources.Process, ccv3.Warnings, error)
	GetApplicationProcesses(appGUID string) ([]resources.Process, ccv3.Warnings, error)
	GetApplicationRevisions(appGUID string, query ...ccv3.Query) ([]resources.Revision, ccv3.Warnings, error)
	GetApplicationRevisionsDeployed(appGUID string) ([]resources.Revision, ccv3.Warnings, error)
	GetApplicationRoutes(appGUID string) ([]resources.Route, ccv3.Warnings, error)
	GetApplicationTasks(appGUID string, query ...ccv3.Query) ([]resources.Task, ccv3.Warnings, error)
	GetApplications(query ...ccv3.Query) ([]resources.Application, ccv3.Warnings, error)
	GetBuild(guid string) (resources.Build, ccv3.Warnings, error)
	GetBuildpacks(query ...ccv3.Query) ([]resources.Buildpack, ccv3.Warnings, error)
	GetDefaultDomain(orgGuid string) (resources.Domain, ccv3.Warnings, error)
	GetDeployment(guid string) (resources.Deployment, ccv3.Warnings, error)
	GetDeployments(query ...ccv3.Query) ([]resources.Deployment, ccv3.Warnings, error)
	GetDomain(GUID string) (resources.Domain, ccv3.Warnings, error)
	GetDomains(query ...ccv3.Query) ([]resources.Domain, ccv3.Warnings, error)
	GetDroplet(guid string) (resources.Droplet, ccv3.Warnings, error)
	GetDroplets(query ...ccv3.Query) ([]resources.Droplet, ccv3.Warnings, error)
	GetEnvironmentVariableGroup(group constant.EnvironmentVariableGroupName) (resources.EnvironmentVariables, ccv3.Warnings, error)
	GetEvents(query ...ccv3.Query) ([]ccv3.Event, ccv3.Warnings, error)
	GetFeatureFlag(featureFlagName string) (resources.FeatureFlag, ccv3.Warnings, error)
	GetFeatureFlags() ([]resources.FeatureFlag, ccv3.Warnings, error)
	GetInfo() (ccv3.Info, ccv3.ResourceLinks, ccv3.Warnings, error)
	GetIsolationSegment(guid string) (resources.IsolationSegment, ccv3.Warnings, error)
	GetIsolationSegmentOrganizations(isolationSegmentGUID string) ([]resources.Organization, ccv3.Warnings, error)
	GetIsolationSegments(query ...ccv3.Query) ([]resources.IsolationSegment, ccv3.Warnings, error)
	GetNewApplicationProcesses(appGUID string, deploymentGUID string) ([]resources.Process, ccv3.Warnings, error)
	GetOrganization(orgGUID string) (resources.Organization, ccv3.Warnings, error)
	GetOrganizationDefaultIsolationSegment(orgGUID string) (resources.Relationship, ccv3.Warnings, error)
	GetOrganizationDomains(orgGUID string, query ...ccv3.Query) ([]resources.Domain, ccv3.Warnings, error)
	GetOrganizationQuota(quotaGUID string) (resources.OrganizationQuota, ccv3.Warnings, error)
	GetOrganizationQuotas(query ...ccv3.Query) ([]resources.OrganizationQuota, ccv3.Warnings, error)
	GetOrganizations(query ...ccv3.Query) ([]resources.Organization, ccv3.Warnings, error)
	GetPackage(guid string) (resources.Package, ccv3.Warnings, error)
	GetPackages(query ...ccv3.Query) ([]resources.Package, ccv3.Warnings, error)
	GetPackageDroplets(packageGUID string, query ...ccv3.Query) ([]resources.Droplet, ccv3.Warnings, error)
	GetProcess(processGUID string) (resources.Process, ccv3.Warnings, error)
	GetProcesses(query ...ccv3.Query) ([]resources.Process, ccv3.Warnings, error)
	GetProcessInstances(processGUID string) ([]ccv3.ProcessInstance, ccv3.Warnings, error)
	GetProcessSidecars(processGUID string) ([]resources.Sidecar, ccv3.Warnings, error)
	GetRoles(query ...ccv3.Query) ([]resources.Role, ccv3.IncludedResources, ccv3.Warnings, error)
	RootResponse() (ccv3.Info, ccv3.Warnings, error)
	GetRouteDestinations(routeGUID string) ([]resources.RouteDestination, ccv3.Warnings, error)
	GetRoutes(query ...ccv3.Query) ([]resources.Route, ccv3.Warnings, error)
	GetRunningSecurityGroups(spaceGUID string, queries ...ccv3.Query) ([]resources.SecurityGroup, ccv3.Warnings, error)
	GetSecurityGroups(query ...ccv3.Query) ([]resources.SecurityGroup, ccv3.Warnings, error)
	GetServiceBrokers(query ...ccv3.Query) ([]resources.ServiceBroker, ccv3.Warnings, error)
	GetServiceInstances(query ...ccv3.Query) ([]resources.ServiceInstance, ccv3.Warnings, error)
	GetServiceOfferings(query ...ccv3.Query) ([]resources.ServiceOffering, ccv3.Warnings, error)
	GetServiceOfferingByNameAndBroker(serviceOfferingName, serviceBrokerName string) (resources.ServiceOffering, ccv3.Warnings, error)
	GetServicePlans(query ...ccv3.Query) ([]resources.ServicePlan, ccv3.Warnings, error)
	GetServicePlansWithOfferings(query ...ccv3.Query) ([]ccv3.ServiceOfferingWithPlans, ccv3.Warnings, error)
	GetServicePlansWithSpaceAndOrganization(query ...ccv3.Query) ([]ccv3.ServicePlanWithSpaceAndOrganization, ccv3.Warnings, error)
	GetSpaceFeature(spaceGUID string, featureName string) (bool, ccv3.Warnings, error)
	GetSpaceIsolationSegment(spaceGUID string) (resources.Relationship, ccv3.Warnings, error)
	GetSpaceQuota(spaceQuotaGUID string) (resources.SpaceQuota, ccv3.Warnings, error)
	GetSpaces(query ...ccv3.Query) ([]resources.Space, ccv3.IncludedResources, ccv3.Warnings, error)
	GetSpaceQuotas(query ...ccv3.Query) ([]resources.SpaceQuota, ccv3.Warnings, error)
	GetSSHEnabled(appGUID string) (ccv3.SSHEnabled, ccv3.Warnings, error)
	GetAppFeature(appGUID string, featureName string) (resources.ApplicationFeature, ccv3.Warnings, error)
	GetStacks(query ...ccv3.Query) ([]resources.Stack, ccv3.Warnings, error)
	GetStagingSecurityGroups(spaceGUID string, queries ...ccv3.Query) ([]resources.SecurityGroup, ccv3.Warnings, error)
	GetTask(guid string) (resources.Task, ccv3.Warnings, error)
	GetUser(userGUID string) (resources.User, ccv3.Warnings, error)
	GetUsers(query ...ccv3.Query) ([]resources.User, ccv3.Warnings, error)
	MapRoute(routeGUID string, appGUID string) (ccv3.Warnings, error)
	PollJob(jobURL ccv3.JobURL) (ccv3.Warnings, error)
	PurgeServiceOffering(serviceOfferingGUID string) (ccv3.Warnings, error)
	ResourceMatch(resources []ccv3.Resource) ([]ccv3.Resource, ccv3.Warnings, error)
	SetApplicationDroplet(appGUID string, dropletGUID string) (resources.Relationship, ccv3.Warnings, error)
	SharePrivateDomainToOrgs(domainGuid string, sharedOrgs ccv3.SharedOrgs) (ccv3.Warnings, error)
	ShareServiceInstanceToSpaces(serviceInstanceGUID string, spaceGUIDs []string) (resources.RelationshipList, ccv3.Warnings, error)
	TargetCF(settings ccv3.TargetSettings) (ccv3.Info, ccv3.Warnings, error)
	UnbindSecurityGroupRunningSpace(securityGroupGUID string, spaceGUID string) (ccv3.Warnings, error)
	UnbindSecurityGroupStagingSpace(securityGroupGUID string, spaceGUID string) (ccv3.Warnings, error)
	UnmapRoute(routeGUID string, destinationGUID string) (ccv3.Warnings, error)
	UnsharePrivateDomainFromOrg(domainGUID string, sharedOrgGUID string) (ccv3.Warnings, error)
	UpdateAppFeature(appGUID string, enabled bool, featureName string) (ccv3.Warnings, error)
	UpdateApplication(app resources.Application) (resources.Application, ccv3.Warnings, error)
	UpdateApplicationName(newAppName string, appGUID string) (resources.Application, ccv3.Warnings, error)
	UpdateApplicationApplyManifest(appGUID string, rawManifest []byte) (ccv3.JobURL, ccv3.Warnings, error)
	UpdateApplicationEnvironmentVariables(appGUID string, envVars resources.EnvironmentVariables) (resources.EnvironmentVariables, ccv3.Warnings, error)
	UpdateApplicationRestart(appGUID string) (resources.Application, ccv3.Warnings, error)
	UpdateApplicationStart(appGUID string) (resources.Application, ccv3.Warnings, error)
	UpdateApplicationStop(appGUID string) (resources.Application, ccv3.Warnings, error)
	UpdateBuildpack(buildpack resources.Buildpack) (resources.Buildpack, ccv3.Warnings, error)
	UpdateEnvironmentVariableGroup(group constant.EnvironmentVariableGroupName, envVars resources.EnvironmentVariables) (resources.EnvironmentVariables, ccv3.Warnings, error)
	UpdateFeatureFlag(flag resources.FeatureFlag) (resources.FeatureFlag, ccv3.Warnings, error)
	UpdateOrganization(org resources.Organization) (resources.Organization, ccv3.Warnings, error)
	UpdateOrganizationDefaultIsolationSegmentRelationship(orgGUID string, isolationSegmentGUID string) (resources.Relationship, ccv3.Warnings, error)
	UpdateOrganizationQuota(orgQuota resources.OrganizationQuota) (resources.OrganizationQuota, ccv3.Warnings, error)
	UpdateProcess(process resources.Process) (resources.Process, ccv3.Warnings, error)
	UpdateResourceMetadata(resource string, resourceGUID string, metadata resources.Metadata) (ccv3.JobURL, ccv3.Warnings, error)
	UpdateSecurityGroupRunningSpace(securityGroupGUID string, spaceGUIDs []string) (ccv3.Warnings, error)
	UpdateSecurityGroupStagingSpace(securityGroupGUID string, spaceGUIDs []string) (ccv3.Warnings, error)
	UpdateSecurityGroup(securityGroup resources.SecurityGroup) (resources.SecurityGroup, ccv3.Warnings, error)
	UpdateSpace(space resources.Space) (resources.Space, ccv3.Warnings, error)
	UpdateSpaceApplyManifest(spaceGUID string, rawManifest []byte) (ccv3.JobURL, ccv3.Warnings, error)
	UpdateSpaceFeature(spaceGUID string, enabled bool, featureName string) (ccv3.Warnings, error)
	UpdateSpaceIsolationSegmentRelationship(spaceGUID string, isolationSegmentGUID string) (resources.Relationship, ccv3.Warnings, error)
	UpdateSpaceQuota(spaceQuota resources.SpaceQuota) (resources.SpaceQuota, ccv3.Warnings, error)
	UnsetSpaceQuota(spaceQuotaGUID, spaceGUID string) (ccv3.Warnings, error)
	UpdateServiceBroker(serviceBrokerGUID string, serviceBroker resources.ServiceBroker) (ccv3.JobURL, ccv3.Warnings, error)
	UpdateTaskCancel(taskGUID string) (resources.Task, ccv3.Warnings, error)
	UploadBitsPackage(pkg resources.Package, matchedResources []ccv3.Resource, newResources io.Reader, newResourcesLength int64) (resources.Package, ccv3.Warnings, error)
	UploadBuildpack(buildpackGUID string, buildpackPath string, buildpack io.Reader, buildpackLength int64) (ccv3.JobURL, ccv3.Warnings, error)
	UploadDropletBits(dropletGUID string, dropletPath string, droplet io.Reader, dropletLength int64) (ccv3.JobURL, ccv3.Warnings, error)
	UploadPackage(pkg resources.Package, zipFilepath string) (resources.Package, ccv3.Warnings, error)

	servicePlanVisibilityClient
}

type servicePlanVisibilityClient interface {
	GetServicePlanVisibility(servicePlanGUID string) (resources.ServicePlanVisibility, ccv3.Warnings, error)
	UpdateServicePlanVisibility(servicePlanGUID string, visibility resources.ServicePlanVisibility) (resources.ServicePlanVisibility, ccv3.Warnings, error)
	DeleteServicePlanVisibility(servicePlanGUID, organizationGUID string) (ccv3.Warnings, error)
}

// TODO: Split this enormous interface
