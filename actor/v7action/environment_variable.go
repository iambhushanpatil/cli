package v7action

import (
	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	"code.cloudfoundry.org/cli/types"
)

// EnvironmentVariableGroups represents all environment variables for application
type EnvironmentVariableGroups ccv3.Environment

// EnvironmentVariableGroup represents a CC environment variable group (e.g. staging or running)
type EnvironmentVariableGroup ccv3.EnvironmentVariables

// EnvironmentVariablePair represents an environment variable and its value
// on an application
type EnvironmentVariablePair struct {
	Key   string
	Value string
}

// GetEnvironmentVariableGroup returns the values of an environment variable group.
func (actor *Actor) GetEnvironmentVariableGroup(group constant.EnvironmentVariableGroupName) (EnvironmentVariableGroup, Warnings, error) {
	ccEnvGroup, warnings, err := actor.CloudControllerClient.GetEnvironmentVariableGroup(group)
	return EnvironmentVariableGroup(ccEnvGroup), Warnings(warnings), err
}

// GetEnvironmentVariablesByApplicationNameAndSpace returns the environment
// variables for an application.
func (actor *Actor) GetEnvironmentVariablesByApplicationNameAndSpace(appName string, spaceGUID string) (EnvironmentVariableGroups, Warnings, error) {
	app, warnings, appErr := actor.GetApplicationByNameAndSpace(appName, spaceGUID)
	if appErr != nil {
		return EnvironmentVariableGroups{}, warnings, appErr
	}

	ccEnvGroups, v3Warnings, apiErr := actor.CloudControllerClient.GetApplicationEnvironment(app.GUID)
	warnings = append(warnings, v3Warnings...)
	return EnvironmentVariableGroups(ccEnvGroups), warnings, apiErr
}

// SetEnvironmentVariableByApplicationNameAndSpace adds an
// EnvironmentVariablePair to an application. It must be restarted for changes
// to take effect.
func (actor *Actor) SetEnvironmentVariableByApplicationNameAndSpace(appName string, spaceGUID string, envPair EnvironmentVariablePair) (Warnings, error) {
	app, warnings, err := actor.GetApplicationByNameAndSpace(appName, spaceGUID)
	if err != nil {
		return warnings, err
	}

	_, v3Warnings, apiErr := actor.CloudControllerClient.UpdateApplicationEnvironmentVariables(
		app.GUID,
		ccv3.EnvironmentVariables{
			envPair.Key: {Value: envPair.Value, IsSet: true},
		})
	warnings = append(warnings, v3Warnings...)
	return warnings, apiErr
}

func (actor *Actor) SetEnvironmentVariableGroup(group constant.EnvironmentVariableGroupName, envVars ccv3.EnvironmentVariables) (Warnings, error) {
	warnings := ccv3.Warnings{}
	if len(envVars) == 0 {
		existingEnvVars, getWarnings, err := actor.CloudControllerClient.GetEnvironmentVariableGroup(constant.StagingEnvironmentVariableGroup)

		if err != nil {
			return Warnings(getWarnings), err
		}

		warnings = append(warnings, getWarnings...)

		for k := range existingEnvVars {
			envVars[k] = types.FilteredString{Value: "", IsSet: false}
		}

	}
	_, updateWarnings, err := actor.CloudControllerClient.UpdateEnvironmentVariableGroup(group, envVars) //ccv3.EnvironmentVariables{}

	return Warnings(append(warnings, updateWarnings...)), err
}

// UnsetEnvironmentVariableByApplicationNameAndSpace removes an environment
// variable from an application. It must be restarted for changes to take
// effect.
func (actor *Actor) UnsetEnvironmentVariableByApplicationNameAndSpace(appName string, spaceGUID string, environmentVariableName string) (Warnings, error) {
	app, warnings, appErr := actor.GetApplicationByNameAndSpace(appName, spaceGUID)
	if appErr != nil {
		return warnings, appErr
	}
	envGroups, getWarnings, getErr := actor.CloudControllerClient.GetApplicationEnvironment(app.GUID)
	warnings = append(warnings, getWarnings...)
	if getErr != nil {
		return warnings, getErr
	}

	if _, ok := envGroups.EnvironmentVariables[environmentVariableName]; !ok {
		return warnings, actionerror.EnvironmentVariableNotSetError{EnvironmentVariableName: environmentVariableName}
	}

	_, patchWarnings, patchErr := actor.CloudControllerClient.UpdateApplicationEnvironmentVariables(
		app.GUID,
		ccv3.EnvironmentVariables{
			environmentVariableName: {Value: "", IsSet: false},
		})
	warnings = append(warnings, patchWarnings...)
	return warnings, patchErr
}
