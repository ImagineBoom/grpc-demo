///
//  Generated code. Do not modify.
//  source: four_methods.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'four_methods.pb.dart' as $0;
export 'four_methods.pb.dart';

class DemoClient extends $grpc.Client {
  static final _$unaryRPC = $grpc.ClientMethod<$0.DemoRequest, $0.DemoResponse>(
      '/protobuf.Demo/UnaryRPC',
      ($0.DemoRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.DemoResponse.fromBuffer(value));
  static final _$serverStreamRPC =
      $grpc.ClientMethod<$0.DemoRequest, $0.DemoResponse>(
          '/protobuf.Demo/ServerStreamRPC',
          ($0.DemoRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.DemoResponse.fromBuffer(value));
  static final _$clientStreamRPC =
      $grpc.ClientMethod<$0.DemoRequest, $0.DemoResponse>(
          '/protobuf.Demo/ClientStreamRPC',
          ($0.DemoRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.DemoResponse.fromBuffer(value));
  static final _$biStreamRPC =
      $grpc.ClientMethod<$0.DemoRequest, $0.DemoResponse>(
          '/protobuf.Demo/BiStreamRPC',
          ($0.DemoRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.DemoResponse.fromBuffer(value));

  DemoClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.DemoResponse> unaryRPC($0.DemoRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$unaryRPC, request, options: options);
  }

  $grpc.ResponseStream<$0.DemoResponse> serverStreamRPC($0.DemoRequest request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$serverStreamRPC, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseFuture<$0.DemoResponse> clientStreamRPC(
      $async.Stream<$0.DemoRequest> request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$clientStreamRPC, request, options: options)
        .single;
  }

  $grpc.ResponseStream<$0.DemoResponse> biStreamRPC(
      $async.Stream<$0.DemoRequest> request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$biStreamRPC, request, options: options);
  }
}

abstract class DemoServiceBase extends $grpc.Service {
  $core.String get $name => 'protobuf.Demo';

  DemoServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.DemoRequest, $0.DemoResponse>(
        'UnaryRPC',
        unaryRPC_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.DemoRequest.fromBuffer(value),
        ($0.DemoResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DemoRequest, $0.DemoResponse>(
        'ServerStreamRPC',
        serverStreamRPC_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $0.DemoRequest.fromBuffer(value),
        ($0.DemoResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DemoRequest, $0.DemoResponse>(
        'ClientStreamRPC',
        clientStreamRPC,
        true,
        false,
        ($core.List<$core.int> value) => $0.DemoRequest.fromBuffer(value),
        ($0.DemoResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DemoRequest, $0.DemoResponse>(
        'BiStreamRPC',
        biStreamRPC,
        true,
        true,
        ($core.List<$core.int> value) => $0.DemoRequest.fromBuffer(value),
        ($0.DemoResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.DemoResponse> unaryRPC_Pre(
      $grpc.ServiceCall call, $async.Future<$0.DemoRequest> request) async {
    return unaryRPC(call, await request);
  }

  $async.Stream<$0.DemoResponse> serverStreamRPC_Pre(
      $grpc.ServiceCall call, $async.Future<$0.DemoRequest> request) async* {
    yield* serverStreamRPC(call, await request);
  }

  $async.Future<$0.DemoResponse> unaryRPC(
      $grpc.ServiceCall call, $0.DemoRequest request);
  $async.Stream<$0.DemoResponse> serverStreamRPC(
      $grpc.ServiceCall call, $0.DemoRequest request);
  $async.Future<$0.DemoResponse> clientStreamRPC(
      $grpc.ServiceCall call, $async.Stream<$0.DemoRequest> request);
  $async.Stream<$0.DemoResponse> biStreamRPC(
      $grpc.ServiceCall call, $async.Stream<$0.DemoRequest> request);
}
