# Go API client for swagger

MangaDex is an ad-free manga reader offering high-quality images!  Here is some generic stuff about the API #
Authentication You can login with `/auth/login` endpoints. It will return a JWT that remains for 15min and that have a
4h refresh token. # Rating limits We are using rating limits in order to avoid too much calls on our endpoints, here is
how is it configured:  | Endpoint | Calls | Time frame |
|-------------------------|------------------|---------------------------| | `/account/create` | 1 | 60 minutes |
| `/account/activate/{code}` | 30 | 60 minutes | | `/account/activate/resend`, `/account/recover`
, `/account/recover/{code}` | 5 | 60 minutes | | `/auth/login` | 30 | 60 minutes | | `/auth/refresh` | 30 | 60 minutes |
| `/chapter/{id}/read` | 300 | 10 minutes | | `/upload/begin`, `/upload/{id}`, `/upload/{id}/commit` | 30 | 1 minute |
| `PUT /chapter/{id}` | 10 | 1 minute | | `DELETE /chapter/{id}` | 10 | 1 minute | | `POST /manga` | 10 | 60 minutes |
| `PUT /manga/{id}` | 10 | 1 minute | | `DELETE /manga/{id}` | 10 | 10 minutes | | `POST /group` | 10 | 60 minutes |
| `PUT /group/{id}` | 10 | 1 minute | | `DELETE /group/{id}` | 10 | 10 minutes | | `POST /author` | 10 | 60 minutes |
| `PUT /author` | 10 | 1 minutes | | `DELETE /author/{id}` | 10 | 10 minutes | | `POST /captcha/solve` | 10 | 10 minutes
| You can get details about how your rate limit is going by reading following headers:  | Header | Description |
|-------------------------|-----------------------------------------------------------| | X-RateLimit-Limit | Number of
max requests allowed in the current time period | | X-RateLimit-Remaining | Number of remaining requests in the current
time period | | X-RateLimit-Retry-After | Timestamp of end of current time period as UNIX timestamp | # Captchas Some
endpoints may require captchas to proceed, in order to slow down automated malicious traffic. Regular users might see a
couple of captchas, based on the frequency of write requests or on certain endpoints like user signup. Once an endpoint
decides that a captcha needs to be solved, a 403 Forbidden response will be returned, where the error title
is `captcha_required_exception`. The sitekey needed for recaptcha to function is provided in both
the `X-Captcha-Sitekey` header field, as well as in the error context, specified as the `siteKey` parameter. The captcha
result of the client can either be passed into the repeated original request with the `X-Captcha-Result` header or
alternatively the `POST /captcha/solve` endpoint can be called to solve this captcha. The time a solved captcha is
remembered varies across different endpoints and can also be influenced by individual client behavior. Authentication is
not required for the `POST /captcha/solve` endpoint, captchas are tracked separately for client ip and user id. If you
are logged in, you want to send the session token so you validate the captcha for your client ip and user id at the same
time, but it is not required. # Chapter pages processing ## Pages processing When you fetch a chapter response you'll
have 4 fields that you need for pages processing:  | Field | Type | Description |
|-----------------------------|----------|--------------------| | `data.id`                   | `string` | API
identifier | | `data.attributes.hash`      | `string` | MD@H identifier | | `data.attributes.data`      | `array`  |
High quality pages | | `data.attributes.dataSaver` | `array`  | Low quality pages | From this point you miss one thing:
a MD@H backend server to get images from, to get that make a request to `GET /at-home/server/{data.id}`, it will return
the server url to use. Then to build the pages, you have to build your url as
following: `https://{md@h server node}/data/{data.attributes.hash}/{data.attributes.data}`  Or if you want to use the
low quality files: `https://{md@h server node}/data-saver/{data.attributes.hash}/{data.attributes.dataSaver}`  Here is
an example of what the url could looks like:  | Field | Value | |-----------------------------|-------| | `data.id`
| `000002b1-e8de-4281-9781-8e81e869f579` | | `data.attributes.hash`      | `caad0c22434276b9e3e56a78fe2e7993` |
| `data.attributes.data`      | `[\"x1-a87ae6522fa5c244fd76985c7d953ccf3975bec66ce9b8e813549e642b38a47a.png\", ...]` |
| `data.attributes.dataSaver` | `[\"x1-a1d3047dfccd77b3117a86ccf19a9c5403e09baec6a78893ed1d3825d2c71256.jpg\", ...]` |
As a \"fake\" MD@H node we'll use `https://s2.mangadex.org/` server. So for high quality we'll have an URL like
that: https://s2.mangadex.org/data/caad0c22434276b9e3e56a78fe2e7993/x1-a87ae6522fa5c244fd76985c7d953ccf3975bec66ce9b8e813549e642b38a47a.png
And for low
quality: https://s2.mangadex.org/data-saver/caad0c22434276b9e3e56a78fe2e7993/x1-a1d3047dfccd77b3117a86ccf19a9c5403e09baec6a78893ed1d3825d2c71256.jpg
## Report In order to make everything works well, we keep statistics over MD@H nodes and how they perform. In order to
keep theses statistics you have to post data for each page you fetch from a MD@H node. Here is an
example: ```curl POST https://api.mangadex.network/report { \"url\": \"https://s2.mangadex.org/data/caad0c22434276b9e3e56a78fe2e7993/x1-a87ae6522fa5c244fd76985c7d953ccf3975bec66ce9b8e813549e642b38a47a.png\", \"success\": true, \"bytes\": 800000, // size of the loaded image \"duration\": 213, // miliseconds to load the image \"cached\": false, // X-Cache header of the MDAH node == 'HIT' ? } ```
# Static data ## Manga publication demographic | Value | Description | |------------------|---------------------------|
| shonen | Manga is a Shonen | | shoujo | Manga is a Shoujo | | josei | Manga is a Josei | | seinen | Manga is a Seinen
| ## Manga status | Value | Description | |------------------|---------------------------| | ongoing | Manga is still
going on | | completed | Manga is completed | | hiatus | Manga is paused | | abandoned | Manga has been abandoned | ##
Manga reading status | Value | |------------------| | reading | | on_hold | | plan\\_to\\_read | | dropped | | re\\_
reading | | completed | ## Manga content rating | Value | Description | |------------------|---------------------------|
| safe | Safe content | | suggestive | Suggestive content | | erotica | Erotica content | | pornographic | Pornographic
content | ## CustomList visibility | Value | Description | |------------------|---------------------------| | public |
CustomList is public | | private | CustomList is private | ## Relationship types | Value | Description |
|------------------|--------------------------------| | manga | Manga resource | | chapter | Chapter resource | | author
| Author resource | | artist | Author resource (drawers only) | | scanlation_group | ScanlationGroup resource | | tag |
Tag resource | | user | User resource | | custom_list | CustomList resource |

