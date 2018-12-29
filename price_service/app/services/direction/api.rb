module Direction
  module Api
    module_function

    def calculate(start_point_lat:, start_point_lng:, end_point_lat:, end_point_lng:)
      request = DirectionPb::Calculate::Request.new(
        startPoint: DirectionPb::Calculate::Point.new(
          lat: start_point_lat.to_f,
          lng: start_point_lng.to_f
        ),
        endPoint: DirectionPb::Calculate::Point.new(
          lat: end_point_lat.to_f,
          lng: end_point_lng.to_f
        )
      )

      response = stub.calculate(request)
      response.to_h
    end

    def stub
      @stub ||= DirectionPb::Direction::Stub.new(
        Rails.application.config.grpc_server_url,
        :this_channel_is_insecure
      )
    end
  end
end
