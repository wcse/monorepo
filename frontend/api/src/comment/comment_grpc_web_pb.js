/**
 * @fileoverview gRPC-Web generated client stub for comment
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.comment = require('./comment_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.comment.commentClient =
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
proto.comment.commentPromiseClient =
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
 *   !proto.comment.PostRequest,
 *   !proto.comment.PostReply>}
 */
const methodDescriptor_comment_Post = new grpc.web.MethodDescriptor(
  '/comment.comment/Post',
  grpc.web.MethodType.UNARY,
  proto.comment.PostRequest,
  proto.comment.PostReply,
  /**
   * @param {!proto.comment.PostRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.comment.PostReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.comment.PostRequest,
 *   !proto.comment.PostReply>}
 */
const methodInfo_comment_Post = new grpc.web.AbstractClientBase.MethodInfo(
  proto.comment.PostReply,
  /**
   * @param {!proto.comment.PostRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.comment.PostReply.deserializeBinary
);


/**
 * @param {!proto.comment.PostRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.comment.PostReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.comment.PostReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.comment.commentClient.prototype.post =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/comment.comment/Post',
      request,
      metadata || {},
      methodDescriptor_comment_Post,
      callback);
};


/**
 * @param {!proto.comment.PostRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.comment.PostReply>}
 *     Promise that resolves to the response
 */
proto.comment.commentPromiseClient.prototype.post =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/comment.comment/Post',
      request,
      metadata || {},
      methodDescriptor_comment_Post);
};


module.exports = proto.comment;

