/*
 * Copyright (c) 2023 ErSauravAdhikari and GoCaproverAPI contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
 
package crapi

const (
	ResourceOneMb  int64 = 1048576
	ResourceOneCpu int64 = 1000000000

	URLLoginPath                 = "/api/v2/login"
	URLAppListPath               = "/api/v2/user/apps/appDefinitions"
	URLAppRegisterPath           = "/api/v2/user/apps/appDefinitions/register"
	URLUpdateAppPath             = "/api/v2/user/apps/appDefinitions/update"
	URLAppTriggerBuild           = "/api/v2/user/apps/webhooks/triggerbuild"
	URLEnableBaseDomainSslPath   = "/api/v2/user/apps/appDefinitions/enablebasedomainssl"
	URLAddCustomDomainPath       = "/api/v2/user/apps/appDefinitions/customdomain"
	URLEnableCustomDomainSslPath = "/api/v2/user/apps/appDefinitions/enablecustomdomainssl"
	URLAppBuildLog               = "/api/v2/user/apps/appData"
	URLAppDeletePath             = "/api/v2/user/apps/appDefinitions/delete"
)