## Overview

This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project. By using
the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API
client.

- API version: 5.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation

Put the package under your project folder and add the following in import:

```golang
import "./mangadex"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.mangadex.org*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountApi* | [**GetAccountActivateCode**](docs/AccountApi.md#getaccountactivatecode) | **
Get** /account/activate/{code} | Activate account
*AccountApi* | [**PostAccountActivateResend**](docs/AccountApi.md#postaccountactivateresend) | **
Post** /account/activate/resend | Resend Activation code
*AccountApi* | [**PostAccountCreate**](docs/AccountApi.md#postaccountcreate) | **Post** /account/create | Create Account
*AccountApi* | [**PostAccountRecover**](docs/AccountApi.md#postaccountrecover) | **
Post** /account/recover | Recover given Account
*AccountApi* | [**PostAccountRecoverCode**](docs/AccountApi.md#postaccountrecovercode) | **
Post** /account/recover/{code} | Complete Account recover
*AtHomeApi* | [**GetAtHomeServerChapterId**](docs/AtHomeApi.md#getathomeserverchapterid) | **
Get** /at-home/server/{chapterId} | Get MD@Home node URL
*AuthApi* | [**GetAuthCheck**](docs/AuthApi.md#getauthcheck) | **Get** /auth/check | Check token
*AuthApi* | [**PostAuthLogin**](docs/AuthApi.md#postauthlogin) | **Post** /auth/login | Login
*AuthApi* | [**PostAuthLogout**](docs/AuthApi.md#postauthlogout) | **Post** /auth/logout | Logout
*AuthApi* | [**PostAuthRefresh**](docs/AuthApi.md#postauthrefresh) | **Post** /auth/refresh | Refresh token
*AuthorApi* | [**DeleteAuthorId**](docs/AuthorApi.md#deleteauthorid) | **Delete** /author/{id} | Delete Author
*AuthorApi* | [**GetAuthor**](docs/AuthorApi.md#getauthor) | **Get** /author | Author list
*AuthorApi* | [**GetAuthorId**](docs/AuthorApi.md#getauthorid) | **Get** /author/{id} | Get Author
*AuthorApi* | [**PostAuthor**](docs/AuthorApi.md#postauthor) | **Post** /author | Create Author
*AuthorApi* | [**PutAuthorId**](docs/AuthorApi.md#putauthorid) | **Put** /author/{id} | Update Author
*CaptchaApi* | [**PostCaptchaSolve**](docs/CaptchaApi.md#postcaptchasolve) | **Post** /captcha/solve | Solve Captcha
*ChapterApi* | [**ChapterIdRead**](docs/ChapterApi.md#chapteridread) | **Post** /chapter/{id}/read | Mark Chapter read
*ChapterApi* | [**ChapterIdUnread**](docs/ChapterApi.md#chapteridunread) | **
Delete** /chapter/{id}/read | Mark Chapter unread
*ChapterApi* | [**GetChapter**](docs/ChapterApi.md#getchapter) | **Get** /chapter | Chapter list
*ChapterApi* | [**GetChapterId**](docs/ChapterApi.md#getchapterid) | **Get** /chapter/{id} | Get Chapter
*ChapterApi* | [**PutChapterId**](docs/ChapterApi.md#putchapterid) | **Put** /chapter/{id} | Update Chapter
*CustomListApi* | [**DeleteListId**](docs/CustomListApi.md#deletelistid) | **Delete** /list/{id} | Delete CustomList
*CustomListApi* | [**DeleteMangaIdListListId**](docs/CustomListApi.md#deletemangaidlistlistid) | **
Delete** /manga/{id}/list/{listId} | Remove Manga in CustomList
*CustomListApi* | [**GetListId**](docs/CustomListApi.md#getlistid) | **Get** /list/{id} | Get CustomList
*CustomListApi* | [**GetListIdFeed**](docs/CustomListApi.md#getlistidfeed) | **
Get** /list/{id}/feed | CustomList Manga feed
*CustomListApi* | [**GetUserIdList**](docs/CustomListApi.md#getuseridlist) | **
Get** /user/{id}/list | Get User&#x27;s CustomList list
*CustomListApi* | [**GetUserList**](docs/CustomListApi.md#getuserlist) | **
Get** /user/list | Get logged User CustomList list
*CustomListApi* | [**PostList**](docs/CustomListApi.md#postlist) | **Post** /list | Create CustomList
*CustomListApi* | [**PostMangaIdListListId**](docs/CustomListApi.md#postmangaidlistlistid) | **
Post** /manga/{id}/list/{listId} | Add Manga in CustomList
*CustomListApi* | [**PutListId**](docs/CustomListApi.md#putlistid) | **Put** /list/{id} | Update CustomList
*FeedApi* | [**GetListIdFeed**](docs/FeedApi.md#getlistidfeed) | **Get** /list/{id}/feed | CustomList Manga feed
*FeedApi* | [**GetUserFollowsMangaFeed**](docs/FeedApi.md#getuserfollowsmangafeed) | **
Get** /user/follows/manga/feed | Get logged User followed Manga feed
*InfrastructureApi* | [**PingGet**](docs/InfrastructureApi.md#pingget) | **Get** /ping | Ping the server
*LegacyApi* | [**PostLegacyMapping**](docs/LegacyApi.md#postlegacymapping) | **
Post** /legacy/mapping | Legacy ID mapping
*MangaApi* | [**DeleteChapterId**](docs/MangaApi.md#deletechapterid) | **Delete** /chapter/{id} | Delete Chapter
*MangaApi* | [**DeleteMangaId**](docs/MangaApi.md#deletemangaid) | **Delete** /manga/{id} | Delete Manga
*MangaApi* | [**DeleteMangaIdFollow**](docs/MangaApi.md#deletemangaidfollow) | **
Delete** /manga/{id}/follow | Unfollow Manga
*MangaApi* | [**DeleteMangaIdListListId**](docs/MangaApi.md#deletemangaidlistlistid) | **
Delete** /manga/{id}/list/{listId} | Remove Manga in CustomList
*MangaApi* | [**GetMangaChapterReadmarkers**](docs/MangaApi.md#getmangachapterreadmarkers) | **
Get** /manga/{id}/read | Manga read markers
*MangaApi* | [**GetMangaId**](docs/MangaApi.md#getmangaid) | **Get** /manga/{id} | View Manga
*MangaApi* | [**GetMangaIdFeed**](docs/MangaApi.md#getmangaidfeed) | **Get** /manga/{id}/feed | Manga feed
*MangaApi* | [**GetMangaRandom**](docs/MangaApi.md#getmangarandom) | **Get** /manga/random | Get a random Manga
*MangaApi* | [**GetMangaStatus**](docs/MangaApi.md#getmangastatus) | **
Get** /manga/status | Get all Manga reading status for logged User
*MangaApi* | [**GetMangaTag**](docs/MangaApi.md#getmangatag) | **Get** /manga/tag | Tag list
*MangaApi* | [**GetSearchManga**](docs/MangaApi.md#getsearchmanga) | **Get** /manga | Manga list
*MangaApi* | [**GetUserFollowsManga**](docs/MangaApi.md#getuserfollowsmanga) | **
Get** /user/follows/manga | Get logged User followed Manga list
*MangaApi* | [**GetUserFollowsMangaFeed**](docs/MangaApi.md#getuserfollowsmangafeed) | **
Get** /user/follows/manga/feed | Get logged User followed Manga feed
*MangaApi* | [**PostManga**](docs/MangaApi.md#postmanga) | **Post** /manga | Create Manga
*MangaApi* | [**PostMangaIdFollow**](docs/MangaApi.md#postmangaidfollow) | **Post** /manga/{id}/follow | Follow Manga
*MangaApi* | [**PostMangaIdListListId**](docs/MangaApi.md#postmangaidlistlistid) | **
Post** /manga/{id}/list/{listId} | Add Manga in CustomList
*MangaApi* | [**PostMangaIdStatus**](docs/MangaApi.md#postmangaidstatus) | **
Post** /manga/{id}/status | Update Manga reading status
*MangaApi* | [**PutMangaId**](docs/MangaApi.md#putmangaid) | **Put** /manga/{id} | Update Manga
*ScanlationGroupApi* | [**DeleteGroupId**](docs/ScanlationGroupApi.md#deletegroupid) | **
Delete** /group/{id} | Delete Scanlation Group
*ScanlationGroupApi* | [**DeleteGroupIdFollow**](docs/ScanlationGroupApi.md#deletegroupidfollow) | **
Delete** /group/{id}/follow | Unfollow Scanlation Group
*ScanlationGroupApi* | [**GetGroupId**](docs/ScanlationGroupApi.md#getgroupid) | **
Get** /group/{id} | View Scanlation Group
*ScanlationGroupApi* | [**GetSearchGroup**](docs/ScanlationGroupApi.md#getsearchgroup) | **
Get** /group | Scanlation Group list
*ScanlationGroupApi* | [**GetUserFollowsGroup**](docs/ScanlationGroupApi.md#getuserfollowsgroup) | **
Get** /user/follows/group | Get logged User followed Groups
*ScanlationGroupApi* | [**PostGroup**](docs/ScanlationGroupApi.md#postgroup) | **Post** /group | Create Scanlation Group
*ScanlationGroupApi* | [**PostGroupIdFollow**](docs/ScanlationGroupApi.md#postgroupidfollow) | **
Post** /group/{id}/follow | Follow Scanlation Group
*ScanlationGroupApi* | [**PutGroupId**](docs/ScanlationGroupApi.md#putgroupid) | **
Put** /group/{id} | Update Scanlation Group
*SearchApi* | [**GetAuthor**](docs/SearchApi.md#getauthor) | **Get** /author | Author list
*SearchApi* | [**GetChapter**](docs/SearchApi.md#getchapter) | **Get** /chapter | Chapter list
*SearchApi* | [**GetSearchGroup**](docs/SearchApi.md#getsearchgroup) | **Get** /group | Scanlation Group list
*SearchApi* | [**GetSearchManga**](docs/SearchApi.md#getsearchmanga) | **Get** /manga | Manga list
*UserApi* | [**GetUserFollowsGroup**](docs/UserApi.md#getuserfollowsgroup) | **
Get** /user/follows/group | Get logged User followed Groups
*UserApi* | [**GetUserFollowsManga**](docs/UserApi.md#getuserfollowsmanga) | **
Get** /user/follows/manga | Get logged User followed Manga list
*UserApi* | [**GetUserFollowsUser**](docs/UserApi.md#getuserfollowsuser) | **
Get** /user/follows/user | Get logged User followed User list
*UserApi* | [**GetUserId**](docs/UserApi.md#getuserid) | **Get** /user/{id} | Get User
*UserApi* | [**GetUserMe**](docs/UserApi.md#getuserme) | **Get** /user/me | Logged User details

## Documentation For Models

- [AccountActivateResponse](docs/AccountActivateResponse.md)
- [Author](docs/Author.md)
- [AuthorAttributes](docs/AuthorAttributes.md)
- [AuthorCreate](docs/AuthorCreate.md)
- [AuthorEdit](docs/AuthorEdit.md)
- [AuthorList](docs/AuthorList.md)
- [AuthorResponse](docs/AuthorResponse.md)
- [Body](docs/Body.md)
- [Chapter](docs/Chapter.md)
- [ChapterAttributes](docs/ChapterAttributes.md)
- [ChapterEdit](docs/ChapterEdit.md)
- [ChapterList](docs/ChapterList.md)
- [ChapterRequest](docs/ChapterRequest.md)
- [ChapterResponse](docs/ChapterResponse.md)
- [CheckResponse](docs/CheckResponse.md)
- [CreateAccount](docs/CreateAccount.md)
- [CreateScanlationGroup](docs/CreateScanlationGroup.md)
- [CustomList](docs/CustomList.md)
- [CustomListAttributes](docs/CustomListAttributes.md)
- [CustomListCreate](docs/CustomListCreate.md)
- [CustomListEdit](docs/CustomListEdit.md)
- [CustomListList](docs/CustomListList.md)
- [CustomListResponse](docs/CustomListResponse.md)
- [ErrorResponse](docs/ErrorResponse.md)
- [InlineResponse200](docs/InlineResponse200.md)
- [InlineResponse2001](docs/InlineResponse2001.md)
- [InlineResponse2002](docs/InlineResponse2002.md)
- [InlineResponse2003](docs/InlineResponse2003.md)
- [Login](docs/Login.md)
- [LoginResponse](docs/LoginResponse.md)
- [LoginResponseToken](docs/LoginResponseToken.md)
- [LogoutResponse](docs/LogoutResponse.md)
- [Manga](docs/Manga.md)
- [MangaAttributes](docs/MangaAttributes.md)
- [MangaCreate](docs/MangaCreate.md)
- [MangaEdit](docs/MangaEdit.md)
- [MangaList](docs/MangaList.md)
- [MangaRequest](docs/MangaRequest.md)
- [MangaResponse](docs/MangaResponse.md)
- [MappingId](docs/MappingId.md)
- [MappingIdAttributes](docs/MappingIdAttributes.md)
- [MappingIdBody](docs/MappingIdBody.md)
- [MappingIdResponse](docs/MappingIdResponse.md)
- [ModelError](docs/ModelError.md)
- [Order](docs/Order.md)
- [Order1](docs/Order1.md)
- [RecoverCompleteBody](docs/RecoverCompleteBody.md)
- [RefreshResponse](docs/RefreshResponse.md)
- [RefreshToken](docs/RefreshToken.md)
- [Relationship](docs/Relationship.md)
- [Response](docs/Response.md)
- [ScanlationGroup](docs/ScanlationGroup.md)
- [ScanlationGroupAttributes](docs/ScanlationGroupAttributes.md)
- [ScanlationGroupEdit](docs/ScanlationGroupEdit.md)
- [ScanlationGroupList](docs/ScanlationGroupList.md)
- [ScanlationGroupResponse](docs/ScanlationGroupResponse.md)
- [ScanlationGroupResponseRelationships](docs/ScanlationGroupResponseRelationships.md)
- [SendAccountActivationCode](docs/SendAccountActivationCode.md)
- [Tag](docs/Tag.md)
- [TagAttributes](docs/TagAttributes.md)
- [TagResponse](docs/TagResponse.md)
- [UpdateMangaStatus](docs/UpdateMangaStatus.md)
- [User](docs/User.md)
- [UserAttributes](docs/UserAttributes.md)
- [UserList](docs/UserList.md)
- [UserResponse](docs/UserResponse.md)

## Documentation For Authorization

## Bearer

## Author

mangadexstaff@gmail.com
