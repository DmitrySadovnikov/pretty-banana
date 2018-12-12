module Price
  class Calculate < ActiveInteraction::Base
    TARIFF = {
      currency: 'RUB',
      min_price: 500.0,
      order_price: 250.0,
      minute_price: 20.0,
      km_price: 20.0
    }.freeze

    private

    hash :start_point do
      decimal :lat
      decimal :lng
    end

    hash :end_point do
      decimal :lat
      decimal :lng
    end

    def execute
      result =
        TARIFF[:order_price] +
        TARIFF[:minute_price] * time_in_minutes +
        TARIFF[:km_price] * distance_in_km

      result = result.round(0)
      result < TARIFF[:min_price] ? TARIFF[:min_price] : result
    end

    def direction_service_response
      @direction_service_response ||= Direction::Api.calculate(
        start_point_lat: start_point[:lat],
        start_point_lng: start_point[:lng],
        end_point_lat: end_point[:lat],
        end_point_lng: end_point[:lng]
      )
    end

    def time_in_minutes
      (direction_service_response[:time] / 6000.0).to_d.round(2)
    end

    def distance_in_km
      (direction_service_response[:distance].to_d / 1000.0).round(2)
    end
  end
end
