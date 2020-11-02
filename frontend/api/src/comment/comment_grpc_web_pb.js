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


var validate_validate_pb = require('../validate/validate_pb.js')
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
 *   !proto.comment.CommentRequest,
 *   !proto.comment.CommentReply>}
 */
const methodDescriptor_comment_Write = new grpc.web.MethodDescriptor(
  '/comment.comment/Write',
  grpc.web.MethodType.UNARY,
  proto.comment.CommentRequest,
  proto.comment.CommentReply,
  /**
   * @param {!proto.comment.CommentRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.comment.CommentReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.comment.CommentRequest,
 *   !proto.comment.CommentReply>}
 */
const methodInfo_comment_Write = new grpc.web.AbstractClientBase.MethodInfo(
  proto.comment.CommentReply,
  /**
   * @param {!proto.comment.CommentRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.comment.CommentReply.deserializeBinary
);


/**
 * @param {!proto.comment.CommentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.comment.CommentReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.comment.CommentReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.comment.commentClient.prototype.write =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/comment.comment/Write',
      request,
      metadata || {},
      methodDescriptor_comment_Write,
      callback);
};


/**
 * @param {!proto.comment.CommentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.comment.CommentReply>}
 *     Promise that resolves to the response
 */
proto.comment.commentPromiseClient.prototype.write =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/comment.comment/Write',
      request,
      metadata || {},
      methodDescriptor_comment_Write);
};


module.exports = proto.comment;

