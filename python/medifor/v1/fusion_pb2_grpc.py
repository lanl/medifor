# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from medifor.v1 import analytic_pb2 as medifor_dot_v1_dot_analytic__pb2
from medifor.v1 import fusion_pb2 as medifor_dot_v1_dot_fusion__pb2


class FuserStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.FuseImageManipulation = channel.unary_unary(
        '/mediforproto.Fuser/FuseImageManipulation',
        request_serializer=medifor_dot_v1_dot_fusion__pb2.FuseImageManipulationRequest.SerializeToString,
        response_deserializer=medifor_dot_v1_dot_analytic__pb2.ImageManipulation.FromString,
        )
    self.FuseImageSplice = channel.unary_unary(
        '/mediforproto.Fuser/FuseImageSplice',
        request_serializer=medifor_dot_v1_dot_fusion__pb2.FuseImageSpliceRequest.SerializeToString,
        response_deserializer=medifor_dot_v1_dot_analytic__pb2.ImageSplice.FromString,
        )
    self.FuseVideoManipulation = channel.unary_unary(
        '/mediforproto.Fuser/FuseVideoManipulation',
        request_serializer=medifor_dot_v1_dot_fusion__pb2.FuseVideoManipulationRequest.SerializeToString,
        response_deserializer=medifor_dot_v1_dot_analytic__pb2.VideoManipulation.FromString,
        )
    self.FuseImageCameraMatch = channel.unary_unary(
        '/mediforproto.Fuser/FuseImageCameraMatch',
        request_serializer=medifor_dot_v1_dot_fusion__pb2.FuseImageCameraMatchRequest.SerializeToString,
        response_deserializer=medifor_dot_v1_dot_analytic__pb2.ImageCameraMatch.FromString,
        )
    self.FuseVideoCameraMatch = channel.unary_unary(
        '/mediforproto.Fuser/FuseVideoCameraMatch',
        request_serializer=medifor_dot_v1_dot_fusion__pb2.FuseVideoCameraMatchRequest.SerializeToString,
        response_deserializer=medifor_dot_v1_dot_analytic__pb2.VideoCameraMatch.FromString,
        )
    self.Kill = channel.unary_unary(
        '/mediforproto.Fuser/Kill',
        request_serializer=medifor_dot_v1_dot_analytic__pb2.Empty.SerializeToString,
        response_deserializer=medifor_dot_v1_dot_analytic__pb2.Empty.FromString,
        )


class FuserServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def FuseImageManipulation(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def FuseImageSplice(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def FuseVideoManipulation(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def FuseImageCameraMatch(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def FuseVideoCameraMatch(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Kill(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_FuserServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'FuseImageManipulation': grpc.unary_unary_rpc_method_handler(
          servicer.FuseImageManipulation,
          request_deserializer=medifor_dot_v1_dot_fusion__pb2.FuseImageManipulationRequest.FromString,
          response_serializer=medifor_dot_v1_dot_analytic__pb2.ImageManipulation.SerializeToString,
      ),
      'FuseImageSplice': grpc.unary_unary_rpc_method_handler(
          servicer.FuseImageSplice,
          request_deserializer=medifor_dot_v1_dot_fusion__pb2.FuseImageSpliceRequest.FromString,
          response_serializer=medifor_dot_v1_dot_analytic__pb2.ImageSplice.SerializeToString,
      ),
      'FuseVideoManipulation': grpc.unary_unary_rpc_method_handler(
          servicer.FuseVideoManipulation,
          request_deserializer=medifor_dot_v1_dot_fusion__pb2.FuseVideoManipulationRequest.FromString,
          response_serializer=medifor_dot_v1_dot_analytic__pb2.VideoManipulation.SerializeToString,
      ),
      'FuseImageCameraMatch': grpc.unary_unary_rpc_method_handler(
          servicer.FuseImageCameraMatch,
          request_deserializer=medifor_dot_v1_dot_fusion__pb2.FuseImageCameraMatchRequest.FromString,
          response_serializer=medifor_dot_v1_dot_analytic__pb2.ImageCameraMatch.SerializeToString,
      ),
      'FuseVideoCameraMatch': grpc.unary_unary_rpc_method_handler(
          servicer.FuseVideoCameraMatch,
          request_deserializer=medifor_dot_v1_dot_fusion__pb2.FuseVideoCameraMatchRequest.FromString,
          response_serializer=medifor_dot_v1_dot_analytic__pb2.VideoCameraMatch.SerializeToString,
      ),
      'Kill': grpc.unary_unary_rpc_method_handler(
          servicer.Kill,
          request_deserializer=medifor_dot_v1_dot_analytic__pb2.Empty.FromString,
          response_serializer=medifor_dot_v1_dot_analytic__pb2.Empty.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'mediforproto.Fuser', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
