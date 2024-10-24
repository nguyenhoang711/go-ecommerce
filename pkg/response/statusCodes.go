package response

const (
	CODE_CONTINUE = 100

	CODE_SWITCHING_PROTOCOLS = 101

	CODE_PROCESSING = 102

	CODE_OK = 200

	CODE_CREATED = 201

	CODE_ACCEPTED = 202

	CODE_NON_AUTHORITATIVE_INFORMATION = 203

    CODE_NO_CONTENT = 204

    CODE_RESET_CONTENT = 205

    CODE_PARTIAL_CONTENT = 206

    CODE_MULTI_STATUS = 207

    CODE_MULTIPLE_CHOICES = 300

	CODE_MOVED_PERMANENTLY = 301
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.4.3
     *
     * This response code means that URI of requested resource has been changed temporarily. New changes in the URI might be made in the future. Therefore, this same URI should be used by the client in future requests.
     */
	CODE_MOVED_TEMPORARILY = 302
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.4.4
     *
     * Server sent this response to directing client to get requested resource to another URI with an GET request.
     */
	CODE_SEE_OTHER = 303
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7232#section-4.1
     *
     * This is used for caching purposes. It is telling to client that response has not been modified. So, client can continue to use same cached version of response.
     */
	CODE_NOT_MODIFIED = 304
    /**
     * @deprecated
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.4.6
     *
     * Was defined in a previous version of the HTTP specification to indicate that a requested response must be accessed by a proxy. It has been deprecated due to security concerns regarding in-band configuration of a proxy.
     */
    CODE_USE_PROXY = 305
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.4.7
     *
     * Server sent this response to directing client to get requested resource to another URI with same method that used prior request. This has the same semantic than the 302 Found HTTP response code, with the exception that the user agent must not change the HTTP method used = if a POST was used in the first request, a POST must be used in the second request.
     */
    CODE_TEMPORARY_REDIRECT = 307
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7538#section-3
     *
     * This means that the resource is now permanently located at another URI, specified by the Location = HTTP Response header. This has the same semantics as the 301 Moved Permanently HTTP response code, with the exception that the user agent must not change the HTTP method used = if a POST was used in the first request, a POST must be used in the second request.
     */
    CODE_PERMANENT_REDIRECT = 308
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.1
     *
     * This response means that server could not understand the request due to invalid syntax.
     */
    CODE_BAD_REQUEST = 400
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7235#section-3.1
     *
     * Although the HTTP standard specifies "unauthorized", semantically this response means "unauthenticated". That is, the client must authenticate itself to get the requested response.
     */
    CODE_UNAUTHORIZED = 401
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.2
     *
     * This response code is reserved for future use. Initial aim for creating this code was using it for digital payment systems however this is not used currently.
     */
    CODE_PAYMENT_REQUIRED = 402
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.3
     *
     * The client does not have access rights to the content, i.e. they are unauthorized, so server is rejecting to give proper response. Unlike 401, the client's identity is known to the server.
     */
    CODE_FORBIDDEN = 403
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.4
     *
     * The server can not find requested resource. In the browser, this means the URL is not recognized. In an API, this can also mean that the endpoint is valid but the resource itself does not exist. Servers may also send this response instead of 403 to hide the existence of a resource from an unauthorized client. This response code is probably the most famous one due to its frequent occurence on the web.
     */
    CODE_NOT_FOUND = 404
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.5
     *
     * The request method is known by the server but has been disabled and cannot be used. For example, an API may forbid DELETE-ing a resource. The two mandatory methods, GET and HEAD, must never be disabled and should not return this error code.
     */
    CODE_METHOD_NOT_ALLOWED = 405
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.6
     *
     * This response is sent when the web server, after performing server-driven content negotiation, doesn't find any content following the criteria given by the user agent.
     */
    CODE_NOT_ACCEPTABLE = 406
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7235#section-3.2
     *
     * This is similar to 401 but authentication is needed to be done by a proxy.
     */
    CODE_PROXY_AUTHENTICATION_REQUIRED = 407
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.7
     *
     * This response is sent on an idle connection by some servers, even without any previous request by the client. It means that the server would like to shut down this unused connection. This response is used much more since some browsers, like Chrome, Firefox 27+, or IE9, use HTTP pre-connection mechanisms to speed up surfing. Also note that some servers merely shut down the connection without sending this message.
     */
    CODE_REQUEST_TIMEOUT = 408
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.8
     *
     * This response is sent when a request conflicts with the current state of the server.
     */
    CODE_CONFLICT = 409
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.9
     *
     * This response would be sent when the requested content has been permenantly deleted from server, with no forwarding address. Clients are expected to remove their caches and links to the resource. The HTTP specification intends this status code to be used for "limited-time, promotional services". APIs should not feel compelled to indicate resources that have been deleted with this status code.
     */
    CODE_GONE = 410
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.10
     *
     * The server rejected the request because the Content-Length header field is not defined and the server requires it.
     */
    CODE_LENGTH_REQUIRED = 411
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7232#section-4.2
     *
     * The client has indicated preconditions in its headers which the server does not meet.
     */
    CODE_PRECONDITION_FAILED = 412
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.11
     *
     * Request entity is larger than limits defined by server; the server might close the connection or return an Retry-After header field.
     */
    CODE_REQUEST_TOO_LONG = 413
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.12
     *
     * The URI requested by the client is longer than the server is willing to interpret.
     */
    CODE_REQUEST_URI_TOO_LONG = 414
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.13
     *
     * The media format of the requested data is not supported by the server, so the server is rejecting the request.
     */
    CODE_UNSUPPORTED_MEDIA_TYPE = 415
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7233#section-4.4
     *
     * The range specified by the Range header field in the request can't be fulfilled; it's possible that the range is outside the size of the target URI's data.
     */
    CODE_REQUESTED_RANGE_NOT_SATISFIABLE = 416
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.5.14
     *
     * This response code means the expectation indicated by the Expect request header field can't be met by the server.
     */
    CODE_EXPECTATION_FAILED = 417
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc2324#section-2.3.2
     *
     * Any attempt to brew coffee with a teapot should result in the error code "418 I'm a teapot". The resulting entity body MAY be short and stout.
     */
    CODE_IM_A_TEAPOT = 418
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc2518#section-10.6
     *
     * The 507 (Insufficient Storage) status code means the method could not be performed on the resource because the server is unable to store the representation needed to successfully complete the request. This condition is considered to be temporary. If the request which received this status code was the result of a user action, the request MUST NOT be repeated until it is requested by a separate user action.
     */
    CODE_INSUFFICIENT_SPACE_ON_RESOURCE = 419
    /**
     * @deprecated
     * Official Documentation @ https =//tools.ietf.org/rfcdiff?difftype=--hwdiff&url2=draft-ietf-webdav-protocol-06.txt
     *
     * A deprecated response used by the Spring Framework when a method has failed.
     */
    CODE_METHOD_FAILURE = 420
    /**
     * Official Documentation @ https =//datatracker.ietf.org/doc/html/rfc7540#section-9.1.2
     *
     * Defined in the specification of HTTP/2 to indicate that a server is not able to produce a response for the combination of scheme and authority that are included in the request URI.
     */
    CODE_MISDIRECTED_REQUEST = 421
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc2518#section-10.3
     *
     * The request was well-formed but was unable to be followed due to semantic errors.
     */
    CODE_UNPROCESSABLE_ENTITY = 422
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc2518#section-10.4
     *
     * The resource that is being accessed is locked.
     */
    CODE_LOCKED = 423
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc2518#section-10.5
     *
     * The request failed due to failure of a previous request.
     */
    CODE_FAILED_DEPENDENCY = 424
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc6585#section-3
     *
     * The origin server requires the request to be conditional. Intended to prevent the 'lost update' problem, where a client GETs a resource's state, modifies it, and PUTs it back to the server, when meanwhile a third party has modified the state on the server, leading to a conflict.
     */
    CODE_PRECONDITION_REQUIRED = 428
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc6585#section-4
     *
     * The user has sent too many requests in a given amount of time ("rate limiting").
     */
    CODE_TOO_MANY_REQUESTS = 429
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc6585#section-5
     *
     * The server is unwilling to process the request because its header fields are too large. The request MAY be resubmitted after reducing the size of the request header fields.
     */
    CODE_REQUEST_HEADER_FIELDS_TOO_LARGE = 431
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7725
     *
     * The user-agent requested a resource that cannot legally be provided, such as a web page censored by a government.
     */
    CODE_UNAVAILABLE_FOR_LEGAL_REASONS = 451
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.6.1
     *
     * The server encountered an unexpected condition that prevented it from fulfilling the request.
     */
    CODE_INTERNAL_SERVER_ERROR = 500
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.6.2
     *
     * The request method is not supported by the server and cannot be handled. The only methods that servers are required to support (and therefore that must not return this code) are GET and HEAD.
     */
    CODE_NOT_IMPLEMENTED = 501
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.6.3
     *
     * This error response means that the server, while working as a gateway to get a response needed to handle the request, got an invalid response.
     */
    CODE_BAD_GATEWAY = 502
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.6.4
     *
     * The server is not ready to handle the request. Common causes are a server that is down for maintenance or that is overloaded. Note that together with this response, a user-friendly page explaining the problem should be sent. This responses should be used for temporary conditions and the Retry-After = HTTP header should, if possible, contain the estimated time before the recovery of the service. The webmaster must also take care about the caching-related headers that are sent along with this response, as these temporary condition responses should usually not be cached.
     */
    CODE_SERVICE_UNAVAILABLE = 503
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.6.5
     *
     * This error response is given when the server is acting as a gateway and cannot get a response in time.
     */
    CODE_GATEWAY_TIMEOUT = 504
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc7231#section-6.6.6
     *
     * The HTTP version used in the request is not supported by the server.
     */
    CODE_HTTP_VERSION_NOT_SUPPORTED = 505
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc2518#section-10.6
     *
     * The server has an internal configuration error = the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper end point in the negotiation process.
     */
    CODE_INSUFFICIENT_STORAGE = 507
    /**
     * Official Documentation @ https =//tools.ietf.org/html/rfc6585#section-6
     *
     * The 511 status code indicates that the client needs to authenticate to gain network access.
     */
    CODE_NETWORK_AUTHENTICATION_REQUIRED = 511
)