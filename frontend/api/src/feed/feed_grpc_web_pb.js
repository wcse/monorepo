/**
 * @fileoverview gRPC-Web generated client stub for feed
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.feed = require('./feed_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.feed.FeedClient =
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
proto.feed.FeedPromiseClient =
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
 *   !proto.feed.PostRequest,
 *   !proto.feed.PostReply>}
 */
const methodDescriptor_Feed_Post = new grpc.web.MethodDescriptor(
  '/feed.Feed/Post',
  grpc.web.MethodType.UNARY,
  proto.feed.PostRequest,
  proto.feed.PostReply,
  /**
   * @param {!proto.feed.PostRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.feed.PostReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.feed.PostRequest,
 *   !proto.feed.PostReply>}
 */
const methodInfo_Feed_Post = new grpc.web.AbstractClientBase.MethodInfo(
  proto.feed.PostReply,
  /**
   * @param {!proto.feed.PostRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.feed.PostReply.deserializeBinary
);


/**
 * @param {!proto.feed.PostRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.feed.PostReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.feed.PostReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.feed.FeedClient.prototype.post =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/feed.Feed/Post',
      request,
      metadata || {},
      methodDescriptor_Feed_Post,
      callback);
};


/**
 * @param {!proto.feed.PostRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.feed.PostReply>}
 *     Promise that resolves to the response
 */
proto.feed.FeedPromiseClient.prototype.post =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/feed.Feed/Post',
      request,
      metadata || {},
      methodDescriptor_Feed_Post);
};


module.exports = proto.feed;

