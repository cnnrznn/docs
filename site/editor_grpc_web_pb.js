/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = require('./editor_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.EditorClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.EditorPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.JoinRequest,
 *   !proto.JoinResponse>}
 */
const methodDescriptor_Editor_Join = new grpc.web.MethodDescriptor(
  '/Editor/Join',
  grpc.web.MethodType.UNARY,
  proto.JoinRequest,
  proto.JoinResponse,
  /**
   * @param {!proto.JoinRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.JoinResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.JoinRequest,
 *   !proto.JoinResponse>}
 */
const methodInfo_Editor_Join = new grpc.web.AbstractClientBase.MethodInfo(
  proto.JoinResponse,
  /**
   * @param {!proto.JoinRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.JoinResponse.deserializeBinary
);


/**
 * @param {!proto.JoinRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.JoinResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.JoinResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.EditorClient.prototype.join =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Editor/Join',
      request,
      metadata || {},
      methodDescriptor_Editor_Join,
      callback);
};


/**
 * @param {!proto.JoinRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.JoinResponse>}
 *     A native promise that resolves to the response
 */
proto.EditorPromiseClient.prototype.join =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Editor/Join',
      request,
      metadata || {},
      methodDescriptor_Editor_Join);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.LeaveRequest,
 *   !proto.LeaveResponse>}
 */
const methodDescriptor_Editor_Leave = new grpc.web.MethodDescriptor(
  '/Editor/Leave',
  grpc.web.MethodType.UNARY,
  proto.LeaveRequest,
  proto.LeaveResponse,
  /**
   * @param {!proto.LeaveRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.LeaveResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.LeaveRequest,
 *   !proto.LeaveResponse>}
 */
const methodInfo_Editor_Leave = new grpc.web.AbstractClientBase.MethodInfo(
  proto.LeaveResponse,
  /**
   * @param {!proto.LeaveRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.LeaveResponse.deserializeBinary
);


/**
 * @param {!proto.LeaveRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.LeaveResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.LeaveResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.EditorClient.prototype.leave =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Editor/Leave',
      request,
      metadata || {},
      methodDescriptor_Editor_Leave,
      callback);
};


/**
 * @param {!proto.LeaveRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.LeaveResponse>}
 *     A native promise that resolves to the response
 */
proto.EditorPromiseClient.prototype.leave =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Editor/Leave',
      request,
      metadata || {},
      methodDescriptor_Editor_Leave);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Nil,
 *   !proto.DocState>}
 */
const methodDescriptor_Editor_State = new grpc.web.MethodDescriptor(
  '/Editor/State',
  grpc.web.MethodType.UNARY,
  proto.Nil,
  proto.DocState,
  /**
   * @param {!proto.Nil} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.DocState.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.Nil,
 *   !proto.DocState>}
 */
const methodInfo_Editor_State = new grpc.web.AbstractClientBase.MethodInfo(
  proto.DocState,
  /**
   * @param {!proto.Nil} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.DocState.deserializeBinary
);


/**
 * @param {!proto.Nil} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.DocState)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.DocState>|undefined}
 *     The XHR Node Readable Stream
 */
proto.EditorClient.prototype.state =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Editor/State',
      request,
      metadata || {},
      methodDescriptor_Editor_State,
      callback);
};


/**
 * @param {!proto.Nil} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.DocState>}
 *     A native promise that resolves to the response
 */
proto.EditorPromiseClient.prototype.state =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Editor/State',
      request,
      metadata || {},
      methodDescriptor_Editor_State);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Op,
 *   !proto.Nil>}
 */
const methodDescriptor_Editor_Send = new grpc.web.MethodDescriptor(
  '/Editor/Send',
  grpc.web.MethodType.UNARY,
  proto.Op,
  proto.Nil,
  /**
   * @param {!proto.Op} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Nil.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.Op,
 *   !proto.Nil>}
 */
const methodInfo_Editor_Send = new grpc.web.AbstractClientBase.MethodInfo(
  proto.Nil,
  /**
   * @param {!proto.Op} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Nil.deserializeBinary
);


/**
 * @param {!proto.Op} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.Nil)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.Nil>|undefined}
 *     The XHR Node Readable Stream
 */
proto.EditorClient.prototype.send =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Editor/Send',
      request,
      metadata || {},
      methodDescriptor_Editor_Send,
      callback);
};


/**
 * @param {!proto.Op} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.Nil>}
 *     A native promise that resolves to the response
 */
proto.EditorPromiseClient.prototype.send =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Editor/Send',
      request,
      metadata || {},
      methodDescriptor_Editor_Send);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Version,
 *   !proto.Op>}
 */
const methodDescriptor_Editor_Recv = new grpc.web.MethodDescriptor(
  '/Editor/Recv',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.Version,
  proto.Op,
  /**
   * @param {!proto.Version} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Op.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.Version,
 *   !proto.Op>}
 */
const methodInfo_Editor_Recv = new grpc.web.AbstractClientBase.MethodInfo(
  proto.Op,
  /**
   * @param {!proto.Version} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Op.deserializeBinary
);


/**
 * @param {!proto.Version} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.Op>}
 *     The XHR Node Readable Stream
 */
proto.EditorClient.prototype.recv =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/Editor/Recv',
      request,
      metadata || {},
      methodDescriptor_Editor_Recv);
};


/**
 * @param {!proto.Version} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.Op>}
 *     The XHR Node Readable Stream
 */
proto.EditorPromiseClient.prototype.recv =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/Editor/Recv',
      request,
      metadata || {},
      methodDescriptor_Editor_Recv);
};


module.exports = proto;

