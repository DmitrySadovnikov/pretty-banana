# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: direction.proto for package 'DirectionPb'

require 'grpc'
require 'direction_pb'

module DirectionPb
  module Direction
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'DirectionPb.Direction'

      rpc :Calculate, Calculate::Request, Calculate::Response
    end

    Stub = Service.rpc_stub_class
  end
end
