import 'package:grpc/grpc.dart';
import 'package:four_methods_pb/protobuf/four_methods.pbgrpc.dart';

class Client{
  late DemoClient stub;

  Future<void> main(List<String> args) async{
    final channel = ClientChannel(
      '127.0.0.1',
      port: 58888,
      options: const ChannelOptions(credentials: ChannelCredentials.insecure())
    );
    stub = DemoClient(
      channel,
      options: CallOptions(timeout: Duration(seconds: 30))
    );
    try {
      // await runUnaryRPC();
      // await serverStreamRPC();
      // await clientStreamRPC();
      await biStreamRPC();
    } catch (e) {
      print('Caught error: $e');
    }
    await channel.shutdown();
  }

  Future<void> runUnaryRPC() async {
    DemoResponse res = await stub.unaryRPC(DemoRequest(data: "UnaryRPC-req"));
    print(res.data);
  }

  Future<void> serverStreamRPC() async {
    try{
      await for( var res in stub.serverStreamRPC(DemoRequest(data: "ServerStreamRPC"))){
        print(res);
      }
    }
    catch(e){
      print('ERROR: $e');
    }
  }

  Future<void> clientStreamRPC() async {
    Stream<DemoRequest> generateReq() async* {
      for (var i = 0; i < 10; i++) {
        var demoReq = DemoRequest(data: i.toString());
        yield demoReq;
      }
    }

    var res = await stub.clientStreamRPC(generateReq());
    print(res);
  }

  Future<void> biStreamRPC() async {
    Stream<DemoRequest> generateReq() async* {
      for (var i = 0; i < 10; i++) {
        var demoReq = DemoRequest(data: i.toString());
        yield demoReq;
      }
    }
    var res = stub.biStreamRPC(generateReq());
    await for (var r in res ){
      print(r);
    }
  }
}

// 执行客户端
Future<void> main(List<String> args) async {
  await Client().main(args);
}