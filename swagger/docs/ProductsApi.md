# \ProductsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigurationsGet**](ProductsApi.md#ConfigurationsGet) | **Get** /configurations | Get system configurations.
[**ConfigurationsPut**](ProductsApi.md#ConfigurationsPut) | **Put** /configurations | Modify system configurations.
[**ConfigurationsResetPost**](ProductsApi.md#ConfigurationsResetPost) | **Post** /configurations/reset | Reset system configurations.
[**EmailPingPost**](ProductsApi.md#EmailPingPost) | **Post** /email/ping | Test connection and authentication with email server.
[**InternalSyncregistryPost**](ProductsApi.md#InternalSyncregistryPost) | **Post** /internal/syncregistry | Sync repositories from registry to DB.
[**JobsReplicationGet**](ProductsApi.md#JobsReplicationGet) | **Get** /jobs/replication | List filters jobs according to the policy and repository
[**JobsReplicationIdDelete**](ProductsApi.md#JobsReplicationIdDelete) | **Delete** /jobs/replication/{id} | Delete specific ID job.
[**JobsReplicationIdLogGet**](ProductsApi.md#JobsReplicationIdLogGet) | **Get** /jobs/replication/{id}/log | Get job logs.
[**JobsScanIdLogGet**](ProductsApi.md#JobsScanIdLogGet) | **Get** /jobs/scan/{id}/log | Get job logs.
[**LdapPingPost**](ProductsApi.md#LdapPingPost) | **Post** /ldap/ping | Ping available ldap service.
[**LdapUsersImportPost**](ProductsApi.md#LdapUsersImportPost) | **Post** /ldap/users/import | Import selected available ldap users.
[**LdapUsersSearchPost**](ProductsApi.md#LdapUsersSearchPost) | **Post** /ldap/users/search | Search available ldap users.
[**LogsGet**](ProductsApi.md#LogsGet) | **Get** /logs | Get recent logs of the projects which the user is a member of
[**PoliciesReplicationGet**](ProductsApi.md#PoliciesReplicationGet) | **Get** /policies/replication | List filters policies by name and project_id
[**PoliciesReplicationIdEnablementPut**](ProductsApi.md#PoliciesReplicationIdEnablementPut) | **Put** /policies/replication/{id}/enablement | Put modifies enablement of the policy.
[**PoliciesReplicationIdGet**](ProductsApi.md#PoliciesReplicationIdGet) | **Get** /policies/replication/{id} | Get replication policy.
[**PoliciesReplicationIdPut**](ProductsApi.md#PoliciesReplicationIdPut) | **Put** /policies/replication/{id} | Put modifies name, description, target and enablement of policy.
[**PoliciesReplicationPost**](ProductsApi.md#PoliciesReplicationPost) | **Post** /policies/replication | Post creates a policy
[**ProjectsGet**](ProductsApi.md#ProjectsGet) | **Get** /projects | List projects
[**ProjectsHead**](ProductsApi.md#ProjectsHead) | **Head** /projects | Check if the project name user provided already exists.
[**ProjectsPost**](ProductsApi.md#ProjectsPost) | **Post** /projects | Create a new project.
[**ProjectsProjectIdDelete**](ProductsApi.md#ProjectsProjectIdDelete) | **Delete** /projects/{project_id} | Delete project by projectID
[**ProjectsProjectIdGet**](ProductsApi.md#ProjectsProjectIdGet) | **Get** /projects/{project_id} | Return specific project detail infomation
[**ProjectsProjectIdLogsGet**](ProductsApi.md#ProjectsProjectIdLogsGet) | **Get** /projects/{project_id}/logs | Get access logs accompany with a relevant project.
[**ProjectsProjectIdMembersGet**](ProductsApi.md#ProjectsProjectIdMembersGet) | **Get** /projects/{project_id}/members/ | Return a project&#39;s relevant role members.
[**ProjectsProjectIdMembersPost**](ProductsApi.md#ProjectsProjectIdMembersPost) | **Post** /projects/{project_id}/members/ | Add project role member accompany with relevant project and user.
[**ProjectsProjectIdMembersUserIdDelete**](ProductsApi.md#ProjectsProjectIdMembersUserIdDelete) | **Delete** /projects/{project_id}/members/{user_id} | Delete project role members accompany with relevant project and user.
[**ProjectsProjectIdMembersUserIdGet**](ProductsApi.md#ProjectsProjectIdMembersUserIdGet) | **Get** /projects/{project_id}/members/{user_id} | Return role members accompany with relevant project and user.
[**ProjectsProjectIdMembersUserIdPut**](ProductsApi.md#ProjectsProjectIdMembersUserIdPut) | **Put** /projects/{project_id}/members/{user_id} | Update project role members accompany with relevant project and user.
[**ProjectsProjectIdPublicityPut**](ProductsApi.md#ProjectsProjectIdPublicityPut) | **Put** /projects/{project_id}/publicity | Update properties for a selected project.
[**RepositoriesGet**](ProductsApi.md#RepositoriesGet) | **Get** /repositories | Get repositories accompany with relevant project and repo name.
[**RepositoriesRepoNameDelete**](ProductsApi.md#RepositoriesRepoNameDelete) | **Delete** /repositories/{repo_name} | Delete a repository.
[**RepositoriesRepoNameSignaturesGet**](ProductsApi.md#RepositoriesRepoNameSignaturesGet) | **Get** /repositories/{repo_name}/signatures | Get signature information of a repository
[**RepositoriesRepoNameTagsGet**](ProductsApi.md#RepositoriesRepoNameTagsGet) | **Get** /repositories/{repo_name}/tags | Get tags of a relevant repository.
[**RepositoriesRepoNameTagsTagDelete**](ProductsApi.md#RepositoriesRepoNameTagsTagDelete) | **Delete** /repositories/{repo_name}/tags/{tag} | Delete a tag in a repository.
[**RepositoriesRepoNameTagsTagGet**](ProductsApi.md#RepositoriesRepoNameTagsTagGet) | **Get** /repositories/{repo_name}/tags/{tag} | Get the tag of the repository.
[**RepositoriesRepoNameTagsTagManifestGet**](ProductsApi.md#RepositoriesRepoNameTagsTagManifestGet) | **Get** /repositories/{repo_name}/tags/{tag}/manifest | Get manifests of a relevant repository.
[**RepositoriesRepoNameTagsTagScanPost**](ProductsApi.md#RepositoriesRepoNameTagsTagScanPost) | **Post** /repositories/{repo_name}/tags/{tag}/scan | Scan the image.
[**RepositoriesRepoNameTagsTagVulnerabilityDetailsGet**](ProductsApi.md#RepositoriesRepoNameTagsTagVulnerabilityDetailsGet) | **Get** /repositories/{repo_name}/tags/{tag}/vulnerability/details | Get vulnerability details of the image.
[**RepositoriesTopGet**](ProductsApi.md#RepositoriesTopGet) | **Get** /repositories/top | Get public repositories which are accessed most.
[**SearchGet**](ProductsApi.md#SearchGet) | **Get** /search | Search for projects and repositories
[**StatisticsGet**](ProductsApi.md#StatisticsGet) | **Get** /statistics | Get projects number and repositories number relevant to the user
[**SysteminfoGet**](ProductsApi.md#SysteminfoGet) | **Get** /systeminfo | Get general system info
[**SysteminfoGetcertGet**](ProductsApi.md#SysteminfoGetcertGet) | **Get** /systeminfo/getcert | Get default root certificate under OVA deployment.
[**SysteminfoVolumesGet**](ProductsApi.md#SysteminfoVolumesGet) | **Get** /systeminfo/volumes | Get system volume info (total/free size).
[**TargetsGet**](ProductsApi.md#TargetsGet) | **Get** /targets | List filters targets by name.
[**TargetsIdDelete**](ProductsApi.md#TargetsIdDelete) | **Delete** /targets/{id} | Delete specific replication&#39;s target.
[**TargetsIdGet**](ProductsApi.md#TargetsIdGet) | **Get** /targets/{id} | Get replication&#39;s target.
[**TargetsIdPingPost**](ProductsApi.md#TargetsIdPingPost) | **Post** /targets/{id}/ping | Ping target.
[**TargetsIdPoliciesGet**](ProductsApi.md#TargetsIdPoliciesGet) | **Get** /targets/{id}/policies/ | List the target relevant policies.
[**TargetsIdPut**](ProductsApi.md#TargetsIdPut) | **Put** /targets/{id} | Update replication&#39;s target.
[**TargetsPingPost**](ProductsApi.md#TargetsPingPost) | **Post** /targets/ping | Ping validates target.
[**TargetsPost**](ProductsApi.md#TargetsPost) | **Post** /targets | Create a new replication target.
[**UsersCurrentGet**](ProductsApi.md#UsersCurrentGet) | **Get** /users/current | Get current user info.
[**UsersGet**](ProductsApi.md#UsersGet) | **Get** /users | Get registered users of Harbor.
[**UsersPost**](ProductsApi.md#UsersPost) | **Post** /users | Creates a new user account.
[**UsersUserIdDelete**](ProductsApi.md#UsersUserIdDelete) | **Delete** /users/{user_id} | Mark a registered user as be removed.
[**UsersUserIdGet**](ProductsApi.md#UsersUserIdGet) | **Get** /users/{user_id} | Get a user&#39;s profile.
[**UsersUserIdPasswordPut**](ProductsApi.md#UsersUserIdPasswordPut) | **Put** /users/{user_id}/password | Change the password on a user that already exists.
[**UsersUserIdPut**](ProductsApi.md#UsersUserIdPut) | **Put** /users/{user_id} | Update a registered user to change his profile.
[**UsersUserIdSysadminPut**](ProductsApi.md#UsersUserIdSysadminPut) | **Put** /users/{user_id}/sysadmin | Update a registered user to change to be an administrator of Harbor.


# **ConfigurationsGet**
> Configurations ConfigurationsGet()

Get system configurations.

This endpoint is for retrieving system configurations that only provides for admin user. 


### Parameters
This endpoint does not need any parameter.

### Return type

[**Configurations**](Configurations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigurationsPut**
> ConfigurationsPut($configurations)

Modify system configurations.

This endpoint is for modifying system configurations that only provides for admin user. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurations** | [**Configurations**](Configurations.md)| The configuration map can contain a subset of the attributes of the schema, which are to be updated. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigurationsResetPost**
> ConfigurationsResetPost()

Reset system configurations.

Reset system configurations from environment variables. Can only be accessed by admin user. 


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmailPingPost**
> EmailPingPost($settings)

Test connection and authentication with email server.

Test connection and authentication with email server.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settings** | [**EmailServerSetting**](EmailServerSetting.md)| Email server settings, if some of the settings are not assigned, they will be read from system configuration. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternalSyncregistryPost**
> InternalSyncregistryPost()

Sync repositories from registry to DB.

This endpoint is for syncing all repositories of registry with database.  


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobsReplicationGet**
> []JobStatus JobsReplicationGet($policyId, $num, $endTime, $startTime, $repository, $status, $page, $pageSize)

List filters jobs according to the policy and repository

This endpoint let user list filters jobs according to the policy and repository. (if start_time and end_time are both null, list jobs of last 10 days) 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyId** | **int32**| The ID of the policy that triggered this job. | 
 **num** | **int32**| The return list length number. | [optional] 
 **endTime** | **int64**| The end time of jobs done. (Timestamp) | [optional] 
 **startTime** | **int64**| The start time of jobs. (Timestamp) | [optional] 
 **repository** | **string**| The respond jobs list filter by repository name. | [optional] 
 **status** | **string**| The respond jobs list filter by status. | [optional] 
 **page** | **int32**| The page nubmer, default is 1. | [optional] 
 **pageSize** | **int32**| The size of per page, default is 10, maximum is 100. | [optional] 

### Return type

[**[]JobStatus**](JobStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobsReplicationIdDelete**
> JobsReplicationIdDelete($id)

Delete specific ID job.

This endpoint is aimed to remove specific ID job from jobservice. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| Delete job ID. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobsReplicationIdLogGet**
> JobsReplicationIdLogGet($id)

Get job logs.

This endpoint let user search job logs filtered by specific ID. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| Relevant job ID | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobsScanIdLogGet**
> JobsScanIdLogGet($id)

Get job logs.

This endpoint let user get scan job logs filtered by specific ID. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| Relevant job ID | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapPingPost**
> LdapPingPost($ldapconf)

Ping available ldap service.

This endpoint ping the available ldap service for test related configuration parameters.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapconf** | [**LdapConf**](LdapConf.md)| ldap configuration. support input ldap service configuration. If it&#39;s a empty request, will load current configuration from the system. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapUsersImportPost**
> LdapUsersImportPost($uidList)

Import selected available ldap users.

This endpoint adds the selected available ldap users to harbor based on related configuration parameters from the system. System will try to guess the user email address and realname, add to harbor user information.  If have errors when import user, will return the list of importing failed uid and the failed reason. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uidList** | [**LdapImportUsers**](LdapImportUsers.md)| The uid listed for importing. This list will check users validity of ldap service based on configuration from the system. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapUsersSearchPost**
> []LdapUsers LdapUsersSearchPost($username, $ldapConf)

Search available ldap users.

This endpoint searches the available ldap users based on related configuration parameters. Support searched by input ladp configuration, load configuration from the system and specific filter. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| Registered user ID | [optional] 
 **ldapConf** | [**LdapConf**](LdapConf.md)| ldap search configuration. ldapconf field can input ldap service configuration. If this item are blank, will load default configuration will load current configuration from the system. | [optional] 

### Return type

[**[]LdapUsers**](LdapUsers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogsGet**
> []AccessLog LogsGet($username, $repository, $tag, $operation, $beginTimestamp, $endTimestamp, $page, $pageSize)

Get recent logs of the projects which the user is a member of

This endpoint let user see the recent operation logs of the projects which he is member of  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| Username of the operator. | [optional] 
 **repository** | **string**| The name of repository | [optional] 
 **tag** | **string**| The name of tag | [optional] 
 **operation** | **string**| The operation | [optional] 
 **beginTimestamp** | **string**| The begin timestamp | [optional] 
 **endTimestamp** | **string**| The end timestamp | [optional] 
 **page** | **int32**| The page nubmer, default is 1. | [optional] 
 **pageSize** | **int32**| The size of per page, default is 10, maximum is 100. | [optional] 

### Return type

[**[]AccessLog**](AccessLog.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesReplicationGet**
> []RepPolicy PoliciesReplicationGet($name, $projectId)

List filters policies by name and project_id

This endpoint let user list filters policies by name and project_id, if name and project_id are nil, list returns all policies 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The replication&#39;s policy name. | [optional] 
 **projectId** | **int64**| Relevant project ID. | [optional] 

### Return type

[**[]RepPolicy**](RepPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesReplicationIdEnablementPut**
> PoliciesReplicationIdEnablementPut($id, $enabledflag)

Put modifies enablement of the policy.

This endpoint let user update policy enablement flag. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| policy ID | 
 **enabledflag** | [**RepPolicyEnablementReq**](RepPolicyEnablementReq.md)| The policy enablement flag. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesReplicationIdGet**
> RepPolicy PoliciesReplicationIdGet($id)

Get replication policy.

This endpoint let user search replication policy by specific ID. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| policy ID | 

### Return type

[**RepPolicy**](RepPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesReplicationIdPut**
> PoliciesReplicationIdPut($id, $policyupdate)

Put modifies name, description, target and enablement of policy.

This endpoint let user update policy name, description, target and enablement. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| policy ID | 
 **policyupdate** | [**RepPolicyUpdate**](RepPolicyUpdate.md)| Update policy name, description, target and enablement. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesReplicationPost**
> PoliciesReplicationPost($policyinfo)

Post creates a policy

This endpoint let user creates a policy, and if it is enabled, the replication will be triggered right now. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyinfo** | [**RepPolicyPost**](RepPolicyPost.md)| Create new policy. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsGet**
> []Project ProjectsGet($name, $public, $owner, $page, $pageSize)

List projects

This endpoint returns all projects created by Harbor, and can be filtered by project name. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of project. | [optional] 
 **public** | **bool**| The project is public or private. | [optional] 
 **owner** | **string**| The name of project owner. | [optional] 
 **page** | **int32**| The page nubmer, default is 1. | [optional] 
 **pageSize** | **int32**| The size of per page, default is 10, maximum is 100. | [optional] 

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsHead**
> ProjectsHead($projectName)

Check if the project name user provided already exists.

This endpoint is used to check if the project name user provided already exist. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectName** | **string**| Project name for checking exists. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsPost**
> ProjectsPost($project)

Create a new project.

This endpoint is for user to create a new project. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**ProjectReq**](ProjectReq.md)| New created project. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdDelete**
> ProjectsProjectIdDelete($projectId)

Delete project by projectID

This endpoint is aimed to delete project by project ID. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int64**| Project ID of project which will be deleted. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdGet**
> Project ProjectsProjectIdGet($projectId)

Return specific project detail infomation

This endpoint returns specific project information by project ID. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int64**| Project ID for filtering results. | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdLogsGet**
> []AccessLog ProjectsProjectIdLogsGet($projectId, $username, $repository, $tag, $operation, $beginTimestamp, $endTimestamp, $page, $pageSize)

Get access logs accompany with a relevant project.

This endpoint let user search access logs filtered by operations and date time ranges. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int64**| Relevant project ID | 
 **username** | **string**| Username of the operator. | [optional] 
 **repository** | **string**| The name of repository | [optional] 
 **tag** | **string**| The name of tag | [optional] 
 **operation** | **string**| The operation | [optional] 
 **beginTimestamp** | **string**| The begin timestamp | [optional] 
 **endTimestamp** | **string**| The end timestamp | [optional] 
 **page** | **int32**| The page nubmer, default is 1. | [optional] 
 **pageSize** | **int32**| The size of per page, default is 10, maximum is 100. | [optional] 

### Return type

[**[]AccessLog**](AccessLog.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersGet**
> []User ProjectsProjectIdMembersGet($projectId)

Return a project's relevant role members.

This endpoint is for user to search a specified project's relevant role members. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int64**| Relevant project ID. | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersPost**
> ProjectsProjectIdMembersPost($projectId, $roles)

Add project role member accompany with relevant project and user.

This endpoint is for user to add project role member accompany with relevant project and user. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int64**| Relevant project ID. | 
 **roles** | [**RoleParam**](RoleParam.md)| Role members for adding to relevant project. Only one role is supported in the role list. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersUserIdDelete**
> ProjectsProjectIdMembersUserIdDelete($projectId, $userId)

Delete project role members accompany with relevant project and user.

This endpoint is aimed to remove project role members already added to the relevant project and user. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int64**| Relevant project ID. | 
 **userId** | **int32**| Relevant user ID. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersUserIdGet**
> []Role ProjectsProjectIdMembersUserIdGet($projectId, $userId)

Return role members accompany with relevant project and user.

This endpoint is for user to get role members accompany with relevant project and user.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int64**| Relevant project ID | 
 **userId** | **int32**| Relevant user ID | 

### Return type

[**[]Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersUserIdPut**
> ProjectsProjectIdMembersUserIdPut($projectId, $userId, $roles)

Update project role members accompany with relevant project and user.

This endpoint is for user to update current project role members accompany with relevant project and user. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int64**| Relevant project ID. | 
 **userId** | **int32**| Relevant user ID. | 
 **roles** | [**RoleParam**](RoleParam.md)| Updates for roles and username. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdPublicityPut**
> ProjectsProjectIdPublicityPut($projectId, $project)

Update properties for a selected project.

This endpoint is aimed to toggle a project publicity status. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int64**| Selected project ID. | 
 **project** | [**Project**](Project.md)| Updates of project. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesGet**
> []Repository RepositoriesGet($projectId, $q, $page, $pageSize)

Get repositories accompany with relevant project and repo name.

This endpoint let user search repositories accompanying with relevant project ID and repo name. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32**| Relevant project ID. | 
 **q** | **string**| Repo name for filtering results. | [optional] 
 **page** | **int32**| The page nubmer, default is 1. | [optional] 
 **pageSize** | **int32**| The size of per page, default is 10, maximum is 100. | [optional] 

### Return type

[**[]Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameDelete**
> RepositoriesRepoNameDelete($repoName)

Delete a repository.

This endpoint let user delete a repository with name. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repoName** | **string**| The name of repository which will be deleted. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameSignaturesGet**
> []RepoSignature RepositoriesRepoNameSignaturesGet($repoName)

Get signature information of a repository

This endpoint aims to retrieve signature information of a repository, the data is from the nested notary instance of Harbor. If the repository does not have any signature information in notary, this API will return an empty list with response code 200, instead of 404 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repoName** | **string**| repository name. | 

### Return type

[**[]RepoSignature**](RepoSignature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsGet**
> []DetailedTag RepositoriesRepoNameTagsGet($repoName)

Get tags of a relevant repository.

This endpoint aims to retrieve tags from a relevant repository. If deployed with Notary, the signature property of response represents whether the image is singed or not. If the property is null, the image is unsigned. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repoName** | **string**| Relevant repository name. | 

### Return type

[**[]DetailedTag**](DetailedTag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagDelete**
> RepositoriesRepoNameTagsTagDelete($repoName, $tag)

Delete a tag in a repository.

This endpoint let user delete tags with repo name and tag. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repoName** | **string**| The name of repository which will be deleted. | 
 **tag** | **string**| Tag of a repository. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagGet**
> DetailedTag RepositoriesRepoNameTagsTagGet($repoName, $tag)

Get the tag of the repository.

This endpoint aims to retrieve the tag of the repository. If deployed with Notary, the signature property of response represents whether the image is singed or not. If the property is null, the image is unsigned. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repoName** | **string**| Relevant repository name. | 
 **tag** | **string**| Tag of the repository. | 

### Return type

[**DetailedTag**](DetailedTag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagManifestGet**
> Manifest RepositoriesRepoNameTagsTagManifestGet($repoName, $tag, $version)

Get manifests of a relevant repository.

This endpoint aims to retreive manifests from a relevant repository. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repoName** | **string**| Repository name | 
 **tag** | **string**| Tag name | 
 **version** | **string**| The version of manifest, valid value are \&quot;v1\&quot; and \&quot;v2\&quot;, default is \&quot;v2\&quot; | [optional] 

### Return type

[**Manifest**](Manifest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagScanPost**
> RepositoriesRepoNameTagsTagScanPost($repoName, $tag)

Scan the image.

Trigger jobservice to call Clair API to scan the image identified by the repo_name and tag.  Only project admins have permission to scan images under the project. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repoName** | **string**| Repository name | 
 **tag** | **string**| Tag name | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagVulnerabilityDetailsGet**
> []VulnerabilityItem RepositoriesRepoNameTagsTagVulnerabilityDetailsGet($repoName, $tag)

Get vulnerability details of the image.

Call Clair API to get the vulnerability based on the previous successful scan. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repoName** | **string**| Repository name | 
 **tag** | **string**| Tag name | 

### Return type

[**[]VulnerabilityItem**](VulnerabilityItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesTopGet**
> []Repository RepositoriesTopGet($count)

Get public repositories which are accessed most.

This endpoint aims to let users see the most popular public repositories 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32**| The number of the requested public repositories, default is 10 if not provided. | [optional] 

### Return type

[**[]Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchGet**
> []Search SearchGet($q)

Search for projects and repositories

The Search endpoint returns information about the projects and repositories offered at public status or related to the current logged in user. The response includes the project and repository list in a proper display order. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string**| Search parameter for project and repository name. | 

### Return type

[**[]Search**](Search.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsGet**
> StatisticMap StatisticsGet()

Get projects number and repositories number relevant to the user

This endpoint is aimed to statistic all of the projects number and repositories number relevant to the logined user, also the public projects number and repositories number. If the user is admin, he can also get total projects number and total repositories number. 


### Parameters
This endpoint does not need any parameter.

### Return type

[**StatisticMap**](StatisticMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminfoGet**
> interface{} SysteminfoGet()

Get general system info

This API is for retrieving general system info, this can be called by anonymous request. 


### Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminfoGetcertGet**
> SysteminfoGetcertGet()

Get default root certificate under OVA deployment.

This endpoint is for downloading a default root certificate that only provides for admin user under OVA deployment. 


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminfoVolumesGet**
> interface{} SysteminfoVolumesGet()

Get system volume info (total/free size).

This endpoint is for retrieving system volume info that only provides for admin user. 


### Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsGet**
> []RepTarget TargetsGet($name)

List filters targets by name.

This endpoint let user list filters targets by name, if name is nil, list returns all targets. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The replication&#39;s target name. | [optional] 

### Return type

[**[]RepTarget**](RepTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsIdDelete**
> TargetsIdDelete($id)

Delete specific replication's target.

This endpoint is for to delete specific replication's target. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The replication&#39;s target ID. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsIdGet**
> RepTarget TargetsIdGet($id)

Get replication's target.

This endpoint is for get specific replication's target.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The replication&#39;s target ID. | 

### Return type

[**RepTarget**](RepTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsIdPingPost**
> TargetsIdPingPost($id)

Ping target.

This endpoint is for ping target. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The replication&#39;s target ID. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsIdPoliciesGet**
> []RepPolicy TargetsIdPoliciesGet($id)

List the target relevant policies.

This endpoint list policies filter with specific replication's target ID. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The replication&#39;s target ID. | 

### Return type

[**[]RepPolicy**](RepPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsIdPut**
> TargetsIdPut($id, $repoTarget)

Update replication's target.

This endpoint is for update specific replication's target. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The replication&#39;s target ID. | 
 **repoTarget** | [**PutTarget**](PutTarget.md)| Updates of replication&#39;s target. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsPingPost**
> TargetsPingPost($target)

Ping validates target.

This endpoint is for ping validates whether the target is reachable and whether the credential is valid. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **target** | [**PingTarget**](PingTarget.md)| The target object. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsPost**
> TargetsPost($reptarget)

Create a new replication target.

This endpoint is for user to create a new replication target. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reptarget** | [**RepTargetPost**](RepTargetPost.md)| New created replication target. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersCurrentGet**
> User UsersCurrentGet()

Get current user info.

This endpoint is to get the current user infomation. 


### Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGet**
> []User UsersGet($username, $email, $page, $pageSize)

Get registered users of Harbor.

This endpoint is for user to search registered users, support for filtering results with username.Notice, by now this operation is only for administrator. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| Username for filtering results. | [optional] 
 **email** | **string**| Email for filtering results. | [optional] 
 **page** | **int32**| The page nubmer, default is 1. | [optional] 
 **pageSize** | **int32**| The size of per page. | [optional] 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPost**
> UsersPost($user)

Creates a new user account.

This endpoint is to create a user if the user does not already exist. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md)| New created user. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdDelete**
> UsersUserIdDelete($userId)

Mark a registered user as be removed.

This endpoint let administrator of Harbor mark a registered user as be removed.It actually won't be deleted from DB. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| User ID for marking as to be removed. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdGet**
> UsersUserIdGet($userId)

Get a user's profile.

Get user's profile with user id. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| Registered user ID | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdPasswordPut**
> UsersUserIdPasswordPut($userId, $password)

Change the password on a user that already exists.

This endpoint is for user to update password. Users with the admin role can change any user's password. Guest users can change only their own password. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| Registered user ID. | 
 **password** | [**Password**](Password.md)| Password to be updated. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdPut**
> UsersUserIdPut($userId, $profile)

Update a registered user to change his profile.

This endpoint let a registered user change his profile. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| Registered user ID | 
 **profile** | [**UserProfile**](UserProfile.md)| Only email, realname and comment can be modified. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdSysadminPut**
> UsersUserIdSysadminPut($userId, $hasAdminRole)

Update a registered user to change to be an administrator of Harbor.

This endpoint let a registered user change to be an administrator of Harbor. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| Registered user ID | 
 **hasAdminRole** | [**HasAdminRole**](HasAdminRole.md)| Toggle a user to admin or not. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

