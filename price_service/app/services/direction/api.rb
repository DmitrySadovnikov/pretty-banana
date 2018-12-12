module Direction
  module Api
    module_function

    def calculate(start_point_lat:, start_point_lng:, end_point_lat:, end_point_lng:)
      url = "#{Rails.application.config.direction_service_url}/api/v1/directions/calculate"
      body = {
        start_point: {
          lat: start_point_lat.to_f,
          lng: start_point_lng.to_f
        },
        end_point: {
          lat: end_point_lat.to_f,
          lng: end_point_lng.to_f
        }
      }
      options[:body] = body.to_json
      HTTParty.post(url, options).to_h.deep_symbolize_keys
    end

    def options
      @options ||= {
        headers: headers,
        query: {}
      }
    end

    def headers
      {
        'Content-Type' => 'application/json'
      }
    end
  end
end
