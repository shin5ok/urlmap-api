# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/healthcheck.proto](#proto_healthcheck-proto)
    - [HealthCheckRequest](#grpc-health-v1-HealthCheckRequest)
    - [HealthCheckResponse](#grpc-health-v1-HealthCheckResponse)
  
    - [HealthCheckResponse.ServingStatus](#grpc-health-v1-HealthCheckResponse-ServingStatus)
  
    - [Health](#grpc-health-v1-Health)
  
- [proto/urlmap.proto](#proto_urlmap-proto)
    - [ArrayRedirectData](#urlmap-ArrayRedirectData)
    - [OrgUrl](#urlmap-OrgUrl)
    - [RedirectData](#urlmap-RedirectData)
    - [RedirectData.ValidDate](#urlmap-RedirectData-ValidDate)
    - [RedirectInfo](#urlmap-RedirectInfo)
    - [RedirectPath](#urlmap-RedirectPath)
    - [User](#urlmap-User)
    - [Users](#urlmap-Users)
  
    - [Redirection](#urlmap-Redirection)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto_healthcheck-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/healthcheck.proto



<a name="grpc-health-v1-HealthCheckRequest"></a>

### HealthCheckRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service | [string](#string) |  |  |






<a name="grpc-health-v1-HealthCheckResponse"></a>

### HealthCheckResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [HealthCheckResponse.ServingStatus](#grpc-health-v1-HealthCheckResponse-ServingStatus) |  |  |





 


<a name="grpc-health-v1-HealthCheckResponse-ServingStatus"></a>

### HealthCheckResponse.ServingStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| SERVING | 1 |  |
| NOT_SERVING | 2 |  |
| SERVICE_UNKNOWN | 3 | Used only by the Watch method. |


 

 


<a name="grpc-health-v1-Health"></a>

### Health


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Check | [HealthCheckRequest](#grpc-health-v1-HealthCheckRequest) | [HealthCheckResponse](#grpc-health-v1-HealthCheckResponse) | If the requested service is unknown, the call will fail with status NOT_FOUND. |
| Watch | [HealthCheckRequest](#grpc-health-v1-HealthCheckRequest) | [HealthCheckResponse](#grpc-health-v1-HealthCheckResponse) stream | Performs a watch for the serving status of the requested service. The server will immediately send back a message indicating the current serving status. It will then subsequently send a new message whenever the service&#39;s serving status changes.

If the requested service is unknown when the call is received, the server will send a message setting the serving status to SERVICE_UNKNOWN but will *not* terminate the call. If at some future point, the serving status of the service becomes known, the server will send a new message with the service&#39;s serving status.

If the call terminates with status UNIMPLEMENTED, then clients should assume this method is not supported and should not retry the call. If the call terminates with any other status (including OK), clients should retry the call with appropriate exponential backoff. |

 



<a name="proto_urlmap-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/urlmap.proto



<a name="urlmap-ArrayRedirectData"></a>

### ArrayRedirectData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| redirects | [RedirectData](#urlmap-RedirectData) | repeated |  |






<a name="urlmap-OrgUrl"></a>

### OrgUrl



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org | [string](#string) |  |  |
| notify_to | [string](#string) |  |  |
| slack_url | [string](#string) |  |  |
| email | [string](#string) |  |  |






<a name="urlmap-RedirectData"></a>

### RedirectData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| redirect | [RedirectInfo](#urlmap-RedirectInfo) |  |  |






<a name="urlmap-RedirectData-ValidDate"></a>

### RedirectData.ValidDate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| begin | [string](#string) |  |  |
| end | [string](#string) |  |  |






<a name="urlmap-RedirectInfo"></a>

### RedirectInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  |  |
| redirectPath | [string](#string) |  |  |
| org | [string](#string) |  |  |
| host | [string](#string) |  |  |
| comment | [string](#string) |  |  |
| active | [int32](#int32) |  |  |






<a name="urlmap-RedirectPath"></a>

### RedirectPath



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  |  |
| notify_to | [string](#string) |  |  |






<a name="urlmap-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  |  |
| notify_to | [string](#string) |  |  |
| slack_url | [string](#string) |  |  |
| email | [string](#string) |  |  |






<a name="urlmap-Users"></a>

### Users



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#urlmap-User) | repeated |  |





 

 

 


<a name="urlmap-Redirection"></a>

### Redirection


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetOrgByPath | [RedirectPath](#urlmap-RedirectPath) | [OrgUrl](#urlmap-OrgUrl) |  |
| GetInfoByUser | [User](#urlmap-User) | [ArrayRedirectData](#urlmap-ArrayRedirectData) |  |
| SetInfo | [RedirectData](#urlmap-RedirectData) | [OrgUrl](#urlmap-OrgUrl) |  |
| SetUser | [User](#urlmap-User) | [User](#urlmap-User) |  |
| RemoveUser | [User](#urlmap-User) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| ListUsers | [.google.protobuf.Empty](#google-protobuf-Empty) | [Users](#urlmap-Users) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

