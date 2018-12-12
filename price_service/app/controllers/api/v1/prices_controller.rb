module Api
  module V1
    class PricesController < ApplicationController
      def calculate
        render json: { price: Price::Calculate.run!(params).to_f }, status: :ok
      end
    end
  end
end
