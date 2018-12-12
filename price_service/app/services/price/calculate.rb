module Price
  class Calculate < ActiveInteraction::Base
    DEFAULT_TARIFF = YAML.load_file('db/data/tariff.yml').deep_symbolize_keys

    private

    hash :tariff, default: {} do
      decimal :min_price,     default: DEFAULT_TARIFF[:min_price]
      decimal :order_price,   default: DEFAULT_TARIFF[:order_price]
      decimal :minute_price,  default: DEFAULT_TARIFF[:minute_price]
      decimal :km_price,      default: DEFAULT_TARIFF[:km_price]
    end

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
        tariff[:order_price] +
        tariff[:minute_price] * time_in_minutes +
        tariff[:km_price] * distance_in_km

      result = result.round(0)
      result < tariff[:min_price] ? tariff[:min_price] : result
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
